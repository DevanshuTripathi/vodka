"use client";

import Link from "next/link";
import ThemeToggle from "./ThemeToggle";

export default function Navbar() {
  return (
    <nav className="sticky top-0 z-40 border-b border-slate-200 bg-white/90 text-slate-950 backdrop-blur-xl transition-colors dark:border-slate-800 dark:bg-slate-950/90 dark:text-white">
      <div className="max-w-7xl mx-auto px-4 sm:px-6">
        <div className="flex flex-col gap-4 py-4 md:flex-row md:items-center md:justify-between">
          <Link href="/" className="inline-flex items-center gap-3">
            <span className="grid h-9 w-9 place-items-center rounded-xl bg-slate-950 text-sm font-bold text-white shadow-sm dark:bg-white dark:text-slate-950">
              V
            </span>
            <span className="text-2xl font-bold tracking-tight">Vodka</span>
          </Link>

          <div className="flex flex-wrap items-center justify-center gap-3 text-sm font-medium md:justify-start">
            <Link href="/" className="rounded-full px-3 py-2 transition-colors duration-200 hover:bg-slate-100 hover:text-slate-950 dark:hover:bg-slate-900 dark:hover:text-white">
              Home
            </Link>
            <a href="#features" className="rounded-full px-3 py-2 transition-colors duration-200 hover:bg-slate-100 hover:text-slate-950 dark:hover:bg-slate-900 dark:hover:text-white">
              Features
            </a>
            <a href="#about" className="rounded-full px-3 py-2 transition-colors duration-200 hover:bg-slate-100 hover:text-slate-950 dark:hover:bg-slate-900 dark:hover:text-white">
              About
            </a>
            <Link href="/docs" className="rounded-full px-3 py-2 transition-colors duration-200 hover:bg-slate-100 hover:text-slate-950 dark:hover:bg-slate-900 dark:hover:text-white">
              Documentation
            </Link>
            <a href="https://github.com/DevanshuTripathi/vodka" target="_blank" rel="noopener noreferrer" className="rounded-full px-3 py-2 transition-colors duration-200 hover:bg-slate-100 hover:text-slate-950 dark:hover:bg-slate-900 dark:hover:text-white">
              GitHub
            </a>
            <ThemeToggle />
          </div>
        </div>
      </div>
    </nav>
  );
}
