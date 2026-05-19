package wait

import (
	"context"
	"time"
)

// Implement interface
var (
	_ Strategy        = (*MultiStrategy)(nil)
	_ StrategyTimeout = (*MultiStrategy)(nil)
)

type MultiStrategy struct {
	// all Strategies should have a startupTimeout to avoid waiting infinitely
	timeout  *time.Duration
	deadline *time.Duration

	// additional properties
	Strategies []Strategy
}

// WithStartupTimeoutDefault sets the default timeout for all inner wait strategies
func (ms *MultiStrategy) WithStartupTimeoutDefault(timeout time.Duration) *MultiStrategy {
	_ = "STUB: not implemented"
	return nil
}

// WithStartupTimeout sets a time.Duration which limits all wait strategies
//
// Deprecated: use WithDeadline
func (ms *MultiStrategy) WithStartupTimeout(timeout time.Duration) Strategy {
	_ = "STUB: not implemented"
	return *new(Strategy)
}

// WithDeadline sets a time.Duration which limits all wait strategies
func (ms *MultiStrategy) WithDeadline(deadline time.Duration) *MultiStrategy {
	_ = "STUB: not implemented"
	return nil
}

func ForAll(strategies ...Strategy) *MultiStrategy { _ = "STUB: not implemented"; return nil }

func (ms *MultiStrategy) Timeout() *time.Duration {
	_ = "STUB: not implemented"

	// String returns a human-readable description of the wait strategy.
	return nil
}

func (ms *MultiStrategy) String() string { _ = "STUB: not implemented"; return "" }

// Always include "all of:" prefix to make it clear this is a MultiStrategy
// even when there's only one strategy after filtering out nils

func (ms *MultiStrategy) WaitUntilReady(ctx context.Context, target StrategyTarget) error {
	_ = "STUB: not implemented"
	return nil
}

// A module could be appending strategies after part of the container initialization,
// and use wait.ForAll on a not initialized strategy.
// In this case, we just skip the nil strategy.

// Set default Timeout when strategy implements StrategyTimeout
