export default function WhyVodkaCards() {
  const cards = [
    {
      icon: "⚡",
      title: "Fast backend iteration",
      desc: "Built-in hot reload watches your .go files and restarts instantly.",
    },
    {
      icon: "⚛️",
      title: "React + Vite integration",
      desc: "Scaffold a full-stack app with frontend and backend together.",
    },
    {
      icon: "🧩",
      title: "Lightweight routing",
      desc: "Radix Tree router delivers fast route matching.",
    },
    {
      icon: "🔐",
      title: "Auth helpers built in",
      desc: "Bearer auth and JWT helpers included out of the box.",
    },
    {
      icon: "✅",
      title: "Struct-tag validation",
      desc: "Declare validation rules and let Vodka handle the rest.",
    },
    {
      icon: "🚀",
      title: "Developer-first defaults",
      desc: "Minimal boilerplate and sensible defaults.",
    },
  ];

  return (
    <div className="grid md:grid-cols-2 gap-4 mt-8">
      {" "}
      {cards.map((card) => (
        <div key={card.title} className="bg-slate-900 rounded-xl p-7">
          {" "}
          <div className="text-2xl mb-3">{card.icon}</div>
          <h3 className="text-lg font-semibold text-white mb-2">
            {card.title}
          </h3>{" "}
          <p className="text-sm text-slate-400 leading-6">{card.desc}</p>{" "}
        </div>
      ))}
    </div>
  );
}
