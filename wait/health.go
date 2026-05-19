package wait

import (
	"context"
	"time"
)

// Implement interface
var (
	_ Strategy        = (*HealthStrategy)(nil)
	_ StrategyTimeout = (*HealthStrategy)(nil)
)

// HealthStrategy will wait until the container becomes healthy
type HealthStrategy struct {
	// all Strategies should have a startupTimeout to avoid waiting infinitely
	timeout *time.Duration

	// additional properties
	PollInterval time.Duration
}

// NewHealthStrategy constructs with polling interval of 100 milliseconds and startup timeout of 60 seconds by default
func NewHealthStrategy() *HealthStrategy { _ = "STUB: not implemented"; return nil }

// fluent builders for each property
// since go has neither covariance nor generics, the return type must be the type of the concrete implementation
// this is true for all properties, even the "shared" ones like startupTimeout

// WithStartupTimeout can be used to change the default startup timeout
func (ws *HealthStrategy) WithStartupTimeout(startupTimeout time.Duration) *HealthStrategy {
	_ = "STUB: not implemented"
	return nil
}

// WithPollInterval can be used to override the default polling interval of 100 milliseconds
func (ws *HealthStrategy) WithPollInterval(pollInterval time.Duration) *HealthStrategy {
	_ = "STUB: not implemented"
	return nil
}

// ForHealthCheck is the default construction for the fluid interface.
//
// For Example:
//
//	wait.
//		ForHealthCheck().
//		WithPollInterval(1 * time.Second)
func ForHealthCheck() *HealthStrategy { _ = "STUB: not implemented"; return nil }

func (ws *HealthStrategy) Timeout() *time.Duration {
	_ = "STUB: not implemented"

	// String returns a human-readable description of the wait strategy.
	return nil
}

func (ws *HealthStrategy) String() string { _ = "STUB: not implemented"; return "" }

// WaitUntilReady implements Strategy.WaitUntilReady
func (ws *HealthStrategy) WaitUntilReady(ctx context.Context, target StrategyTarget) error {
	_ = "STUB: not implemented"
	return nil
}
