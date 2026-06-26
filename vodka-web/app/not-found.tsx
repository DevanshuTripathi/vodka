"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";

export default function NotFound() {
  const pathname = usePathname();

  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-slate-100 text-slate-900 px-6 text-center">

      {/* Main 404 */}
      <h1 className="mt-4 text-7xl font-bold">
        404
      </h1>

      <p className="mt-4 text-2xl font-semibold">
        Page not found
      </p>

      <p className="mt-3 text-slate-600 max-w-md">
        The page you are looking for doesn’t exist or has been moved.
      </p>

      {/* Requested text */}
      <p className="mt-6 text-sm text-slate-500">
        Requested: <span className="font-mono">{pathname}</span>
      </p>

      {/* Buttons */}
      <div className="mt-8 flex gap-4">
        <Link
          href="/"
          className="inline-flex items-center justify-center rounded-2xl bg-slate-900 px-6 py-3 text-white text-sm font-semibold hover:bg-slate-700 transition"
        >
          Go Home
        </Link>

        <button
          onClick={() => window.history.back()}
          className="inline-flex items-center justify-center rounded-2xl border border-slate-900 px-6 py-3 text-sm font-semibold hover:bg-slate-900 hover:text-white transition"
        >
          Go Back
        </button>
      </div>
    </div>
  );
}