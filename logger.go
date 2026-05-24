package vodka

import (
	"encoding/json"
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

func wrapResponseWriter(c *Context) *responseWriter {
	rw := &responseWriter{ResponseWriter: c.Writer, status: http.StatusOK}
	c.Writer = rw
	return rw
}

func Logger() HandlerFunc {
	return func(c *Context) {
		start := time.Now()
		rw := wrapResponseWriter(c)
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

type jsonLogEntry struct {
	Method    string `json:"method"`
	Path      string `json:"path"`
	Status    int    `json:"status"`
	Latency   string `json:"latency"`
	IP        string `json:"ip"`
	Timestamp string `json:"timestamp"`
	RequestID string `json:"request_id,omitempty"`
}

func JSONLogger() HandlerFunc {
	return func(c *Context) {
		start := time.Now()
		rw := wrapResponseWriter(c)
		c.Next()

		entry := jsonLogEntry{
			Method:    c.Request.Method,
			Path:      c.Request.URL.Path,
			Status:    rw.status,
			Latency:   time.Since(start).String(),
			IP:        c.ClientIP(),
			Timestamp: time.Now().UTC().Format(time.RFC3339),
			RequestID: c.Request.Header.Get("X-Request-ID"),
		}

		payload, err := json.Marshal(entry)
		if err != nil {
			log.Printf("vodka: JSON logger encode error: %v", err)
			return
		}

		log.Println(string(payload))
	}
}
