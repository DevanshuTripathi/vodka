export default function RoadmapSection() {
  return (
    <section
      id="roadmap"
      className="bg-white border border-slate-200 rounded-lg p-8 my-8"
    >
      <h2 className="text-xl font-bold mb-6">Roadmap</h2>

      <p className="text-slate-600 mb-6">
        Upcoming improvements planned for the Vodka ecosystem.
      </p>

      <div className="space-y-3">
        <div className="flex items-center gap-3">
          <span className="text-green-600">✓</span>
          <span>CLI scaffolding</span>
        </div>

        <div className="flex items-center gap-3">
          <span className="text-green-600">✓</span>
          <span>SPA integration</span>
        </div>

        <div className="flex items-center gap-3">
          <span className="text-slate-400">○</span>
          <span>Plugin ecosystem</span>
        </div>

        <div className="flex items-center gap-3">
          <span className="text-slate-400">○</span>
          <span>Documentation improvements</span>
        </div>

        <div className="flex items-center gap-3">
          <span className="text-slate-400">○</span>
          <span>Additional templates</span>
        </div>
      </div>
    </section>
  );
}
