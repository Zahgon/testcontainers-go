package testcontainers

import (
	"context"
	"time"
)

// TerminateOptions is a type that holds the options for terminating a container.
type TerminateOptions struct {
	ctx         context.Context
	stopTimeout *time.Duration
	volumes     []string
}

// TerminateOption is a type that represents an option for terminating a container.
type TerminateOption func(*TerminateOptions)

// NewTerminateOptions returns a fully initialised TerminateOptions.
// Defaults: StopTimeout: 10 seconds.
func NewTerminateOptions(ctx context.Context, opts ...TerminateOption) *TerminateOptions {
	_ = "STUB: not implemented"
	return nil
}

// Context returns the context to use during a Terminate.
func (o *TerminateOptions) Context() context.Context {
	_ = "STUB: not implemented"

	// StopTimeout returns the stop timeout to use during a Terminate.
	return *new(context.Context)
}

func (o *TerminateOptions) StopTimeout() *time.Duration { _ = "STUB: not implemented"; return nil }

// Cleanup performs any clean up needed
func (o *TerminateOptions) Cleanup() error {
	_ = "STUB: not implemented"
	// TODO: simplify this when when perform the client refactor.
	return nil
}

// Best effort to remove all volumes.

// StopContext returns a TerminateOption that sets the context.
// Default: context.Background().
func StopContext(ctx context.Context) TerminateOption {
	_ = "STUB: not implemented"
	return *new(TerminateOption)
}

// StopTimeout returns a TerminateOption that sets the timeout.
// Default: See [Container.Stop].
func StopTimeout(timeout time.Duration) TerminateOption {
	_ = "STUB: not implemented"
	return *new(TerminateOption)
}

// RemoveVolumes returns a TerminateOption that sets additional volumes to remove.
// This is useful when the container creates named volumes that should be removed
// which are not removed by default.
// Default: nil.
func RemoveVolumes(volumes ...string) TerminateOption {
	_ = "STUB: not implemented"
	return *new(TerminateOption)
}

// TerminateContainer calls [Container.Terminate] on the container if it is not nil.
//
// This should be called as a defer directly after [GenericContainer](...)
// or a modules Run(...) to ensure the container is terminated when the
// function ends.
func TerminateContainer(container Container, options ...TerminateOption) error {
	_ = "STUB: not implemented"
	return nil
}

// isNil returns true if val is nil or a nil instance false otherwise.
func isNil(val any) bool { _ = "STUB: not implemented"; return false }
