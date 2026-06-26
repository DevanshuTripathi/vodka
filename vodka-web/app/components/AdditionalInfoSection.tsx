export default function AdditionalInfoSection() {
  const items = [
  {
    title: "Performance",
    desc: "Vodka uses a Radix Tree router optimized for fast route matching, low memory usage, and efficient middleware execution.",
  },
  {
    title: "Production Build",
    desc: "Build your frontend with Vite and deploy the backend using Vodka for a complete production-ready setup.",
  },
  {
    title: "Philosophy",
    desc: "Fast development workflow, minimal boilerplate, strong developer experience, and practical defaults.",
  },
  {
    title: "Contributing",
    desc: "Contributions, issues, and feature requests are welcome. Feel free to open issues or submit pull requests.",
  },
  {
    title: "License",
    desc: "Vodka is released under the MIT License and is free to use in personal and commercial projects.",
  },
  {
    title: "Community",
    desc: "Join discussions, share feedback, and help improve the Vodka ecosystem.",
  },
  ];

  return (
    <section
      id="additional-info"
      className="border border-slate-200 rounded-xl p-8 my-8 bg-white"
    >
      <h2 className="text-2xl font-bold mb-6">Additional Information</h2>

      <div className="grid md:grid-cols-3 gap-4">
        {items.map((item) => (
          <div
            key={item.title}
            className="border border-slate-200 rounded-xl p-5 bg-slate-50"
          >
            <h3 className="font-semibold text-lg mb-3">{item.title}</h3>

            <p className="text-sm text-slate-600 leading-6">{item.desc}</p>
          </div>
        ))}
      </div>
    </section>
  );
}
