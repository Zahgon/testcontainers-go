package wait

import (
	"context"
	"time"
)

// Implement interface
var (
	_ Strategy        = (*ExitStrategy)(nil)
	_ StrategyTimeout = (*ExitStrategy)(nil)
)

// ExitStrategy will wait until container exit
type ExitStrategy struct {
	// all Strategies should have a timeout to avoid waiting infinitely
	timeout *time.Duration

	// additional properties
	PollInterval time.Duration
}

// NewExitStrategy constructs with polling interval of 100 milliseconds without timeout by default
func NewExitStrategy() *ExitStrategy { _ = "STUB: not implemented"; return nil }

// fluent builders for each property
// since go has neither covariance nor generics, the return type must be the type of the concrete implementation
// this is true for all properties, even the "shared" ones

// WithExitTimeout can be used to change the default exit timeout
func (ws *ExitStrategy) WithExitTimeout(exitTimeout time.Duration) *ExitStrategy {
	_ = "STUB: not implemented"
	return nil
}

// WithPollInterval can be used to override the default polling interval of 100 milliseconds
func (ws *ExitStrategy) WithPollInterval(pollInterval time.Duration) *ExitStrategy {
	_ = "STUB: not implemented"
	return nil
}

// ForExit is the default construction for the fluid interface.
//
// For Example:
//
//	wait.
//		ForExit().
//		WithPollInterval(1 * time.Second)
func ForExit() *ExitStrategy { _ = "STUB: not implemented"; return nil }

func (ws *ExitStrategy) Timeout() *time.Duration {
	_ = "STUB: not implemented"

	// String returns a human-readable description of the wait strategy.
	return nil
}

func (ws *ExitStrategy) String() string { _ = "STUB: not implemented"; return "" }

// WaitUntilReady implements Strategy.WaitUntilReady
func (ws *ExitStrategy) WaitUntilReady(ctx context.Context, target StrategyTarget) error {
	_ = "STUB: not implemented"
	return nil
}
