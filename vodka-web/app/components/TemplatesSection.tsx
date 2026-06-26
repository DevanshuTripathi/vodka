export default function TemplatesSection() {
  return (
    <section
      id="templates"
      className="bg-white border border-slate-200 rounded-lg p-8 my-8"
    >
      <h2 className="text-xl font-bold mb-6">Templates</h2>

      <p className="text-slate-600 mb-6">
        Start quickly with preconfigured project templates.
      </p>

      <pre className="bg-slate-950 text-white rounded-md p-4 overflow-x-auto mb-6">
        {`vodka new my-app --template react

vodka new my-app --template vue

vodka new my-app --template svelte`}
      </pre>

      <p className="text-slate-600">
        Templates include frontend tooling, routing, build configuration, and
        development workflows.
      </p>
    </section>
  );
}