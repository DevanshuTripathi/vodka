package vodka

import (
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
    http.ResponseWriter
    status      int
    wroteHeader bool
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.status = code
    rw.wroteHeader = true
    rw.ResponseWriter.WriteHeader(code)
}

func Logger() HandlerFunc {
	return func(c *Context) {
		start := time.Now()
		rw := &responseWriter{ResponseWriter: c.Writer, status: http.StatusOK}
		c.Writer = rw
		c.Next()
		log.Printf(
			Blue+"%s %s %d %s"+Reset,
			c.Request.Method,
			c.Request.URL.Path,
			rw.status,
			time.Since(start),
		)
	}
}