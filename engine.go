package vodka

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
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
	srv        *http.Server
	shutdownCh chan struct{}
	router     *httprouter.Router
	WSConfig   *WSConfig
	*RouterGroup
}

// creates a new router
func NewRouter() *Engine {
	router := httprouter.New()
	router.HandleOPTIONS = false
	engine := &Engine{
		router:   router,
		WSConfig: DefaultWSConfig()}

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
	e.router.ServeHTTP(w, req)
}

func (rg *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	absolutePath := rg.prefix + comp

	handlers := make([]HandlerFunc, 0, len(rg.middlewares)+1)
	handlers = append(handlers, rg.middlewares...)
	handlers = append(handlers, handler)

	rg.engine.router.Handle(method, absolutePath, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		c := &Context{
			Writer:   w,
			Request:  r,
			Params:   params,
			handlers: handlers,
			index:    -1,
		}

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

type RunConfig struct {
	Addr             string        // overrides the addr argument if set
	GracefulShutdown bool          // set to true to wait for active requests before closing
	GracefulTimeout  time.Duration // max time to wait for in-flight requests; defaults to 10s

}

// app.Run(":8080", vodka.RunConfig{GracefulTimeout: 10 * time.Second})
func (e *Engine) Run(addr string, cfgs ...RunConfig) error {
	if addr == "" {
		addr = ":8080"
	}

	// No config or graceful shutdown not requested — plain ListenAndServe
	if len(cfgs) == 0 || !cfgs[0].GracefulShutdown {
		log.Printf(Green+"Pouring Vodka on %s\n"+Reset, addr)
		return http.ListenAndServe(addr, e)
	}

	cfg := cfgs[0]
	if cfg.Addr != "" {
		addr = cfg.Addr
	}

	e.srv = &http.Server{
		Addr:    addr,
		Handler: e,
	}

	e.shutdownCh = make(chan struct{})

	// Register for OS shutdown signals (Ctrl+C and Docker/k8s stop)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	serverErr := make(chan error, 1)
	go func() {
		log.Printf(Green+"Pouring Vodka on %s\n"+Reset, addr)
		if err := e.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErr <- err
		}
	}()

	// Block until OS signal, context cancellation, or server error
	select {
	case err := <-serverErr:
		return err
	case <-e.shutdownCh:
		log.Printf(Yellow + "[Vodka] Shutdown triggered — draining requests...\n" + Reset)
	case sig := <-quit:
		log.Printf(Yellow+"[Vodka] Signal received: %v — draining requests...\n"+Reset, sig)
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.GracefulTimeout)
	defer cancel()

	if err := e.srv.Shutdown(shutdownCtx); err != nil {
		log.Printf(Red+"[Vodka] Forced shutdown after timeout: %v\n"+Reset, err)
		return err
	}

	log.Printf(Green + "[Vodka] Server shut down cleanly.\n" + Reset)
	return nil
}

// Shutdown gracefully stops the server, waiting for active requests to finish.
// The context controls how long to wait before force-closing.
// Useful in tests to trigger shutdown without sending OS signals.
func (e *Engine) Shutdown(ctx context.Context) error {
	if e.shutdownCh != nil {
		close(e.shutdownCh)
	}
	if e.srv == nil {
		return nil
	}
	return e.srv.Shutdown(ctx)
}
