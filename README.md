# Vodka

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

# Vodka

**A modern Go web framework focused on developer experience, full-stack workflow, and fast iteration.**

Vodka is a lightweight, high-performance HTTP framework for Go that combines clean routing, middleware chaining, validation, authentication utilities, and a powerful CLI for building modern full-stack applications.

Unlike traditional Go frameworks that focus only on request handling, Vodka focuses heavily on developer experience:

- ⚡ Fast routing with Radix Tree architecture
- 🔥 Built-in hot reload for Go backends
- ⚛️ Vite + React full-stack scaffolding
- 🧩 Middleware chaining system
- 🔐 Authentication helpers and JWT validation
- ✅ Request validation support
- 🛠️ Clean and ergonomic API design
- 🚀 One-command development workflow

---

# Installation

## Install the Vodka CLI

```bash
go install github.com/DevanshuTripathi/vodka/cmd/vodka@latest
```

Make sure your Go bin directory is in your system PATH.

### Linux / macOS

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### Windows

Add the following directory to Environment Variables:

```text
%USERPROFILE%\go\bin
```

---

# Quick Start

## Create a Full-Stack App

```bash
vodka create my-app
```

This generates:

- A Go backend powered by Vodka
- A Vite + React frontend
- Preconfigured development workflow
- SPA serving support for production deployments

---

## Install Frontend Dependencies

```bash
cd my-app
cd frontend && npm install
cd ..
```

---

## Start Development Mode

```bash
vodka run dev
```

This starts:

- The Vite frontend dev server
- The Vodka backend
- Automatic Go hot reload
- Concurrent frontend/backend workflow

Edit `.jsx` files → frontend updates instantly.

Edit `.go` files → backend rebuilds automatically.

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

# Using Vodka for APIs

## Create a Go Module

```bash
mkdir backend-app
cd backend-app
go mod init app
go get github.com/DevanshuTripathi/vodka
```

---

## Run with Hot Reload

Inside any Go project using Vodka:

```bash
vodka
```

Vodka automatically:

- Watches `.go` files
- Rebuilds your backend
- Restarts the server instantly

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
| Hot Reload | ✅ |
| Vite + React Scaffolding | ✅ |
| SPA Serving | ✅ |
| Panic Recovery Middleware | ✅ |
| Logger Middleware | ✅ |
| CORS Middleware | ✅ |
| Context Storage | ✅ |

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

### JSON Response

```go
c.JSON(200, vodka.M{
    "message": "hello",
})
```

### Query Parameters

```go
name := c.Query("name")
```

### URL Parameters

```go
id := c.Param("id")
```

### Bind JSON

```go
var user User
c.BindJSON(&user)
```

### Error Handling

```go
c.Error(400, errors.New("invalid request"))
```

---

# Middleware

Vodka uses a middleware chain architecture similar to modern backend frameworks.

Middlewares can:

- Modify requests
- Attach values to context
- Handle authentication
- Log requests
- Recover from panics
- Intercept errors

---

## Custom Middleware Example

```go
func Logger() vodka.HandlerFunc {
	return func(c *vodka.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)

		log.Printf(
			"[%s] %s %v",
			c.Request.Method,
			c.Request.URL.Path,
			latency, 
		)
	}
}

app.Use(Logger())
```

---

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

If a route does not match an API endpoint, Vodka automatically serves the frontend application, enabling seamless React Router support in production.

---

# Philosophy

Vodka is designed around a few core ideas:

- Fast development workflow
- Minimal boilerplate
- Strong developer experience
- Clean APIs
- Modern full-stack integration
- Practical defaults without excessive abstraction

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
