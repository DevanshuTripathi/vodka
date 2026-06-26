export default function SPASupportSection() {
  return (
    <section
      id="spa-support"
      className="bg-white border border-slate-200 rounded-lg p-8 my-8"
    >
      <h2 className="text-xl font-bold mb-6">SPA Support</h2>

      <p className="text-slate-600 mb-6">
        Serve Single Page Applications directly from your Vodka backend.
      </p>

      <pre className="bg-slate-950 text-white rounded-md p-4 overflow-x-auto mb-6">
        {`app.Static("/", "./frontend/dist")

app.NoRoute(func(c *vodka.Context) {
  c.File("./frontend/dist/index.html")
})`}
      </pre>

      <p className="text-slate-600">
        This enables React, Vue, Svelte, and other SPA frameworks to work
        seamlessly with backend routes.
      </p>
    </section>
  );
}
