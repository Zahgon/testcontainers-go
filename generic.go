package testcontainers

import (
	"context"
	"errors"
	"sync"

	"github.com/testcontainers/testcontainers-go/log"
)

var (
	reuseContainerMx  sync.Mutex
	ErrReuseEmptyName = errors.New("with reuse option a container name mustn't be empty")
)

// GenericContainerRequest represents parameters to a generic container
type GenericContainerRequest struct {
	ContainerRequest              // embedded request for provider
	Started          bool         // whether to auto-start the container
	ProviderType     ProviderType // which provider to use, Docker if empty
	Logger           log.Logger   // provide a container specific Logging - use default global logger if empty
	Reuse            bool         // reuse an existing container if it exists or create a new one. a container name mustn't be empty
}

// Deprecated: will be removed in the future.
// GenericNetworkRequest represents parameters to a generic network
type GenericNetworkRequest struct {
	NetworkRequest              // embedded request for provider
	ProviderType   ProviderType // which provider to use, Docker if empty
}

// Deprecated: use network.New instead
// GenericNetwork creates a generic network with parameters
func GenericNetwork(ctx context.Context, req GenericNetworkRequest) (Network, error) {
	_ = "STUB: not implemented"
	return *new(Network), nil
}

// GenericContainer creates a generic container with parameters
func GenericContainer(ctx context.Context, req GenericContainerRequest) (Container, error) {
	_ = "STUB: not implemented"
	return *new(Container), nil
}

// Ensure there is always a non-nil logger by default

// we must protect the reusability of the container in the case it's invoked
// in a parallel execution, via ParallelContainers or t.Parallel()

// At this point `c` might not be nil. Give the caller an opportunity to call Destroy on the container.
// TODO: Remove this debugging.

// Debugging information for rate limiting.

// GenericProvider represents an abstraction for container and network providers
type GenericProvider interface {
	ContainerProvider
	NetworkProvider
	ImageProvider
}

// GenericLabels returns a map of labels that can be used to identify resources
// created by this library. This includes the standard LabelSessionID if the
// reaper is enabled, otherwise this is excluded to prevent resources being
// incorrectly reaped.
func GenericLabels() map[string]string { _ = "STUB: not implemented"; return nil }

// AddGenericLabels adds the generic labels to target.
func AddGenericLabels(target map[string]string) { _ = "STUB: not implemented"; return }

// Run is a convenience function that creates a new container and starts it.
// It calls the GenericContainer function and returns a concrete DockerContainer type.
func Run(ctx context.Context, img string, opts ...ContainerCustomizer) (*DockerContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
