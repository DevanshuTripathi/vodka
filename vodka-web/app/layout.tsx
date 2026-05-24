import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Vodka",
  description: "A fast, focused Go web framework for modern full-stack development.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className="h-full antialiased">
      <body className="min-h-full flex flex-col">{children}</body>
    </html>
  );
}
