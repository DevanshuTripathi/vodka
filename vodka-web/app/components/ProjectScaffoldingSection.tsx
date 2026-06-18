export default function ProjectScaffoldingSection() {
  return (
    <section
      id="project-scaffolding"
      className="border border-slate-300 rounded-xl p-8 my-8 bg-white shadow-sm"
    >
      <h2 className="text-2xl font-bold mb-6">Project Scaffolding</h2>

      <p className="text-slate-600 mb-6">
        Vodka can generate full-stack applications with frontend and backend
        configured out of the box.
      </p>

      <div className="bg-slate-950 text-white rounded-lg p-4 overflow-x-auto">
        <pre>{`vodka new my-app --template react`}</pre>
      </div>

      <div className="mt-6 border border-slate-200 rounded-lg p-4 bg-slate-50">
        <h3 className="text-lg font-semibold mb-3">Generated Structure</h3>
        <pre className="text-sm text-slate-700 overflow-x-auto">
          {`my-app/
├── controllers/
├── routes/
├── frontend/
├── main.go
├── go.mod
└── vodka.config.json`}
        </pre>
      </div>
    </section>
  );
}
