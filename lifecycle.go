package vodka

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

// StartHook represents a function executed when the application starts.
type StartHook func() error

// ShutdownHook represents a function executed when the application shuts down.
type ShutdownHook func(context.Context) error

type lifecycleHook struct {
	priority int
	order    int
	fn       ShutdownHook
}

// LifecycleManager manages startup and shutdown hooks.
type LifecycleManager struct {
	startupHooks  []StartHook
	shutdownHooks []lifecycleHook
	timeout       time.Duration
	mu            sync.Mutex
}

// NewLifecycleManager creates a new LifecycleManager with a default timeout.
func NewLifecycleManager() *LifecycleManager {
	return &LifecycleManager{
		startupHooks:  make([]StartHook, 0),
		shutdownHooks: make([]lifecycleHook, 0),
		timeout:       30 * time.Second,
	}
}

// RegisterStart registers a new startup hook.
func (lm *LifecycleManager) RegisterStart(fn StartHook) {
	lm.mu.Lock()
	defer lm.mu.Unlock()
	lm.startupHooks = append(lm.startupHooks, fn)
}

// RegisterShutdown registers a new shutdown hook with the given priority.
func (lm *LifecycleManager) RegisterShutdown(priority int, fn ShutdownHook) {
	lm.mu.Lock()
	defer lm.mu.Unlock()
	order := len(lm.shutdownHooks)
	lm.shutdownHooks = append(lm.shutdownHooks, lifecycleHook{
		priority: priority,
		order:    order,
		fn:       fn,
	})
}

// SetTimeout configures the maximum duration allowed for shutdown.
func (lm *LifecycleManager) SetTimeout(timeout time.Duration) {
	lm.mu.Lock()
	defer lm.mu.Unlock()
	lm.timeout = timeout
}

// runStartupHooks runs all startup hooks in registration order.
func (lm *LifecycleManager) runStartupHooks() error {
	lm.mu.Lock()
	hooks := make([]StartHook, len(lm.startupHooks))
	copy(hooks, lm.startupHooks)
	lm.mu.Unlock()

	for _, hook := range hooks {
		if err := hook(); err != nil {
			return err
		}
	}
	return nil
}

// ShutdownError represents a collection of errors encountered during shutdown.
type ShutdownError struct {
	Errors []error
}

// Error formats the aggregated errors.
func (e *ShutdownError) Error() string {
	if len(e.Errors) == 0 {
		return ""
	}
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%d shutdown errors:\n", len(e.Errors)))
	for i, err := range e.Errors {
		sb.WriteString(fmt.Sprintf("- %v", err))
		if i < len(e.Errors)-1 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

// Unwrap supports Go 1.20+ multi-error unwrapping.
func (e *ShutdownError) Unwrap() []error {
	return e.Errors
}
