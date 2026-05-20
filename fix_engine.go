//go:build ignore

package main

import (
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("engine.go")
	s := string(data)

	// 1. Remove BaseContext from RunConfig
	s = strings.ReplaceAll(s,
		"        GracefulTimeout  time.Duration // max time to wait for in-flight requests; defaults to 10s\n        BaseContext     context.Context // cancelling this context triggers graceful shutdown (useful in tests)\n}",
		"        GracefulTimeout  time.Duration // max time to wait for in-flight requests; defaults to 10s\n}")

	// 2. Fix GracefulShutdown comment
	s = strings.ReplaceAll(s,
		"        GracefulShutdown bool          // if true, enables graceful shutdown on SIGINT/SIGTERM",
		"        GracefulShutdown bool          // set to true to wait for active requests before closing")

	// 3. Fix stale comment above Run
	s = strings.ReplaceAll(s,
		"//      app.Run(\":8080\", vodka.RunConfig{GracefulTimeout: 10 * time.Second})\nfunc (e *Engine) Run",
		"// Run starts the HTTP server on the given address.\n// Without a RunConfig it works exactly as before.\n// Pass RunConfig{GracefulShutdown: true} to handle Ctrl+C and SIGTERM cleanly.\nfunc (e *Engine) Run")

	// 4. Fix stale internal comment
	s = strings.ReplaceAll(s,
		"        // No config or zero timeout → original behavior, no changes",
		"        // No config or graceful shutdown not requested — plain ListenAndServe")

	// 5. Remove baseCtx block
	s = strings.ReplaceAll(s,
		"        baseCtx := cfg.BaseContext\n        if baseCtx == nil {\n                baseCtx = context.Background()\n        }\n",
		"")

	// 6. Add comment above e.srv
	s = strings.ReplaceAll(s,
		"        e.srv = &http.Server{",
		"        // Store the server on the engine so Shutdown() can reach it later\n        e.srv = &http.Server{")

	// 7. Fix shutdownCh comment
	s = strings.ReplaceAll(s,
		"        e.shutdownCh = make(chan struct{})\n        // Register for OS shutdown signals (Ctrl+C and Docker/k8s stop)",
		"        // shutdownCh lets Shutdown() tell Run() to start draining without needing an OS signal\n        e.shutdownCh = make(chan struct{})\n\n        // Listen for Ctrl+C (SIGINT) or container stop (SIGTERM)")

	// 8. Fix select block comment and remove baseCtx.Done case
	s = strings.ReplaceAll(s,
		"        // Block until OS signal, context cancellation, or server error\n        select {\n        case err := <-serverErr:\n                return err\n        case <-e.shutdownCh:\n                log.Printf(Yellow+\"[Vodka] Shutdown triggered — draining requests...\\n\"+Reset)\n        case sig := <-quit:\n                log.Printf(Yellow+\"[Vodka] Signal received: %v — draining requests...\\n\"+Reset, sig)\n        case <-baseCtx.Done():\n                log.Printf(Yellow+\"[Vodka] Shutdown triggered — draining requests...\\n\"+Reset)\n        }",
		"        // Wait here until something tells us to shut down\n        select {\n        case err := <-serverErr:\n                // Server failed to start or crashed unexpectedly\n                return err\n        case <-e.shutdownCh:\n                // Shutdown() was called directly, e.g. from a test\n                log.Printf(Yellow+\"[Vodka] Shutdown triggered — draining requests...\\n\"+Reset)\n        case sig := <-quit:\n                // OS sent a stop signal (Ctrl+C or docker stop)\n                log.Printf(Yellow+\"[Vodka] Signal received: %v — draining requests...\\n\"+Reset, sig)\n        }")

	// 9. Add comment above shutdownCtx
	s = strings.ReplaceAll(s,
		"        shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.GracefulTimeout)",
		"        // Give active requests up to GracefulTimeout to finish before force-closing\n        shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.GracefulTimeout)")

	// 10. Fix duplicate Shutdown() comment
	s = strings.ReplaceAll(s,
		"// Shutdown gracefully stops the server, waiting for active requests to finish.\n// The context controls how long to wait before force-closing.\n// Useful in tests to trigger shutdown without sending OS signals.\n// Shutdown gracefully stops the server, waiting for active requests to finish.\n// The context controls how long to wait before force-closing.\n// Useful in tests to trigger shutdown without sending OS signals.\nfunc (e *Engine) Shutdown",
		"// Shutdown gracefully stops the running server.\n// It signals Run() to stop accepting new requests and waits for active ones to finish.\n// The context sets the deadline — if requests take too long, the server force-closes.\nfunc (e *Engine) Shutdown")

	os.WriteFile("engine.go", []byte(s), 0644)
}
