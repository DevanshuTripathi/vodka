import { fetchREADME } from "../lib/github";
import MarkdownRenderer from "../components/MarkdownRenderer";
import Navbar from "../components/Navbar";

export const revalidate = 3600;

export default async function DocsPage() {
  const readmeContent = await fetchREADME();

  return (
    <div>
      <Navbar />
      <div className="min-h-screen bg-slate-50 text-slate-950 transition-colors dark:bg-slate-950 dark:text-slate-100">
        <main className="flex-1 p-8 max-w-4xl mx-auto">
          <MarkdownRenderer content={readmeContent} />
        </main>
      </div>
    </div>
  );
}
