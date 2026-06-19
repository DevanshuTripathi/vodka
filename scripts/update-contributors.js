const fs = require("fs");

const OWNER = "DevanshuTripathi";
const REPO = "vodka";

async function main() {
    const res = await fetch(
        `https://api.github.com/repos/${OWNER}/${REPO}/contributors`
    );

    const contributors = await res.json();

    const html = contributors
        .map(
            (c) => `
<a href="${c.html_url}">
  <img src="${c.avatar_url}" width="80px;" alt="${c.login}"/><br/>
  <sub><b>${c.login}</b></sub>
</a>
`
        )
        .join("\n");

    const readme = fs.readFileSync("README.md", "utf8");

    const updated = readme.replace(
        /<!-- CONTRIBUTORS:START -->([\s\S]*?)<!-- CONTRIBUTORS:END -->/,
        `<!-- CONTRIBUTORS:START -->\n${html}\n<!-- CONTRIBUTORS:END -->`
    );

    fs.writeFileSync("README.md", updated);
}

main().catch(console.error);