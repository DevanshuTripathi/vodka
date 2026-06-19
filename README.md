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


Contributions, issues, and feature requests are welcome.

If you find bugs or have suggestions, feel free to open an issue or submit a pull request.

---

# Our Contributors

<!-- CONTRIBUTORS:START -->

<a href="https://github.com/DevanshuTripathi">
  <img src="https://avatars.githubusercontent.com/u/74350750?v=4" width="80px;" alt="DevanshuTripathi"/><br/>
  <sub><b>DevanshuTripathi</b></sub>
</a>


<a href="https://github.com/hariom888">
  <img src="https://avatars.githubusercontent.com/u/188325558?v=4" width="80px;" alt="hariom888"/><br/>
  <sub><b>hariom888</b></sub>
</a>


<a href="https://github.com/Sounak-star">
  <img src="https://avatars.githubusercontent.com/u/180566751?v=4" width="80px;" alt="Sounak-star"/><br/>
  <sub><b>Sounak-star</b></sub>
</a>


<a href="https://github.com/Utkarsh0uchiha">
  <img src="https://avatars.githubusercontent.com/u/117010422?v=4" width="80px;" alt="Utkarsh0uchiha"/><br/>
  <sub><b>Utkarsh0uchiha</b></sub>
</a>


<a href="https://github.com/Srejoye">
  <img src="https://avatars.githubusercontent.com/u/177326090?v=4" width="80px;" alt="Srejoye"/><br/>
  <sub><b>Srejoye</b></sub>
</a>


<a href="https://github.com/anshggss">
  <img src="https://avatars.githubusercontent.com/u/144388813?v=4" width="80px;" alt="anshggss"/><br/>
  <sub><b>anshggss</b></sub>
</a>


<a href="https://github.com/jk1507">
  <img src="https://avatars.githubusercontent.com/u/187723724?v=4" width="80px;" alt="jk1507"/><br/>
  <sub><b>jk1507</b></sub>
</a>


<a href="https://github.com/Mritunjay2005">
  <img src="https://avatars.githubusercontent.com/u/138350127?v=4" width="80px;" alt="Mritunjay2005"/><br/>
  <sub><b>Mritunjay2005</b></sub>
</a>


<a href="https://github.com/grishabhatia">
  <img src="https://avatars.githubusercontent.com/u/178971480?v=4" width="80px;" alt="grishabhatia"/><br/>
  <sub><b>grishabhatia</b></sub>
</a>


<a href="https://github.com/DevdharManpuria">
  <img src="https://avatars.githubusercontent.com/u/175636139?v=4" width="80px;" alt="DevdharManpuria"/><br/>
  <sub><b>DevdharManpuria</b></sub>
</a>


<a href="https://github.com/Anand-240">
  <img src="https://avatars.githubusercontent.com/u/188213371?v=4" width="80px;" alt="Anand-240"/><br/>
  <sub><b>Anand-240</b></sub>
</a>


<a href="https://github.com/ghosthouse7">
  <img src="https://avatars.githubusercontent.com/u/229156343?v=4" width="80px;" alt="ghosthouse7"/><br/>
  <sub><b>ghosthouse7</b></sub>
</a>


<a href="https://github.com/i-OmSharma">
  <img src="https://avatars.githubusercontent.com/u/116426515?v=4" width="80px;" alt="i-OmSharma"/><br/>
  <sub><b>i-OmSharma</b></sub>
</a>


<a href="https://github.com/rajsingh1301">
  <img src="https://avatars.githubusercontent.com/u/195672131?v=4" width="80px;" alt="rajsingh1301"/><br/>
  <sub><b>rajsingh1301</b></sub>
</a>


<a href="https://github.com/Aish-kul16">
  <img src="https://avatars.githubusercontent.com/u/227191125?v=4" width="80px;" alt="Aish-kul16"/><br/>
  <sub><b>Aish-kul16</b></sub>
</a>


<a href="https://github.com/bhavyanjain3004">
  <img src="https://avatars.githubusercontent.com/u/127675863?v=4" width="80px;" alt="bhavyanjain3004"/><br/>
  <sub><b>bhavyanjain3004</b></sub>
</a>


<a href="https://github.com/EchoOfCode">
  <img src="https://avatars.githubusercontent.com/u/210523620?v=4" width="80px;" alt="EchoOfCode"/><br/>
  <sub><b>EchoOfCode</b></sub>
</a>


<a href="https://github.com/OdaloV">
  <img src="https://avatars.githubusercontent.com/u/147773886?v=4" width="80px;" alt="OdaloV"/><br/>
  <sub><b>OdaloV</b></sub>
</a>


<a href="https://github.com/dimple9079">
  <img src="https://avatars.githubusercontent.com/u/219795540?v=4" width="80px;" alt="dimple9079"/><br/>
  <sub><b>dimple9079</b></sub>
</a>


<a href="https://github.com/dynamo-pentester">
  <img src="https://avatars.githubusercontent.com/u/145715786?v=4" width="80px;" alt="dynamo-pentester"/><br/>
  <sub><b>dynamo-pentester</b></sub>
</a>


<a href="https://github.com/mzl2233">
  <img src="https://avatars.githubusercontent.com/u/109468061?v=4" width="80px;" alt="mzl2233"/><br/>
  <sub><b>mzl2233</b></sub>
</a>

<!-- CONTRIBUTORS:END -->

---
# License

MIT License
