package mixers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/DevanshuTripathi/vodka"
)

func captureLog(t *testing.T) *bytes.Buffer {
	t.Helper()
	buf := &bytes.Buffer{}
	log.SetOutput(buf)
	t.Cleanup(func() { log.SetOutput(nil) })
	return buf
}

func parseLogEntry(t *testing.T, buf *bytes.Buffer) jsonLogEntry {
	t.Helper()
	line := buf.String()
	idx := strings.Index(line, "{")
	if idx == -1 {
		t.Fatalf("no JSON object found in log output: %q", line)
	}
	var entry jsonLogEntry
	if err := json.Unmarshal([]byte(strings.TrimSpace(line[idx:])), &entry); err != nil {
		t.Fatalf("failed to parse log entry: %v\nraw: %q", err, line)
	}
	return entry
}

func TestJSONLogger_LogsCorrectFields(t *testing.T) {

	buf := captureLog(t)

	app := vodka.NewRouter()
	app.Use(JSONLogger())
	app.GET("/test", func(c *vodka.Context) {
		c.JSON(http.StatusOK, vodka.M{"ok": true})
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	entry := parseLogEntry(t, buf)

	if entry.Method != http.MethodGet {
		t.Errorf("expected method GET, got %q", entry.Method)
	}
	if entry.Path != "/test" {
		t.Errorf("expected path /test, got %q", entry.Path)
	}
	if entry.Status != http.StatusOK {
		t.Errorf("expected status 200, got %d", entry.Status)
	}
	if entry.Latency == "" {
		t.Error("expected non-empty latency")
	}
	if entry.Timestamp == "" {
		t.Error("expected non-empty timestamp")
	}
	if entry.IP == "" {
		t.Error("expected non-empty IP")
	}
}

func TestJSONLogger_LogsCorrectStatusCode(t *testing.T) {

	buf := captureLog(t)

	app := vodka.NewRouter()
	app.Use(JSONLogger())
	app.GET("/notfound", func(c *vodka.Context) {
		c.JSON(http.StatusNotFound, vodka.M{"error": "not found"})
	})

	req := httptest.NewRequest(http.MethodGet, "/notfound", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	entry := parseLogEntry(t, buf)
	if entry.Status != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", entry.Status)
	}
}

func TestJSONLogger_DefaultsTo200WhenNoWriteHeader(t *testing.T) {

	buf := captureLog(t)

	app := vodka.NewRouter()
	app.Use(JSONLogger())
	app.GET("/silent", func(c *vodka.Context) {
		c.Writer.Write([]byte("ok")) 
	})

	req := httptest.NewRequest(http.MethodGet, "/silent", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	entry := parseLogEntry(t, buf)
	if entry.Status != http.StatusOK {
		t.Errorf("expected default status 200, got %d", entry.Status)
	}
}

func TestJSONLogger_TimestampIsRFC3339(t *testing.T) {

	buf := captureLog(t)

	app := vodka.NewRouter()
	app.Use(JSONLogger())
	app.GET("/ts", func(c *vodka.Context) {
		c.JSON(http.StatusOK, vodka.M{})
	})

	req := httptest.NewRequest(http.MethodGet, "/ts", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	entry := parseLogEntry(t, buf)
	if _, err := time.Parse(time.RFC3339, entry.Timestamp); err != nil {
		t.Errorf("timestamp %q is not valid RFC3339: %v", entry.Timestamp, err)
	}
}

func TestJSONLogger_OutputIsValidJSON(t *testing.T) {

	buf := captureLog(t)

	app := vodka.NewRouter()
	app.Use(JSONLogger())
	app.GET("/json", func(c *vodka.Context) {
		c.JSON(http.StatusOK, vodka.M{"msg": "hello"})
	})

	req := httptest.NewRequest(http.MethodGet, "/json", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	line := buf.String()
	idx := strings.Index(line, "{")
	if idx == -1 {
		t.Fatalf("no JSON found in log: %q", line)
	}
	var raw map[string]any
	if err := json.Unmarshal([]byte(strings.TrimSpace(line[idx:])), &raw); err != nil {
		t.Errorf("log output is not valid JSON: %v\nraw: %q", err, line)
	}
}

func TestJSONLogger_LogsClientIP(t *testing.T) {

	buf := captureLog(t)

	app := vodka.NewRouter()
	app.Use(JSONLogger())
	app.GET("/ip", func(c *vodka.Context) {
		c.JSON(http.StatusOK, vodka.M{})
	})

	req := httptest.NewRequest(http.MethodGet, "/ip", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	entry := parseLogEntry(t, buf)
	if entry.IP == "" {
		t.Error("expected non-empty IP in log entry")
	}
	if strings.Contains(entry.IP, ":") {
		t.Errorf("IP should not contain port, got %q", entry.IP)
	}
}

func TestJSONLogger_DoesNotAffectResponse(t *testing.T) {

	captureLog(t)

	app := vodka.NewRouter()
	app.Use(JSONLogger())
	app.GET("/body", func(c *vodka.Context) {
		c.JSON(http.StatusCreated, vodka.M{"created": true})
	})

	req := httptest.NewRequest(http.MethodGet, "/body", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected response 201, got %d", rec.Code)
	}
	expected := `{"created":true}` + "\n"
	if rec.Body.String() != expected {
		t.Errorf("expected body %q, got %q", expected, rec.Body.String())
	}
}