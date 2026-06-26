"use client";

import Link from "next/link";
import { motion } from "framer-motion";

const containerVariants = {
  hidden: {},
  show: {
    transition: {
      staggerChildren: 0.15,
    },
  },
};

const itemVariants = {
  hidden: {
    opacity: 0,
    y: 30,
  },
  show: {
    opacity: 1,
    y: 0,
    transition: {
      duration: 0.55,
      ease: [0.22, 1, 0.36, 1],
    },
  },
};

export default function Contact() {
  return (
    <section
      id="contact"
      className="relative overflow-hidden bg-white py-24 sm:py-28"
    >
        <div className="mx-auto mb-20 flex max-w-7xl items-center gap-6 px-6 lg:px-8">
        <div className="h-px flex-1 bg-slate-200" />

        <span className="text-sm font-medium uppercase tracking-[0.25em] text-slate-400">
          Contact
        </span>

        <div className="h-px flex-1 bg-slate-200" />
      </div>
      {/* Background */}
      <div className="absolute inset-0 -z-10 overflow-hidden">
        <div className="absolute left-0 top-0 h-72 w-72 rounded-full bg-slate-200/40 blur-3xl" />
        <div className="absolute right-0 bottom-0 h-96 w-96 rounded-full bg-slate-300/30 blur-3xl" />
      </div>

      <div className="mx-auto max-w-7xl px-6 lg:px-8">

        {/* Heading */}

        <motion.div
          initial={{ opacity: 0, y: 40 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          transition={{
            duration: 0.7,
            ease: [0.22, 1, 0.36, 1],
          }}
          className="mx-auto max-w-3xl text-center"
        >
          <span className="rounded-full border border-slate-300 bg-slate-100 px-4 py-1 text-sm font-medium text-slate-700">
            Contact Us
          </span>

          <h2 className="mt-6 text-4xl font-bold tracking-tight text-slate-950 sm:text-5xl">
            We'd Love To Hear From You
          </h2>

          <p className="mt-6 text-lg leading-8 text-slate-600">
            Whether you have a question, found a bug, want to request a
            feature, or simply want to contribute to Vodka, we're always happy
            to connect with developers.
          </p>
        </motion.div>

        {/* Grid */}

        <div className="mt-20 grid gap-10 lg:grid-cols-[400px_1fr]">

          {/* Left Side */}

          <motion.div
            variants={containerVariants}
            initial="hidden"
            whileInView="show"
            viewport={{ once: true }}
            className="space-y-6"
          >

            {/* GitHub */}

            <motion.div
              variants={itemVariants}
              whileHover={{
                y: -6,
                transition: {
                  duration: 0.2,
                },
              }}
              className="rounded-3xl border border-slate-200 bg-white p-7 shadow-sm"
            >
              <motion.div
                whileHover={{
                  scale: 1.08,
                  rotate: 8,
                }}
                className="flex h-14 w-14 items-center justify-center rounded-2xl bg-slate-950 text-white"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 24 24"
                  fill="currentColor"
                  className="h-7 w-7"
                >
                  <path d="M12 .5C5.65.5.5 5.65.5 12A11.5 11.5 0 0 0 8.35 22.9c.58.1.79-.25.79-.56v-2.03c-3.2.7-3.88-1.54-3.88-1.54-.52-1.33-1.28-1.69-1.28-1.69-1.05-.72.08-.71.08-.71 1.16.08 1.77 1.19 1.77 1.19 1.03 1.76 2.7 1.25 3.36.96.1-.75.4-1.25.73-1.54-2.56-.29-5.25-1.28-5.25-5.72 0-1.26.45-2.3 1.18-3.11-.12-.29-.51-1.46.11-3.05 0 0 .96-.31 3.15 1.19a10.9 10.9 0 0 1 5.74 0c2.19-1.5 3.15-1.19 3.15-1.19.62 1.59.23 2.76.11 3.05.73.81 1.18 1.85 1.18 3.11 0 4.45-2.69 5.43-5.26 5.71.42.36.79 1.08.79 2.18v3.23c0 .31.21.67.8.56A11.5 11.5 0 0 0 23.5 12C23.5 5.65 18.35.5 12 .5Z"/>
                </svg>
              </motion.div>

              <h3 className="mt-6 text-xl font-semibold text-slate-950">
                GitHub Repository
              </h3>

              <p className="mt-3 leading-7 text-slate-600">
                Explore the source code, contribute to the framework,
                report bugs, or request new features.
              </p>

              <a
                href="https://github.com/DevanshuTripathi/vodka"
                target="_blank"
                rel="noopener noreferrer"
                className="mt-6 inline-flex items-center gap-2 font-medium text-slate-950 transition hover:gap-3"
              >
                Visit Repository

                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  strokeWidth={2}
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  className="h-4 w-4"
                >
                  <path d="M5 12h14" />
                  <path d="m12 5 7 7-7 7" />
                </svg>
              </a>
            </motion.div>
                        {/* Documentation */}

            <motion.div
              variants={itemVariants}
              whileHover={{
                y: -6,
                transition: { duration: 0.2 },
              }}
              className="rounded-3xl border border-slate-200 bg-white p-7 shadow-sm"
            >
              <motion.div
                whileHover={{ scale: 1.08, rotate: 8 }}
                className="flex h-14 w-14 items-center justify-center rounded-2xl bg-slate-950 text-white"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  strokeWidth={2}
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  className="h-7 w-7"
                >
                  <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20" />
                  <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5V4.5A2.5 2.5 0 0 1 6.5 2Z" />
                </svg>
              </motion.div>

              <h3 className="mt-6 text-xl font-semibold text-slate-950">
                Documentation
              </h3>

              <p className="mt-3 leading-7 text-slate-600">
                Browse our guides, examples and API documentation to get started
                quickly with Vodka.
              </p>

              <Link
                href="/docs"
                className="mt-6 inline-flex items-center gap-2 font-medium text-slate-950 transition hover:gap-3"
              >
                Read Documentation

                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  strokeWidth={2}
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  className="h-4 w-4"
                >
                  <path d="M5 12h14" />
                  <path d="m12 5 7 7-7 7" />
                </svg>
              </Link>
            </motion.div>

            {/* Report Issue */}

            <motion.div
              variants={itemVariants}
              whileHover={{
                y: -6,
                transition: { duration: 0.2 },
              }}
              className="rounded-3xl border border-slate-200 bg-white p-7 shadow-sm"
            >
              <motion.div
                whileHover={{ scale: 1.08, rotate: 8 }}
                className="flex h-14 w-14 items-center justify-center rounded-2xl bg-slate-950 text-white"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  strokeWidth={2}
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  className="h-7 w-7"
                >
                  <path d="M8 2l1 2" />
                  <path d="M16 2l-1 2" />
                  <path d="M9 7h6" />
                  <rect x="7" y="7" width="10" height="12" rx="5" />
                  <path d="M5 13H2" />
                  <path d="M22 13h-3" />
                </svg>
              </motion.div>

              <h3 className="mt-6 text-xl font-semibold text-slate-950">
                Report an Issue
              </h3>

              <p className="mt-3 leading-7 text-slate-600">
                Found a bug or have an idea? Open an issue on GitHub and help us
                improve the framework.
              </p>
            </motion.div>

          </motion.div>

          {/* Contact Form */}

          <motion.div
            initial={{ opacity: 0, x: 40 }}
            whileInView={{ opacity: 1, x: 0 }}
            viewport={{ once: true }}
            transition={{
              duration: 0.6,
              ease: [0.22, 1, 0.36, 1],
            }}
            className="rounded-[2rem] border border-slate-200 bg-white p-8 shadow-xl"
          >
            <h3 className="text-3xl font-bold text-slate-950">
              Send us a Message
            </h3>

            <p className="mt-3 text-slate-600">
              Fill out the form below and we'll get back to you as soon as
              possible.
            </p>

            <form className="mt-10 space-y-6">

              <div className="grid gap-6 md:grid-cols-2">

                <div>
                  <label className="mb-2 block text-sm font-medium text-slate-700">
                    Full Name
                  </label>

                  <input
                    type="text"
                    placeholder="John Doe"
                    className="w-full rounded-xl border border-slate-300 px-4 py-3 outline-none transition focus:border-slate-900"
                  />
                </div>

                <div>
                  <label className="mb-2 block text-sm font-medium text-slate-700">
                    Email Address
                  </label>

                  <input
                    type="email"
                    placeholder="john@example.com"
                    className="w-full rounded-xl border border-slate-300 px-4 py-3 outline-none transition focus:border-slate-900"
                  />
                </div>

              </div>

              <div>
                <label className="mb-2 block text-sm font-medium text-slate-700">
                  Subject
                </label>

                <input
                  type="text"
                  placeholder="Feature Request"
                  className="w-full rounded-xl border border-slate-300 px-4 py-3 outline-none transition focus:border-slate-900"
                />
              </div>

              <div>
                <label className="mb-2 block text-sm font-medium text-slate-700">
                  Message
                </label>

                <textarea
                  rows={6}
                  placeholder="Tell us more about your question or feedback..."
                  className="w-full resize-none rounded-xl border border-slate-300 px-4 py-3 outline-none transition focus:border-slate-900"
                />
              </div>
                            <motion.button
                whileHover={{ scale: 1.02 }}
                whileTap={{ scale: 0.98 }}
                transition={{
                  duration: 0.2,
                  ease: [0.22, 1, 0.36, 1],
                }}
                type="submit"
                className="inline-flex w-full items-center justify-center gap-2 rounded-xl bg-slate-950 px-6 py-4 font-semibold text-white shadow-lg shadow-slate-900/20 transition-colors hover:bg-slate-800 focus:outline-none focus-visible:ring-2 focus-visible:ring-slate-500 focus-visible:ring-offset-2"
              >
                Send Message

                <motion.svg
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  strokeWidth={2}
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  className="h-5 w-5"
                  whileHover={{ x: 3 }}
                >
                  <path d="M22 2 11 13" />
                  <path d="m22 2-7 20-4-9-9-4Z" />
                </motion.svg>
              </motion.button>

              <p className="text-center text-sm text-slate-500">
                We usually respond within{" "}
                <span className="font-semibold text-slate-700">
                  24–48 hours
                </span>
                .
              </p>
            </form>
          </motion.div>
        </div>

        {/* Bottom CTA */}

        <motion.div
          initial={{ opacity: 0, y: 30 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          transition={{
            duration: 0.6,
            delay: 0.15,
            ease: [0.22, 1, 0.36, 1],
          }}
          className="mt-24 rounded-[2rem] border border-slate-200 bg-gradient-to-r from-slate-900 via-slate-800 to-slate-900 px-8 py-12 text-center text-white shadow-xl"
        >
          <h3 className="text-3xl font-bold">
            Want to Contribute?
          </h3>

          <p className="mx-auto mt-4 max-w-2xl text-slate-300">
            Vodka is an open-source framework built by developers, for
            developers. Every contribution—whether it's code, documentation, bug
            reports, or feature ideas—helps make the project even better.
          </p>

          <motion.a
            whileHover={{
              scale: 1.04,
            }}
            whileTap={{
              scale: 0.97,
            }}
            href="https://github.com/DevanshuTripathi/vodka"
            target="_blank"
            rel="noopener noreferrer"
            className="mt-8 inline-flex items-center gap-2 rounded-xl bg-white px-6 py-3 font-semibold text-slate-950 transition hover:bg-slate-100"
          >
            Contribute on GitHub

            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="currentColor"
              className="h-5 w-5"
            >
              <path d="M12 .5C5.65.5.5 5.65.5 12A11.5 11.5 0 0 0 8.35 22.9c.58.1.79-.25.79-.56v-2.03c-3.2.7-3.88-1.54-3.88-1.54-.52-1.33-1.28-1.69-1.28-1.69-1.05-.72.08-.71.08-.71 1.16.08 1.77 1.19 1.77 1.19 1.03 1.76 2.7 1.25 3.36.96.1-.75.4-1.25.73-1.54-2.56-.29-5.25-1.28-5.25-5.72 0-1.26.45-2.3 1.18-3.11-.12-.29-.51-1.46.11-3.05 0 0 .96-.31 3.15 1.19a10.9 10.9 0 0 1 5.74 0c2.19-1.5 3.15-1.19 3.15-1.19.62 1.59.23 2.76.11 3.05.73.81 1.18 1.85 1.18 3.11 0 4.45-2.69 5.43-5.26 5.71.42.36.79 1.08.79 2.18v3.23c0 .31.21.67.8.56A11.5 11.5 0 0 0 23.5 12C23.5 5.65 18.35.5 12 .5Z" />
            </svg>
          </motion.a>
        </motion.div>
      </div>
    </section>
  );
}