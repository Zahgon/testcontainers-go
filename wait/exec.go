package wait

import (
	"context"
	"io"
	"time"
)

// Implement interface
var (
	_ Strategy        = (*ExecStrategy)(nil)
	_ StrategyTimeout = (*ExecStrategy)(nil)
)

type ExecStrategy struct {
	// all Strategies should have a startupTimeout to avoid waiting infinitely
	timeout *time.Duration
	cmd     []string

	// additional properties
	ExitCodeMatcher func(exitCode int) bool
	ResponseMatcher func(body io.Reader) bool
	PollInterval    time.Duration
}

// NewExecStrategy constructs an Exec strategy ...
func NewExecStrategy(cmd []string) *ExecStrategy { _ = "STUB: not implemented"; return nil }

func defaultExitCodeMatcher(exitCode int) bool { _ = "STUB: not implemented"; return false }

// WithStartupTimeout can be used to change the default startup timeout
func (ws *ExecStrategy) WithStartupTimeout(startupTimeout time.Duration) *ExecStrategy {
	_ = "STUB: not implemented"
	return nil
}

func (ws *ExecStrategy) WithExitCode(exitCode int) *ExecStrategy {
	_ = "STUB: not implemented"
	return nil
}

func (ws *ExecStrategy) WithExitCodeMatcher(exitCodeMatcher func(exitCode int) bool) *ExecStrategy {
	_ = "STUB: not implemented"
	return nil
}

func (ws *ExecStrategy) WithResponseMatcher(matcher func(body io.Reader) bool) *ExecStrategy {
	_ = "STUB: not implemented"
	return nil
}

// WithPollInterval can be used to override the default polling interval of 100 milliseconds
func (ws *ExecStrategy) WithPollInterval(pollInterval time.Duration) *ExecStrategy {
	_ = "STUB: not implemented"
	return nil
}

// ForExec is a convenience method to assign ExecStrategy
func ForExec(cmd []string) *ExecStrategy { _ = "STUB: not implemented"; return nil }

func (ws *ExecStrategy) Timeout() *time.Duration {
	_ = "STUB: not implemented"

	// String returns a human-readable description of the wait strategy.
	return nil
}

func (ws *ExecStrategy) String() string { _ = "STUB: not implemented"; return "" }

// Only show the command name and argument count to avoid exposing sensitive data

func (ws *ExecStrategy) WaitUntilReady(ctx context.Context, target StrategyTarget) error {
	_ = "STUB: not implemented"
	return nil
}
