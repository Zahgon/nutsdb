package core

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

// ComponentLifecycle provides shared lifecycle management for components.
// It encapsulates common patterns: context management, goroutine tracking,
// and idempotent stop behavior.
//
// Usage:
//
//	type MyComponent struct {
//	    lifecycle ComponentLifecycle
//	    // ... business fields ...
//	}
//
//	func (c *MyComponent) Start(ctx context.Context) error {
//	    if err := c.lifecycle.Start(ctx); err != nil {
//	        return err
//	    }
//	    c.lifecycle.Go(func(ctx context.Context) {
//	        // worker goroutine
//	    })
//	    return nil
//	}
//
//	func (c *MyComponent) Stop(timeout time.Duration) error {
//	    return c.lifecycle.Stop(timeout)
//	}
type ComponentLifecycle struct {
	// ctx is the component's context, derived from parent context
	ctx context.Context

	// cancel cancels the component's context
	cancel context.CancelFunc

	// wg tracks all goroutines started via Go()
	wg sync.WaitGroup

	// started indicates if Start() has been called
	started atomic.Bool

	// stopped indicates if Stop() has been called
	stopped atomic.Bool

	// mu protects state transitions
	mu sync.Mutex
}

// ComponentState represents the lifecycle state of a component
type ComponentState int

const (
	ComponentStateCreated ComponentState = iota
	ComponentStateRunning
	ComponentStateStopped
)

// ErrAlreadyStarted is returned when Start() is called on an already started component
var ErrAlreadyStarted = error(&componentError{msg: "component already started"})

// ErrAlreadyStopped is returned when Stop() is called on an already stopped component
var ErrAlreadyStopped = error(&componentError{msg: "component already stopped"})

// ErrStopTimeout is returned when Stop() times out waiting for goroutines
var ErrStopTimeout = error(&componentError{msg: "component stop timeout"})

type componentError struct {
	msg string
}

func (e *componentError) Error() string {
	_ = "STUB: not implemented"

	// Start initializes the component lifecycle with the given parent context.
	// Returns ErrAlreadyStarted if the component has already been started.
	// Returns ErrAlreadyStopped if the component has already been stopped.
	return ""
}

func (cl *ComponentLifecycle) Start(parentCtx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Check stopped first to handle start-after-stop case

// Create child context from parent

// Stop stops the component and waits for all goroutines to finish.
// It is idempotent - calling Stop() multiple times is safe.
// Returns ErrStopTimeout if goroutines don't finish within the timeout.
func (cl *ComponentLifecycle) Stop(timeout time.Duration) error {
	_ = "STUB: not implemented"

	// Idempotent check - if already stopped, return nil
	return nil
}

// Mark as stopped

// Cancel context to signal goroutines

// Wait for all goroutines with timeout

// Go starts a goroutine managed by the lifecycle.
// The goroutine will receive the component's context and should respect cancellation.
// Panics in the goroutine are recovered and logged.
func (cl *ComponentLifecycle) Go(fn func(ctx context.Context)) { _ = "STUB: not implemented"; return }

// Context returns the component's context.
// This context is canceled when Stop() is called or when the parent context is canceled.
func (cl *ComponentLifecycle) Context() context.Context {
	_ = "STUB: not implemented"

	// IsRunning returns true if the component is currently running.
	return *new(context.Context)
}

func (cl *ComponentLifecycle) IsRunning() bool { _ = "STUB: not implemented"; return false }

// IsStopped returns true if the component has been stopped.
func (cl *ComponentLifecycle) IsStopped() bool { _ = "STUB: not implemented"; return false }

// GetState returns the current state of the component.
func (cl *ComponentLifecycle) GetState() ComponentState {
	_ = "STUB: not implemented"
	return *new(ComponentState)
}
