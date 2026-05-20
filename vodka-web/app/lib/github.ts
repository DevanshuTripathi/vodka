// lib/github.ts
const GITHUB_OWNER = "DevanshuTripathi";
const GITHUB_REPO = "vodka";
const GITHUB_API_BASE = "https://api.github.com";

export async function fetchREADME(): Promise<string> {
  try {
    const response = await fetch(
      `${GITHUB_API_BASE}/repos/${GITHUB_OWNER}/${GITHUB_REPO}/contents/README.md`,
      {
        headers: {
          "Accept": "application/vnd.github.v3.raw",
        },
      }
    );

    if (!response.ok) {
      throw new Error(`Failed to fetch README: ${response.statusText}`);
    }

    return await response.text();
  } catch (error) {
    console.error("Error fetching README:", error);
    return "# Documentation\n\nFailed to load documentation.";
  }
}