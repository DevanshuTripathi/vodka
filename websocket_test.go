package vodka

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

func wsURL(s *httptest.Server) string {
	return "ws" + strings.TrimPrefix(s.URL, "http")
}

func TestWSEcho(t *testing.T) {
	app := NewRouter()
	app.AllowWSOrigins([]string{"*"})

	app.WS("/ws", func(c *WSContext) {
		msgType, msg, err := c.ReadMessage()
		if err != nil {
			t.Errorf("ReadMessage error: %v", err)
			return
		}
		c.WriteMessage(msgType, msg)
	})

	s := httptest.NewServer(app)
	defer s.Close()

	conn, _, err := websocket.DefaultDialer.Dial(wsURL(s)+"/ws", nil)
	if err != nil {
		t.Fatalf("dial error: %v", err)
	}
	defer conn.Close()

	want := "hello"
	conn.WriteMessage(websocket.TextMessage, []byte(want))

	_, got, err := conn.ReadMessage()
	if err != nil {
		t.Fatalf("read error: %v", err)
	}
	if string(got) != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestWSJSONRoundTrip(t *testing.T) {
	app := NewRouter()
	app.AllowWSOrigins([]string{"*"})

	type Msg struct {
		Value string `json:"value"`
	}

	app.WS("/ws/json", func(c *WSContext) {
		var m Msg
		if err := c.ReadJSON(&m); err != nil {
			return
		}
		c.WriteJSON(Msg{Value: "pong:" + m.Value})
	})

	s := httptest.NewServer(app)
	defer s.Close()

	conn, _, err := websocket.DefaultDialer.Dial(wsURL(s)+"/ws/json", nil)
	if err != nil {
		t.Fatalf("dial error: %v", err)
	}
	defer conn.Close()

	conn.WriteJSON(Msg{Value: "ping"})

	var reply Msg
	if err := conn.ReadJSON(&reply); err != nil {
		t.Fatalf("ReadJSON error: %v", err)
	}
	if reply.Value != "pong:ping" {
		t.Errorf("got %q, want %q", reply.Value, "pong:ping")
	}
}

func TestWSMiddlewareAbort(t *testing.T) {
	app := NewRouter()
	app.AllowWSOrigins([]string{"*"})

	auth := func(c *Context) {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		c.Abort()
	}

	api := app.Group("/api", auth)
	api.WS("/ws", func(c *WSContext) {
		t.Error("handler should not run when middleware aborts")
	})

	s := httptest.NewServer(app)
	defer s.Close()

	_, resp, _ := websocket.DefaultDialer.Dial(wsURL(s)+"/api/ws", nil)
	if resp == nil || resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("expected 401, got %v", resp)
	}
}

func TestWSContextSetGet(t *testing.T) {
	c := &WSContext{Keys: make(map[string]any)}
	c.Set("user", "sounak")

	val, exists := c.Get("user")
	if !exists {
		t.Error("key should exist")
	}
	if val != "sounak" {
		t.Errorf("got %v, want sounak", val)
	}

	_, exists = c.Get("missing")
	if exists {
		t.Error("missing key should not exist")
	}
}

func TestWSContextQuery(t *testing.T) {
	req, _ := http.NewRequest("GET", "/ws?room=general", nil)
	c := &WSContext{Request: req}

	if got := c.Query("room"); got != "general" {
		t.Errorf("got %q, want %q", got, "general")
	}
	if got := c.Query("missing"); got != "" {
		t.Errorf("got %q, want empty", got)
	}
}

func TestAllowWSOriginsBlocks(t *testing.T) {
	app := NewRouter()
	app.AllowWSOrigins([]string{"https://allowed.com"})

	app.WS("/ws", func(c *WSContext) {})

	s := httptest.NewServer(app)
	defer s.Close()

	header := http.Header{"Origin": {"https://blocked.com"}}
	_, resp, _ := websocket.DefaultDialer.Dial(wsURL(s)+"/ws", header)
	if resp == nil || resp.StatusCode != http.StatusForbidden {
		t.Errorf("expected 403 for blocked origin, got %v", resp)
	}
}

func TestAllowWSOriginsAllows(t *testing.T) {
	app := NewRouter()
	app.AllowWSOrigins([]string{"https://allowed.com"})

	app.WS("/ws", func(c *WSContext) {
		c.WriteMessage(websocket.TextMessage, []byte("ok"))
	})

	s := httptest.NewServer(app)
	defer s.Close()

	header := http.Header{"Origin": {"https://allowed.com"}}
	conn, _, err := websocket.DefaultDialer.Dial(wsURL(s)+"/ws", header)
	if err != nil {
		t.Fatalf("expected connection to succeed: %v", err)
	}
	defer conn.Close()

	_, msg, err := conn.ReadMessage()
	if err != nil || string(msg) != "ok" {
		t.Errorf("got %q, want ok", msg)
	}
}
