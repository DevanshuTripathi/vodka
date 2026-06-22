package vodka

import (
	"context"
	"errors"
	"net/http"
	"os"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestStartupHooks(t *testing.T) {
	// 1. Single startup hook
	t.Run("SingleHook", func(t *testing.T) {
		app := NewRouter()
		called := false
		app.OnStart(func() error {
			called = true
			return nil
		})

		err := app.lifecycle.runStartupHooks()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !called {
			t.Fatal("expected startup hook to be called")
		}
	})

	// 2. Multiple startup hooks (order of execution)
	t.Run("MultipleHooksOrder", func(t *testing.T) {
		app := NewRouter()
		var order []string
		var mu sync.Mutex

		app.OnStart(func() error {
			mu.Lock()
			order = append(order, "first")
			mu.Unlock()
			return nil
		})
		app.OnStart(func() error {
			mu.Lock()
			order = append(order, "second")
			mu.Unlock()
			return nil
		})

		err := app.lifecycle.runStartupHooks()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expected := []string{"first", "second"}
		if !reflect.DeepEqual(order, expected) {
			t.Fatalf("expected order %v, got %v", expected, order)
		}
	})

	// 3. Startup hook failure
	t.Run("StartupFailure", func(t *testing.T) {
		app := NewRouter()
		app.OnStart(func() error {
			return errors.New("startup failed")
		})

		// Verify Run returns the startup error and aborts
		err := app.Run(":invalid_port_or_addr")
		if err == nil || err.Error() != "startup failed" {
			t.Fatalf("expected 'startup failed' error, got: %v", err)
		}
	})
}

func TestShutdownHooksPriority(t *testing.T) {
	app := NewRouter()
	var order []string
	var mu sync.Mutex

	app.OnShutdownWithPriority(50, func(ctx context.Context) error {
		mu.Lock()
		order = append(order, "priority 50 (first registration)")
		mu.Unlock()
		return nil
	})

	app.OnShutdownWithPriority(100, func(ctx context.Context) error {
		mu.Lock()
		order = append(order, "priority 100")
		mu.Unlock()
		return nil
	})

	app.OnShutdownWithPriority(50, func(ctx context.Context) error {
		mu.Lock()
		order = append(order, "priority 50 (second registration)")
		mu.Unlock()
		return nil
	})

	app.OnShutdown(func(ctx context.Context) error {
		mu.Lock()
		order = append(order, "default priority (0)")
		mu.Unlock()
		return nil
	})

	app.OnShutdownWithPriority(-10, func(ctx context.Context) error {
		mu.Lock()
		order = append(order, "priority -10")
		mu.Unlock()
		return nil
	})

	// Create a dummy http server
	srv := &http.Server{}
	err := app.shutdown(srv)
	if err != nil {
		t.Fatalf("unexpected error during shutdown: %v", err)
	}

	expected := []string{
		"priority 100",
		"priority 50 (first registration)",
		"priority 50 (second registration)",
		"default priority (0)",
		"priority -10",
	}

	if !reflect.DeepEqual(order, expected) {
		t.Fatalf("expected shutdown execution order %v, got %v", expected, order)
	}
}

func TestShutdownTimeout(t *testing.T) {
	app := NewRouter()
	app.SetShutdownTimeout(50 * time.Millisecond)

	app.OnShutdown(func(ctx context.Context) error {
		select {
		case <-time.After(150 * time.Millisecond):
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	})

	// The second hook should not run because context gets cancelled
	secondCalled := false
	app.OnShutdown(func(ctx context.Context) error {
		secondCalled = true
		return nil
	})

	srv := &http.Server{}
	err := app.shutdown(srv)
	if err == nil {
		t.Fatal("expected error due to timeout, got nil")
	}

	// The error should mention shutdown timeout exceeded/context deadline exceeded
	if !strings.Contains(err.Error(), "context deadline exceeded") && !strings.Contains(err.Error(), "shutdown timeout exceeded") {
		t.Fatalf("expected timeout error message, got: %v", err)
	}

	if secondCalled {
		t.Fatal("expected second hook not to run because context was cancelled")
	}
}

func TestErrorAggregation(t *testing.T) {
	app := NewRouter()

	app.OnShutdown(func(ctx context.Context) error {
		return errors.New("Database close failed")
	})

	app.OnShutdown(func(ctx context.Context) error {
		return errors.New("Worker stop failed")
	})

	srv := &http.Server{}
	err := app.shutdown(srv)
	if err == nil {
		t.Fatal("expected aggregated errors, got nil")
	}

	expectedMsg := "2 shutdown errors:\n- Database close failed\n- Worker stop failed"
	if err.Error() != expectedMsg {
		t.Fatalf("expected formatted error:\n%q\ngot:\n%q", expectedMsg, err.Error())
	}

	// Verify Go 1.20+ Unwrap compatibility
	unwrapper, ok := err.(interface{ Unwrap() []error })
	if !ok {
		t.Fatal("expected error to implement Unwrap() []error")
	}

	errs := unwrapper.Unwrap()
	if len(errs) != 2 {
		t.Fatalf("expected 2 errors, got %d: %v", len(errs), errs)
	}

	if errs[0].Error() != "Database close failed" || errs[1].Error() != "Worker stop failed" {
		t.Fatalf("unexpected unwrapped errors: %v", errs)
	}
}

func TestSignalHandling(t *testing.T) {
	app := NewRouter()

	shutdownCalled := false
	app.OnShutdown(func(ctx context.Context) error {
		shutdownCalled = true
		return nil
	})

	// Run the server in a goroutine
	errChan := make(chan error, 1)
	go func() {
		errChan <- app.Run(":0")
	}()

	// Wait a moment for server to start listening
	time.Sleep(100 * time.Millisecond)

	// Send SIGINT to ourselves
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatalf("failed to find current process: %v", err)
	}

	err = p.Signal(os.Interrupt)
	if err != nil {
		t.Fatalf("failed to send interrupt signal: %v", err)
	}

	// Wait for Run to return
	select {
	case runErr := <-errChan:
		if runErr != nil {
			t.Fatalf("Run returned unexpected error: %v", runErr)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("timeout waiting for server to shutdown gracefully")
	}

	if !shutdownCalled {
		t.Fatal("expected shutdown hooks to be called on signal")
	}
}

func TestShutdownIdempotent(t *testing.T) {
	app := NewRouter()
	callCount := 0

	app.OnShutdown(func(ctx context.Context) error {
		callCount++
		return nil
	})

	srv := &http.Server{}

	if err := app.shutdown(srv); err != nil {
		t.Fatalf("unexpected first shutdown error: %v", err)
	}
	if callCount != 1 {
		t.Fatalf("expected shutdown hook to run once, got %d", callCount)
	}

	if err := app.shutdown(srv); err != nil {
		t.Fatalf("unexpected second shutdown error: %v", err)
	}
	if callCount != 1 {
		t.Fatalf("expected shutdown hook not to run again, got %d", callCount)
	}
}

func TestConcurrentHookRegistration(t *testing.T) {
	app := NewRouter()
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			app.OnStart(func() error { return nil })
			app.OnShutdown(func(ctx context.Context) error { return nil })
		}()
	}
	wg.Wait()

	if err := app.lifecycle.runStartupHooks(); err != nil {
		t.Fatalf("unexpected startup hooks error: %v", err)
	}

	srv := &http.Server{}
	if err := app.shutdown(srv); err != nil {
		t.Fatalf("unexpected shutdown error after concurrent registration: %v", err)
	}
}
