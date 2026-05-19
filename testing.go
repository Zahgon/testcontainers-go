package testcontainers

import (
	"context"
	"regexp"
	"testing"
)

// errAlreadyInProgress is a regular expression that matches the error for a container
// removal that is already in progress.
var errAlreadyInProgress = regexp.MustCompile(`removal of container .* is already in progress`)

// SkipIfProviderIsNotHealthy is a utility function capable of skipping tests
// if the provider is not healthy, or running at all.
// This is a function designed to be used in your test, when Docker is not mandatory for CI/CD.
// In this way tests that depend on Testcontainers won't run if the provider is provisioned correctly.
func SkipIfProviderIsNotHealthy(t *testing.T) { _ = "STUB: not implemented"; return }

// SkipIfDockerDesktop is a utility function capable of skipping tests
// if tests are run using Docker Desktop or another VM-based Docker
// environment (e.g. colima) where host network access is not available.
func SkipIfDockerDesktop(t *testing.T, ctx context.Context) {
	_ = "STUB: not implemented"

	// Colima runs Docker inside a Linux VM, so host networking doesn't work
	// the same way as native Docker on Linux. Detect it via DOCKER_HOST which
	// typically contains the colima socket path.
	return
}

// SkipIfNotDockerDesktop is a utility function capable of skipping tests
// if tests are not run using Docker Desktop.
func SkipIfNotDockerDesktop(t *testing.T, ctx context.Context) { _ = "STUB: not implemented"; return }

// exampleLogConsumer {

// StdoutLogConsumer is a LogConsumer that prints the log to stdout
type StdoutLogConsumer struct{}

// Accept prints the log to stdout
func (lc *StdoutLogConsumer) Accept(l Log) { _ = "STUB: not implemented"; return }

// }

// CleanupContainer is a helper function that schedules the container
// to be stopped / terminated when the test ends.
//
// This should be called as a defer directly after (before any error check)
// of [GenericContainer](...) or a modules Run(...) in a test to ensure the
// container is stopped when the function ends.
//
// before any error check. If container is nil, it's a no-op.
func CleanupContainer(tb testing.TB, ctr Container, options ...TerminateOption) {
	_ = "STUB: not implemented"
	return
}

// CleanupNetwork is a helper function that schedules the network to be
// removed when the test ends.
// This should be the first call after NewNetwork(...) in a test before
// any error check. If network is nil, it's a no-op.
func CleanupNetwork(tb testing.TB, network Network) { _ = "STUB: not implemented"; return }

// noErrorOrIgnored is a helper function that checks if the error is nil or an error
// we can ignore.
func noErrorOrIgnored(tb testing.TB, err error) { _ = "STUB: not implemented"; return }

// causer is an interface that allows to get the cause of an error.
type causer interface {
	Cause() error
}

// wrapErr is an interface that allows to unwrap an error.
type wrapErr interface {
	Unwrap() error
}

// unwrapErrs is an interface that allows to unwrap multiple errors.
type unwrapErrs interface {
	Unwrap() []error
}

// isCleanupSafe reports whether all errors in err's tree are one of the
// following, so can safely be ignored:
//   - nil
//   - not found
//   - already in progress
func isCleanupSafe(err error) bool { _ = "STUB: not implemented"; return false }

// First try with containerd's errdefs

// Terminating a container that is already terminating.

//nolint:errorlint // We need to check for interfaces.

// RequireContainerExec is a helper function that executes a command in a container
// It insures that there is no error during the execution
// Finally returns the output of its execution
func RequireContainerExec(ctx context.Context, t *testing.T, container Container, cmd []string) string {
	_ = "STUB: not implemented"
	return ""
}
