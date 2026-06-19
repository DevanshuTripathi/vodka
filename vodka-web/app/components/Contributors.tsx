"use client";

import { useEffect, useState } from "react";

export default function Contributors() {
  const [contributors, setContributors] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch(
      "https://api.github.com/repos/DevanshuTripathi/vodka/contributors"
    )
      .then((res) => res.json())
      .then((data) => {
        if (Array.isArray(data)) {
          setContributors(data);
        } else {
          setContributors([]);
        }
      })
      .catch(() => setContributors([]))
      .finally(() => setLoading(false));
  }, []);

  return (
    <section className="bg-slate-100 py-20 sm:py-24">
      <div className="max-w-7xl mx-auto px-6 lg:px-8 text-center">
        <h2 className="text-3xl font-semibold">Contributors</h2>
        <p className="mt-3 text-slate-600">
          People who helped build Vodka
        </p>

        {loading ? (
          <p className="mt-10 text-slate-500">Loading contributors...</p>
        ) : contributors.length === 0 ? (
          <p className="mt-10 text-slate-500">
            No contributors found yet.
          </p>
        ) : (
          <div className="mt-10 grid gap-6 grid-cols-2 sm:grid-cols-3 md:grid-cols-5">
            {contributors.map((user) => (
              <a
                key={user.id}
                href={user.html_url}
                target="_blank"
                rel="noopener noreferrer"
                className="flex flex-col items-center hover:scale-105 transition"
              >
                <img
                  src={user.avatar_url}
                  alt={user.login}
                  className="w-16 h-16 rounded-full"
                />
                <p className="mt-2 text-sm">{user.login}</p>
              </a>
            ))}
          </div>
        )}
      </div>
    </section>
  );
}