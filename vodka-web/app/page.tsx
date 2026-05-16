import Link from "next/link";

const features = [
  { icon: "⚡", title: "Fast Routing", desc: "Radix-tree based router via httprouter for zero-allocation route matching." },
  { icon: "🔗", title: "Middleware Chain", desc: "Composable middleware with abort support. Use global or group-scoped middleware." },
  { icon: "🔌", title: "WebSocket Support", desc: "First-class WS support with WSContext, lifecycle handling, and origin whitelisting." },
  { icon: "📡", title: "SSE Support", desc: "Server-Sent Events with SSEContext for real-time one-way data streaming." },
  { icon: "🔒", title: "JWT & Auth", desc: "Built-in JWT utilities and bearer token authentication via mixers." },
  { icon: "🛡️", title: "Rate Limiting", desc: "Per-IP token bucket rate limiting out of the box." },
];

export default function Home() {
  return (
    <div className="flex flex-col gap-24">
      {/* Hero */}
      <section className="flex flex-col items-center text-center gap-6 pt-16">
        <div className="text-6xl">🍸</div>
        <h1 className="text-5xl font-bold tracking-tight text-white">
          Vodka
        </h1>
        <p className="text-xl text-zinc-400 max-w-xl">
          A modern, fast, and developer-friendly Go web framework.
          Build REST APIs, WebSockets, and SSE streams with clean, minimal code.
        </p>
        <div className="flex gap-4 mt-2">
          <Link
            href="/docs"
            className="px-6 py-3 bg-white text-zinc-950 rounded-lg font-semibold hover:bg-zinc-200 transition-colors"
          >
            Get Started
          </Link>
          <a
            href="https://github.com/DevanshuTripathi/vodka"
            target="_blank"
            rel="noopener noreferrer"
            className="px-6 py-3 border border-zinc-700 rounded-lg font-semibold text-zinc-300 hover:border-zinc-500 hover:text-white transition-colors"
          >
            GitHub →
          </a>
        </div>

        {/* Quick install */}
        <div className="mt-4 bg-zinc-900 border border-zinc-800 rounded-lg px-6 py-3 font-mono text-sm text-zinc-300">
          go get github.com/DevanshuTripathi/vodka
        </div>
      </section>

      {/* Code snippet */}
      <section className="bg-zinc-900 border border-zinc-800 rounded-xl p-6 font-mono text-sm leading-7">
        <pre className="text-zinc-300 overflow-x-auto">{`package main

import "github.com/DevanshuTripathi/vodka"

func main() {
    app := vodka.DefaultRouter()

    app.GET("/ping", func(c *vodka.Context) {
        c.JSON(200, vodka.M{"message": "pong"})
    })

    app.WS("/ws", func(c *vodka.WSContext) {
        for {
            t, msg, err := c.ReadMessage()
            if err != nil { return }
            c.WriteMessage(t, msg)
        }
    })

    app.Run(":8080")
}`}</pre>
      </section>

      {/* Features */}
      <section>
        <h2 className="text-2xl font-bold text-white mb-8 text-center">Everything you need</h2>
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
          {features.map((f) => (
            <div key={f.title} className="bg-zinc-900 border border-zinc-800 rounded-xl p-5 flex flex-col gap-2">
              <span className="text-2xl">{f.icon}</span>
              <h3 className="font-semibold text-white">{f.title}</h3>
              <p className="text-sm text-zinc-400">{f.desc}</p>
            </div>
          ))}
        </div>
      </section>
    </div>
  );
}
