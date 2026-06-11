package vodka

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	engine      *Engine
}

// httprouter wrapper
type Engine struct {
	router         *httprouter.Router
	WSConfig       *WSConfig
	trustedProxies []*net.IPNet
	templates      map[string]*template.Template
	templatesMu    sync.RWMutex
	*RouterGroup
	lifecycle      *LifecycleManager
}

// creates a new router
func NewRouter() *Engine {
	router := httprouter.New()
	router.HandleOPTIONS = false
	engine := &Engine{
		router:    router,
		WSConfig:  DefaultWSConfig(),
		lifecycle: NewLifecycleManager(),
	}

	engine.RouterGroup = &RouterGroup{
		prefix:      "",
		middlewares: make([]HandlerFunc, 0),
		engine:      engine,
	}

	return engine
}

func DefaultRouter() *Engine {
	engine := NewRouter()
	engine.Use(Logger(), Recovery(), ErrorHandler())
	return engine
}

func (rg *RouterGroup) Group(prefix string, middlewares ...HandlerFunc) *RouterGroup {
	newMiddlewares := make([]HandlerFunc, len(rg.middlewares), len(rg.middlewares)+len(middlewares))
	copy(newMiddlewares, rg.middlewares)
	newMiddlewares = append(newMiddlewares, middlewares...)

	return &RouterGroup{
		prefix:      rg.prefix + prefix,
		middlewares: newMiddlewares,
		engine:      rg.engine,
	}
}

func (rg *RouterGroup) Use(middlewares ...HandlerFunc) {
	rg.middlewares = append(rg.middlewares, middlewares...)
}

// Runs the http server
// Runs the http server with startup/shutdown hook management and graceful shutdown
func (e *Engine) Run(addr string) error {
	if addr == "" {
		addr = ":8080"
	}

	// 1. Execute startup hooks before serving
	if err := e.lifecycle.runStartupHooks(); err != nil {
		return err
	}

	log.Printf(Green+"Pouring Vodka on %s\n"+Reset, addr)

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	srv := &http.Server{
		Addr:    addr,
		Handler: e,
	}

	serverErr := make(chan error, 1)
	go func() {
		if err := srv.Serve(ln); err != nil && err != http.ErrServerClosed {
			serverErr <- err
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErr:
		// If server failed to start/serve, perform shutdown to clean up hooks
		shutdownErr := e.shutdown(srv)
		if shutdownErr != nil {
			return fmt.Errorf("server error: %v; shutdown errors: %v", err, shutdownErr)
		}
		return err
	case sig := <-quit:
		log.Printf(Yellow+"Received signal: %v. Initiating graceful shutdown...\n"+Reset, sig)
		return e.shutdown(srv)
	}
}

// shutdown handles graceful HTTP server shutdown and executes registered shutdown hooks.
func (e *Engine) shutdown(srv *http.Server) error {
	e.lifecycle.mu.Lock()
	timeout := e.lifecycle.timeout
	hooks := make([]lifecycleHook, len(e.lifecycle.shutdownHooks))
	copy(hooks, e.lifecycle.shutdownHooks)
	e.lifecycle.mu.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var shutdownErrors []error

	// Step 2: Gracefully finish active requests
	if err := srv.Shutdown(ctx); err != nil {
		shutdownErrors = append(shutdownErrors, fmt.Errorf("server shutdown failed: %w", err))
	}

	// Step 3: Sort shutdown hooks by priority (descending), preserving registration order for equal priorities.
	sort.SliceStable(hooks, func(i, j int) bool {
		if hooks[i].priority != hooks[j].priority {
			return hooks[i].priority > hooks[j].priority
		}
		return hooks[i].order < hooks[j].order
	})

	// Step 4: Execute registered shutdown hooks sequentially
	for _, hook := range hooks {
		if err := ctx.Err(); err != nil {
			shutdownErrors = append(shutdownErrors, fmt.Errorf("shutdown timeout exceeded: %w", err))
			break
		}
		if err := hook.fn(ctx); err != nil {
			shutdownErrors = append(shutdownErrors, err)
		}
	}

	if len(shutdownErrors) > 0 {
		return &ShutdownError{Errors: shutdownErrors}
	}

	return nil
}

func (e *Engine) OnStart(fn func() error) {
	e.lifecycle.RegisterStart(fn)
}

func (e *Engine) OnShutdown(fn func(context.Context) error) {
	e.lifecycle.RegisterShutdown(0, fn)
}

func (e *Engine) OnShutdownWithPriority(priority int, fn func(context.Context) error) {
	e.lifecycle.RegisterShutdown(priority, fn)
}

func (e *Engine) SetShutdownTimeout(timeout time.Duration) {
	e.lifecycle.SetTimeout(timeout)
}

// LoadHTMLGlob parses and caches templates from a glob pattern
func (e *Engine) LoadHTMLGlob(pattern string) error {
	e.templatesMu.Lock()
	defer e.templatesMu.Unlock()

	if e.templates == nil {
		e.templates = make(map[string]*template.Template)
	}

	tmpl, err := template.ParseGlob(pattern)
	if err != nil {
		return err
	}

	// Store parsed templates
	for _, t := range tmpl.Templates() {
		e.templates[t.Name()] = t
	}
	return nil
}

// LoadHTMLFiles parses and caches specific template files
func (e *Engine) LoadHTMLFiles(files ...string) error {
	e.templatesMu.Lock()
	defer e.templatesMu.Unlock()

	if e.templates == nil {
		e.templates = make(map[string]*template.Template)
	}

	for _, file := range files {
		tmpl, err := template.ParseFiles(file)
		if err != nil {
			return err
		}
		for _, t := range tmpl.Templates() {
			e.templates[t.Name()] = t
		}
	}
	return nil
}

// getTemplate returns a cached template by name
func (e *Engine) getTemplate(name string) *template.Template {
	e.templatesMu.RLock()
	defer e.templatesMu.RUnlock()
	if e.templates == nil {
		return nil
	}
	return e.templates[name]
}

// Serve Static files
func (rg *RouterGroup) Static(relativePath string, root string) {
	urlPattern := path.Join(relativePath, "/*filepath")

	fileServer := http.FileServer(http.Dir(root))

	rg.engine.router.GET(urlPattern, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		filepath := params.ByName("filepath")
		fullPath := path.Join(root, filepath)

		// Check if the requested file actually exists on the disk
		info, err := os.Stat(fullPath)

		// If the file doesn't exist OR it's a directory, serve index.html (React's entry point)
		if os.IsNotExist(err) || info.IsDir() {
			http.ServeFile(w, r, path.Join(root, "index.html"))
			return
		}

		// Otherwise, serve the actual file (css, js, images)
		// We use StripPrefix so /static/js/main.js looks in ./public/js/main.js
		http.StripPrefix(rg.prefix+relativePath, fileServer).ServeHTTP(w, r)
	})
}

