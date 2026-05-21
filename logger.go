package vodka

import (
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	if rw.status == 0 {
		rw.status = http.StatusOK
	}
	return rw.ResponseWriter.Write(b)
}

func Logger() HandlerFunc {
	return func(c *Context) {
		start := time.Now()
		rw := &responseWriter{ResponseWriter: c.Writer}
		c.Writer = rw
		c.Next()
		if rw.status == 0 {
			rw.status = http.StatusOK
		}
		log.Printf(
			Blue+"%s %s %d %s"+Reset,
			c.Request.Method,
			c.Request.URL.Path,
			rw.status,
			time.Since(start),
		)
	}
}