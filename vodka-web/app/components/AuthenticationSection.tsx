export default function AuthenticationSection() {
  return (
    <section
      id="authentication"
      className="bg-white border border-slate-200 rounded-lg p-8 my-8"
    >
      <h2 className="text-xl font-bold mb-6">Authentication</h2>

      <p className="text-slate-600 mb-6">
        Protect routes and verify user access using middleware.
      </p>

      <pre className="bg-slate-950 text-white rounded-md p-4 overflow-x-auto mb-6">
        {`app.Use(AuthMiddleware())

app.GET("/dashboard", DashboardHandler)`}
      </pre>

      <p className="text-slate-600">
        Authentication middleware can validate JWTs, sessions, API keys, or
        custom credentials before allowing access.
      </p>
    </section>
  );
}
