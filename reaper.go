package testcontainers

import (
	"context"
	"errors"
	"net"
	"sync"

	"github.com/cenkalti/backoff/v4"
	"github.com/moby/moby/api/types/network"

	"github.com/testcontainers/testcontainers-go/internal/config"
)

const (
	// Deprecated: it has been replaced by the internal core.LabelLang
	TestcontainerLabel = "org.testcontainers.golang"
	// Deprecated: it has been replaced by the internal core.LabelSessionID
	TestcontainerLabelSessionID = TestcontainerLabel + ".sessionId"
	// Deprecated: it has been replaced by the internal core.LabelReaper
	TestcontainerLabelIsReaper = TestcontainerLabel + ".reaper"
)

var (
	// Deprecated: it has been replaced by an internal value
	ReaperDefaultImage = config.ReaperDefaultImage

	// defaultReaperPort is the default port that the reaper listens on if not
	// overridden by the RYUK_PORT environment variable.
	defaultReaperPort = network.MustParsePort("8080/tcp")

	// errReaperNotFound is returned when no reaper container is found.
	errReaperNotFound = errors.New("reaper not found")

	// errReaperDisabled is returned if a reaper is requested but the
	// config has it disabled.
	errReaperDisabled = errors.New("reaper disabled")

	// spawner is the singleton instance of reaperSpawner.
	spawner = &reaperSpawner{}

	// reaperAck is the expected response from the reaper container.
	reaperAck = []byte("ACK\n")
)

// ReaperProvider represents a provider for the reaper to run itself with
// The ContainerProvider interface should usually satisfy this as well, so it is pluggable
type ReaperProvider interface {
	RunContainer(ctx context.Context, req ContainerRequest) (Container, error)
	Config() TestcontainersConfig
}

