package mixers

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/DevanshuTripathi/vodka"
)

// timeoutResponseWriter wraps ResponseWriter to block writes after timeout.
type timeoutResponseWriter struct {
	http.ResponseWriter
	mu         sync.Mutex
	timedOut   bool
	deadHeader http.Header // dummy header map to return after timeout
}

// Header returns a dummy map after timeout so late header writes go nowhere.
func (tw *timeoutResponseWriter) Header() http.Header {
	tw.mu.Lock()
	defer tw.mu.Unlock()
	if tw.timedOut {
		return tw.deadHeader
	}
	return tw.ResponseWriter.Header()
}

// WriteHeader discards the status code if timeout already claimed the response.
func (tw *timeoutResponseWriter) WriteHeader(code int) {
	tw.mu.Lock()
	defer tw.mu.Unlock()
	if tw.timedOut {
		return
	}
	tw.ResponseWriter.WriteHeader(code)
}

// Write discards the body if timeout already claimed the response.
func (tw *timeoutResponseWriter) Write(b []byte) (int, error) {
	tw.mu.Lock()
	defer tw.mu.Unlock()
	if tw.timedOut {
		return 0, nil
	}
	return tw.ResponseWriter.Write(b)
}

// markTimedOut atomically claims the response for the timeout path.
// Returns true if it won the race, false if the handler already wrote.
func (tw *timeoutResponseWriter) markTimedOut() bool {
	tw.mu.Lock()
	defer tw.mu.Unlock()
	if tw.timedOut {
		return false
	}
	tw.timedOut = true
	return true
}

// Timeout returns a middleware that cancels the request context and responds
// with HTTP 408 if the handler chain does not finish within d.
// Usage: app.Use(mixers.Timeout(5 * time.Second))
func Timeout(d time.Duration) vodka.HandlerFunc {
	return func(c *vodka.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), d)
		defer cancel()

		// replace the request's context with our timeout context
		c.Request = c.Request.WithContext(ctx)

		tw := &timeoutResponseWriter{
			ResponseWriter: c.Writer,
			deadHeader:     make(http.Header),
		}
		c.Writer = tw

		// closed by AfterFunc when it exits, win or lose.
		afuncDone := make(chan struct{})

		// true when AfterFunc won the race and wrote 408.
		claimed := false

		timer := time.AfterFunc(d, func() {
			defer close(afuncDone)

			if tw.markTimedOut() {
				claimed = true

				rw := tw.ResponseWriter
				rw.Header().Set("Content-Type", "application/json")
				rw.WriteHeader(http.StatusRequestTimeout)
				rw.Write([]byte("{\"error\":\"request timeout\"}\n")) //nolint:errcheck
			}
		})

		// defer guarantees this runs even if the handler panics, preventing a
		// race where AfterFunc writes concurrently with Logger reading rw.status.
		defer func() {
			if !timer.Stop() {
				<-afuncDone
				if claimed {
					c.Abort()
				}
			}
		}()

		c.Next()
	}
}