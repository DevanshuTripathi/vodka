export default function QuickStartSection() {
  return (
    <section id="quick-start"className="border border-slate-300 rounded-xl p-8 my-8 bg-white shadow-sm">
      <h2 className="text-2xl font-bold mb-6">Quick Start</h2>

      <p className="text-slate-600 mb-6">
        Create a new Vodka project and start developing immediately.
      </p>

      <div className="bg-slate-950 text-white rounded-lg p-4 overflow-x-auto">
        <pre>{`vodka new my-app

cd my-app

vodka dev`}</pre>
      </div>

      <p className="text-slate-600 mt-6">
        Your application will be available at:
      </p>

      <div className="bg-slate-100 border border-slate-300 rounded-lg p-3 mt-3 font-mono">
        http://localhost:8080
      </div>
    </section>
  );
}
