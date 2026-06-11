# Vodka

![Go Version](https://img.shields.io/badge/Go-1.24+-blue)
![License](https://img.shields.io/badge/license-MIT-green)
![Status](https://img.shields.io/badge/status-active-success)

```text
__/\\\________/\\\_______/\\\\\_______/\\\\\\\\\\\\_____/\\\________/\\\_____/\\\\\\\\\____        
 _\/\\\_______\/\\\_____/\\\///\\\____\/\\\////////\\\__\/\\\_____/\\\//____/\\\\\\\\\\\\\__       
  _\//\\\______/\\\____/\\\/__\///\\\__\/\\\______\//\\\_\/\\\__/\\\//______/\\\/////////\\\_      
   __\//\\\____/\\\____/\\\______\//\\\_\/\\\_______\/\\\_\/\\\\\\//\\\_____\/\\\_______\/\\\_     
    ___\//\\\__/\\\____\/\\\_______\/\\\_\/\\\_______\/\\\_\/\\\//_\//\\\____\/\\\\\\\\\\\\\\\_    
     ____\//\\\/\\\_____\//\\\______/\\\__\/\\\_______\/\\\_\/\\\____\//\\\___\/\\\/////////\\\_   
      _____\//\\\\\_______\///\\\__/\\\____\/\\\_______/\\\__\/\\\_____\//\\\__\/\\\_______\/\\\_  
       ______\//\\\__________\///\\\\\/_____\/\\\\\\\\\\\\/___\/\\\______\//\\\_\/\\\_______\/\\\_ 
        _______\///_____________\/////_______\////////////_____\///________\///__\///________\///__
```

## A modern Go web framework focused on developer experience, full-stack workflow, and rapid iteration.

Vodka is a lightweight and high-performance HTTP framework for Go that combines clean routing, middleware chaining, authentication utilities, validation support, and a powerful CLI for building modern full-stack applications.

Unlike traditional Go frameworks that focus only on request handling, Vodka heavily emphasizes developer experience and fast iteration.

---

# Demo

## Project Scaffolding

![CLI Demo](./assets/demo.gif)

---

## Full Stack Workflow

![Workflow](./assets/workflow.gif)

---

# Table of Contents

- [Features](#features)
- [Why Vodka](#why-vodka)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Generated Backend Structure](#generated-backend-structure)
- [Minimal API Example](#minimal-api-example)
- [Using Vodka for APIs](#using-vodka-for-apis)
- [Core Concepts](#core-concepts)
- [Graceful Shutdown & Lifecycle Manager](#graceful-shutdown--lifecycle-manager)
- [Middleware](#middleware)
- [Validation](#validation)
- [Authentication](#authentication)
- [SPA Support](#full-stack-spa-support)
- [Production Build](#production-build)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)

---

# Features

| Feature | Included |
|---|---|
| Radix Tree Routing | ✅ |
| Middleware Chaining | ✅ |
| Route Groups | ✅ |
| JSON Binding | ✅ |
| Request Validation | ✅ |
| JWT Validation Helpers | ✅ |
| Bearer Auth Middleware | ✅ |
| Vite + React Scaffolding | ✅ |
| SPA Serving | ✅ |
| Panic Recovery Middleware | ✅ |
| Logger Middleware | ✅ |
| CORS Middleware | ✅ |
| Context Storage | ✅ |
| HTML Template Rendering | ✅ |

---

# Why Vodka?

Vodka combines:

- ⚡ Fast backend iteration
- ⚛️ React + Vite integration
- 🧩 Lightweight routing
- 🔐 Authentication helpers
- ✅ Validation support
- 🚀 Developer-first workflow

without requiring heavy configuration or excessive boilerplate.

---

# Prerequisites

Make sure the following tools are installed before using Vodka:

- Go 1.24+
- Node.js 20.19+ or newer
- npm

Verify installation:

```bash
go version
node -v
npm -v
```

---

# Installation

## Install the Vodka CLI

```bash
go install github.com/DevanshuTripathi/vodka/cmd/vodka@latest
```

Make sure your Go bin directory is added to your system `PATH`.

---

## Linux / macOS

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

---

## Windows

Add the following directory to Environment Variables:

```text
%USERPROFILE%\go\bin
```

---

# Quick Start

## Create a Full-Stack App

```bash
vodka create <project-name> [location] [--minimal]
```

Basic usage:

```bash
vodka create my-app
```

You can also specify a target directory to scaffold the project into:

```bash
vodka create my-app /path/to/parent-dir
```

The project will be created at `<location>/<name>` (e.g. `/path/to/parent-dir/my-app`). If no location is provided, the current directory is used.

```bash
vodka create my-app --minimal
vodka create my-app /path/to/dir --minimal
```

This generates:

- A Go backend powered by Vodka
- A React + Vite frontend
- Preconfigured development workflow
- SPA support for production deployments


---

## Generated Project Structure

```text
my-app/
├── controllers/
├── routes/
├── frontend/
├── main.go
├── go.mod
└── vodka.config.json
```

---

## Install Frontend Dependencies

After scaffolding the project, install frontend dependencies manually:

```bash
cd my-app

cd frontend
npm install

cd ..
```

---

## Start Development Mode

```bash
vodka run dev
```

This starts:

- Vite frontend dev server
- Vodka backend server
- Concurrent frontend/backend workflow

---

# Generated Backend Structure

Vodka organizes backend code into simple and clean layers.

```text
backend/
├── controllers/
│   └── ping.go
├── routes/
│   └── routes.go
├── main.go
└── go.mod
```

---

## controllers/

Contains request handlers and business logic.

```go
func Pong(c *vodka.Context) {
	c.String(200, "Pong!")
}
```

---

## routes/

Registers API routes.

```go
app.GET("/ping", controllers.Pong)
```

---

## main.go

Bootstraps the Vodka engine and middleware configuration.

```go
app := vodka.DefaultRouter()

routes.Setup(app)

app.Run(":8080")
```

---

# Minimal API Example

```go
package main

import (
	"log"

	"github.com/DevanshuTripathi/vodka"
)

func main() {
	app := vodka.DefaultRouter()

	app.GET("/ping", func(c *vodka.Context) {
		c.JSON(200, vodka.M{
			"message": "pong!",
		})
	})

	if err := app.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
```

---

## Test the API

```bash
curl http://localhost:8080/ping
```

Response:

```json
{
  "message": "pong!"
}
```

---

# Using Vodka for APIs

## Create a Go Module

```bash
mkdir backend-app

cd backend-app

go mod init app

go get github.com/DevanshuTripathi/vodka
```

---

## Run Backend Server

```bash
vodka
```
Vodka automatically:

- Watches `.go` files
- Rebuilds your backend
- Restarts the server instantly

---

# Core Concepts

## Engine

`vodka.Engine` is the central router and application instance.

It uses a Radix Tree-based routing architecture for fast request matching and low overhead.

```go
app := vodka.DefaultRouter()
```

---

## Context

`vodka.Context` wraps Go's `http.Request` and `http.ResponseWriter` into a clean and ergonomic API.

---

### JSON Response

```go
c.JSON(200, vodka.M{
	"message": "hello",
})
```

---

### Query Parameters

```go
name := c.Query("name")
```

---

### URL Parameters

```go
id := c.Param("id")
```

---

### Bind JSON

```go
var user User
c.BindJSON(&user)
```

---

### Error Handling

```go
c.Error(400, errors.New("invalid request"))
```

---

# Graceful Shutdown & Lifecycle Manager

Vodka features a production-ready application lifecycle management system that handles:
- **Startup Hooks**: Run code before the HTTP server starts.
- **Shutdown Hooks**: Run code sequentially when the application receives termination signals (`SIGINT`, `SIGTERM`).
- **Priority-Based Execution**: Control the shutdown sequence of your resources.
- **Graceful HTTP Server Shutdown**: Stop accepting new connections and finish active requests.
- **Configurable Shutdown Timeout**: Cancel remaining work if the timeout expires.
- **Error Aggregation**: Collect and report errors from all hooks.

---

## Startup Hooks

Startup hooks allow executing initialization code (e.g., establishing database connections, seeding data) before the server begins serving requests. If any startup hook returns an error, server startup is immediately aborted.

```go
app.OnStart(func() error {
    return initializeDatabase()
})
```

---

## Shutdown Hooks & Priority

Shutdown hooks are executed sequentially when a termination signal is received. You can register hooks with a priority value:

- Higher priority hooks execute first.
- If priorities are equal, hooks execute in registration order.
- The default priority is `0` when calling `OnShutdown`.

```go
// Priority 100: runs first
app.OnShutdownWithPriority(100, func(ctx context.Context) error {
    return closeDatabase()
})

// Priority 50: runs second
app.OnShutdownWithPriority(50, func(ctx context.Context) error {
    return stopWorkers()
})

// Default priority (0): runs last
app.OnShutdown(func(ctx context.Context) error {
    return cleanupTempFiles()
})
```

---

## Timeout Configuration

By default, Vodka allows up to `30 seconds` for the entire shutdown sequence (including finishing active HTTP requests and running all shutdown hooks). You can customize this timeout:

```go
app.SetShutdownTimeout(45 * time.Second)
```

---

## Production Example

Here is a full production-ready example demonstrating database cleanup, worker queue termination, and graceful server shutdown:

```go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/DevanshuTripathi/vodka"
	_ "github.com/lib/pq"
)

func main() {
	app := vodka.DefaultRouter()

	var db *sql.DB

	// 1. Register Startup Hook to initialize database
	app.OnStart(func() error {
		var err error
		db, err = sql.Open("postgres", "postgres://user:pass@localhost/db?sslmode=disable")
		if err != nil {
			return fmt.Errorf("failed to open database: %w", err)
		}
		
		// Verify connection
		if err := db.Ping(); err != nil {
			return fmt.Errorf("failed to ping database: %w", err)
		}
		log.Println("Database connection established")
		return nil
	})

	// 2. Start workers
	workerCtx, cancelWorkers := context.WithCancel(context.Background())
	app.OnStart(func() error {
		go runBackgroundWorkers(workerCtx)
		log.Println("Background workers started")
		return nil
	})

	// 3. Configure Shutdown Timeout
	app.SetShutdownTimeout(15 * time.Second)

	// 4. Register Shutdown Hooks in priority order
	
	// Stop background workers first (Priority 100)
	app.OnShutdownWithPriority(100, func(ctx context.Context) error {
		log.Println("Stopping background workers...")
		cancelWorkers()
		return nil
	})

	// Close database connection (Priority 50)
	app.OnShutdownWithPriority(50, func(ctx context.Context) error {
		log.Println("Closing database connection...")
		if db != nil {
			return db.Close()
		}
		return nil
	})

	// Simple route
	app.GET("/", func(c *vodka.Context) {
		c.String(200, "Hello, Graceful Vodka!")
	})

	// 5. Start the server (handles SIGINT/SIGTERM automatically)
	if err := app.Run(":8080"); err != nil {
		log.Fatalf("Server stopped with error: %v", err)
	}
}

func runBackgroundWorkers(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Workers stopped")
			return
		default:
			// Perform background work
			time.Sleep(1 * time.Second)
		}
	}
}
```

---

# Middleware

Vodka middleware is simply a `vodka.HandlerFunc`.

```go
func(*vodka.Context)
```

Middlewares can:

- Modify requests
- Attach values to context
- Authenticate users
- Log requests
- Recover from panics
- Handle errors

Middlewares are registered using:

```go
app.Use(...)
```

or:

```go
group.Use(...)
```

---

## How Middleware Works

Each incoming request is wrapped in a `*vodka.Context` and the framework builds a handler chain.

The chain includes:

- Group middlewares
- Route middlewares
- Final route handler

`c.Next()` continues execution to the next middleware or handler.

If you omit `c.Next()`, the request chain stops immediately.

---

## Custom Middleware Example

```go
func RequestTimer() vodka.HandlerFunc {
	return func(c *vodka.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)

		log.Printf(
			"[RequestTimer] %s %s %v",
			c.Request.Method,
			c.Request.URL.Path,
			latency,
		)
	}
}

func main() {
	app := vodka.DefaultRouter()

	app.Use(
		vodka.Logger(),
		vodka.Recovery(),
		vodka.ErrorHandler(),
	)

	app.Use(RequestTimer())

	api := app.Group("/api", AdminOnly())

	api.GET("/dashboard", func(c *vodka.Context) {
		userRole, _ := c.Get("user_role")

		c.JSON(200, vodka.M{
			"role": userRole,
		})
	})
}
```

---

## Request ID Middleware

Track requests across your logs using automatic request ID generation. Every request gets a unique ID that's automatically added to response headers and stored in the context.

```go
app := vodka.DefaultRouter()

// Add request ID middleware — generates unique ID for each request
app.Use(mixers.RequestID())

app.GET("/api/users", func(c *vodka.Context) {
	requestID, _ := c.Get("request-id")
	log.Printf("[%v] Fetching users", requestID)
	c.JSON(200, vodka.M{"users": []string{"Alice", "Bob"}})
})

// Client receives: X-Request-ID: 550e8400-e29b-41d4-a716-446655440000
```

### Custom Header Name

Use a different header name if needed (context key is always `"request-id"`):

```go
app.Use(mixers.RequestIDWithHeader("X-Correlation-ID"))
```

---
# Template Rendering

Vodka supports Go's native `html/template` package for server-side HTML rendering.

## Basic Usage

```go
app := vodka.DefaultRouter()

// Load all templates from templates folder
app.LoadHTMLGlob("templates/*.html")

app.GET("/user", func(c *vodka.Context) {
    c.HTML(200, "user.html", vodka.M{
        "name": "John Doe",
        "email": "john@example.com",
    })
})

## Template Files

Create a templates/user.html file:

```html
<!DOCTYPE html>
<html>
<head>
    <title>User Profile</title>
</head>
<body>
    <h1>Hello, {{.name}}!</h1>
    <p>Email: {{.email}}</p>
</body>
</html>  ```

## Methods
LoadHTMLGlob(pattern string) error - Load templates using glob pattern

LoadHTMLFiles(files ...string) error - Load specific template files

c.HTML(code int, name string, data interface{}) - Render HTML template

Templates work seamlessly with Vodka's hot-reload in development mode.


# Validation

Vodka supports request validation using struct tags.

```go
type User struct {
	Email    string `validate:"required,email"`
	Password string `validate:"min=8"`
}
```

---

# Authentication

Vodka includes built-in Bearer authentication middleware and JWT validation helpers.

You can also provide custom validation logic:

```go
app.Use(vodka.BearerAuth(contextKey, func(token string) (any, bool) {
	return validateToken(token)
}))
```

---

# Full-Stack SPA Support

Projects generated using `vodka create` come with SPA serving preconfigured.

```go
app.ServeSPA("./frontend/dist")
```

If a route does not match an API endpoint, Vodka automatically serves the frontend application.

This enables seamless React Router support in production.

---

# Performance

Vodka uses a Radix Tree router optimized for:

- Fast route matching
- Low memory overhead
- Efficient middleware execution

---

# Production Build

## Build Frontend

```bash
cd frontend
npm run build
```

---

## Build Backend

```bash
vodka
```

---

# Philosophy

Vodka is designed around a few core ideas:

- Fast development workflow
- Minimal boilerplate
- Strong developer experience
- Clean APIs
- Modern full-stack integration
- Practical defaults

---

# Roadmap

- More middleware packages
- Enhanced validation system
- Improved benchmarking suite
- Expanded CLI tooling
- Testing utilities

---

# Contributing

Contributions, issues, and feature requests are welcome.

If you find bugs or have suggestions, feel free to open an issue or submit a pull request.

---

# License

MIT License
