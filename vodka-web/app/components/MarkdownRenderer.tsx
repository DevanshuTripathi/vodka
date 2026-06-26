"use client";

import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";

interface MarkdownRendererProps {
  content: string;
}

export default function MarkdownRenderer({ content }: MarkdownRendererProps) {
  return (
    <div className="prose prose-sm dark:prose-invert max-w-none p-4">
      <ReactMarkdown
        remarkPlugins={[remarkGfm]}
        components={{
          h1: ({ children, ...props }) => {
            const text = String(children)
              .toLowerCase()
              .replace(/[^\w\s]/g, "")
              .replace(/\s+/g, "-");

            return (
              <h1 id={text} className="text-2xl font-extrabold mb-6" {...props}>
                {children}
              </h1>
            );
          },
          h2: ({ children, ...props }) => {
            const text = String(children)
              .toLowerCase()
              .replace(/[^\w\s]/g, "")
              .replace(/\s+/g, "-");

            return (
              <h2
                id={text}
                className="text-xl font-semibold mt-8 mb-3"
                {...props}
              >
                {children}
              </h2>
            );
          },
          h3: ({ ...props }) => (
            <h3 className="text-xl font-semibold mt-6 mb-3" {...props} />
          ),
          hr: () => null,
          pre: ({ ...props }) => (
            <pre
              className="bg-slate-900 border border-slate-800 text-gray-100 p-3 rounded-xl overflow-x-auto my-6"
              {...props}
            />
          ),
          code: ({ className, children, ...props }) => (
            <code
              className={`${className ?? ""} text-white`}
              {...props}
            >
              {children}
            </code>
          ),
          a: ({ ...props }) => (
            <a
              className="text-blue-600 dark:text-blue-400 hover:underline"
              {...props}
            />
          ),
          img: ({ ...props }) => (
            <img
              className="rounded-xl my-6 max-h-[450px] w-auto mx-auto border border-slate-800"
              {...props}
            />
          ),
          table: ({ children, ...props }) => {
            console.log("TABLE DATA:", children);

            if (Array.isArray(children)) {
              console.log("THEAD:", children[0]);
              console.log("TBODY:", children[1]);
            }

            return (
              <div className="overflow-x-auto my-6">
                <table
                  className="w-full border-collapse border border-slate-800 rounded-xl overflow-hidden"
                  {...props}
                >
                  {children}
                </table>
              </div>
            );
          },

          thead: ({ ...props }) => (
            <thead className="bg-slate-900" {...props} />
          ),

          th: ({ ...props }) => (
            <th
              className="border border-slate-800 px-4 py-3 text-left font-semibold"
              {...props}
            />
          ),

          td: ({ ...props }) => (
            <td className="border border-slate-800 px-4 py-3" {...props} />
          ),
        }}
      >
        {content}
      </ReactMarkdown>
    </div>
  );
}