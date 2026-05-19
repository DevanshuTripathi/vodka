package vodka

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ShutdownConfig struct {
	Timeout time.Duration
}

func (e *Engine) RunGracefully(addr string, cfg ShutdownConfig) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)
	return e.runGracefully(addr, cfg, quit)
}

func (e *Engine) runGracefully(addr string, cfg ShutdownConfig, quit <-chan os.Signal) error {
	if addr == "" {
		addr = ":8080"
	}
	if cfg.Timeout == 0 {
		cfg.Timeout = 10 * time.Second
	}

	srv := &http.Server{
		Addr:    addr,
		Handler: e,
	}

	serverErr := make(chan error, 1)
	go func() {
		log.Printf(Green+"Pouring Vodka on %s\n"+Reset, addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErr <- err
		}
	}()

	select {
	case err := <-serverErr:
		return fmt.Errorf("server error: %w", err)
	case sig := <-quit:
		log.Printf(Yellow+"Received signal %s — draining in-flight requests (timeout: %s)...\n"+Reset, sig, cfg.Timeout)
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("graceful shutdown timed out: %w", err)
	}

	log.Println(Green + "Server exited cleanly." + Reset)
	return nil
}
