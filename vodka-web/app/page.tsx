import Link from "next/link";
import Navbar from "./components/Navbar";

export default function Home() {
  return (
    <div className="min-h-screen bg-white text-slate-950 transition-colors dark:bg-slate-950 dark:text-slate-100">
      <Navbar />

      <main>
        <section className="bg-linear-to-br from-white via-slate-100 to-slate-200 py-20 text-slate-950 transition-colors dark:from-slate-950 dark:via-slate-900 dark:to-slate-800 dark:text-slate-100 sm:py-24">
          <div className="max-w-7xl mx-auto px-6 lg:px-8">
            <div className="grid gap-12 lg:grid-cols-[minmax(0,1fr)_420px] items-center">
              <div className="space-y-8">
                <div className="inline-flex rounded-full border border-slate-300 bg-white/70 px-4 py-1 text-sm font-medium text-slate-700 shadow-sm transition-colors dark:border-slate-700 dark:bg-slate-900/70 dark:text-slate-300">
                  Built for Go developers and full-stack teams
                </div>
                <div className="space-y-6">
                  <h1 className="text-4xl font-semibold tracking-tight sm:text-5xl lg:text-6xl">
                    A fast, focused Go framework for modern developer workflows.
                  </h1>
                  <p className="max-w-2xl text-lg leading-8 text-slate-600 dark:text-slate-300 sm:text-xl">
                    Vodka keeps the surface area small while giving you the tooling and structure to build full-stack apps quickly.
                  </p>
                </div>
                <div className="flex flex-col gap-4 sm:flex-row sm:items-center">
                  <Link href="/docs" className="inline-flex w-full items-center justify-center rounded-2xl bg-slate-950 px-7 py-3 text-sm font-semibold text-white shadow-sm shadow-slate-900/20 transition hover:-translate-y-0.5 hover:bg-slate-800 dark:bg-white dark:text-slate-950 dark:shadow-black/30 dark:hover:bg-slate-200 sm:w-auto">
                    Get Started
                  </Link>
                  <a href="https://github.com/DevanshuTripathi/vodka" target="_blank" rel="noopener noreferrer" className="inline-flex w-full items-center justify-center rounded-2xl border border-slate-950 px-7 py-3 text-sm font-semibold text-slate-950 transition hover:-translate-y-0.5 hover:bg-slate-950 hover:text-white dark:border-slate-400 dark:text-slate-100 dark:hover:border-white dark:hover:bg-white dark:hover:text-slate-950 sm:w-auto">
                    View on GitHub
                  </a>
                </div>
              </div>

              <div className="rounded-4xl border border-slate-200 bg-white/90 p-8 shadow-lg shadow-slate-900/5 backdrop-blur-xl transition-colors dark:border-slate-800 dark:bg-slate-900/80 dark:shadow-black/30">
                <div className="rounded-3xl bg-slate-950/95 p-6 text-white ring-1 ring-white/10 dark:bg-slate-900">
                  <p className="text-sm uppercase tracking-[0.3em] text-slate-400">Minimal by design</p>
                  <h2 className="mt-4 text-2xl font-semibold">Developer-first defaults</h2>
                  <p className="mt-3 text-sm leading-6 text-slate-300">
                    Clean routing, composable middleware, and seamless SPA readiness so you can iterate fast without sacrificing control.
                  </p>
                  <div className="mt-6 grid gap-4 sm:grid-cols-2">
                    <div className="rounded-3xl border border-white/10 bg-slate-950/70 p-4 transition hover:border-white/20 hover:bg-slate-900">
                      <p className="text-sm text-slate-400">Routing</p>
                      <p className="mt-2 font-semibold">Zero-allocation radix tree</p>
                    </div>
                    <div className="rounded-3xl border border-white/10 bg-slate-950/70 p-4 transition hover:border-white/20 hover:bg-slate-900">
                      <p className="text-sm text-slate-400">Middleware</p>
                      <p className="mt-2 font-semibold">Composable request flow</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>

        <section id="features" className="bg-slate-950 py-20 text-white transition-colors dark:bg-slate-900 sm:py-24">
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
                <article key={feature.title} className="group rounded-3xl border border-white/10 bg-white/5 p-6 shadow-sm shadow-black/10 transition hover:-translate-y-1 hover:border-slate-500/40 hover:bg-white/10">
                  <h3 className="text-xl font-semibold text-white">{feature.title}</h3>
                  <p className="mt-3 text-sm leading-6 text-slate-300">{feature.description}</p>
                </article>
              ))}
            </div>
          </div>
        </section>

        <section id="about" className="bg-white py-20 text-slate-950 transition-colors dark:bg-slate-950 dark:text-slate-100 sm:py-24">
          <div className="max-w-7xl mx-auto px-6 lg:px-8">
            <div className="grid gap-12 lg:grid-cols-2 lg:items-center">
              <div className="space-y-6">
                <p className="text-sm uppercase tracking-[0.3em] text-slate-500 dark:text-slate-400">About Vodka</p>
                <h2 className="text-3xl font-semibold sm:text-4xl">
                  Built for fast iteration, clear APIs, and dependable production apps.
                </h2>
                <p className="text-base leading-8 text-slate-600 dark:text-slate-300 sm:text-lg">
                  Vodka is designed around developer experience: fast startup, straightforward route definitions, and a minimal API surface that keeps your codebase readable and easy to evolve.
                </p>
              </div>
              <div className="rounded-3xl border border-slate-200 bg-slate-50 p-8 shadow-sm transition-colors dark:border-slate-800 dark:bg-slate-900/70">
                <div className="space-y-5">
                  <div>
                    <h3 className="text-lg font-semibold text-slate-950 dark:text-white">Philosophy</h3>
                    <p className="mt-3 text-sm leading-7 text-slate-600 dark:text-slate-300">
                      Focus on the essentials: clear routing, composable middleware, and integrated support for both API and frontend workflows.
                    </p>
                  </div>
                  <div>
                    <h3 className="text-lg font-semibold text-slate-950 dark:text-white">Developer Experience</h3>
                    <p className="mt-3 text-sm leading-7 text-slate-600 dark:text-slate-300">
                      Spend less time wiring boilerplate and more time building features, with predictable behavior and familiar Go idioms.
                    </p>
                  </div>
                  <div>
                    <h3 className="text-lg font-semibold text-slate-950 dark:text-white">Rapid Iteration</h3>
                    <p className="mt-3 text-sm leading-7 text-slate-600 dark:text-slate-300">
                      Lightweight tooling and fast reload-friendly structure let you iterate quickly across backend and frontend layers.
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>
      </main>

      <footer className="border-t border-slate-200 bg-slate-50 py-8 text-slate-600 transition-colors dark:border-slate-800 dark:bg-slate-950 dark:text-slate-300">
        <div className="max-w-7xl mx-auto px-6 lg:px-8">
          <div className="flex flex-col gap-6 items-start justify-between sm:flex-row sm:items-center">
            <div>
              <p className="text-sm font-semibold text-slate-950 dark:text-white">Vodka</p>
              <p className="mt-2 text-sm text-slate-500 dark:text-slate-400">Minimal Go web framework for fast full-stack development.</p>
            </div>
            <div className="flex flex-wrap gap-4 text-sm">
              <Link href="/docs" className="transition hover:text-slate-950 dark:hover:text-white">Docs</Link>
              <a href="https://github.com/DevanshuTripathi/vodka" target="_blank" rel="noopener noreferrer" className="transition hover:text-slate-950 dark:hover:text-white">GitHub</a>
            </div>
          </div>
        </div>
      </footer>
    </div>
  );
}
