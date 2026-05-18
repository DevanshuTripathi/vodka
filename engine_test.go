package vodka

import (
	"context"
	"net/http"
	"testing"
	"time"
)

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
		errCh <- app.RunGraceful(":19090", 5*time.Second)
	}()

	// Give the server a moment to start
	time.Sleep(100 * time.Millisecond)

	// Shutdown programmatically — no OS signal needed
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.Shutdown(ctx); err != nil {
		t.Fatalf("shutdown error: %v", err)
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
		errCh <- app.RunGraceful(":19091", 2*time.Second)
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

	// Let the request reach the handler, then shutdown programmatically
	time.Sleep(50 * time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_ = app.Shutdown(ctx)

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
