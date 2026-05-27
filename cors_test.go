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