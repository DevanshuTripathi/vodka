package vodka

import (
	"context"
	"net/http"
	"sync"
	"testing"
	"time"
)

// TestRun_GracefulShutdown_CleanExit verifies the server exits without error
// when Shutdown is called with no requests in-flight.
func TestRun_GracefulShutdown_CleanExit(t *testing.T) {
	app := NewRouter()
	app.GET("/ping", func(c *Context) {
		c.String(http.StatusOK, "pong")
	})

	errCh := make(chan error, 1)
	go func() {
		errCh <- app.Run(":19001", RunConfig{
			GracefulShutdown: true,
			GracefulTimeout:  5 * time.Second,
		})
	}()

	waitForServer(t, "http://localhost:19001/ping")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	app.Shutdown(ctx)

	select {
	case err := <-errCh:
		if err != nil {
			t.Fatalf("expected clean shutdown, got: %v", err)
		}
	case <-time.After(6 * time.Second):
		t.Fatal("server did not shut down in time")
	}
}

// TestRun_GracefulShutdown_InFlightRequestsComplete verifies that a request
// already being processed is allowed to finish before the server closes.
func TestRun_GracefulShutdown_InFlightRequestsComplete(t *testing.T) {
	app := NewRouter()
	app.GET("/slow", func(c *Context) {
		time.Sleep(400 * time.Millisecond)
		c.String(http.StatusOK, "done")
	})

	errCh := make(chan error, 1)
	go func() {
		errCh <- app.Run(":19002", RunConfig{
			GracefulShutdown: true,
			GracefulTimeout:  5 * time.Second,
		})
	}()

	waitForServer(t, "http://localhost:19002/slow")

	requestDone := make(chan int, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		resp, err := http.Get("http://localhost:19002/slow")
		if err != nil {
			t.Errorf("in-flight request failed: %v", err)
			return
		}
		resp.Body.Close()
		requestDone <- resp.StatusCode
	}()

	// Let the request reach the handler, then trigger shutdown
	time.Sleep(80 * time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	app.Shutdown(ctx)

	select {
	case code := <-requestDone:
		if code != http.StatusOK {
			t.Errorf("expected 200, got %d", code)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("in-flight request did not complete before server exited")
	}

	wg.Wait()

	select {
	case err := <-errCh:
		if err != nil {
			t.Fatalf("expected clean shutdown, got: %v", err)
		}
	case <-time.After(6 * time.Second):
		t.Fatal("server did not exit after draining requests")
	}
}

// TestRun_GracefulShutdown_TimeoutEnforced verifies the server force-closes
// and returns an error when requests exceed the shutdown timeout.
func TestRun_GracefulShutdown_TimeoutEnforced(t *testing.T) {
	app := NewRouter()
	app.GET("/hang", func(c *Context) {
		time.Sleep(10 * time.Second)
		c.String(http.StatusOK, "too late")
	})

	errCh := make(chan error, 1)
	go func() {
		errCh <- app.Run(":19003", RunConfig{
			GracefulShutdown: true,
			GracefulTimeout:  200 * time.Millisecond,
		})
	}()

	waitForServer(t, "http://localhost:19003/hang")

	go http.Get("http://localhost:19003/hang") //nolint:errcheck

	time.Sleep(80 * time.Millisecond)

	// Very short context — server should force-close before handler finishes
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	app.Shutdown(ctx)

	select {
	case err := <-errCh:
		if err == nil {
			t.Fatal("expected timeout error, got nil")
		}
	case <-time.After(3 * time.Second):
		t.Fatal("server did not respect the shutdown timeout")
	}
}

// TestRun_PlainBehaviorWithoutConfig verifies that calling Run without a
// RunConfig still works exactly as before — no graceful shutdown logic runs.
func TestRun_PlainBehaviorWithoutConfig(t *testing.T) {
	app := NewRouter()
	app.GET("/ping", func(c *Context) {
		c.String(http.StatusOK, "pong")
	})

	go app.Run(":19004") //nolint:errcheck

	waitForServer(t, "http://localhost:19004/ping")

	resp, err := http.Get("http://localhost:19004/ping")
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
}

// waitForServer polls until the server is ready or 3 seconds pass.
// Prevents tests from triggering shutdown before the server is listening.
func waitForServer(t *testing.T, url string) {
	t.Helper()
	deadline := time.Now().Add(3 * time.Second)
	for time.Now().Before(deadline) {
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
			return
		}
		time.Sleep(20 * time.Millisecond)
	}
	t.Fatalf("server at %s did not become ready in time", url)
}
