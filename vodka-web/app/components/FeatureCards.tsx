export default function FeatureCards() {
  const features = [
    "Radix Tree Routing",
    "Middleware Chaining",
    "Route Groups",
    "JSON Binding",
    "Request Validation",
    "JWT Validation Helpers",
    "Bearer Auth Middleware",
    "Vite + React Scaffolding",
    "SPA Serving",
    "Panic Recovery",
    "Logger Middleware",
    "CORS Middleware",
    "Context Storage",
    "HTML Template Rendering",
  ];

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 gap-4">
      {features.map((feature) => (
        <div
          key={feature}
          className="group rounded-xl border border-slate-200 bg-slate-50 p-4 hover:bg-white hover:border-slate-300 transition-all duration-200 min-h-[80px]"
        >
          <div className="flex items-center gap-3">
            {" "}
            <span className="text-green-500 font-semibold text-base">✓</span>
            <span className="font-medium text-slate-900">{feature}</span>{" "}
          </div>
        </div>
      ))}
    </div>
  );
}
