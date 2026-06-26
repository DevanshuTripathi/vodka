export default function MinimalAPISection() {
  return (
    <section
      id="minimal-api"
      className="border border-slate-200 rounded-xl p-8 my-8 bg-white"
    >
      <h2 className="text-2xl font-bold mb-6">Minimal API Example</h2>

      <p className="text-slate-600 mb-6">
        Create a simple API endpoint with Vodka in just a few lines.
      </p>

      <div className="bg-slate-950 text-white rounded-lg p-4 overflow-x-auto">
        <pre>{`package main

import (
    "github.com/DevanshuTripathi/vodka"
)

func main() {
    app := vodka.DefaultRouter()

    app.GET("/ping", func(c *vodka.Context) {
        c.JSON(200, vodka.M{
            "message": "pong!",
        })
    })

    app.Run(":8080")
}`}</pre>
      </div>

      <div className="mt-6">
        <h3 className="text-lg font-semibold mb-3">Test the API</h3>

        <div className="bg-slate-950 text-white rounded-lg p-4 overflow-x-auto">
          <pre>{`curl http://localhost:8080/ping`}</pre>
        </div>
      </div>

      <div className="mt-6">
        <h3 className="text-lg font-semibold mb-3">Response</h3>

        <div className="bg-slate-950 text-white rounded-lg p-4 overflow-x-auto">
          <pre>{`{
  "message": "pong!"
}`}</pre>
        </div>
      </div>
    </section>
  );
}
