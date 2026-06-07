"use client";

import Link from "next/link";
import { motion } from "framer-motion";

const containerVariants = {
  hidden: { opacity: 0 },
  visible: {
    opacity: 1,
    transition: {
      staggerChildren: 0.15,
      delayChildren: 0.2,
    },
  },
};

const itemVariants = {
  hidden: { opacity: 0, y: 24 },
  visible: {
    opacity: 1,
    y: 0,
    transition: { duration: 0.6, ease: "easeOut" as const },
  },
};

export default function Hero() {
  return (
    <section className="relative overflow-hidden bg-linear-to-br from-white via-slate-100 to-slate-200 py-20 text-slate-950 sm:py-24">
      <div
        className="pointer-events-none absolute right-0 top-1/2 -translate-y-1/2 size-[600px]"
        aria-hidden="true"
      >
        <div className="size-full rounded-full bg-[radial-gradient(circle_at_center,oklch(0.65_0.12_240/0.2),transparent_70%)]" />
      </div>

      <motion.div
        className="mx-auto max-w-7xl px-6 lg:px-8"
        variants={containerVariants}
        initial="hidden"
        animate="visible"
      >
        <div className="grid items-center gap-12 lg:grid-cols-[minmax(0,1fr)_420px]">
          <div className="space-y-8">
            <motion.div variants={itemVariants}>
              <div className="inline-flex rounded-full border border-slate-300 bg-slate-900/5 px-4 py-1 text-sm font-medium text-slate-700">
                Built for Go developers and full-stack teams
              </div>
            </motion.div>

            <div className="space-y-6">
              <motion.h1
                variants={itemVariants}
                className="text-4xl font-semibold tracking-tight sm:text-5xl lg:text-6xl"
              >
                A fast, focused Go framework for modern developer workflows.
              </motion.h1>
              <motion.p
                variants={itemVariants}
                className="max-w-2xl text-lg leading-8 text-slate-600 sm:text-xl"
              >
                Vodka keeps the surface area small while giving you the tooling
                and structure to build full-stack apps quickly.
              </motion.p>
            </div>

            <motion.div
              variants={itemVariants}
              className="flex flex-col gap-4 pt-4 sm:flex-row sm:items-center"
            >
              <motion.div
                whileHover={{ scale: 1.04 }}
                whileTap={{ scale: 0.97 }}
                transition={{ type: "spring", stiffness: 400, damping: 17 }}
              >
                <Link
                  href="/docs"
                  className="inline-flex w-full items-center justify-center rounded-2xl bg-slate-950 px-7 py-3 text-sm font-semibold text-white transition-colors duration-300 hover:bg-slate-700 sm:w-auto"
                >
                  Get Started
                </Link>
              </motion.div>

              <motion.div
                whileHover={{ scale: 1.04 }}
                whileTap={{ scale: 0.97 }}
                transition={{ type: "spring", stiffness: 400, damping: 17 }}
              >
                <a
                  href="https://github.com/DevanshuTripathi/vodka"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="relative inline-flex w-full items-center justify-center overflow-hidden rounded-2xl border border-slate-950 px-7 py-3 text-sm font-semibold sm:w-auto"
                >
                  <motion.span
                    className="absolute inset-0 rounded-2xl bg-slate-950"
                    initial={{ x: "-100%" }}
                    whileHover={{ x: 0 }}
                    transition={{ duration: 0.35, ease: "easeOut" }}
                  />
                  <motion.span
                    className="relative z-10"
                    initial={{ color: "#0f172a" }}
                    whileHover={{ color: "#ffffff" }}
                    transition={{ duration: 0.35 }}
                  >
                    View on GitHub
                  </motion.span>
                </a>
              </motion.div>
            </motion.div>
          </div>

          <motion.div
            variants={itemVariants}
            animate={{ y: [0, -8, 0] }}
            transition={{ duration: 4, repeat: Infinity, ease: "easeInOut" }}
            className="rounded-4xl border border-slate-200 bg-white/90 p-8 shadow-lg shadow-slate-900/5 backdrop-blur-xl"
          >
            <div className="rounded-3xl bg-slate-950/95 p-6 text-white ring-1 ring-white/10">
              <p className="text-sm uppercase tracking-[0.3em] text-slate-400">
                Minimal by design
              </p>
              <h2 className="mt-4 text-2xl font-semibold">
                Developer-first defaults
              </h2>
              <p className="mt-3 text-sm leading-6 text-slate-300">
                Clean routing, composable middleware, and seamless SPA readiness
                so you can iterate fast without sacrificing control.
              </p>
              <div className="mt-6 grid gap-4 sm:grid-cols-2">
                <motion.div
                  whileHover={{ scale: 1.02 }}
                  transition={{ type: "spring", stiffness: 400, damping: 20 }}
                  className="rounded-3xl border border-white/10 bg-slate-950/70 p-4"
                >
                  <p className="text-sm text-slate-400">Routing</p>
                  <p className="mt-2 font-semibold">
                    Zero-allocation radix tree
                  </p>
                </motion.div>
                <motion.div
                  whileHover={{ scale: 1.02 }}
                  transition={{ type: "spring", stiffness: 400, damping: 20 }}
                  className="rounded-3xl border border-white/10 bg-slate-950/70 p-4"
                >
                  <p className="text-sm text-slate-400">Middleware</p>
                  <p className="mt-2 font-semibold">
                    Composable request flow
                  </p>
                </motion.div>
              </div>
            </div>
          </motion.div>
        </div>
      </motion.div>
    </section>
  );
}
