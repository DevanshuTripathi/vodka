export default function MiddlewareSection() {
  return (
    <section
      id="middleware"
      className="border border-slate-300 rounded-xl p-8 my-8 bg-white shadow-sm"
    >
      <h2 className="text-2xl font-bold mb-6">Middleware</h2>

      <p className="text-slate-600 mb-6">
        Middleware allows you to run logic before and after request handlers.
      </p>

      <div className="bg-slate-950 text-white rounded-lg p-4 overflow-x-auto">
        <pre>{`app.Use(vodka.Logger())
app.Use(vodka.Recovery())

app.GET("/", handler)`}</pre>
      </div>

      <div className="mt-8">
        <h3 className="text-lg font-semibold mb-2">Custom Middleware</h3>

        <div className="bg-slate-950 text-white rounded-lg p-4 overflow-x-auto">
          <pre>{`func Logger() vodka.HandlerFunc {
  return func(c *vodka.Context) {
    c.Next()
  }
}`}</pre>
        </div>
      </div>
    </section>
  );
}
