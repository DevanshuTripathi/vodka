const fs = require("fs");

const OWNER = "DevanshuTripathi";
const REPO = "vodka";

async function getJSON(url) {
    const res = await fetch(url);
    return res.json();
}

async function main() {
    const contributors = await getJSON(
        `https://api.github.com/repos/${OWNER}/${REPO}/contributors`
    );

    const users = await Promise.all(
        contributors.map(async (c) => {
            const u = await getJSON(`https://api.github.com/users/${c.login}`);

            return {
                login: c.login,
                avatar: c.avatar_url,
                url: c.html_url,
                bio: u.bio || "GitHub Contributor",
                followers: u.followers,
                repos: u.public_repos,
            };
        })
    );

    const html = `
<div class="contributors-grid">
${users
            .map(
                (u) => `
  <a class="card" href="${u.url}" target="_blank">
    <img src="${u.avatar}" />
    <div class="name">${u.login}</div>
    <div class="bio">${u.bio}</div>

    <div class="stats">
      <span>⭐ ${u.repos}</span>
      <span>👥 ${u.followers}</span>
    </div>
  </a>
`
            )
            .join("")}
</div>
`;

    const readme = fs.readFileSync("README.md", "utf8");

    const updated = readme.replace(
        /<!-- CONTRIBUTORS:START -->([\s\S]*?)<!-- CONTRIBUTORS:END -->/,
        `<!-- CONTRIBUTORS:START -->\n${html}\n<!-- CONTRIBUTORS:END -->`
    );

    fs.writeFileSync("README.md", updated);
}

main().catch(console.error);