import Link from "next/link";

export default function Navbar() {
  return (
    <nav className="border-b border-zinc-800 bg-zinc-950/80 backdrop-blur sticky top-0 z-50">
      <div className="max-w-6xl mx-auto px-4 h-14 flex items-center justify-between">
        <Link href="/" className="text-xl font-bold tracking-tight text-white">
          🍸 Vodka
        </Link>
        <div className="flex items-center gap-6 text-sm text-zinc-400">
          <Link href="/docs" className="hover:text-white transition-colors">Docs</Link>
          <Link href="/releases" className="hover:text-white transition-colors">Releases</Link>
          <a
            href="https://github.com/DevanshuTripathi/vodka"
            target="_blank"
            rel="noopener noreferrer"
            className="hover:text-white transition-colors"
          >
            GitHub
          </a>
        </div>
      </div>
    </nav>
  );
}
