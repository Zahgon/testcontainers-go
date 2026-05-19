package wait

import (
	"context"
	"regexp"
	"time"
)

// Implement interface
var (
	_ Strategy        = (*LogStrategy)(nil)
	_ StrategyTimeout = (*LogStrategy)(nil)
)

// PermanentError is a special error that will stop the wait and return an error.
type PermanentError struct {
	err error
}

// Error implements the error interface.
func (e *PermanentError) Error() string { _ = "STUB: not implemented"; return "" }

// NewPermanentError creates a new PermanentError.
func NewPermanentError(err error) *PermanentError { _ = "STUB: not implemented"; return nil }

// LogStrategy will wait until a given log entry shows up in the docker logs
type LogStrategy struct {
	// all Strategies should have a startupTimeout to avoid waiting infinitely
	timeout *time.Duration

	// additional properties
	Log          string
	IsRegexp     bool
	Occurrence   int
	PollInterval time.Duration

	// check is the function that will be called to check if the log entry is present.
	check func([]byte) error

	// submatchCallback is a callback that will be called with the sub matches of the regexp.
	submatchCallback func(pattern string, matches [][][]byte) error

	// re is the optional compiled regexp.
	re *regexp.Regexp

	// log byte slice version of [LogStrategy.Log] used for count checks.
	log []byte
}

// NewLogStrategy constructs with polling interval of 100 milliseconds and startup timeout of 60 seconds by default
func NewLogStrategy(log string) *LogStrategy { _ = "STUB: not implemented"; return nil }

// fluent builders for each property
// since go has neither covariance nor generics, the return type must be the type of the concrete implementation
// this is true for all properties, even the "shared" ones like startupTimeout

// AsRegexp can be used to change the default behavior of the log strategy to use regexp instead of plain text
func (ws *LogStrategy) AsRegexp() *LogStrategy { _ = "STUB: not implemented"; return nil }

// Submatch configures a function that will be called with the result of
// [regexp.Regexp.FindAllSubmatch], allowing the caller to process the results.
// If the callback returns nil, the strategy will be considered successful.
// Returning a [PermanentError] will stop the wait and return an error, otherwise
// it will retry until the timeout is reached.
// [LogStrategy.Occurrence] is ignored if this option is set.
func (ws *LogStrategy) Submatch(callback func(pattern string, matches [][][]byte) error) *LogStrategy {
	_ = "STUB: not implemented"
	return nil
}

// WithStartupTimeout can be used to change the default startup timeout
func (ws *LogStrategy) WithStartupTimeout(timeout time.Duration) *LogStrategy {
	_ = "STUB: not implemented"
	return nil
}

// WithPollInterval can be used to override the default polling interval of 100 milliseconds
func (ws *LogStrategy) WithPollInterval(pollInterval time.Duration) *LogStrategy {
	_ = "STUB: not implemented"
	return nil
}

func (ws *LogStrategy) WithOccurrence(o int) *LogStrategy {
	_ = "STUB: not implemented"
	// the number of occurrence needs to be positive
	return nil
}

// ForLog is the default construction for the fluid interface.
//
// For Example:
//
//	wait.
//		ForLog("some text").
//		WithPollInterval(1 * time.Second)
func ForLog(log string) *LogStrategy { _ = "STUB: not implemented"; return nil }

func (ws *LogStrategy) Timeout() *time.Duration {
	_ = "STUB: not implemented"

	// String returns a human-readable description of the wait strategy.
	return nil
}

func (ws *LogStrategy) String() string { _ = "STUB: not implemented"; return "" }

// WaitUntilReady implements Strategy.WaitUntilReady
func (ws *LogStrategy) WaitUntilReady(ctx context.Context, target StrategyTarget) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: fix as this will wait for timeout if the logs are not available.

// TODO: fix as this will wait for timeout if the logs are not readable.

// Log length hasn't changed so we're not making progress.

// checkCount checks if the log entry is present in the logs using a string count.
func (ws *LogStrategy) checkCount(b []byte) error { _ = "STUB: not implemented"; return nil }

// checkRegexp checks if the log entry is present in the logs using a regexp count.
func (ws *LogStrategy) checkRegexp(b []byte) error { _ = "STUB: not implemented"; return nil }

// checkSubmatch checks if the log entry is present in the logs using a regexp sub match callback.
func (ws *LogStrategy) checkSubmatch(b []byte) error { _ = "STUB: not implemented"; return nil }