// SPA fallback
func (e *Engine) ServeSPA(root string) {
	fs := http.FileServer(http.Dir(root))

	e.router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Clean path to prevent directory traversal attacks
		cleanedPath := filepath.Clean(r.URL.Path)
		absolutePath := filepath.Join(root, cleanedPath)

		// Check if the actual file (like a .js or .css file) exists
		info, err := os.Stat(absolutePath)
		if os.IsNotExist(err) || info.IsDir() {
			// File not found? It's probably a React frontend route (like /dashboard).
			// Serve index.html and let React handle the routing.
			http.ServeFile(w, r, filepath.Join(root, "index.html"))
			return
		}

		fs.ServeHTTP(w, r)
	})

}

// ServeHTTP intercepts every request before it hits the router
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodOptions {
		e.handlePreFlight(w, req)
		return
	}
	e.router.ServeHTTP(w, req)
}

func (e *Engine) handlePreFlight(w http.ResponseWriter, req *http.Request) {
	c := contextPool.Get().(*Context)
	// Passing e.middlewares so global middlewares like AllowCORS execute
	c.Initialize(w, req, nil, e.middlewares, e)

	defer func() {
		c.Reset()
		contextPool.Put(c)
	}()

	c.Next()

	// If the middleware didn't abort (e.g. no CORS middleware), fallback to 204
	if !c.isAborted {
		c.Writer.WriteHeader(http.StatusNoContent)
	}
}

func (rg *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	absolutePath := rg.prefix + comp

	handlers := make([]HandlerFunc, 0, len(rg.middlewares)+1)
	handlers = append(handlers, rg.middlewares...)
	handlers = append(handlers, handler)

	rg.engine.router.Handle(method, absolutePath, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

		if len(handlers) == 1 {
			c := contextPool.Get().(*Context)

			c.Writer = w
			c.Request = r
			c.Params = params
			c.engine = rg.engine

			handlers[0](c)

			c.Reset()
			contextPool.Put(c)
			return
		}

		c := contextPool.Get().(*Context)
		c.Initialize(w, r, params, handlers, rg.engine)

		defer func() {
			c.Reset()
			contextPool.Put(c)
		}()

		c.Next()
	})
}

func (rg *RouterGroup) GET(path string, handler HandlerFunc) {
	rg.addRoute(http.MethodGet, path, handler)
}

