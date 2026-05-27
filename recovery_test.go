package vodka

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRecovery_NoPanicWritesNormally(t *testing.T) {
	app := DefaultRouter()

	app.GET("/ok", func(c *Context) {
		c.JSON(http.StatusOK, M{"status": "ok"})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ok", nil)
	app.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestRecovery_PanicWrite(t *testing.T) {
	app := DefaultRouter()

	app.GET("/panic", func(c *Context) {
		panic("panic")
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/panic", nil)
	app.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected 500, Got %d", w.Code)
	}

}

func TestRecovery_PanicAfterWrite(t *testing.T) {
	app := DefaultRouter()

	app.GET("/panic-after-write", func(c *Context) {
		c.JSON(http.StatusOK, M{"data": "partial"})
		panic("panic after write")
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/panic-after-write", nil)
	app.ServeHTTP(w, req)

	// Headers already sent, Recovery should NOT append a second body
	body := w.Body.String()
	if strings.Count(body, "{") > 1 {
		t.Errorf("expected single JSON body, got multiple concatenated: %s", body)
	}
}
