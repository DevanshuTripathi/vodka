package vodka

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestJSONLoggerEmitsStructuredLog(t *testing.T) {
	app := NewRouter()

	var logBuffer bytes.Buffer
	originalWriter := log.Writer()
	originalFlags := log.Flags()
	log.SetOutput(&logBuffer)
	log.SetFlags(0)
	t.Cleanup(func() {
		log.SetOutput(originalWriter)
		log.SetFlags(originalFlags)
	})

	app.Use(JSONLogger())
	app.GET("/users", func(c *Context) {
		c.String(http.StatusCreated, "ok")
	})

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.RemoteAddr = "203.0.113.10:4312"
	req.Header.Set("X-Request-ID", "req-42")
	w := httptest.NewRecorder()

	app.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", w.Code)
	}

	line := strings.TrimSpace(logBuffer.String())
	if line == "" {
		t.Fatal("expected JSON logger output, got empty log")
	}

	var entry map[string]any
	if err := json.Unmarshal([]byte(line), &entry); err != nil {
		t.Fatalf("expected valid JSON log output, got error: %v", err)
	}

	if entry["method"] != http.MethodGet {
		t.Fatalf("expected method GET, got %v", entry["method"])
	}

	if entry["path"] != "/users" {
		t.Fatalf("expected path /users, got %v", entry["path"])
	}

	if status, ok := entry["status"].(float64); !ok || int(status) != http.StatusCreated {
		t.Fatalf("expected status 201, got %v", entry["status"])
	}

	if entry["ip"] != "203.0.113.10" {
		t.Fatalf("expected ip 203.0.113.10, got %v", entry["ip"])
	}

	if entry["request_id"] != "req-42" {
		t.Fatalf("expected request_id req-42, got %v", entry["request_id"])
	}

	latency, ok := entry["latency"].(string)
	if !ok || latency == "" {
		t.Fatalf("expected non-empty latency, got %v", entry["latency"])
	}

	timestamp, ok := entry["timestamp"].(string)
	if !ok || timestamp == "" {
		t.Fatalf("expected non-empty timestamp, got %v", entry["timestamp"])
	}

	if _, err := time.Parse(time.RFC3339, timestamp); err != nil {
		t.Fatalf("expected RFC3339 timestamp, got %q: %v", timestamp, err)
	}
}

