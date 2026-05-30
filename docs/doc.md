# Vodka Documentation

Vodka is a developer friendly Go framework for full stack applications.

---

## Table of Contents

1. [Installation](#installation)
2. [Quick Start](#quick-start)
3. [Router](#router)
4. [Routing](#routing)
5. [Responses](#responses)
6. [Request Binding](#request-binding)
7. [Middleware](#middleware)
8. [Error Handling](#error-handling)
9. [Context](#context)
10. [Built-in Middleware](#built-in-middleware)
11. [Mixers (Plugin Middleware)](#mixers-plugin-middleware)
12. [Server-Sent Events (SSE)](#server-sent-events-sse)
13. [WebSocket](#websocket)
14. [Static Files and SPA](#static-files-and-spa)
15. [Trusted Proxies](#trusted-proxies)

---

## Installation

Install the Vodka package:

```bash
go get github.com/DevanshuTripathi/vodka
```

Install the Vodka CLI:

```bash
go install github.com/DevanshuTripathi/vodka/cmd/vodka@latest
```

---

## Quick Start

### Create a new project with the Vodka CLI

```bash
vodka create myapp
cd myapp
vodka run
```

### Manually

```go
package main

import (
    "log"
    "github.com/DevanshuTripathi/vodka"
)

func main() {
    app := vodka.DefaultRouter()

    app.GET("/ping", func(c *vodka.Context) {
        c.String(200, "pong")
    })

    app.GET("/hello/:name", func(c *vodka.Context) {
        name := c.Param("name")
        c.JSON(200, vodka.M{
            "message": "Greetings!",
            "name":    name,
        })
    })

    if err := app.Run(":8080"); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
```

---

## Router

### DefaultRouter vs NewRouter

| | `DefaultRouter()` | `NewRouter()` |
|---|---|---|
| Logger | ✅ included | ❌ not included |
| Recovery | ✅ included | ❌ not included |
| ErrorHandler | ✅ included | ❌ not included |
| Use case | Most apps | Full manual control |

```go
// DefaultRouter — batteries included
app := vodka.DefaultRouter()

// NewRouter — bare bones, add only what you need
app := vodka.NewRouter()
app.Use(vodka.Logger(), vodka.Recovery(), vodka.ErrorHandler())
```

### Running the Server

```go
if err := app.Run(":8080"); err != nil {
    log.Fatalf("Server failed to start: %v", err)
}
```

---

## Routing

### HTTP Methods

```go
app.GET("/path", handler)
app.POST("/path", handler)
app.PUT("/path", handler)
app.PATCH("/path", handler)
app.DELETE("/path", handler)
app.HEAD("/path", handler)
```

### URL Parameters

```go
app.GET("/users/:id", func(c *vodka.Context) {
    id := c.Param("id")
    c.JSON(200, vodka.M{"id": id})
})
```

### Query Parameters

```go
app.GET("/search", func(c *vodka.Context) {
    q := c.Query("q")               // returns "" if missing
    page := c.DefaultQuery("page", "1") // returns "1" if missing

    // Typed helpers
    limit, err := c.QueryInt("limit")
    active, err := c.QueryBool("active")

    c.JSON(200, vodka.M{"q": q, "page": page, "limit": limit})
})
```

### Route Groups

Route groups let you prefix routes and apply middleware to a subset of routes only.

```go
app := vodka.NewRouter()
app.Use(vodka.Logger(), vodka.Recovery())

// All routes under /api get ErrorHandler middleware
api := app.Group("/api", vodka.ErrorHandler())

api.GET("/users/:id", func(c *vodka.Context) {
    id := c.Param("userId")
    c.JSON(200, vodka.M{"userId": id, "name": "generic username"})
})

// Add middleware to an existing group later
api.Use(someOtherMiddleware())
```

---

## Responses

### JSON Response

```go
app.GET("/user", func(c *vodka.Context) {
    c.JSON(200, vodka.M{
        "id":   1,
        "name": "Alice",
    })
})
```

### String Response

```go
app.GET("/ping", func(c *vodka.Context) {
    c.String(200, "pong")
})
```

### vodka.M Shorthand

`vodka.M` is a shorthand for `map[string]any`. Use it anywhere you'd write a JSON object.

```go
vodka.M{
    "success": true,
    "data": vodka.M{
        "id": 42,
    },
}
```

---

## Request Binding

### BindJSON

Binds a JSON request body to a struct and validates fields using `validate` struct tags.

```go
type User struct {
    Email    string `json:"email"    validate:"required,email"`
    Password string `json:"password" validate:"min=8"`
}

app.POST("/create", func(c *vodka.Context) {
    var user User
    if err := c.BindJSON(&user); err != nil {
        c.Error(400, err)
        return
    }
    c.JSON(200, vodka.M{"email": user.Email})
})
```

### Form Values

```go
app.POST("/form", func(c *vodka.Context) {
    name := c.FormValue("name")
    c.JSON(200, vodka.M{"name": name})
})
```

### File Upload

```go
app.POST("/upload", func(c *vodka.Context) {
    fileHeader, err := c.FormFile("file")
    if err != nil {
        c.Error(400, err)
        return
    }

    if err := c.SaveUploadedFile(fileHeader, "./uploads/"+fileHeader.Filename); err != nil {
        c.Error(500, err)
        return
    }

    c.JSON(200, vodka.M{
        "filename": fileHeader.Filename,
        "size":     fileHeader.Size,
    })
})
```

---

## Middleware

### Using Built-in Middleware

```go
app := vodka.NewRouter()
app.Use(vodka.Logger(), vodka.Recovery(), vodka.ErrorHandler())
```

### Writing Custom Middleware

A middleware is a `HandlerFunc` that calls `c.Next()` to pass control to the next handler.

```go
func RequestTimer() vodka.HandlerFunc {
    return func(c *vodka.Context) {
        start := time.Now()
        c.Next()
        log.Printf("[%s] %s %v", c.Request.Method, c.Request.URL.Path, time.Since(start))
    }
}

app.Use(RequestTimer())
```

### c.Next() and c.Abort()

- `c.Next()` — continues execution down the middleware chain
- `c.Abort()` — stops the chain, no further handlers are called
- `c.Error(status, err)` — records an error and aborts the chain

```go
func AuthGuard() vodka.HandlerFunc {
    return func(c *vodka.Context) {
        token := c.Request.Header.Get("Authorization")
        if token == "" {
            c.Error(401, errors.New("unauthorized"))
            return
        }
        c.Next()
    }
}
```

---

## Error Handling

### c.Error()

`c.Error()` appends an error to the context and aborts the chain. `ErrorHandler` middleware picks it up and sends the JSON response.

```go
app.GET("/resource", func(c *vodka.Context) {
    if somethingWentWrong {
        c.Error(500, errors.New("something went wrong"))
        return
    }
    c.JSON(200, vodka.M{"status": "ok"})
})
```

### ErrorHandler Middleware

`ErrorHandler` must be present (either via `DefaultRouter` or added manually) to render errors as JSON responses.

```go
api := app.Group("/api", vodka.ErrorHandler())
```

Error response format:

```json
{
  "success": false,
  "message": "something went wrong"
}
```

---

## Context

### Keys — Set / Get

Store and retrieve arbitrary values scoped to the current request.

```go
// Set a value in middleware
c.Set("userID", 42)

// Get it in a handler
userID, exists := c.Get("userID")
```

### Cookies

```go
// Set a cookie
c.SetCookie("session", "abc123", 3600)

// Read a cookie
value, err := c.Cookie("session")

// Clear a cookie
c.ClearCookie("session")
```

### ClientIP

Returns the real client IP, respecting trusted proxy headers.

```go
ip := c.ClientIP()
```

### Copy (Async Use)

Use `c.Copy()` to safely pass context to a goroutine after the request handler returns.

```go
app.GET("/async", func(c *vodka.Context) {
    cp := c.Copy()
    go func() {
        log.Println("async task for", cp.Request.URL.Path)
    }()
    c.JSON(200, vodka.M{"status": "accepted"})
})
```

---

## Built-in Middleware

### Logger

Logs method, path, status code, and latency for every request.

```go
app.Use(vodka.Logger())
// Output: GET /api/users 200 1.2ms
```

### Recovery

Catches panics in handlers, logs them, and returns a `500 Internal Server Error` response.

```go
app.Use(vodka.Recovery())
```

### CORS

```go
// Allow all origins
app.Use(vodka.AllowCORS([]string{"*"}))

// Allow specific origins
app.Use(vodka.AllowCORS([]string{"https://myapp.com", "https://admin.myapp.com"}))
```

---

## Mixers (Plugin Middleware)

Import: `github.com/DevanshuTripathi/vodka/mixers`

### BearerAuth + JWTValidator

Validates a `Bearer` token from the `Authorization` header and stores decoded claims in context.

```go
import "github.com/DevanshuTripathi/vodka/mixers"

jwtMiddleware := mixers.BearerAuth("claims", mixers.JWTValidator("your-secret-key"))

secure := app.Group("/api/secure", jwtMiddleware)

secure.GET("/profile", func(c *vodka.Context) {
    claims, _ := c.Get("claims")
    c.JSON(200, vodka.M{"claims": claims})
})
```

### GenerateJWT

```go
token, err := mixers.GenerateJWT("your-secret-key", map[string]any{
    "username": "alice",
    "role":     "admin",
}, 24*time.Hour)
```

### RateLimiter

```go
limiter := mixers.NewRateLimiter(2.0, 10) // rate: 2 req/s, burst: 10
app.Use(mixers.RateLimiter(limiter))
```

### Gzip / GzipWithLevel

```go
// Default compression
api := app.Group("/api", mixers.Gzip())

// Custom compression level
api := app.Group("/api", mixers.GzipWithLevel(gzip.BestSpeed))
```

Only compresses responses when the client sends `Accept-Encoding: gzip`.

### RequestID

Generates a unique `X-Request-ID` header for every request and makes it available in context.

```go
app.Use(mixers.RequestID())

app.GET("/", func(c *vodka.Context) {
    requestID, _ := c.Get("request-id")
    c.JSON(200, vodka.M{"request_id": requestID})
})

// Custom header name
app.Use(mixers.RequestIDWithHeader("X-Correlation-ID"))
```

---

## Server-Sent Events (SSE)

SSE allows the server to push events to the client over a persistent HTTP connection.

### Registering SSE Routes

```go
app.SSE("/events/clock", func(c *vodka.SSEContext) {
    for {
        select {
        case <-c.Done():
            return // client disconnected
        default:
            c.Send("tick", vodka.M{"time": time.Now().Format(time.RFC3339)})
            time.Sleep(time.Second)
        }
    }
})
```

### Sending Events

```go
// Send a named event with a JSON payload
c.Send("eventName", vodka.M{"key": "value"})

// Send plain data
c.SendData("hello")

// Send a comment (keepalive)
c.SendComment("ping")
```

### URL Parameters in SSE

```go
app.SSE("/events/feed/:topic", func(c *vodka.SSEContext) {
    topic := c.Param("topic")
    for {
        select {
        case <-c.Done():
            return
        default:
            c.Send("message", vodka.M{"topic": topic})
            time.Sleep(2 * time.Second)
        }
    }
})
```

### Client Disconnect Detection

`c.Done()` returns a channel that is closed when the client disconnects. Always select on it to avoid goroutine leaks.

```go
case <-c.Done():
    return
```

### SSELogger Mixer

```go
app.SSE("/events", mixers.SSELogger(func(c *vodka.SSEContext) {
    // your handler
}))
```

---

## WebSocket

### Registering WS Routes

```go
app.WS("/ws/echo", func(c *vodka.WSContext) {
    for {
        msgType, msg, err := c.ReadMessage()
        if err != nil {
            return
        }
        c.WriteMessage(msgType, msg)
    }
})
```

### JSON over WebSocket

```go
app.WS("/ws/json", func(c *vodka.WSContext) {
    type Ping struct{ Message string `json:"message"` }
    type Pong struct{ Reply   string `json:"reply"` }

    for {
        var p Ping
        if err := c.ReadJSON(&p); err != nil {
            return
        }
        c.WriteJSON(Pong{Reply: "pong: " + p.Message})
    }
})
```

### Origin Control

```go
app.AllowWSOrigins([]string{
    "http://localhost:3000",
    "https://myapp.com",
})
```

### WSLogger Mixer

```go
app.WS("/ws/echo", mixers.WSLogger(func(c *vodka.WSContext) {
    // your handler
}))
```

---

## Static Files and SPA

Serve a single-page application (SPA) where all unmatched routes fall back to `index.html`.

```go
app.ServeSPA("./public")
```

Place your built frontend assets in `./public`. All API routes registered before `ServeSPA` take priority.

---

## Trusted Proxies

Configure trusted proxies so `ClientIP()` correctly resolves the real client IP from `X-Forwarded-For` headers.

```go
if err := app.SetTrustedProxies([]string{
    "10.0.0.0/8",
    "172.16.0.0/12",
    "192.168.0.0/16",
}); err != nil {
    log.Fatalf("invalid proxy config: %v", err)
}
```

When a request comes from a trusted proxy, `c.ClientIP()` reads the leftmost non-trusted IP from `X-Forwarded-For`. When the request does not come from a trusted proxy, `RemoteAddr` is used directly.