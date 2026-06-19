import Link from "next/link";
import Navbar from "./components/Navbar";
import Contributors from "./components/Contributors";

export default function Home() {
  return (
    <div className="min-h-screen bg-white text-slate-950">
      <Navbar />

      <main>
        <section className="bg-linear-to-br from-white via-slate-100 to-slate-200 text-slate-950 py-20 sm:py-24">
          <div className="max-w-7xl mx-auto px-6 lg:px-8">
            <div className="grid gap-12 lg:grid-cols-[minmax(0,1fr)_420px] items-center">
              <div className="space-y-8">
                <div className="inline-flex rounded-full border border-slate-300 bg-slate-900/5 px-4 py-1 text-sm font-medium text-slate-700">
                  Built for Go developers and full-stack teams
                </div>
                <div className="space-y-6">
                  <h1 className="text-4xl font-semibold tracking-tight sm:text-5xl lg:text-6xl">
                    A fast, focused Go framework for modern developer workflows.
                  </h1>
                  <p className="max-w-2xl text-lg leading-8 text-slate-600 sm:text-xl">
                    Vodka keeps the surface area small while giving you the tooling and structure to build full-stack apps quickly.
                  </p>
                </div>
                <div className="flex flex-col gap-4 sm:flex-row sm:items-center">
                  <Link
                    href="/docs"
                    className="inline-flex w-full items-center justify-center rounded-2xl bg-slate-950 px-7 py-3 text-sm font-semibold text-white transition hover:bg-slate-800 sm:w-auto"
                  >
                    Get Started
                  </Link>
                  <a
                    href="https://github.com/DevanshuTripathi/vodka"
                    target="_blank"
                    rel="noopener noreferrer"
                    className="inline-flex w-full items-center justify-center rounded-2xl border border-slate-950 px-7 py-3 text-sm font-semibold text-slate-950 transition hover:bg-slate-950 hover:text-white sm:w-auto"
                  >
                    View on GitHub
                  </a>
                </div>
              </div>

              <div className="rounded-4xl border border-slate-200 bg-white/90 p-8 shadow-lg shadow-slate-900/5 backdrop-blur-xl">
                <div className="rounded-3xl bg-slate-950/95 p-6 text-white ring-1 ring-white/10">
                  <p className="text-sm uppercase tracking-[0.3em] text-slate-400">
                    Minimal by design
                  </p>
                  <h2 className="mt-4 text-2xl font-semibold">
                    Developer-first defaults
                  </h2>
                  <p className="mt-3 text-sm leading-6 text-slate-300">
                    Clean routing, composable middleware, and seamless SPA readiness so you can iterate fast without sacrificing control.
                  </p>
                  <div className="mt-6 grid gap-4 sm:grid-cols-2">
                    <div className="rounded-3xl border border-white/10 bg-slate-950/70 p-4">
                      <p className="text-sm text-slate-400">Routing</p>
                      <p className="mt-2 font-semibold">
                        Zero-allocation radix tree
                      </p>
                    </div>
                    <div className="rounded-3xl border border-white/10 bg-slate-950/70 p-4">
                      <p className="text-sm text-slate-400">Middleware</p>
                      <p className="mt-2 font-semibold">
                        Composable request flow
                      </p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>

        <section id="features" className="bg-slate-950 text-white py-20 sm:py-24">
          <div className="max-w-7xl mx-auto px-6 lg:px-8">
            <div className="max-w-3xl text-center mx-auto">
              <p className="text-sm uppercase tracking-[0.3em] text-slate-500">
                Features
              </p>
              <h2 className="mt-4 text-3xl font-semibold sm:text-4xl">
                Everything you need for modern Go web development.
              </h2>
              <p className="mt-4 text-base leading-7 text-slate-400 sm:text-lg">
                Solid defaults, clear APIs, and production-friendly design let you build APIs, middleware, and full-stack apps with confidence.
              </p>
            </div>

            <div className="mt-12 grid gap-6 sm:grid-cols-2 xl:grid-cols-3">
              {[
                {
                  title: "Fast Routing",
                  description:
                    "Radix tree based router for zero-allocation route matching",
                },
                {
                  title: "Middleware Chaining",
                  description: "Composable middleware with abort support",
                },
                {
                  title: "Authentication",
                  description:
                    "Built-in JWT validation helpers and Bearer auth",
                },
                {
                  title: "Request Validation",
                  description:
                    "Support for request validation using struct tags",
                },
                {
                  title: "React + Vite Integration",
                  description:
                    "Full-stack scaffolding with frontend and backend",
                },
                {
                  title: "SPA Support",
                  description:
                    "Seamless single page application serving in production",
                },
              ].map((feature) => (
                <article
                  key={feature.title}
                  className="group rounded-3xl border border-white/10 bg-white/5 p-6 transition hover:-translate-y-1 hover:border-slate-500/40 hover:bg-white/10"
                >
                  <h3 className="text-xl font-semibold text-white">
                    {feature.title}
                  </h3>
                  <p className="mt-3 text-sm leading-6 text-slate-300">
                    {feature.description}
                  </p>
                </article>
              ))}
            </div>
          </div>
        </section>

        <section id="about" className="bg-white text-slate-950 py-20 sm:py-24">
          <div className="max-w-7xl mx-auto px-6 lg:px-8">
            <div className="grid gap-12 lg:grid-cols-2 lg:items-center">
              <div className="space-y-6">
                <p className="text-sm uppercase tracking-[0.3em] text-slate-500">
                  About Vodka
                </p>
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
                    <h3 className="text-lg font-semibold text-slate-950">
                      Philosophy
                    </h3>
                    <p className="mt-3 text-sm leading-7 text-slate-600">
                      Focus on the essentials: clear routing, composable middleware, and integrated support for both API and frontend workflows.
                    </p>
                  </div>

                  <div>
                    <h3 className="text-lg font-semibold text-slate-950">
                      Developer Experience
                    </h3>
                    <p className="mt-3 text-sm leading-7 text-slate-600">
                      Spend less time wiring boilerplate and more time building features, with predictable behavior and familiar Go idioms.
                    </p>
                  </div>

                  <div>
                    <h3 className="text-lg font-semibold text-slate-950">
                      Rapid Iteration
                    </h3>
                    <p className="mt-3 text-sm leading-7 text-slate-600">
                      Lightweight tooling and fast reload-friendly structure let you iterate quickly across backend and frontend layers.
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>

        {/* CONTRIBUTORS SECTION ADDED */}
        <section id="contributors">
          <Contributors />
        </section>
      </main>



<footer className="bg-slate-950 border-t border-slate-800">
  <div className="max-w-7xl mx-auto px-6 lg:px-8 py-12">
    <div className="flex flex-col gap-10 lg:flex-row lg:justify-between">
      {/* Brand */}
      <div className="max-w-md">
        <h2 className="text-3xl font-bold text-white">Vodka</h2>

        <p className="mt-4 text-slate-400 leading-relaxed">
          A fast, focused Go framework for modern developer workflows.
        </p>

        <p className="mt-3 text-slate-500">
          Minimal by design. Built for performance. Made for developers.
        </p>
      </div>

      {/* Navigation */}
      <div className="lg:pt-3">
        <p className="mb-5 text-xs font-semibold uppercase tracking-[0.2em] text-slate-500">
          Navigation
        </p>

        <div className="flex flex-wrap gap-x-10 gap-y-5 text-lg">
          <Link
            href="/"
            className="font-medium text-slate-300 transition hover:text-white"
          >
            Home
          </Link>

          <Link
            href="/#features"
            className="font-medium text-slate-300 transition hover:text-white"
          >
            Features
          </Link>

          <Link
            href="/#about"
            className="font-medium text-slate-300 transition hover:text-white"
          >
            About
          </Link>

          <Link
            href="/docs"
            className="font-medium text-slate-300 transition hover:text-white"
          >
            Documentation
          </Link>

          <a
            href="https://github.com/DevanshuTripathi/vodka"
            target="_blank"
            rel="noopener noreferrer"
            className="font-medium text-slate-300 transition hover:text-white"
          >
            GitHub
          </a>
        </div>
      </div>
    </div>

    <div className="mt-10 flex flex-col gap-4 border-t border-slate-800 pt-6 sm:flex-row sm:items-center sm:justify-between">
      <p className="text-sm text-slate-500">
        © {new Date().getFullYear()} Vodka. Built with Go.
      </p>
    </div>
  </div>
</footer>



    </div>
  );
}