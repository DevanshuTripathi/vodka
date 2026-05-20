"use client";

import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";

interface MarkdownRendererProps {
  content: string;
}

export default function MarkdownRenderer({ content }: MarkdownRendererProps) {
  return (
    <div className="prose prose-sm dark:prose-invert max-w-none p-8">
      <ReactMarkdown
        remarkPlugins={[remarkGfm]}
        components={{
          h1: ({ node, ...props }) => (
            <h1 className="text-4xl font-bold mt-8 mb-4" {...props} />
          ),
          h2: ({ node, ...props }) => (
            <h2 className="text-3xl font-bold mt-6 mb-3" {...props} />
          ),
          h3: ({ node, ...props }) => (
            <h3 className="text-2xl font-bold mt-4 mb-2" {...props} />
          ),
          code: ({ node, inline, ...props }) => (
            inline ? (
              <code className="bg-gray-200 dark:bg-gray-800 px-2 py-1 rounded" {...props} />
            ) : (
              <pre className="bg-gray-900 text-gray-100 p-4 rounded-lg overflow-x-auto">
                <code {...props} />
              </pre>
            )
          ),
          a: ({ node, ...props }) => (
            <a className="text-blue-600 dark:text-blue-400 hover:underline" {...props} />
          ),
        }}
      >
        {content}
      </ReactMarkdown>
    </div>
  );
}