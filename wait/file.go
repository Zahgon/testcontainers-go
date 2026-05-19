package wait

import (
	"context"
	"io"
	"time"
)

var (
	_ Strategy        = (*FileStrategy)(nil)
	_ StrategyTimeout = (*FileStrategy)(nil)
)

// FileStrategy waits for a file to exist in the container.
type FileStrategy struct {
	timeout      *time.Duration
	file         string
	pollInterval time.Duration
	matcher      func(io.Reader) error
}

// NewFileStrategy constructs an FileStrategy strategy.
func NewFileStrategy(file string) *FileStrategy { _ = "STUB: not implemented"; return nil }

// WithStartupTimeout can be used to change the default startup timeout
func (ws *FileStrategy) WithStartupTimeout(startupTimeout time.Duration) *FileStrategy {
	_ = "STUB: not implemented"
	return nil
}

// WithPollInterval can be used to override the default polling interval of 100 milliseconds
func (ws *FileStrategy) WithPollInterval(pollInterval time.Duration) *FileStrategy {
	_ = "STUB: not implemented"
	return nil
}

// WithMatcher can be used to consume the file content.
// The matcher can return an errdefs.ErrNotFound to indicate that the file is not ready.
// Any other error will be considered a failure.
// Default: nil, will only wait for the file to exist.
func (ws *FileStrategy) WithMatcher(matcher func(io.Reader) error) *FileStrategy {
	_ = "STUB: not implemented"
	return nil
}

// ForFile is a convenience method to assign FileStrategy
func ForFile(file string) *FileStrategy { _ = "STUB: not implemented"; return nil }

// Timeout returns the timeout for the strategy
func (ws *FileStrategy) Timeout() *time.Duration {
	_ = "STUB: not implemented"

	// String returns a human-readable description of the wait strategy.
	return nil
}

func (ws *FileStrategy) String() string { _ = "STUB: not implemented"; return "" }

// WaitUntilReady waits until the file exists in the container and copies it to the target.
func (ws *FileStrategy) WaitUntilReady(ctx context.Context, target StrategyTarget) error {
	_ = "STUB: not implemented"
	return nil
}

// Not found, continue polling.

// matchFile tries to copy the file from the container and match it.
func (ws *FileStrategy) matchFile(ctx context.Context, target StrategyTarget) error {
	_ = "STUB: not implemented"
	return nil
}

// No matcher, just check if the file exists.