// Deprecated: it's not possible to create a reaper any more. Compose module uses this method
// to create a reaper for the compose stack.
//
// # NewReaper creates a Reaper with a sessionID to identify containers and a provider to use
//
// The caller must call Connect at least once on the returned Reaper and use the returned
// result otherwise the reaper will be kept open until the process exits.
func NewReaper(ctx context.Context, sessionID string, provider ReaperProvider, _ string) (*Reaper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// reaperContainerNameFromSessionID returns the container name that uniquely
// identifies the container based on the session id.
func reaperContainerNameFromSessionID(sessionID string) string {
	_ = "STUB: not implemented"
	// The session id is 64 characters, so we will not hit the limit of 128
	// characters for container names.
	return ""
}

// reaperSpawner is a singleton that manages the reaper container.
type reaperSpawner struct {
	instance *Reaper
	mtx      sync.Mutex
}

// port returns the port that a new reaper should listen on.
func (r *reaperSpawner) port() network.Port { _ = "STUB: not implemented"; return *new(network.Port) }

// backoff returns a backoff policy for the reaper spawner.
// It will take at most 20 seconds, doing each attempt every 100ms - 250ms.
func (r *reaperSpawner) backoff() *backoff.ExponentialBackOff {
	_ = "STUB: not implemented"
	// We want random intervals between 100ms and 250ms for concurrent executions
	// to not be synchronized: it could be the case that multiple executions of this
	// function happen at the same time (specifically when called from a different test
	// process execution), and we want to avoid that they all try to find the reaper
	// container at the same time.
	return nil
}

// Adjust MaxInterval to compensate for randomization factor which can be added to
// returned interval so we have a maximum of 250ms.

// cleanup terminates the reaper container if set.
func (r *reaperSpawner) cleanup() error { _ = "STUB: not implemented"; return nil }

// cleanupLocked terminates the reaper container if set.
// It must be called with the lock held.
func (r *reaperSpawner) cleanupLocked() error { _ = "STUB: not implemented"; return nil }

// lookupContainer returns a DockerContainer type with the reaper container in the case
// it's found in the running state, and including the labels for sessionID, reaper, and ryuk.
// It will perform a retry with exponential backoff to allow for the container to be started and
// avoid potential false negatives.
func (r *reaperSpawner) lookupContainer(ctx context.Context, sessionID string) (*DockerContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No reaper container not found.

// isRunning returns an error if the container is not running.
func (r *reaperSpawner) isRunning(ctx context.Context, ctr Container) error {
	_ = "STUB: not implemented"
	return nil
}

// Use NotFound error to indicate the container is not running
// and should be recreated.

// retryError returns a permanent error if the error is not considered retryable.
func (r *reaperSpawner) retryError(err error) error { _ = "STUB: not implemented"; return nil }

// Retryable error.

// reaper returns an existing Reaper instance if it exists and is running, otherwise
// a new Reaper instance will be created with a sessionID to identify containers in
// the same test session/program. If connect is true, the reaper will be connected
// to the reaper container.
// Returns an error if config.RyukDisabled is true.
//
// Safe for concurrent calls.
func (r *reaperSpawner) reaper(ctx context.Context, sessionID string, provider ReaperProvider) (*Reaper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// retryLocked returns a function that can be used to create or reuse a reaper container.
// If connect is true, the reaper will be connected to the reaper container.
// It must be called with the lock held.
func (r *reaperSpawner) retryLocked(ctx context.Context, sessionID string, provider ReaperProvider) func() (*Reaper, error) {
	_ = "STUB: not implemented"
	return nil
}

// Ensure that the reaper is terminated if an error occurred.

// Check we can still connect.

// reuseOrCreate returns an existing Reaper instance if it exists, otherwise a new Reaper instance.
func (r *reaperSpawner) reuseOrCreate(ctx context.Context, sessionID string, provider ReaperProvider) (*Reaper, error) {
	_ = "STUB: not implemented"
	return nil,

		// We already have an associated reaper.
		nil
}

// Look for an existing reaper created in the same test session but in a
// different test process execution e.g. when running tests in parallel.

// The reaper container was not found, continue to create a new one.

// A reaper container exists re-use it.

// fromContainer constructs a Reaper from an already running reaper DockerContainer.
func (r *reaperSpawner) fromContainer(ctx context.Context, sessionID string, provider ReaperProvider, dockerContainer *DockerContainer) (*Reaper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reusing an existing container so we determine the port from the container's exposed ports.

// newReaper creates a connected Reaper with a sessionID to identify containers
// and a provider to use.
func (r *reaperSpawner) newReaper(ctx context.Context, sessionID string, provider ReaperProvider) (reaper *Reaper, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Setup reaper-specific labels for the reaper container.

// Attach reaper container to a requested network if it is specified

// Reaper is used to start a sidecar container that cleans up resources
type Reaper struct {
	Provider   ReaperProvider
	SessionID  string
	Endpoint   string
	container  Container
	mtx        sync.Mutex // Protects termSignal.
	termSignal chan bool
}

// Connect connects to the reaper container and sends the labels to it
// so that it can clean up the containers with the same labels.
//
// It returns a channel that can be closed to terminate the connection.
// Returns an error if config.RyukDisabled is true.
func (r *Reaper) Connect() (chan bool, error) { _ = "STUB: not implemented"; return nil, nil }

// close signals the connection to close if needed.
// Safe for concurrent calls.
func (r *Reaper) close() { _ = "STUB: not implemented"; return }

// setOrSignal sets the reapers termSignal field if nil
// otherwise consumes by sending true to it.
// Safe for concurrent calls.
func (r *Reaper) setOrSignal(termSignal chan bool) { _ = "STUB: not implemented"; return }

// Already have an existing connection, close the new one.

// First or new unused termSignal, assign for caller to reuse.

// useTermSignal if termSignal is not nil returns it
// and sets it to nil, otherwise returns nil.
//
// Safe for concurrent calls.
func (r *Reaper) useTermSignal() chan bool { _ = "STUB: not implemented"; return nil }

// Use existing connection.

// connect connects to the reaper container and sends the labels to it
// so that it can clean up the containers with the same labels.
//
// It returns a channel that can be sent true to terminate the connection.
// Returns an error if config.RyukDisabled is true.
func (r *Reaper) connect(ctx context.Context) (chan bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// handshake sends the labels to the reaper container and reads the ACK.
func (r *Reaper) handshake(conn net.Conn) error { _ = "STUB: not implemented"; return nil }

// We have received the ACK so all done.

// Deprecated: internally replaced by core.DefaultLabels(sessionID)
//
// Labels returns the container labels to use so that this Reaper cleans them up
func (r *Reaper) Labels() map[string]string { _ = "STUB: not implemented"; return nil }

// isReaperImage returns true if the image name is the reaper image.
func isReaperImage(name string) bool { _ = "STUB: not implemented"; return false }
