package mixers

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DevanshuTripathi/vodka"
)

func TestNewRateLimiter(t *testing.T) {
	vrl := NewRateLimiter(10, 5)
	if vrl == nil {
		t.Fatal("expected non-nil VodkaRateLimiter")
	}
	if vrl.done == nil {
		t.Fatal("expected done channel to be initialized")
	}
	vrl.Stop()
}

func TestRateLimiterStop(t *testing.T) {
	vrl := NewRateLimiter(10, 5)

	// done channel must be open before Stop
	select {
	case <-vrl.done:
		t.Fatal("done channel should not be closed before Stop()")
	default:
	}

	vrl.Stop()

	// done channel must be closed after Stop, causing cleanup goroutine to exit
	select {
	case <-vrl.done:
		// expected: channel closed
	case <-time.After(100 * time.Millisecond):
		t.Fatal("done channel not closed after Stop()")
	}
}

func TestRateLimiterAllowsUnderLimit(t *testing.T) {
	app := vodka.NewRouter()
	vrl := NewRateLimiter(10, 3)
	defer vrl.Stop()

	app.Use(RateLimiter(vrl))
	app.GET("/test", func(c *vodka.Context) {
		c.String(200, "ok")
	})

	for i := 0; i < 3; i++ {
		req := httptest.NewRequest("GET", "/test", nil)
		req.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		app.ServeHTTP(w, req)
		if w.Code != 200 {
			t.Errorf("request %d: expected 200, got %d", i+1, w.Code)
		}
	}
}

func TestRateLimiterBlocksOverLimit(t *testing.T) {
	app := vodka.NewRouter()
	// rate=0 means no token refill, so burst is strict
	vrl := NewRateLimiter(0, 2)
	defer vrl.Stop()

	app.Use(RateLimiter(vrl))
	app.GET("/test", func(c *vodka.Context) {
		c.String(200, "ok")
	})

	for i := 0; i < 2; i++ {
		req := httptest.NewRequest("GET", "/test", nil)
		req.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		app.ServeHTTP(w, req)
		if w.Code != 200 {
			t.Errorf("request %d: expected 200 within burst, got %d", i+1, w.Code)
		}
	}

	req := httptest.NewRequest("GET", "/test", nil)
	req.RemoteAddr = "127.0.0.1:1234"
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if w.Code != 429 {
		t.Errorf("expected 429 after burst exceeded, got %d", w.Code)
	}
}

func TestRateLimiterStopIsIdempotent(t *testing.T) {
	vrl := NewRateLimiter(10, 5)
	// Calling Stop multiple times must not panic
	vrl.Stop()
	vrl.Stop()
}
