markdown

\# Vodka



!\[Go Version](https://img.shields.io/badge/Go-1.24+-blue)

!\[License](https://img.shields.io/badge/license-MIT-green)

!\[Status](https://img.shields.io/badge/status-active-success)



```text

\_\_/\\\\\\\_\_\_\_\_\_\_\_/\\\\\\\_\_\_\_\_\_\_/\\\\\\\\\\\_\_\_\_\_\_\_/\\\\\\\\\\\\\\\\\\\\\\\\\_\_\_\_\_/\\\\\\\_\_\_\_\_\_\_\_/\\\\\\\_\_\_\_\_/\\\\\\\\\\\\\\\\\\\_\_\_\_        

&#x20;\_\\/\\\\\\\_\_\_\_\_\_\_\\/\\\\\\\_\_\_\_\_/\\\\\\///\\\\\\\_\_\_\_\\/\\\\\\////////\\\\\\\_\_\\/\\\\\\\_\_\_\_\_/\\\\\\//\_\_\_\_/\\\\\\\\\\\\\\\\\\\\\\\\\\\_\_       

&#x20; \_\\//\\\\\\\_\_\_\_\_\_/\\\\\\\_\_\_\_/\\\\\\/\_\_\\///\\\\\\\_\_\\/\\\\\\\_\_\_\_\_\_\\//\\\\\\\_\\/\\\\\\\_\_/\\\\\\//\_\_\_\_\_\_/\\\\\\/////////\\\\\\\_      

&#x20;  \_\_\\//\\\\\\\_\_\_\_/\\\\\\\_\_\_\_/\\\\\\\_\_\_\_\_\_\\//\\\\\\\_\\/\\\\\\\_\_\_\_\_\_\_\\/\\\\\\\_\\/\\\\\\\\\\\\//\\\\\\\_\_\_\_\_\\/\\\\\\\_\_\_\_\_\_\_\\/\\\\\\\_     

&#x20;   \_\_\_\\//\\\\\\\_\_/\\\\\\\_\_\_\_\\/\\\\\\\_\_\_\_\_\_\_\\/\\\\\\\_\\/\\\\\\\_\_\_\_\_\_\_\\/\\\\\\\_\\/\\\\\\//\_\\//\\\\\\\_\_\_\_\\/\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\_    

&#x20;    \_\_\_\_\\//\\\\\\/\\\\\\\_\_\_\_\_\\//\\\\\\\_\_\_\_\_\_/\\\\\\\_\_\\/\\\\\\\_\_\_\_\_\_\_\\/\\\\\\\_\\/\\\\\\\_\_\_\_\\//\\\\\\\_\_\_\\/\\\\\\/////////\\\\\\\_   

&#x20;     \_\_\_\_\_\\//\\\\\\\\\\\_\_\_\_\_\_\_\\///\\\\\\\_\_/\\\\\\\_\_\_\_\\/\\\\\\\_\_\_\_\_\_\_/\\\\\\\_\_\\/\\\\\\\_\_\_\_\_\\//\\\\\\\_\_\\/\\\\\\\_\_\_\_\_\_\_\\/\\\\\\\_  

&#x20;      \_\_\_\_\_\_\\//\\\\\\\_\_\_\_\_\_\_\_\_\_\\///\\\\\\\\\\/\_\_\_\_\_\\/\\\\\\\\\\\\\\\\\\\\\\\\/\_\_\_\\/\\\\\\\_\_\_\_\_\_\\//\\\\\\\_\\/\\\\\\\_\_\_\_\_\_\_\\/\\\\\\\_ 

&#x20;       \_\_\_\_\_\_\_\\///\_\_\_\_\_\_\_\_\_\_\_\_\_\\/////\_\_\_\_\_\_\_\\////////////\_\_\_\_\_\\///\_\_\_\_\_\_\_\_\\///\_\_\\///\_\_\_\_\_\_\_\_\\///\_\_

A modern Go web framework focused on developer experience, full-stack workflow, and rapid iteration.

Vodka is a lightweight and high-performance HTTP framework for Go that combines clean routing, middleware chaining, authentication utilities, validation support, and a powerful CLI for building modern full-stack applications.



Table of Contents

Features



Why Vodka



Installation



Quick Start



Core Concepts



Middleware



Template Rendering



Validation



Authentication



License



Features

Feature	Included

Radix Tree Routing	✅

Middleware Chaining	✅

Route Groups	✅

JSON Binding	✅

Request Validation	✅

JWT Validation Helpers	✅

Bearer Auth Middleware	✅

Vite + React Scaffolding	✅

SPA Serving	✅

Panic Recovery Middleware	✅

Logger Middleware	✅

CORS Middleware	✅

Context Storage	✅

HTML Template Rendering	✅

Why Vodka?

Vodka combines:



⚡ Fast backend iteration



⚛️ React + Vite integration



🧩 Lightweight routing



🔐 Authentication helpers



✅ Validation support



🚀 Developer-first workflow



Installation

bash

go install github.com/DevanshuTripathi/vodka/cmd/vodka@latest

Quick Start

bash

vodka create my-app

cd my-app

vodka run dev

Core Concepts

go

app := vodka.DefaultRouter()



app.GET("/user/:id", func(c \*vodka.Context) {

&#x20;   c.JSON(200, vodka.M{

&#x20;       "id": c.Param("id"),

&#x20;   })

})

Middleware

go

app.Use(vodka.Logger())

app.Use(vodka.Recovery())

Template Rendering

go

app := vodka.DefaultRouter()

app.LoadHTMLGlob("templates/\*.html")



app.GET("/user", func(c \*vodka.Context) {

&#x20;   c.HTML(200, "user.html", vodka.M{

&#x20;       "name": "John Doe",

&#x20;   })

})

Create templates/user.html:



html

<h1>Hello, {{.name}}!</h1>

Validation

go

type User struct {

&#x20;   Email string `validate:"required,email"`

}

Authentication

go

app.Use(vodka.BearerAuth("user", validateToken))

License

MIT License

