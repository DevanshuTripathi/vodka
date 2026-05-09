package vodka

import (
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	engine      *Engine
}

// httprouter wrapper
type Engine struct {
	router *httprouter.Router
	*RouterGroup
}

// creates a new router
func NewRouter() *Engine {
	router := httprouter.New()
	router.HandleOPTIONS = false
	engine := &Engine{
		router: router,
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
	engine.Use(Recovery(), Logger(), ErrorHandler())
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
func (e *Engine) Run(addr string) error {
	if addr == "" {
		addr = ":8080"
	}

	log.Printf(Green+"Pouring Vodka on %s\n"+Reset, addr)

	// Using net/http
	return http.ListenAndServe(addr, e)
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
	// Create a fresh Context for the incoming request
	c := &Context{
		Writer:  w,
		Request: req,
	}

	// 1. Run all Global Middlewares FIRST (like your AllowCORS)
	// This assumes your app.Use() appends to a slice called e.middlewares
	for _, middleware := range e.middlewares {
		middleware(c)

		// If a middleware aborts the request (e.g., CORS returns 204 for OPTIONS),
		// we stop execution immediately. The router is never called.
		if c.isAborted {
			return
		}
	}

	// 2. If the middleware didn't stop the request, pass it to the router
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
