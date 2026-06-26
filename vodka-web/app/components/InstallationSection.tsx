export default function InstallationSection() {
  return (
    <section
      id="installation"
      className="bg-white border border-slate-200 rounded-xl p-8 mt-6 mb-6"
    >
      <h2 className="text-xl font-bold mb-6">Installation</h2>

      <h3 className="text-sm font-semibold mb-2">Prerequisites</h3>

      <p className="text-slate-600 mb-4">
        Make sure the following tools are installed before using Vodka:
      </p>

      <ul className="list-disc pl-6 text-slate-600 mb-6">
        <li>Go 1.24 or newer</li>
        <li>Node.js 20.19 or newer</li>
        <li>npm</li>
      </ul>

      <pre className="bg-slate-950 text-white rounded-md p-4 overflow-x-auto mb-6">
        {`go version
node -v
npm -v`}
      </pre>

      <h3 className="text-lg font-semibold mb-3">Install the Vodka CLI</h3>

      <pre className="bg-slate-950 text-white rounded-md p-4 overflow-x-auto mb-6">
        {`go install github.com/DevanshuTripathi/vodka/cmd/vodka@latest`}
      </pre>

      <p className="text-slate-600 mb-6">
        Make sure your Go bin directory is added to your system PATH.
      </p>

      <h3 className="text-lg font-semibold mb-3">Linux / macOS</h3>

      <pre className="bg-slate-950 text-white rounded-md p-4 overflow-x-auto mb-6">
        {`export PATH=$PATH:$(go env GOPATH)/bin`}
      </pre>

      <h3 className="text-lg font-semibold mb-3">Windows</h3>

      <p className="text-slate-600 mb-3">
        Add this directory to your Environment Variables:
      </p>

      <pre className="bg-slate-950 text-white rounded-md p-4 overflow-x-auto">
        {`%USERPROFILE%\\go\\bin`}
      </pre>
    </section>
  );
}
