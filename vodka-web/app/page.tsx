import Link from "next/link";
import Navbar from "./components/Navbar";
import Hero from "./components/Hero";

export default function Home() {
  return (
    <div className="min-h-screen bg-white text-slate-950">
      <Navbar />

      <main>
        <Hero />

        <section id="features" className="bg-slate-950 text-white py-20 sm:py-24">
          <div className="max-w-7xl mx-auto px-6 lg:px-8">
            <div className="max-w-3xl text-center mx-auto">
              <p className="text-sm uppercase tracking-[0.3em] text-slate-500">Features</p>
              <h2 className="mt-4 text-3xl font-semibold sm:text-4xl">Everything you need for modern Go web development.</h2>
              <p className="mt-4 text-base leading-7 text-slate-400 sm:text-lg">
                Solid defaults, clear APIs, and production-friendly design let you build APIs, middleware, and full-stack apps with confidence.
              </p>
            </div>

            <div className="mt-12 grid gap-6 sm:grid-cols-2 xl:grid-cols-3">
              {[
                { title: "Fast Routing", description: "Radix tree based router for zero-allocation route matching" },
                { title: "Middleware Chaining", description: "Composable middleware with abort support" },
                { title: "Authentication", description: "Built-in JWT validation helpers and Bearer auth" },
                { title: "Request Validation", description: "Support for request validation using struct tags" },
                { title: "React + Vite Integration", description: "Full-stack scaffolding with frontend and backend" },
                { title: "SPA Support", description: "Seamless single page application serving in production" },
              ].map((feature) => (
                <article key={feature.title} className="group rounded-3xl border border-white/10 bg-white/5 p-6 transition hover:-translate-y-1 hover:border-slate-500/40 hover:bg-white/10">
                  <h3 className="text-xl font-semibold text-white">{feature.title}</h3>
                  <p className="mt-3 text-sm leading-6 text-slate-300">{feature.description}</p>
                </article>
              ))}
            </div>
          </div>
        </section>

        <section id="about" className="bg-white text-slate-950 py-20 sm:py-24">
          <div className="max-w-7xl mx-auto px-6 lg:px-8">
            <div className="grid gap-12 lg:grid-cols-2 lg:items-center">
              <div className="space-y-6">
                <p className="text-sm uppercase tracking-[0.3em] text-slate-500">About Vodka</p>
                <h2 className="text-3xl font-semibold sm:text-4xl">
                  Built for fast iteration, clear APIs, and dependable production apps.
                </h2>
                <p className="text-base leading-8 text-slate-600 sm:text-lg">
                  Vodka is designed around developer experience: fast startup, straightforward route definitions, and a minimal API surface that keeps your codebase readable and easy to evolve.
                </p>
              </div>
              <div className="rounded-3xl border border-slate-200 bg-slate-950/5 p-8 shadow-sm">
                <div className="space-y-5">
                  <div>
                    <h3 className="text-lg font-semibold text-slate-950">Philosophy</h3>
                    <p className="mt-3 text-sm leading-7 text-slate-600">
                      Focus on the essentials: clear routing, composable middleware, and integrated support for both API and frontend workflows.
                    </p>
                  </div>
                  <div>
                    <h3 className="text-lg font-semibold text-slate-950">Developer Experience</h3>
                    <p className="mt-3 text-sm leading-7 text-slate-600">
                      Spend less time wiring boilerplate and more time building features, with predictable behavior and familiar Go idioms.
                    </p>
                  </div>
                  <div>
                    <h3 className="text-lg font-semibold text-slate-950">Rapid Iteration</h3>
                    <p className="mt-3 text-sm leading-7 text-slate-600">
                      Lightweight tooling and fast reload-friendly structure let you iterate quickly across backend and frontend layers.
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>
      </main>

      <footer className="bg-slate-950 text-slate-300 border-t border-slate-800 py-8">
        <div className="max-w-7xl mx-auto px-6 lg:px-8">
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
