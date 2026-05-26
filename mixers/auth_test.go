package mixers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DevanshuTripathi/vodka"
)

func TestBearerAuthUsesErrorHandlerPipeline(t *testing.T) {
	tests := []struct {
		name            string
		authHeader      string
		expectedCode    int
		expectedMessage vodka.M
	}{
		{
			name:            "missing bearer header",
			authHeader:      "",
			expectedCode:    http.StatusUnauthorized,
			expectedMessage: vodka.M{"success": false, "message": "Unauthorized: Missing or malformed token"},
		},
		{
			name:            "invalid token",
			authHeader:      "Bearer wrong-token",
			expectedCode:    http.StatusUnauthorized,
			expectedMessage: vodka.M{"success": false, "message": "Unauthorized: Invalid token"},
		},
	}

	app := vodka.DefaultRouter()
	app.GET("/protected", BearerAuth("user", func(token string) (any, bool) {
		return nil, token == "valid-token"
	}), func(c *vodka.Context) {
		c.JSON(http.StatusOK, vodka.M{"success": true})
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/protected", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}

			w := httptest.NewRecorder()
			app.ServeHTTP(w, req)

			if w.Code != tt.expectedCode {
				t.Fatalf("expected status %d, got %d", tt.expectedCode, w.Code)
			}

			var got vodka.M
			if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
				t.Fatalf("failed to decode response body: %v", err)
			}

			if got["success"] != tt.expectedMessage["success"] || got["message"] != tt.expectedMessage["message"] {
				t.Fatalf("expected body %v, got %v", tt.expectedMessage, got)
			}
		})
	}
}
