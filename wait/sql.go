package wait

import (
	"context"
	"time"
)

var (
	_ Strategy        = (*waitForSQL)(nil)
	_ StrategyTimeout = (*waitForSQL)(nil)
)

const defaultForSQLQuery = "SELECT 1"

// ForSQL constructs a new waitForSql strategy for the given driver
func ForSQL(port string, driver string, url func(host string, port string) string) *waitForSQL {
	_ = "STUB: not implemented"
	return nil
}

type waitForSQL struct {
	timeout *time.Duration

	URL            func(host string, port string) string
	Driver         string
	Port           string
	startupTimeout time.Duration
	PollInterval   time.Duration
	query          string
}

// WithStartupTimeout can be used to change the default startup timeout
func (w *waitForSQL) WithStartupTimeout(timeout time.Duration) *waitForSQL {
	_ = "STUB: not implemented"
	return nil
}

// WithPollInterval can be used to override the default polling interval of 100 milliseconds
func (w *waitForSQL) WithPollInterval(pollInterval time.Duration) *waitForSQL {
	_ = "STUB: not implemented"
	return nil
}

// WithQuery can be used to override the default query used in the strategy.
func (w *waitForSQL) WithQuery(query string) *waitForSQL { _ = "STUB: not implemented"; return nil }

func (w *waitForSQL) Timeout() *time.Duration {
	_ = "STUB: not implemented"

	// String returns a human-readable description of the wait strategy.
	return nil
}

func (w *waitForSQL) String() string { _ = "STUB: not implemented"; return "" }

// WaitUntilReady repeatedly tries to run "SELECT 1" or user defined query on the given port using sql and driver.
//
// If it doesn't succeed until the timeout value which defaults to 60 seconds, it will return an error.
func (w *waitForSQL) WaitUntilReady(ctx context.Context, target StrategyTarget) error {
	_ = "STUB: not implemented"
	return nil
}
