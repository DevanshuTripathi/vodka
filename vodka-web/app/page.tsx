import Link from "next/link";
import Navbar from "./components/Navbar";

export default function Home() {
  return (
    <div className="min-h-screen bg-slate-950 text-slate-950">
      <Navbar />

      <main>
        <section className="relative overflow-hidden bg-[radial-gradient(circle_at_top_left,_rgba(59,130,246,0.22),_transparent_28%),radial-gradient(circle_at_top_right,_rgba(168,85,247,0.18),_transparent_32%),linear-gradient(180deg,_#020617_0%,_#0f172a_52%,_#e2e8f0_52%,_#f8fafc_100%)] py-20 sm:py-24">
          <div className="mx-auto max-w-7xl px-6 lg:px-8">
            <div className="grid items-center gap-12 lg:grid-cols-[minmax(0,1fr)_440px]">
              <div className="space-y-8 text-white">
                <div className="inline-flex rounded-full border border-white/15 bg-white/8 px-4 py-1 text-sm font-medium text-slate-200">
                  Built for Go developers and full-stack teams
                </div>
                <div className="space-y-6">
                  <h1 className="max-w-4xl text-4xl font-semibold tracking-tight sm:text-5xl lg:text-6xl">
                    Ship crisp Go backends without dragging your team through framework noise.
                  </h1>
                  <p className="max-w-2xl text-lg leading-8 text-slate-300 sm:text-xl">
                    Vodka keeps the API surface small, the defaults practical, and the full-stack path clear so teams can move from routes to production without a pile of ceremony.
                  </p>
                </div>
                <div className="flex flex-col gap-4 sm:flex-row sm:items-center">
                  <Link href="/docs" className="inline-flex w-full items-center justify-center rounded-2xl bg-white px-7 py-3 text-sm font-semibold text-slate-950 transition hover:bg-slate-200 sm:w-auto">
                    Get Started
                  </Link>
                  <a href="https://github.com/DevanshuTripathi/vodka" target="_blank" rel="noopener noreferrer" className="inline-flex w-full items-center justify-center rounded-2xl border border-white/20 bg-white/5 px-7 py-3 text-sm font-semibold text-white transition hover:bg-white/12 sm:w-auto">
                    View on GitHub
                  </a>
                </div>
                <div className="grid gap-4 pt-2 sm:grid-cols-3">
                  {[
                    { label: "Routing", value: "Radix tree core" },
                    { label: "Middleware", value: "Composable chain" },
                    { label: "Stack", value: "Go + React ready" },
                  ].map((item) => (
                    <div key={item.label} className="rounded-2xl border border-white/10 bg-white/6 p-4">
                      <p className="text-xs uppercase tracking-[0.24em] text-slate-400">{item.label}</p>
                      <p className="mt-2 text-sm font-semibold text-white">{item.value}</p>
                    </div>
                  ))}
                </div>
              </div>

              <div className="rounded-[2rem] border border-white/12 bg-slate-950/70 p-6 text-white shadow-2xl shadow-slate-950/45 backdrop-blur-xl">
                <div className="rounded-[1.6rem] border border-white/8 bg-slate-900/90 p-6 ring-1 ring-white/8">
                  <div className="flex items-center justify-between">
                    <p className="text-sm uppercase tracking-[0.3em] text-slate-400">Landing flow</p>
                    <div className="flex gap-2">
                      <span className="h-2.5 w-2.5 rounded-full bg-emerald-400" />
                      <span className="h-2.5 w-2.5 rounded-full bg-sky-400" />
                      <span className="h-2.5 w-2.5 rounded-full bg-fuchsia-400" />
                    </div>
                  </div>
                  <h2 className="mt-5 text-2xl font-semibold">Developer-first defaults</h2>
                  <p className="mt-3 text-sm leading-6 text-slate-300">
                    A landing page should immediately show how Vodka feels: minimal, sharp, and production-aware. This card turns the core framework story into a quick scan instead of a wall of text.
                  </p>

                  <div className="mt-6 space-y-4">
                    {[
                      {
                        title: "Define routes with intent",
                        body: "Keep request flow readable with clear route registration and middleware composition.",
                      },
                      {
                        title: "Stay fast while scaling",
                        body: "Built-in ergonomics keep local dev simple while leaving room for serious production behavior.",
                      },
                      {
                        title: "Connect frontend confidently",
                        body: "The framework story already extends to full-stack patterns, examples, and docs navigation.",
                      },
                    ].map((item, index) => (
                      <div key={item.title} className="rounded-3xl border border-white/8 bg-white/4 p-4">
                        <div className="flex items-start gap-4">
                          <div className="flex h-9 w-9 shrink-0 items-center justify-center rounded-2xl bg-white/10 text-sm font-semibold text-slate-100">
                            0{index + 1}
                          </div>
                          <div>
                            <p className="font-semibold text-white">{item.title}</p>
                            <p className="mt-2 text-sm leading-6 text-slate-300">{item.body}</p>
                          </div>
                        </div>
                      </div>
                    ))}
                  </div>

                  <div className="mt-6 rounded-3xl border border-sky-400/20 bg-sky-400/10 p-4 text-sm text-sky-100">
                    Maintainer ask: start with the landing page first, then expand the docs site page by page.
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>

        <section id="features" className="bg-slate-50 py-20 text-slate-950 sm:py-24">
          <div className="mx-auto max-w-7xl px-6 lg:px-8">
            <div className="mx-auto max-w-3xl text-center">
              <p className="text-sm uppercase tracking-[0.3em] text-slate-500">Features</p>
              <h2 className="mt-4 text-3xl font-semibold sm:text-4xl">Everything you need for modern Go web development.</h2>
              <p className="mt-4 text-base leading-7 text-slate-600 sm:text-lg">
                Clear APIs, production-friendly defaults, and just enough structure to help teams ship without trapping them in framework ceremony.
              </p>
            </div>

            <div className="mt-12 grid gap-6 sm:grid-cols-2 xl:grid-cols-3">
              {[
                { title: "Fast Routing", description: "Radix tree route matching keeps the core path lean and predictable." },
                { title: "Middleware Chaining", description: "Compose request flow in the order your team actually thinks about it." },
                { title: "Authentication Helpers", description: "JWT validation utilities and request guards keep protected flows practical." },
                { title: "Validation Support", description: "Work with request constraints and struct-friendly patterns without extra clutter." },
                { title: "SPA + Full-stack Paths", description: "Pair backend routes with frontend delivery patterns instead of stitching them later." },
                { title: "Examples That Teach", description: "JWT, file upload, middleware, SSE, and more make the docs immediately usable." },
              ].map((feature) => (
                <article key={feature.title} className="rounded-3xl border border-slate-200 bg-white p-6 shadow-sm transition hover:-translate-y-1 hover:shadow-lg">
                  <h3 className="text-xl font-semibold text-slate-950">{feature.title}</h3>
                  <p className="mt-3 text-sm leading-6 text-slate-600">{feature.description}</p>
                </article>
              ))}
            </div>
          </div>
        </section>

        <section id="workflow" className="bg-white py-20 text-slate-950 sm:py-24">
          <div className="mx-auto grid max-w-7xl gap-10 px-6 lg:grid-cols-[0.95fr_1.05fr] lg:px-8">
            <div className="space-y-5">
              <p className="text-sm uppercase tracking-[0.3em] text-slate-500">Workflow</p>
              <h2 className="text-3xl font-semibold sm:text-4xl">
                The framework story should feel obvious from the first scroll.
              </h2>
              <p className="text-base leading-8 text-slate-600 sm:text-lg">
                Instead of a generic hero plus a loose feature grid, the landing page now tells a sharper story: why Vodka exists, how teams use it, and where to go next.
              </p>
            </div>

            <div className="grid gap-5 sm:grid-cols-3">
              {[
                {
                  title: "Scan the promise",
                  body: "See the framework position, the stack fit, and the docs entry point immediately.",
                },
                {
                  title: "Validate the toolkit",
                  body: "Confirm routing, middleware, auth, and full-stack patterns without digging through the repo first.",
                },
                {
                  title: "Jump into docs",
                  body: "Move cleanly from the landing page into the fetched documentation and examples.",
                },
              ].map((step, index) => (
                <div key={step.title} className="rounded-3xl border border-slate-200 bg-slate-50 p-5">
                  <p className="text-xs font-semibold uppercase tracking-[0.24em] text-slate-500">Step 0{index + 1}</p>
                  <h3 className="mt-4 text-lg font-semibold text-slate-950">{step.title}</h3>
                  <p className="mt-3 text-sm leading-6 text-slate-600">{step.body}</p>
                </div>
              ))}
            </div>
          </div>
        </section>

        <section id="about" className="bg-slate-950 py-20 text-white sm:py-24">
          <div className="mx-auto max-w-7xl px-6 lg:px-8">
            <div className="grid gap-12 lg:grid-cols-2 lg:items-center">
              <div className="space-y-6">
                <p className="text-sm uppercase tracking-[0.3em] text-slate-500">About Vodka</p>
                <h2 className="text-3xl font-semibold sm:text-4xl">
                  Built for fast iteration, clear APIs, and dependable production apps.
                </h2>
                <p className="text-base leading-8 text-slate-300 sm:text-lg">
                  Vodka is designed around developer experience: fast startup, straightforward route definitions, and a minimal API surface that keeps your codebase readable and easy to evolve.
                </p>
              </div>
              <div className="rounded-3xl border border-white/10 bg-white/5 p-8 shadow-sm">
                <div className="space-y-5">
                  <div>
                    <h3 className="text-lg font-semibold text-white">Philosophy</h3>
                    <p className="mt-3 text-sm leading-7 text-slate-300">
                      Focus on the essentials: clear routing, composable middleware, and integrated support for both API and frontend workflows.
                    </p>
                  </div>
                  <div>
                    <h3 className="text-lg font-semibold text-white">Developer Experience</h3>
                    <p className="mt-3 text-sm leading-7 text-slate-300">
                      Spend less time wiring boilerplate and more time building features, with predictable behavior and familiar Go idioms.
                    </p>
                  </div>
                  <div>
                    <h3 className="text-lg font-semibold text-white">Rapid Iteration</h3>
                    <p className="mt-3 text-sm leading-7 text-slate-300">
                      Lightweight tooling and fast reload-friendly structure let you iterate quickly across backend and frontend layers.
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>
      </main>

      <footer className="border-t border-slate-800 bg-slate-950 py-8 text-slate-300">
        <div className="mx-auto max-w-7xl px-6 lg:px-8">
          <div className="flex flex-col gap-6 items-start justify-between sm:flex-row sm:items-center">
            <div>
              <p className="text-sm font-semibold text-white">Vodka</p>
              <p className="mt-2 text-sm text-slate-500">Minimal Go web framework for fast full-stack development.</p>
            </div>
            <div className="flex flex-wrap gap-4 text-sm">
              <Link href="/docs" className="transition hover:text-white">Docs</Link>
              <a href="https://github.com/DevanshuTripathi/vodka" target="_blank" rel="noopener noreferrer" className="transition hover:text-white">GitHub</a>
            </div>
          </div>
        </div>
      </footer>
    </div>
  );
}
