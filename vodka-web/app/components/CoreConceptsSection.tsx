export default function CoreConceptsSection() {
  return (
    <section
      id="core-concepts"
      className="border border-slate-300 rounded-xl p-8 my-8 bg-white shadow-sm"
    >
      <h2 className="text-2xl font-bold mb-6">Core Concepts</h2>

      <div className="space-y-8">
        <div>
          <h3 className="text-lg font-semibold mb-2">Engine</h3>

          <p className="text-slate-600 mb-4">
            The Engine is the central router and application instance that
            handles requests and middleware.
          </p>

          <div className="bg-slate-950 text-white rounded-lg p-4 overflow-x-auto">
            <pre>{`app := vodka.DefaultRouter()`}</pre>
          </div>
        </div>

        <div>
          <h3 className="text-lg font-semibold mb-2">Context</h3>

          <p className="text-slate-600 mb-4">
            Context provides helpers for requests, responses, JSON handling,
            query parameters, and middleware communication.
          </p>

          <div className="bg-slate-950 text-white rounded-lg p-4 overflow-x-auto">
            <pre>{`c.JSON(200, vodka.M{
  "message": "hello",
})`}</pre>
          </div>
        </div>
      </div>
    </section>
  );
}
