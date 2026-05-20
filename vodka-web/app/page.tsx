import Link from "next/link";
import Navbar from "./components/Navbar";

export default function Home() {
  return (
    <div>
      <Navbar />
      <section className="bg-white text-black py-24">
        <div className="max-w-4xl mx-auto px-4 text-center">
          <h1 className="text-6xl font-bold mb-6">Vodka</h1>
          <p className="text-xl mb-8">A modern Go web framework focused on developer experience, full-stack workflow, and rapid iteration.</p>
          <div className="space-x-4">
            <Link href="/docs" className="inline-block bg-black text-white px-8 py-3 rounded-lg font-semibold hover:bg-gray-800">
              Get Started
            </Link>
            <a href="https://github.com/DevanshuTripathi/vodka" target="_blank" rel="noopener noreferrer" className="inline-block border-2 border-black text-black px-8 py-3 rounded-lg font-semibold hover:bg-black hover:text-white">
              View on GitHub
            </a>
          </div>
        </div>
      </section>

      <section className="py-16 bg-black">
  <div className="max-w-4xl mx-auto px-4">
    <h2 className="text-4xl font-bold mb-12 text-center text-white">Features</h2>
    <div className="grid md:grid-cols-3 gap-8">
      <div className="p-6 rounded-lg bg-black border border-gray-800">
        <h3 className="text-xl font-semibold mb-2 text-white">Fast Routing</h3>
        <p className="text-gray-300">Radix tree based router for zero-allocation route matching</p>
      </div>
      <div className="p-6 rounded-lg bg-black border border-gray-800">
        <h3 className="text-xl font-semibold mb-2 text-white">Middleware Chaining</h3>
        <p className="text-gray-300">Composable middleware with abort support</p>
      </div>
      <div className="p-6 rounded-lg bg-black border border-gray-800">
        <h3 className="text-xl font-semibold mb-2 text-white">Authentication</h3>
        <p className="text-gray-300">Built-in JWT validation helpers and Bearer auth</p>
      </div>
      <div className="p-6 rounded-lg bg-black border border-gray-800">
        <h3 className="text-xl font-semibold mb-2 text-white">Request Validation</h3>
        <p className="text-gray-300">Support for request validation using struct tags</p>
      </div>
      <div className="p-6 rounded-lg bg-black border border-gray-800">
        <h3 className="text-xl font-semibold mb-2 text-white">React + Vite Integration</h3>
        <p className="text-gray-300">Full-stack scaffolding with frontend and backend</p>
      </div>
      <div className="p-6 rounded-lg bg-black border border-gray-800">
        <h3 className="text-xl font-semibold mb-2 text-white">SPA Support</h3>
        <p className="text-gray-300">Seamless single page application serving in production</p>
      </div>
    </div>
  </div>
</section>
    </div>
  );
}