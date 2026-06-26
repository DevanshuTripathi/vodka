export default function RequestIDSection() {
  return (
    <section
      id="request-id"
      className="border border-slate-200 rounded-xl p-8 my-8 bg-white"
    >
      <h2 className="text-2xl font-bold mb-6">Request ID Middleware</h2>

      <p className="text-slate-600 mb-6">
        Generate a unique ID for every request and track it across logs and
        services.
      </p>

      <div className="bg-slate-950 text-white rounded-lg p-4 overflow-x-auto">
        <pre>{`app := vodka.DefaultRouter()

app.Use(mixers.RequestID())

app.GET("/api/users", func(c *vodka.Context) {
    requestID, _ := c.Get("request-id")

    c.JSON(200, vodka.M{
        "request_id": requestID,
    })
})`}</pre>
      </div>

      <div className="mt-6 border border-slate-200 rounded-lg p-4 bg-slate-50">
        <p className="font-medium mb-2">Benefits</p>

        <ul className="list-disc pl-5 text-slate-600 space-y-1">
          <li>Track requests across services</li>
          <li>Improve debugging</li>
          <li>Correlate logs easily</li>
        </ul>
      </div>
    </section>
  );
}
