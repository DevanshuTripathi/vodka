package mixers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/DevanshuTripathi/vodka"
)

func TestTimeoutMiddleware_Success(t *testing.T) {

	app := vodka.DefaultRouter()
	app.Use(Timeout(5 * time.Second))

	app.GET("/hello", func(c *vodka.Context) {
		c.JSON(200, vodka.M{"message": "success"})
	})

	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	expected := `{"message":"success"}` + "\n"
	if rec.Body.String() != expected {
		t.Fatalf("expected body %s, got %s", expected, rec.Body.String())
	}
}

func TestTimeoutMiddleware_Timeout(t *testing.T) {

	app := vodka.DefaultRouter()
	app.Use(Timeout(1 * time.Second))

	app.GET("/slow", func(c *vodka.Context) {
		time.Sleep(2 * time.Second)
		c.JSON(200, vodka.M{"message": "too late"})
	})

	req := httptest.NewRequest(http.MethodGet, "/slow", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	if rec.Code != http.StatusRequestTimeout {
		t.Fatalf("expected status 408, got %d", rec.Code)
	}

	expected := `{"error":"request timeout"}` + "\n"
	if rec.Body.String() != expected {
		t.Fatalf("expected body %s, got %s", expected, rec.Body.String())
	}
}

func TestTimeoutMiddleware_ContextCancelled(t *testing.T) {

	app := vodka.DefaultRouter()
	app.Use(Timeout(100 * time.Millisecond))

	ctxWasCancelled := false

	app.GET("/ctx", func(c *vodka.Context) {
		select {
		case <-c.Request.Context().Done():
			ctxWasCancelled = true
		case <-time.After(2 * time.Second):
		}
	})

	req := httptest.NewRequest(http.MethodGet, "/ctx", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	if !ctxWasCancelled {
		t.Fatal("expected request context to be cancelled on timeout, but it was not")
	}
}

func TestTimeoutMiddleware_LateWriteDiscarded(t *testing.T) {

	app := vodka.DefaultRouter()
	app.Use(Timeout(100 * time.Millisecond))

	app.GET("/late", func(c *vodka.Context) {
		<-c.Request.Context().Done()

		time.Sleep(50 * time.Millisecond)
		c.JSON(200, vodka.M{"message": "should not appear"})
	})

	req := httptest.NewRequest(http.MethodGet, "/late", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	if rec.Code != http.StatusRequestTimeout {
		t.Fatalf("expected 408, got %d", rec.Code)
	}

	body := rec.Body.String()
	if strings.Contains(body, "should not appear") {
		t.Fatalf("late write leaked into response body: %q", body)
	}
}

func TestTimeoutMiddleware_ContentTypeIsJSON(t *testing.T) {

	app := vodka.DefaultRouter()
	app.Use(Timeout(100 * time.Millisecond))

	app.GET("/ct", func(c *vodka.Context) {
		time.Sleep(2 * time.Second)
	})

	req := httptest.NewRequest(http.MethodGet, "/ct", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	ct := rec.Header().Get("Content-Type")
	if ct != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %q", ct)
	}
}

func TestTimeoutMiddleware_HandlerPanicsAfterTimeout(t *testing.T) {

	app := vodka.DefaultRouter()
	app.Use(Timeout(100 * time.Millisecond))

	app.GET("/panic", func(c *vodka.Context) {
		<-c.Request.Context().Done()
		panic("something went wrong after timeout")
	})

	req := httptest.NewRequest(http.MethodGet, "/panic", nil)
	rec := httptest.NewRecorder()

	app.ServeHTTP(rec, req)

	if rec.Code != http.StatusRequestTimeout {
		t.Fatalf("expected 408 even after handler panic, got %d", rec.Code)
	}
}

func TestTimeoutMiddleware_MultipleRequests(t *testing.T) {

	app := vodka.DefaultRouter()
	app.Use(Timeout(500 * time.Millisecond))

	app.GET("/multi", func(c *vodka.Context) {
		c.JSON(200, vodka.M{"ok": true})
	})

	for i := 0; i < 10; i++ {
		req := httptest.NewRequest(http.MethodGet, "/multi", nil)
		rec := httptest.NewRecorder()
		app.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("request %d: expected 200, got %d", i+1, rec.Code)
		}
	}
}

func TestTimeoutMiddleware_ConcurrentRequests(t *testing.T) {

	app := vodka.DefaultRouter()
	app.Use(Timeout(200 * time.Millisecond))

	var counter atomic.Int32
	app.GET("/concurrent", func(c *vodka.Context) {
		n := counter.Add(1)
		if n%2 == 0 {
			time.Sleep(1 * time.Second) 
		} else {
			c.JSON(200, vodka.M{"ok": true}) 
		}
	})

	const total = 10
	results := make(chan int, total)

	for i := 0; i < total; i++ {
		go func() {
			req := httptest.NewRequest(http.MethodGet, "/concurrent", nil)
			rec := httptest.NewRecorder()
			app.ServeHTTP(rec, req)
			results <- rec.Code
		}()
	}

	var got200, got408 int
	for i := 0; i < total; i++ {
		code := <-results
		switch code {
		case http.StatusOK:
			got200++
		case http.StatusRequestTimeout:
			got408++
		default:
			t.Errorf("unexpected status code: %d", code)
		}
	}

	if got200 == 0 || got408 == 0 {
		t.Errorf("expected a mix of 200 and 408, got %d×200 and %d×408", got200, got408)
	}
}

func TestTimeoutMiddleware_NextMiddlewareRuns(t *testing.T) {

	app := vodka.DefaultRouter()

	middlewareRan := false
	app.Use(Timeout(500 * time.Millisecond))
	app.Use(func(c *vodka.Context) {
		middlewareRan = true
		c.Next()
	})

	app.GET("/chain", func(c *vodka.Context) {
		c.JSON(200, vodka.M{"ok": true})
	})

	req := httptest.NewRequest(http.MethodGet, "/chain", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	if !middlewareRan {
		t.Fatal("middleware registered after Timeout did not run on a fast request")
	}
}

func TestTimeoutMiddleware_AlreadyCancelledContext(t *testing.T) {

	app := vodka.DefaultRouter()
	app.Use(Timeout(5 * time.Second))

	app.GET("/predead", func(c *vodka.Context) {
		<-c.Request.Context().Done()
	})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	req := httptest.NewRequest(http.MethodGet, "/predead", nil)
	req = req.WithContext(ctx)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	if rec.Code == http.StatusInternalServerError {
		t.Fatalf("unexpected 500 when incoming context was already cancelled")
	}
}