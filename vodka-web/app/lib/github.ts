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
          Accept: "application/vnd.github.v3.raw",
        },
      },
    );

    if (!response.ok) {
      throw new Error(`Failed to fetch README: ${response.statusText}`);
    }
    const content = await response.text();

    let cleanedContent = content.replace(/# Table of Contents[\s\S]*?---/, "");

    cleanedContent = cleanedContent.replace(
      /# Features[\s\S]*?# Why Vodka\?/,
      "# Why Vodka?",
    );
    cleanedContent = cleanedContent.replace(
      /# Why Vodka\?[\s\S]*?# Prerequisites/,
      "# Prerequisites",
    );

    return cleanedContent
      .replaceAll(
        "./assets/demo.gif",
        "https://raw.githubusercontent.com/DevanshuTripathi/vodka/main/assets/demo.gif",
      )
      .replaceAll(
        "./assets/workflow.gif",
        "https://raw.githubusercontent.com/DevanshuTripathi/vodka/main/assets/workflow.gif",
      );
  } catch (error) {
    console.error("Error fetching README:", error);
    return "# Documentation\n\nFailed to load documentation.";
  }
}

export async function fetchREADMESections() {
  const content = await fetchREADME();

  const installationStart = content.indexOf("# Prerequisites");
  const quickStartStart = content.indexOf("# Quick Start");

  const installationContent = content.slice(installationStart, quickStartStart);

  const remainingContent =
    content.slice(0, installationStart) + content.slice(quickStartStart);

  return {
    installationContent,
    remainingContent,
  };
}