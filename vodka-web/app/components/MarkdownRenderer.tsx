"use client";

import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";

interface MarkdownRendererProps {
  content: string;
}

export default function MarkdownRenderer({ content }: MarkdownRendererProps) {
  return (
    <div className="prose prose-slate prose-sm max-w-none rounded-3xl border border-slate-200 bg-white p-8 shadow-sm shadow-slate-900/5 dark:prose-invert dark:border-slate-800 dark:bg-slate-900/60">
      <ReactMarkdown
        remarkPlugins={[remarkGfm]}
        components={{
          h1: ({ ...props }) => (
            <h1 className="text-4xl font-bold mt-8 mb-4" {...props} />
          ),
          h2: ({ ...props }) => (
            <h2 className="text-3xl font-bold mt-6 mb-3" {...props} />
          ),
          h3: ({ ...props }) => (
            <h3 className="text-2xl font-bold mt-4 mb-2" {...props} />
          ),
          code: (props) => {
            const { inline } = props as { inline?: boolean };
            return inline ? (
              <code className="rounded bg-slate-100 px-2 py-1 dark:bg-slate-800" {...props} />
            ) : (
              <pre className="overflow-x-auto rounded-2xl bg-slate-950 p-4 text-slate-100">
                <code {...props} />
              </pre>
            );
          },
          a: ({ ...props }) => (
            <a className="text-blue-600 dark:text-blue-400 hover:underline" {...props} />
          ),
        }}
      >
        {content}
      </ReactMarkdown>
    </div>
  );
}
