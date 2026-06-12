package vodka

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCORSPreflightBug(t *testing.T) {
	app := DefaultRouter()
	app.Use(AllowCORS([]string{"*"}))

	app.POST("/test", func(c *Context) {
		c.String(200, "success")
	})

	req := httptest.NewRequest(http.MethodOptions, "/test", nil)
	req.Header.Set("Origin", "http://localhost:3000")
	req.Header.Set("Access-Control-Request-Method", "POST")

	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if w.Code == http.StatusMethodNotAllowed {
		t.Fatalf("expected CORS middleware to handle preflight, got %d", w.Code)
	}
}

func TestCORSAllowsMaxAge(t *testing.T) {
	app := DefaultRouter()
	app.Use(AllowCORS([]string{"*"}, CORSConfig{MaxAge: 86400}))

	app.POST("/test", func(c *Context) {
		c.String(200, "success")
	})

	req := httptest.NewRequest(http.MethodOptions, "/test", nil)
	req.Header.Set("Origin", "http://localhost:3000")
	req.Header.Set("Access-Control-Request-Method", "POST")

	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatalf("expected 204, got %d", w.Code)
	}

	maxAge := w.Header().Get("Access-Control-Max-Age")
	if maxAge != "86400" {
		t.Fatalf("expected Access-Control-Max-Age to be 86400, got %q", maxAge)
	}
}

func TestCORSMaxAgeZeroOmitsHeader(t *testing.T) {
	app := DefaultRouter()
	app.Use(AllowCORS([]string{"*"}, CORSConfig{MaxAge: 0}))

	app.POST("/test", func(c *Context) {
		c.String(200, "success")
	})

	req := httptest.NewRequest(http.MethodOptions, "/test", nil)
	req.Header.Set("Origin", "http://localhost:3000")
	req.Header.Set("Access-Control-Request-Method", "POST")

	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	maxAge := w.Header().Get("Access-Control-Max-Age")
	if maxAge != "" {
		t.Fatalf("expected no Access-Control-Max-Age header, got %q", maxAge)
	}
}

func TestCORSMaxAgeNotSetOnNonPreflight(t *testing.T) {
	app := DefaultRouter()
	app.Use(AllowCORS([]string{"*"}, CORSConfig{MaxAge: 86400}))

	app.GET("/test", func(c *Context) {
		c.String(200, "success")
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Origin", "http://localhost:3000")

	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	maxAge := w.Header().Get("Access-Control-Max-Age")
	if maxAge != "" {
		t.Fatalf("expected no Access-Control-Max-Age on non-preflight, got %q", maxAge)
	}
}

func TestCORSMaxAgeBlockedForDisallowedOrigin(t *testing.T) {
	app := DefaultRouter()
	app.Use(AllowCORS([]string{"https://allowed.com"}, CORSConfig{MaxAge: 86400}))

	app.POST("/test", func(c *Context) {
		c.String(200, "success")
	})

	req := httptest.NewRequest(http.MethodOptions, "/test", nil)
	req.Header.Set("Origin", "https://evil.com")
	req.Header.Set("Access-Control-Request-Method", "POST")

	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if w.Code != http.StatusForbidden {
		t.Fatalf("expected 403 for disallowed origin, got %d", w.Code)
	}

	maxAge := w.Header().Get("Access-Control-Max-Age")
	if maxAge != "" {
		t.Fatalf("expected no Access-Control-Max-Age for blocked origin, got %q", maxAge)
	}
}
