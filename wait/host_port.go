package wait

import (
	"context"
	"errors"
	"time"

	"github.com/moby/moby/api/types/network"
)

const (
	exitEaccess     = 126 // container cmd can't be invoked (permission denied)
	exitCmdNotFound = 127 // container cmd not found/does not exist or invalid bind-mount
)

// Implement interface
var (
	_ Strategy        = (*HostPortStrategy)(nil)
	_ StrategyTimeout = (*HostPortStrategy)(nil)
)

var (
	errShellNotExecutable = errors.New("/bin/sh command not executable")
	errShellNotFound      = errors.New("/bin/sh command not found")
)

type HostPortStrategy struct {
	// Port is a string containing port number and protocol in the format "80/tcp"
	// which
	Port string
	// all WaitStrategies should have a startupTimeout to avoid waiting infinitely
	timeout      *time.Duration
	PollInterval time.Duration

	// skipInternalCheck is a flag to skip the internal check, which is useful when
	// a shell is not available in the container or when the container doesn't bind
	// the port internally until additional conditions are met.
	skipInternalCheck bool

	// skipExternalCheck is a flag to skip the external check, which, if used with
	// skipInternalCheck, makes strategy waiting only for port mapping completion
	// without accessing port.
	skipExternalCheck bool
}

// NewHostPortStrategy constructs a default host port strategy that waits for the given
// port to be exposed. The default startup timeout is 60 seconds.
func NewHostPortStrategy(port string) *HostPortStrategy { _ = "STUB: not implemented"; return nil }

// fluent builders for each property
// since go has neither covariance nor generics, the return type must be the type of the concrete implementation
// this is true for all properties, even the "shared" ones like startupTimeout

// ForListeningPort returns a host port strategy that waits for the given port
// to be exposed and bound internally the container.
// Alias for `NewHostPortStrategy(port)`.
func ForListeningPort(port string) *HostPortStrategy { _ = "STUB: not implemented"; return nil }

// ForExposedPort returns a host port strategy that waits for the first port
// to be exposed and bound internally the container.
func ForExposedPort() *HostPortStrategy { _ = "STUB: not implemented"; return nil }

// ForMappedPort returns a host port strategy that waits for the given port
// to be mapped without accessing the port itself.
func ForMappedPort(port string) *HostPortStrategy { _ = "STUB: not implemented"; return nil }

// SkipInternalCheck changes the host port strategy to skip the internal check,
// which is useful when a shell is not available in the container or when the
// container doesn't bind the port internally until additional conditions are met.
func (hp *HostPortStrategy) SkipInternalCheck() *HostPortStrategy {
	_ = "STUB: not implemented"
	return nil
}

// SkipExternalCheck changes the host port strategy to skip the external check,
// which, if used with SkipInternalCheck, makes strategy waiting only for port
// mapping completion without accessing port.
func (hp *HostPortStrategy) SkipExternalCheck() *HostPortStrategy {
	_ = "STUB: not implemented"
	return nil
}

// WithStartupTimeout can be used to change the default startup timeout
func (hp *HostPortStrategy) WithStartupTimeout(startupTimeout time.Duration) *HostPortStrategy {
	_ = "STUB: not implemented"
	return nil
}

// WithPollInterval can be used to override the default polling interval of 100 milliseconds
func (hp *HostPortStrategy) WithPollInterval(pollInterval time.Duration) *HostPortStrategy {
	_ = "STUB: not implemented"
	return nil
}

func (hp *HostPortStrategy) Timeout() *time.Duration {
	_ = "STUB: not implemented"

	// String returns a human-readable description of the wait strategy.
	return nil
}

func (hp *HostPortStrategy) String() string { _ = "STUB: not implemented"; return "" }

// detectInternalPort returns the lowest internal port that is currently bound.
// If no internal port is found, it returns the zero nat.Port value which
// can be checked against an empty string.
func (hp *HostPortStrategy) detectInternalPort(ctx context.Context, target StrategyTarget) (network.Port, error) {
	_ = "STUB: not implemented"
	return *new(network.Port), nil
}

// WaitUntilReady implements Strategy.WaitUntilReady
func (hp *HostPortStrategy) WaitUntilReady(ctx context.Context, target StrategyTarget) error {
	_ = "STUB: not implemented"
	return nil
}

// Port is not specified, so we need to detect it.

func externalCheck(ctx context.Context, ipAddress string, port network.Port, target StrategyTarget, waitInterval time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func internalCheck(ctx context.Context, internalPort network.Port, target StrategyTarget) error {
	_ = "STUB: not implemented"
	return nil
}

// Docker has an issue which override exit code 127 to 126 due to:
// https://github.com/moby/moby/issues/45795
// Handle both to ensure compatibility with Docker and Podman for now.

func buildInternalCheckCommand(internalPort uint16) string { _ = "STUB: not implemented"; return "" }
