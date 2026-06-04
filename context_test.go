package vodka

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestAbortStopsChain(t *testing.T) {
	calls := []string{}

	// First handler aborts
	handler1 := func(c *Context) {
		calls = append(calls, "h1")
		c.Abort()
	}

	// This should NOT run
	handler2 := func(c *Context) {
		calls = append(calls, "h2")
	}

	c := &Context{handlers: []HandlerFunc{handler1, handler2}, index: -1}
	c.Next()

	if len(calls) != 1 || calls[0] != "h1" {
		t.Errorf("got %v, want [h1]", calls)
	}
}

func TestKeys(t *testing.T) {
	app := DefaultRouter()

	app.Use(func(c *Context) {
		c.Set("M1", "First")

		c.Next()
	})

	app.Use(func(c *Context) {
		c.Set("M2", "Second")

		c.Next()
	})

	app.Use(func(c *Context) {
		c.Set("M1", "Changed")

		c.Next()
	})

	app.GET("/test", func(c *Context) {
		m1, exists := c.Get("M1")
		if !exists {
			t.Fatal("Value for M1 does not exist")
		}

		if m1.(string) != "Changed" {
			t.Fatalf("expected Changed, got %v", m1)
		}

		m2, exists := c.Get("M2")
		if !exists {
			t.Fatal("Value for M2 does not exist")
		}

		if m2.(string) != "Second" {
			t.Fatalf("expected Second, got %v", m2)
		}
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	app.ServeHTTP(w, req)
}

func TestBindJSON(t *testing.T) {
	type User struct {
		Username string `json:"username"`
		Age      int8   `json:"age"`
	}

	app := DefaultRouter()

	app.POST("/test", func(c *Context) {
		var user User

		c.BindJSON(&user)

		if user.Username != "blufftunic" {
			t.Errorf("got %s, expected blufftunic", user.Username)
		}

		if user.Age != 20 {
			t.Errorf("got %d, expected blufftunic", user.Age)
		}

		c.JSON(200, M{
			"message": "success",
		})
	})

	body, _ := json.Marshal(M{
		"username": "blufftunic",
		"age":      20,
	})

	req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	app.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Wrong Status Code: got %d, expected %d", w.Code, 200)
	}

	var response M

	json.Unmarshal(w.Body.Bytes(), &response)

	if response["message"] != "success" {
		t.Errorf("got %s, expected success", response["message"])
	}
}

func newTestContext(method, target string) *Context {
	return &Context{Request: httptest.NewRequest(method, target, nil)}
}

func TestQuery(t *testing.T) {
	tests := []struct {
		name   string
		target string
		key    string
		want   string
	}{
		{name: "existing key", target: "/?page=2&sort=name", key: "page", want: "2"},
		{name: "another key", target: "/?page=2&sort=name", key: "sort", want: "name"},
		{name: "missing key", target: "/?page=2", key: "missing", want: ""},
		{name: "no query string", target: "/", key: "page", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestContext(http.MethodGet, tt.target)
			if got := c.Query(tt.key); got != tt.want {
				t.Fatalf("Query(%q) = %q, want %q", tt.key, got, tt.want)
			}
		})
	}
}

func TestDefaultQuery(t *testing.T) {
	tests := []struct {
		name         string
		target       string
		key          string
		defaultValue string
		want         string
	}{
		{name: "uses query value", target: "/?limit=50", key: "limit", defaultValue: "10", want: "50"},
		{name: "uses default when missing", target: "/", key: "limit", defaultValue: "10", want: "10"},
		{name: "uses default when empty", target: "/?limit=", key: "limit", defaultValue: "10", want: "10"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestContext(http.MethodGet, tt.target)
			if got := c.DefaultQuery(tt.key, tt.defaultValue); got != tt.want {
				t.Fatalf("DefaultQuery(%q, %q) = %q, want %q", tt.key, tt.defaultValue, got, tt.want)
			}
		})
	}
}

func TestParam(t *testing.T) {
	c := &Context{
		Request: httptest.NewRequest(http.MethodGet, "/users/42", nil),
		Params: httprouter.Params{
			{Key: "id", Value: "42"},
			{Key: "name", Value: "vodka"},
		},
	}

	tests := []struct {
		key  string
		want string
	}{
		{key: "id", want: "42"},
		{key: "name", want: "vodka"},
		{key: "missing", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			if got := c.Param(tt.key); got != tt.want {
				t.Fatalf("Param(%q) = %q, want %q", tt.key, got, tt.want)
			}
		})
	}
}

