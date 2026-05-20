package vodka

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

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
