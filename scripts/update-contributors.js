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
        <table>
        <tr>
            ${users
            .map(
                (u) => `
            <td align="center" width="140">
                <a href="${u.url}">
                <img src="${u.avatar}" width="80" height="80"><br/>
                <b>${u.login}</b><br/>
                <sub>${u.bio ? u.bio.slice(0, 20) : "Contributor"}</sub><br/>
                <sub>⭐ ${u.repos} | 👥 ${u.followers}</sub>
                </a>
            </td>
            `
            )
            .join("")}
        </tr>
        </table>
        `;

    const readme = fs.readFileSync("README.md", "utf8");

    const updated = readme.replace(
        /<!-- CONTRIBUTORS:START -->([\s\S]*?)<!-- CONTRIBUTORS:END -->/,
        `<!-- CONTRIBUTORS:START -->\n${html}\n<!-- CONTRIBUTORS:END -->`
    );

    fs.writeFileSync("README.md", updated);
}

main().catch(console.error);