func TestQueryInt(t *testing.T) {
	tests := []struct {
		name      string
		target    string
		key       string
		want      int
		wantError string
	}{
		{name: "valid int", target: "/?page=10", key: "page", want: 10},
		{name: "invalid int", target: "/?page=abc", key: "page", wantError: `query param "page" is not a valid int`},
		{name: "missing key", target: "/", key: "page", wantError: `query param "page" is not a valid int`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestContext(http.MethodGet, tt.target)
			got, err := c.QueryInt(tt.key)

			if tt.wantError != "" {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				if !strings.Contains(err.Error(), tt.wantError) {
					t.Fatalf("error = %q, want substring %q", err.Error(), tt.wantError)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("QueryInt(%q) = %d, want %d", tt.key, got, tt.want)
			}
		})
	}
}

func TestQueryBool(t *testing.T) {
	tests := []struct {
		name      string
		target    string
		key       string
		want      bool
		wantError string
	}{
		{name: "true", target: "/?active=true", key: "active", want: true},
		{name: "false", target: "/?active=false", key: "active", want: false},
		{name: "invalid bool", target: "/?active=maybe", key: "active", wantError: `query param "active" is not a valid bool`},
		{name: "missing key", target: "/", key: "active", wantError: `query param "active" is not a valid bool`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestContext(http.MethodGet, tt.target)
			got, err := c.QueryBool(tt.key)

			if tt.wantError != "" {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				if !strings.Contains(err.Error(), tt.wantError) {
					t.Fatalf("error = %q, want substring %q", err.Error(), tt.wantError)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("QueryBool(%q) = %v, want %v", tt.key, got, tt.want)
			}
		})
	}
}

func TestParamInt(t *testing.T) {
	tests := []struct {
		name      string
		params    httprouter.Params
		key       string
		want      int
		wantError string
	}{
		{
			name:   "valid int",
			params: httprouter.Params{{Key: "id", Value: "99"}},
			key:    "id",
			want:   99,
		},
		{
			name:      "invalid int",
			params:    httprouter.Params{{Key: "id", Value: "abc"}},
			key:       "id",
			wantError: `param "id" is not a valid int`,
		},
		{
			name:      "missing param",
			params:    httprouter.Params{},
			key:       "id",
			wantError: `param "id" is not a valid int`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Context{
				Request: httptest.NewRequest(http.MethodGet, "/", nil),
				Params:  tt.params,
			}
			got, err := c.ParamInt(tt.key)

			if tt.wantError != "" {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				if !strings.Contains(err.Error(), tt.wantError) {
					t.Fatalf("error = %q, want substring %q", err.Error(), tt.wantError)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("ParamInt(%q) = %d, want %d", tt.key, got, tt.want)
			}
		})
	}
}

func TestParamBool(t *testing.T) {
	tests := []struct {
		name      string
		params    httprouter.Params
		key       string
		want      bool
		wantError string
	}{
		{
			name:   "true",
			params: httprouter.Params{{Key: "enabled", Value: "true"}},
			key:    "enabled",
			want:   true,
		},
		{
			name:   "false",
			params: httprouter.Params{{Key: "enabled", Value: "0"}},
			key:    "enabled",
			want:   false,
		},
		{
			name:      "invalid bool",
			params:    httprouter.Params{{Key: "enabled", Value: "nope"}},
			key:       "enabled",
			wantError: `param "enabled" is not a valid bool`,
		},
		{
			name:      "missing param",
			params:    httprouter.Params{},
			key:       "enabled",
			wantError: `param "enabled" is not a valid bool`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Context{
				Request: httptest.NewRequest(http.MethodGet, "/", nil),
				Params:  tt.params,
			}
			got, err := c.ParamBool(tt.key)

			if tt.wantError != "" {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				if !strings.Contains(err.Error(), tt.wantError) {
					t.Fatalf("error = %q, want substring %q", err.Error(), tt.wantError)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("ParamBool(%q) = %v, want %v", tt.key, got, tt.want)
			}
		})
	}
}

