package vodka

import (
	"net/http"
	"testing"
)

func BenchmarkRouterStatic(b *testing.B) {
	app := NewRouter()
	app.GET("/test", func(c *Context) {})
	runReq(b, app, http.MethodGet, "/test")
}

func BenchmarkSingleRouteParam(b *testing.B) {
	app := NewRouter()
	app.GET("/users/:id", func(c *Context) {})
	runReq(b, app, http.MethodGet, "/users/123")
}

func BenchmarkDeepParam(b *testing.B) {
	app := NewRouter()
	app.GET("/api/v1/projects/:project/user/:userId/objective/:obj/comment/:comment", func(c *Context) {
		_ = c.Param("project")
		_ = c.Param("userId")
		_ = c.Param("obj")
		_ = c.Param("comment")
	})
	runReq(b, app, http.MethodGet, "/api/v1/projects/vodka/user/123/objective/test/comment/67")
}

func BenchmarkMiddlewareChain(b *testing.B) {
	app := NewRouter()

	for i := 0; i < 10; i++ {
		app.Use(func(c *Context) {
			c.Next()
		})
	}

	app.GET("/test", func(c *Context) {
		c.String(200, "tested")
	})

	runReq(b, app, http.MethodGet, "/test")
}

// Mock Writer Struct to create mock requests for benchmarking
type mockWriter struct {
	headers http.Header
}

// Creates a mockWriter with empty headers
func newMockWriter() *mockWriter {
	return &mockWriter{
		http.Header{},
	}
}

func (m *mockWriter) Header() (h http.Header) {
	return m.headers
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockWriter) WriteHeader(int) {}

// Function to make a mock request for benchmarking purposes
func runReq(b *testing.B, e *Engine, method, path string) {
	req, err := http.NewRequest(method, path, nil) // Creates a fake request
	if err != nil {
		panic(err)
	}
	w := newMockWriter()
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		e.ServeHTTP(w, req)
	}
}
