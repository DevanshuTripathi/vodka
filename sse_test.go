package vodka

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// mockFlusher wraps ResponseRecorder to satisfy http.Flusher.
type mockFlusher struct {
	*httptest.ResponseRecorder
	flushed int
}

func (m *mockFlusher) Flush() { m.flushed++ }

func newSSEContext() (*SSEContext, *mockFlusher) {
	rr := httptest.NewRecorder()
	mf := &mockFlusher{ResponseRecorder: rr}
	req, _ := http.NewRequest("GET", "/events?topic=news", nil)
	sc := &SSEContext{
		Writer:  mf,
		flusher: mf,
		Keys:    make(map[string]any),
		Request: req,
	}
	return sc, mf
}

func TestSSESendFormat(t *testing.T) {
	sc, mf := newSSEContext()

	if err := sc.Send("update", M{"value": 1}); err != nil {
		t.Fatalf("Send error: %v", err)
	}

	body := mf.Body.String()
	if !strings.Contains(body, "event: update\n") {
		t.Errorf("missing event line, got: %q", body)
	}
	if !strings.Contains(body, `"value":1`) {
		t.Errorf("missing data, got: %q", body)
	}
	if mf.flushed != 1 {
		t.Errorf("expected 1 flush, got %d", mf.flushed)
	}
}

func TestSSESendDataFormat(t *testing.T) {
	sc, mf := newSSEContext()

	if err := sc.SendData(M{"msg": "hello"}); err != nil {
		t.Fatalf("SendData error: %v", err)
	}

	body := mf.Body.String()
	if !strings.HasPrefix(body, "data:") {
		t.Errorf("expected data: prefix, got: %q", body)
	}
	if strings.Contains(body, "event:") {
		t.Errorf("SendData should not write event line, got: %q", body)
	}
}

func TestSSESendCommentFormat(t *testing.T) {
	sc, mf := newSSEContext()

	if err := sc.SendComment("keep-alive"); err != nil {
		t.Fatalf("SendComment error: %v", err)
	}

	body := mf.Body.String()
	if !strings.HasPrefix(body, ": keep-alive") {
		t.Errorf("expected comment format, got: %q", body)
	}
}

func TestSSEContextSetGet(t *testing.T) {
	sc, _ := newSSEContext()
	sc.Set("role", "admin")

	val, exists := sc.Get("role")
	if !exists {
		t.Error("key should exist")
	}
	if val != "admin" {
		t.Errorf("got %v, want admin", val)
	}

	_, exists = sc.Get("missing")
	if exists {
		t.Error("missing key should not exist")
	}
}

func TestSSEContextQuery(t *testing.T) {
	sc, _ := newSSEContext()

	if got := sc.Query("topic"); got != "news" {
		t.Errorf("got %q, want news", got)
	}
	if got := sc.Query("missing"); got != "" {
		t.Errorf("got %q, want empty", got)
	}
}

func TestSSEHeaders(t *testing.T) {
	app := NewRouter()

	app.SSE("/events", func(c *SSEContext) {
		c.Send("ping", M{"ok": true})
	})

	s := httptest.NewServer(app)
	defer s.Close()

	resp, err := http.Get(s.URL + "/events")
	if err != nil {
		t.Fatalf("request error: %v", err)
	}
	defer resp.Body.Close()

	if ct := resp.Header.Get("Content-Type"); ct != "text/event-stream" {
		t.Errorf("Content-Type got %q, want text/event-stream", ct)
	}
	if cc := resp.Header.Get("Cache-Control"); cc != "no-cache" {
		t.Errorf("Cache-Control got %q, want no-cache", cc)
	}
}

func TestSSEEventStream(t *testing.T) {
	app := NewRouter()

	app.SSE("/events", func(c *SSEContext) {
		c.Send("tick", M{"n": 1})
		c.Send("tick", M{"n": 2})
	})

	s := httptest.NewServer(app)
	defer s.Close()

	resp, err := http.Get(s.URL + "/events")
	if err != nil {
		t.Fatalf("request error: %v", err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	eventCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "event: tick") {
			eventCount++
		}
	}
	if eventCount != 2 {
		t.Errorf("expected 2 tick events, got %d", eventCount)
	}
}

func TestSSEMiddlewareAbort(t *testing.T) {
	app := NewRouter()

	auth := func(c *Context) {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		c.Abort()
	}

	api := app.Group("/api", auth)
	api.SSE("/events", func(c *SSEContext) {
		t.Error("SSE handler should not run when middleware aborts")
	})

	s := httptest.NewServer(app)
	defer s.Close()

	resp, err := http.Get(s.URL + "/api/events")
	if err != nil {
		t.Fatalf("request error: %v", err)
	}
	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", resp.StatusCode)
	}
}