func TestParamAndQueryHelpersViaRouter(t *testing.T) {
	app := DefaultRouter()

	app.GET("/users/:id", func(c *Context) {
		id, err := c.ParamInt("id")
		if err != nil {
			t.Fatalf("ParamInt: %v", err)
		}

		page, err := c.QueryInt("page")
		if err != nil {
			t.Fatalf("QueryInt: %v", err)
		}

		active, err := c.QueryBool("active")
		if err != nil {
			t.Fatalf("QueryBool: %v", err)
		}

		if id != 7 || page != 3 || !active {
			t.Fatalf("got id=%d page=%d active=%v, want id=7 page=3 active=true", id, page, active)
		}

		c.String(http.StatusOK, c.Param("id"))
	})

	req := httptest.NewRequest(http.MethodGet, "/users/7?page=3&active=true", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if w.Body.String() != "7" {
		t.Fatalf("expected body 7, got %q", w.Body.String())
	}
}

func TestContextCopy(t *testing.T) {
	app := NewRouter()

	var cp *Context

	app.GET("/users/:id", func(c *Context) {
		cp = c.Copy()

		id, err := c.ParamInt("id")
		if err != nil {
			t.Errorf("Error using ParamInt: %v", err)
		}

		if id != 67 {
			t.Errorf("Expected 67, got=%d", id)
		}

		topic := c.Query("topic")
		if topic != "test" {
			t.Errorf("Expected test, got=%s", topic)
		}
	})

	req := httptest.NewRequest(http.MethodGet, "/users/67?topic=test", nil)
	w := httptest.NewRecorder()

	app.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Eexpected 200, got %d", w.Code)
	}

	id, err := cp.ParamInt("id")
	if err != nil {
		t.Errorf("Error using ParamInt: %v", err)
	}

	if id != 67 {
		t.Errorf("Expected 67, got=%d", id)
	}

	topic := cp.Query("topic")
	if topic != "test" {
		t.Errorf("Expected test, got=%s", topic)
	}
}

func TestSetCookie(t *testing.T) {
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	c := &Context{Writer: rr, Request: req}

	c.SetCookie("token", "abc123", 3600)

	cookies := rr.Result().Cookies()
	if len(cookies) != 1 {
		t.Fatalf("expected 1 cookie, got %d", len(cookies))
	}
	if cookies[0].Name != "token" || cookies[0].Value != "abc123" {
		t.Errorf("got %s=%s, want token=abc123", cookies[0].Name, cookies[0].Value)
	}
	if cookies[0].MaxAge != 3600 {
		t.Errorf("got MaxAge %d, want 3600", cookies[0].MaxAge)
	}
}

func TestCookie(t *testing.T) {
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{Name: "session", Value: "xyz"})
	c := &Context{Writer: rr, Request: req}

	val, err := c.Cookie("session")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != "xyz" {
		t.Errorf("got %q, want xyz", val)
	}

	_, err = c.Cookie("missing")
	if err == nil {
		t.Error("expected error for missing cookie")
	}
}

func TestClearCookie(t *testing.T) {
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	c := &Context{Writer: rr, Request: req}

	c.ClearCookie("token")

	cookies := rr.Result().Cookies()
	if len(cookies) != 1 {
		t.Fatalf("expected 1 cookie, got %d", len(cookies))
	}
	if cookies[0].MaxAge != -1 {
		t.Errorf("got MaxAge %d, want -1", cookies[0].MaxAge)
	}
}

func TestContextRedirectStatus(t *testing.T) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest(
		http.MethodGet, "/old", nil,
	)

	c := &Context{
		Writer:  w,
		Request: req,
	}

	c.Redirect(http.StatusFound, "/login")

	if w.Code != http.StatusFound {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusFound, w.Code,
		)
	}

}
func TestContextRedirectLocation(t *testing.T) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest(
		http.MethodGet,
		"/old",
		nil,
	)

	c := &Context{
		Writer:  w,
		Request: req,
	}

	c.Redirect(http.StatusFound, "/login")

	location := w.Header().Get("Location")

	if location != "/login" {
		t.Fatalf(
			"expected location /login, got %s",
			location,
		)
	}
}

func TestContextPermanentRedirect(t *testing.T) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest(
		http.MethodGet,
		"/old",
		nil,
	)

	c := &Context{
		Writer:  w,
		Request: req,
	}

	c.Redirect(
		http.StatusMovedPermanently,
		"/new-path",
	)

	if w.Code != http.StatusMovedPermanently {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusMovedPermanently,
			w.Code,
		)
	}

	location := w.Header().Get("Location")

	if location != "/new-path" {
		t.Fatalf(
			"expected location /new-path, got %s",
			location,
		)
	}
}

