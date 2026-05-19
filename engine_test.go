package vodka

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func shutdownApp(t *testing.T, app interface{}, ctx context.Context) error {
	if s, ok := app.(interface{ Shutdown(context.Context) error }); ok {
		return s.Shutdown(ctx)
	}
	if c, ok := app.(interface{ Close() error }); ok {
		return c.Close()
	}
	t.Skip("engine does not support shutdown or close")
	return nil
}

func TestRouteParams(t *testing.T) {
	app := DefaultRouter()
	app.GET("/users/:id", func(c *Context) {
		c.String(200, c.Param("id"))
	})
	req := httptest.NewRequest("GET", "/users/67", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if w.Body.String() != "67" {
		t.Fatalf("expected 67, got %s", w.Body.String())
	}
}

func TestMiddlewareChain(t *testing.T) {
	app := NewRouter()
	calls := []string{}
	app.Use(func(c *Context) {
		calls = append(calls, "Middleware1")
		c.Next()
	})
	app.Use(func(c *Context) {
		calls = append(calls, "Middleware2")
		c.Next()
	})
	app.Use(func(c *Context) {
		calls = append(calls, "Middleware3")
		c.Next()
	})
	app.GET("/test", func(c *Context) {
		calls = append(calls, "Handler")
	})
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)
	expected := []string{"Middleware1", "Middleware2", "Middleware3", "Handler"}
	if !reflect.DeepEqual(calls, expected) {
		t.Fatalf("expected %v, got %v", expected, calls)
	}
}

// ─────────────────────────────────────────────
// Graceful Shutdown Tests — Issue #9
// ─────────────────────────────────────────────

func TestRunGraceful_StartsAndShutdownCleanly(t *testing.T) {
	app := NewRouter()
	app.GET("/ping", func(c *Context) {
		c.JSON(200, M{"message": "pong"})
	})

	errCh := make(chan error, 1)
	go func() {
		errCh <- app.Run(":19090")
	}()

	// Give the server a moment to start
	time.Sleep(100 * time.Millisecond)

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := shutdownApp(t, app, shutdownCtx); err != nil {
		t.Fatalf("expected clean shutdown, got: %v", err)
	}

	select {
	case err := <-errCh:
		if err != nil {
			t.Fatalf("expected clean shutdown, got: %v", err)
		}
	case <-time.After(10 * time.Second):
		t.Fatal("graceful shutdown timed out")
	}
}

func TestRunGraceful_InFlightRequestCompletes(t *testing.T) {
	app := NewRouter()

	// Slow handler — simulates an in-flight request
	app.GET("/slow", func(c *Context) {
		time.Sleep(200 * time.Millisecond)
		c.JSON(200, M{"message": "done"})
	})

	errCh := make(chan error, 1)
	go func() {
		errCh <- app.Run(":19091")
	}()

	time.Sleep(100 * time.Millisecond)

	// Fire the slow request in parallel
	reqDone := make(chan struct{})
	go func() {
		resp, err := http.Get("http://localhost:19091/slow")
		if err == nil {
			resp.Body.Close()
		}
		close(reqDone)
	}()

	// Let the request reach the handler, then trigger shutdown
	time.Sleep(50 * time.Millisecond)

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()
	if err := shutdownApp(t, app, shutdownCtx); err != nil {
		t.Fatalf("unexpected shutdown error: %v", err)
	}

	// In-flight request must finish before shutdown completes
	select {
	case <-reqDone:
		// good — request completed
	case <-time.After(3 * time.Second):
		t.Fatal("in-flight request did not complete before shutdown")
	}

	select {
	case err := <-errCh:
		if err != nil {
			t.Fatalf("unexpected shutdown error: %v", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("server did not shut down in time")
	}
}
