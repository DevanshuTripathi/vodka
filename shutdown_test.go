package vodka

import (
	"net/http"
	"os"
	"sync"
	"syscall"
	"testing"
	"time"
)

func TestRunGracefully_CleanShutdownOnSIGTERM(t *testing.T) {
	app := NewRouter()
	app.GET("/ping", func(c *Context) {
		c.String(http.StatusOK, "pong")
	})

	quit := make(chan os.Signal, 1)
	errCh := make(chan error, 1)
	go func() {
		errCh <- app.runGracefully(":19001", ShutdownConfig{Timeout: 5 * time.Second}, quit)
	}()

	waitForServer(t, "http://localhost:19001/ping")
	quit <- syscall.SIGTERM

	select {
	case err := <-errCh:
		if err != nil {
			t.Fatalf("expected clean shutdown, got: %v", err)
		}
	case <-time.After(6 * time.Second):
		t.Fatal("server did not shut down within the expected time")
	}
}

func TestRunGracefully_InFlightRequestsComplete(t *testing.T) {
	app := NewRouter()
	app.GET("/slow", func(c *Context) {
		time.Sleep(400 * time.Millisecond)
		c.String(http.StatusOK, "done")
	})

	quit := make(chan os.Signal, 1)
	errCh := make(chan error, 1)
	go func() {
		errCh <- app.runGracefully(":19002", ShutdownConfig{Timeout: 5 * time.Second}, quit)
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

	time.Sleep(80 * time.Millisecond)
	quit <- syscall.SIGTERM

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

func TestRunGracefully_TimeoutEnforced(t *testing.T) {
	app := NewRouter()
	app.GET("/hang", func(c *Context) {
		time.Sleep(10 * time.Second)
		c.String(http.StatusOK, "too late")
	})

	quit := make(chan os.Signal, 1)
	errCh := make(chan error, 1)
	go func() {
		errCh <- app.runGracefully(":19003", ShutdownConfig{Timeout: 200 * time.Millisecond}, quit)
	}()

	waitForServer(t, "http://localhost:19003/hang")

	go http.Get("http://localhost:19003/hang") //nolint:errcheck

	time.Sleep(80 * time.Millisecond)
	quit <- syscall.SIGTERM

	select {
	case err := <-errCh:
		if err == nil {
			t.Fatal("expected a timeout error, got nil")
		}
	case <-time.After(3 * time.Second):
		t.Fatal("server did not respect the shutdown timeout deadline")
	}
}

func TestRunGracefully_DefaultTimeout(t *testing.T) {
	app := NewRouter()
	app.GET("/ping", func(c *Context) {
		c.String(http.StatusOK, "pong")
	})

	quit := make(chan os.Signal, 1)
	errCh := make(chan error, 1)
	go func() {
		errCh <- app.runGracefully(":19004", ShutdownConfig{}, quit)
	}()

	waitForServer(t, "http://localhost:19004/ping")
	quit <- syscall.SIGTERM

	select {
	case err := <-errCh:
		if err != nil {
			t.Fatalf("unexpected error with default timeout: %v", err)
		}
	case <-time.After(12 * time.Second):
		t.Fatal("server did not exit with default timeout applied")
	}
}

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
