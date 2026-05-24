package mixers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/DevanshuTripathi/vodka"
)

type jsonResponseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *jsonResponseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

type jsonLogEntry struct {
	Timestamp string `json:"timestamp"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	Status    int    `json:"status"`
	Latency   string `json:"latency"`
	IP        string `json:"ip"`
}


func JSONLogger() vodka.HandlerFunc {
	return func(c *vodka.Context) {

		start := time.Now()

		rw := &jsonResponseWriter{
			ResponseWriter: c.Writer,
			status:         http.StatusOK,
		}
		c.Writer = rw

		c.Next()

		entry := jsonLogEntry{
			Timestamp: time.Now().Format(time.RFC3339),
			Method:    c.Request.Method,
			Path:      c.Request.URL.Path,
			Status:    rw.status,
			Latency:   time.Since(start).String(),
			IP:        c.ClientIP(),
		}

		logJSON, err := json.Marshal(entry)
		if err != nil {
			log.Printf("json logger marshal error: %v", err)
			return
		}

		log.Println(string(logJSON))
	}
}