func (rg *RouterGroup) POST(path string, handler HandlerFunc) {
	rg.addRoute(http.MethodPost, path, handler)
}

func (rg *RouterGroup) PUT(path string, handler HandlerFunc) {
	rg.addRoute(http.MethodPut, path, handler)
}

func (rg *RouterGroup) DELETE(path string, handler HandlerFunc) {
	rg.addRoute(http.MethodDelete, path, handler)
}

func (rg *RouterGroup) PATCH(path string, handler HandlerFunc) {
	rg.addRoute(http.MethodPatch, path, handler)
}

func (rg *RouterGroup) HEAD(path string, handler HandlerFunc) {
	rg.addRoute(http.MethodHead, path, handler)
}

// AllowWSOrigins whitelists the given origins for WebSocket upgrade requests.
// Call this before registering WS routes.
//
//	app.AllowWSOrigins([]string{"https://userapp.com", "https://admin.com"})
func (e *Engine) AllowWSOrigins(origins []string) {
	e.WSConfig.CheckOrigin = func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		if origin == "" {
			return true
		}
		for _, o := range origins {
			if o == "*" || o == origin {
				return true
			}
		}
		return false
	}
}

// SSE registers a Server-Sent Events handler at the given path.
// The response is kept open and events are pushed to the client until it disconnects.
// Group middlewares run before the SSE stream is opened.
func (rg *RouterGroup) SSE(relativePath string, handler SSEHandlerFunc) {
	absolutePath := rg.prefix + relativePath

	middlewares := make([]HandlerFunc, len(rg.middlewares))
	copy(middlewares, rg.middlewares)

	rg.engine.router.GET(absolutePath, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// Run group middlewares (auth, rate limiting, etc.) before opening the stream.
		if len(middlewares) > 0 {
			c := &Context{
				Writer:   w,
				Request:  r,
				Params:   params,
				handlers: middlewares,
				index:    -1,
				engine:   rg.engine,
			}
			c.Next()
			if c.isAborted {
				return
			}
		}

		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "streaming not supported", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("X-Accel-Buffering", "no") // disable nginx buffering

		sc := &SSEContext{
			Writer:  w,
			flusher: flusher,
			Keys:    make(map[string]any),
			Params:  params,
			Request: r,
		}

		handler(sc)
	})
}

// WS registers a WebSocket handler at the given path.
// Group middlewares (auth, rate limiting, etc.) run during the HTTP upgrade phase.
// If any middleware aborts the request, the upgrade is cancelled.
func (rg *RouterGroup) WS(relativePath string, handler WSHandlerFunc) {
	absolutePath := rg.prefix + relativePath

	cfg := rg.engine.WSConfig
	upgrader := websocket.Upgrader{
		ReadBufferSize:   cfg.ReadBufferSize,
		WriteBufferSize:  cfg.WriteBufferSize,
		HandshakeTimeout: cfg.HandshakeTimeout,
		CheckOrigin:      cfg.CheckOrigin,
	}

	// Snapshot group middlewares at registration time.
	middlewares := make([]HandlerFunc, len(rg.middlewares))
	copy(middlewares, rg.middlewares)

	rg.engine.router.GET(absolutePath, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// Run group middlewares in the HTTP phase so auth/rate-limit middleware works.
		if len(middlewares) > 0 {
			c := &Context{
				Writer:   w,
				Request:  r,
				Params:   params,
				handlers: middlewares,
				index:    -1,
				engine:   rg.engine,
			}
			c.Next()
			if c.isAborted {
				return
			}
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf(Red+"WS upgrade failed: %v"+Reset, err)
			return
		}
		defer conn.Close()

		wc := &WSContext{
			Conn:    conn,
			Keys:    make(map[string]any),
			Params:  params,
			Request: r,
		}

		handler(wc)
	})
}

// Sets trusted Proxies
func (e *Engine) SetTrustedProxies(proxies []string) error {
	var trusted []*net.IPNet

	for _, proxy := range proxies {
		if !strings.Contains(proxy, "/") {
			proxy += "/32" // attaches 32 to leave 0 bits, meaning only this ip is trusted
		}

		_, cidr, err := net.ParseCIDR(proxy)
		if err != nil {
			return err
		}

		trusted = append(trusted, cidr)
	}

	e.trustedProxies = trusted
	return nil
}

// helper to check if proxy is trusted
func (e *Engine) isTrustedProxy(ip net.IP) bool {
	for _, trusted := range e.trustedProxies {
		if trusted.Contains(ip) {
			return true
		}
	}

	return false
}