type XMLUser struct {
	XMLName xml.Name `xml:"user"`
	Name    string   `xml:"name"`
}

func TestContextXMLStatus(t *testing.T) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest(
		http.MethodGet,
		"/",
		nil,
	)

	c := &Context{
		Writer:  w,
		Request: req,
	}

	c.XML(http.StatusOK, XMLUser{
		Name: "Utkarsh",
	})

	if w.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			w.Code,
		)
	}
}

func TestContextXMLContentType(t *testing.T) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest(
		http.MethodGet,
		"/",
		nil,
	)

	c := &Context{
		Writer:  w,
		Request: req,
	}

	c.XML(http.StatusOK, XMLUser{
		Name: "Utkarsh",
	})

	contentType := w.Header().Get("Content-Type")

	if contentType != "application/xml" {
		t.Fatalf(
			"expected application/xml, got %s",
			contentType,
		)
	}
}

func TestContextXMLBody(t *testing.T) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest(
		http.MethodGet,
		"/",
		nil,
	)

	c := &Context{
		Writer:  w,
		Request: req,
	}

	c.XML(http.StatusOK, XMLUser{
		Name: "Utkarsh",
	})

	body := w.Body.String()

	if !strings.Contains(
		body,
		"<name>Utkarsh</name>",
	) {
		t.Fatalf(
			"unexpected xml body: %s",
			body,
		)
	}
}

func TestClientIPBasic(t *testing.T) {
	app := DefaultRouter()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.RemoteAddr = "203.0.113.42:1234"

	c := &Context{
		Request: req,
		engine:  app,
	}

	ip := c.ClientIP()

	if ip != "203.0.113.42" {
		t.Fatalf("Expected=203.0.113.42:1234, Got=%s", ip)
	}
}

func TestClientIPWithTrustedProxy(t *testing.T) {
	app := DefaultRouter()

	app.SetTrustedProxies([]string{"10.0.0.5/8"})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.RemoteAddr = "10.0.0.5:1234"
	req.Header.Set("X-Forwarded-For", "203.0.113.42")

	c := &Context{
		Request: req,
		engine:  app,
	}

	ip := c.ClientIP()

	if ip != "203.0.113.42" {
		t.Errorf("Expected=203.0.113.42, Got=%s", ip)
	}
}

func TestClientIPIgnoresFakeXFF(t *testing.T) {
	app := DefaultRouter()

	app.SetTrustedProxies([]string{"10.0.0.0/8"})

	req := httptest.NewRequest("GET", "/", nil)
	req.RemoteAddr = "203.0.113.99:1234"
	req.Header.Set("X-Forwarded-For", "8.8.8.8")

	c := &Context{
		Request: req,
		engine:  app,
	}

	ip := c.ClientIP()

	if ip != "203.0.113.99" {
		t.Fatalf("Should ignore fake XFF, Got=%s", ip)
	}
}

func TestClientIPProxyChain(t *testing.T) {
	app := DefaultRouter()

	app.SetTrustedProxies([]string{"10.0.0.0/8", "172.16.0.0/12"})

	req := httptest.NewRequest("GET", "/", nil)

	req.RemoteAddr = "10.0.0.5:1234"

	req.Header.Set("X-Forwarded-For", "203.0.113.42, 172.16.0.1")

	c := &Context{
		Request: req,
		engine:  app,
	}

	ip := c.ClientIP()

	if ip != "203.0.113.42" {
		t.Fatalf("Expected=203.0.113.42, Got=%s", ip)
	}
}

func TestClientIPInvalidHeader(t *testing.T) {
	app := DefaultRouter()

	app.SetTrustedProxies([]string{
		"10.0.0.0/8",
	})

	req := httptest.NewRequest("GET", "/", nil)

	req.RemoteAddr = "10.0.0.5:1234"

	req.Header.Set(
		"X-Forwarded-For",
		"not-an-ip",
	)

	c := &Context{
		Request: req,
		engine:  app,
	}

	ip := c.ClientIP()

	if ip != "10.0.0.5" {
		t.Fatalf("Expected=10.0.0.5, Got=%s", ip)
	}
}
