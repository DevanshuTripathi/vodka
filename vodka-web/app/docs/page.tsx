import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";

async function getReadme(): Promise<string> {
  const res = await fetch(
    "https://api.github.com/repos/DevanshuTripathi/vodka/contents/README.md",
    {
      headers: { Accept: "application/vnd.github.v3.raw" },
      next: { revalidate: 3600 }, // revalidate every hour
    }
  );
  if (!res.ok) return "# Failed to load documentation. Please try again later.";
  return res.text();
}

export const metadata = {
  title: "Docs — Vodka",
  description: "Vodka framework documentation, synced from the GitHub repository.",
};

export default async function DocsPage() {
  const content = await getReadme();

  return (
    <div className="flex gap-8">
      {/* Sidebar */}
      <aside className="hidden lg:block w-48 shrink-0">
        <div className="sticky top-20 flex flex-col gap-1 text-sm text-zinc-400">
          <p className="text-xs uppercase tracking-widest text-zinc-600 mb-2 font-semibold">On this page</p>
          {["Getting Started", "Routing", "Middleware", "WebSocket", "SSE", "Mixers"].map((item) => (
            <a
              key={item}
              href={`#${item.toLowerCase().replace(/\s/g, "-")}`}
              className="hover:text-white transition-colors py-0.5"
            >
              {item}
            </a>
          ))}
        </div>
      </aside>

      {/* Markdown content */}
      <article className="prose prose-invert prose-zinc max-w-3xl w-full
        prose-headings:font-bold prose-headings:tracking-tight
        prose-code:bg-zinc-800 prose-code:px-1.5 prose-code:py-0.5 prose-code:rounded prose-code:text-zinc-200 prose-code:before:content-none prose-code:after:content-none
        prose-pre:bg-zinc-900 prose-pre:border prose-pre:border-zinc-800
        prose-a:text-blue-400 prose-a:no-underline hover:prose-a:underline
        prose-blockquote:border-zinc-700 prose-blockquote:text-zinc-400
        prose-hr:border-zinc-800
        prose-img:rounded-xl">
        <ReactMarkdown remarkPlugins={[remarkGfm]}>{content}</ReactMarkdown>
      </article>
    </div>
  );
}
