import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";

type Release = {
  id: number;
  tag_name: string;
  name: string;
  published_at: string;
  body: string;
  html_url: string;
};

async function getReleases(): Promise<Release[]> {
  const res = await fetch(
    "https://api.github.com/repos/DevanshuTripathi/vodka/releases",
    {
      headers: { Accept: "application/vnd.github+json" },
      next: { revalidate: 3600 },
    }
  );
  if (!res.ok) return [];
  return res.json();
}

export const metadata = {
  title: "Releases — Vodka",
  description: "Vodka framework release history, synced from GitHub.",
};

export default async function ReleasesPage() {
  const releases = await getReleases();

  return (
    <div className="max-w-3xl mx-auto flex flex-col gap-10">
      <h1 className="text-3xl font-bold text-white">Releases</h1>

      {releases.length === 0 ? (
        <p className="text-zinc-400">No releases yet. Check back soon.</p>
      ) : (
        releases.map((release) => (
          <div key={release.id} className="border border-zinc-800 rounded-xl p-6 flex flex-col gap-4">
            <div className="flex items-center justify-between flex-wrap gap-2">
              <div className="flex items-center gap-3">
                <span className="bg-zinc-800 text-zinc-200 text-sm font-mono px-3 py-1 rounded-full">
                  {release.tag_name}
                </span>
                <h2 className="text-lg font-semibold text-white">{release.name || release.tag_name}</h2>
              </div>
              <span className="text-xs text-zinc-500">
                {new Date(release.published_at).toLocaleDateString("en-US", {
                  year: "numeric", month: "long", day: "numeric",
                })}
              </span>
            </div>

            {release.body && (
              <article className="prose prose-invert prose-zinc prose-sm max-w-none
                prose-code:bg-zinc-800 prose-code:px-1.5 prose-code:py-0.5 prose-code:rounded prose-code:before:content-none prose-code:after:content-none
                prose-a:text-blue-400 prose-hr:border-zinc-800">
                <ReactMarkdown remarkPlugins={[remarkGfm]}>{release.body}</ReactMarkdown>
              </article>
            )}

            <a
              href={release.html_url}
              target="_blank"
              rel="noopener noreferrer"
              className="text-sm text-blue-400 hover:underline self-start"
            >
              View on GitHub →
            </a>
          </div>
        ))
      )}
    </div>
  );
}
