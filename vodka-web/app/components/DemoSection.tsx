export default function DemoSection() {
  return (
    <section
      id="demo"
      className="bg-white border border-slate-200 rounded-xl p-8 my-8"
    >
      <h2 className="text-2xl font-bold mb-6">Demo</h2>

      <div className="space-y-8">
        <div>
          <h3 className="text-xl font-semibold mb-4">Project Scaffolding</h3>

          <img
            src="/assets/demo.gif"
            alt="CLI Demo"
            className="rounded-lg border border-slate-200 w-full"
          />
        </div>

        <div>
          <h3 className="text-xl font-semibold mb-4">Full Stack Workflow</h3>

          <img
            src="/assets/workflow.gif"
            alt="Workflow"
            className="rounded-lg border border-slate-200 w-full"
          />
        </div>
      </div>
    </section>
  );
}
