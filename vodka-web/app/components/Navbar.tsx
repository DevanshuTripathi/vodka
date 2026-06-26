"use client";

import Link from "next/link";

export default function Navbar() {
  return (
    <nav className="sticky top-0 bg-black text-white border-b border-slate-800 ">
      <div className="max-w-7xl mx-auto px-4 sm:px-6">
        <div className="flex flex-col gap-4 py-4 md:flex-row md:items-center md:justify-between">
          <Link href="/">
            <span className="text-2xl font-bold">Vodka</span>
          </Link>

          <div className="flex flex-wrap items-center justify-center gap-3 text-sm font-medium md:justify-start">
            <Link href="/">
              <span className="cursor-pointer rounded-full border border-slate-700 px-4 py-2 transition-colors duration-200 hover:border-slate-500 hover:text-slate-300">Home</span>
            </Link>
            <a href="#features" className="rounded-full border border-slate-700 px-4 py-2 transition-colors duration-200 hover:border-slate-500 hover:text-slate-300">
              Features
            </a>
            <a href="#about" className="rounded-full border border-slate-700 px-4 py-2 transition-colors duration-200 hover:border-slate-500 hover:text-slate-300">
              About 
            </a>
                    <a
          href="#contact"
          className="rounded-full border border-slate-700 px-4 py-2 transition-colors duration-200 hover:border-slate-500 hover:text-slate-300"
        >
          Contact
        </a>
            <a
              href="#contributors"
              className="rounded-full border border-slate-700 px-4 py-2 transition-colors duration-200 hover:border-slate-500 hover:text-slate-300"
            >
              Contributors
            </a>

            <Link href="/docs">
              <span className="cursor-pointer rounded-full border border-slate-700 px-4 py-2 transition-colors duration-200 hover:border-slate-500 hover:text-slate-300">Documentation</span>
            </Link>
            <a href="https://github.com/DevanshuTripathi/vodka" target="_blank" rel="noopener noreferrer" className="rounded-full border border-slate-700 px-4 py-2 transition-colors duration-200 hover:border-slate-500 hover:text-slate-300">
              GitHub
            </a>
          </div>
        </div>
      </div>
    </nav>
  );
}