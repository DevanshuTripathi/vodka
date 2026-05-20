# Vodka Documentation Website

A Next.js-based documentation website for the Vodka framework. The site dynamically fetches and displays documentation from the main Vodka repository.

## Features

- 📚 **Dynamic Content** - Fetches README and documentation from GitHub API
- 🎨 **Modern UI** - Built with Next.js, React, and Tailwind CSS
- 📱 **Responsive Design** - Works perfectly on mobile, tablet, and desktop
- 🌙 **Dark Mode Support** - Beautiful dark mode styling
- ⚡ **Fast Performance** - Server-side rendering for optimal SEO
- 🔄 **Auto-updated** - Content syncs with main repository

## Getting Started

### Prerequisites
- Node.js 18+
- npm or yarn

### Installation

```bash
cd vodka-web
npm install
```

### Development

```bash
npm run dev
```

Visit `http://localhost:3000` in your browser.

### Build

```bash
npm run build
npm start
```

## Project Structure

vodka-web/
├── app/
│   ├── components/        # React components
│   ├── lib/              # Helper functions
│   ├── docs/             # Documentation page
│   ├── page.tsx          # Home page
│   └── layout.tsx        # Root layout
├── public/               # Static assets
└── package.json

## Technologies Used

- **Next.js 16** - React framework
- **TypeScript** - Type safety
- **Tailwind CSS** - Styling
- **React Markdown** - Markdown rendering
- **GitHub API** - Dynamic content fetching

## Future Enhancements

- [ ] Playground integration with Vodka backend
- [ ] Full-text search
- [ ] Dark mode toggle
- [ ] Multiple documentation versions
- [ ] Community contributions page

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT - See LICENSE in the main Vodka repository