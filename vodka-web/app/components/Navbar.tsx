"use client";

import Link from "next/link";

export default function Navbar() {
  return (
    <nav className="bg-black text-white border-b border-gray-800">
      <div className="max-w-7xl mx-auto px-4">
        <div className="flex justify-between h-16 items-center">
          <Link href="/">
            <span className="text-2xl font-bold">Vodka</span>
          </Link>
          <div className="flex gap-6">
            <Link href="/">
              <span className="hover:text-gray-300 cursor-pointer">Home</span>
            </Link>
            <Link href="/docs">
              <span className="hover:text-gray-300 cursor-pointer">Documentation</span>
            </Link>
            <a href="https://github.com/DevanshuTripathi/vodka" target="_blank" rel="noopener noreferrer" className="hover:text-gray-300">
              GitHub
            </a>
          </div>
        </div>
      </div>
    </nav>
  );
}