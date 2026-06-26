export default function TemplateRenderingSection() {
  return (
    <section
      id="template-rendering"
      className="border border-slate-200 rounded-xl p-8 my-8 bg-white"
    >
      <h2 className="text-2xl font-bold mb-6">
        Template Rendering
      </h2>

      <p className="text-slate-600 mb-6">
        Vodka supports Go's native html/template package for server-side
        rendering.
      </p>

      <div className="bg-slate-950 text-white rounded-lg p-4 overflow-x-auto">
        <pre>{`app := vodka.DefaultRouter()

app.LoadHTMLGlob("templates/*.html")

app.GET("/user", func(c *vodka.Context) {
    c.HTML(200, "user.html", vodka.M{
        "name": "John Doe",
        "email": "john@example.com",
    })
})`}</pre>
      </div>

      <div className="mt-6 border border-slate-200 rounded-lg p-4 bg-slate-50">
        <h3 className="text-lg font-semibold mb-3">
          Template Features
        </h3>

        <ul className="list-disc pl-5 text-slate-600 space-y-2">
          <li>Load templates using glob patterns</li>
          <li>Render dynamic data</li>
          <li>Works with Go's html/template</li>
          <li>Hot reload during development</li>
        </ul>
      </div>
    </section>
  );
}