package testcontainers

import (
	"context"

	"github.com/testcontainers/testcontainers-go/log"
)

// possible provider types
const (
	ProviderDefault ProviderType = iota // default will auto-detect provider from DOCKER_HOST environment variable
	ProviderDocker
	ProviderPodman
)

type (
	// ProviderType is an enum for the possible providers
	ProviderType int

	// GenericProviderOptions defines options applicable to all providers
	GenericProviderOptions struct {
		Logger         log.Logger
		defaultNetwork string
	}

	// GenericProviderOption defines a common interface to modify GenericProviderOptions
	// These options can be passed to GetProvider in a variadic way to customize the returned GenericProvider instance
	GenericProviderOption interface {
		ApplyGenericTo(opts *GenericProviderOptions)
	}

	// GenericProviderOptionFunc is a shorthand to implement the GenericProviderOption interface
	GenericProviderOptionFunc func(opts *GenericProviderOptions)

	// DockerProviderOptions defines options applicable to DockerProvider
	DockerProviderOptions struct {
		defaultBridgeNetworkName string
		*GenericProviderOptions
	}

	// DockerProviderOption defines a common interface to modify DockerProviderOptions
	// These can be passed to NewDockerProvider in a variadic way to customize the returned DockerProvider instance
	DockerProviderOption interface {
		ApplyDockerTo(opts *DockerProviderOptions)
	}

	// DockerProviderOptionFunc is a shorthand to implement the DockerProviderOption interface
	DockerProviderOptionFunc func(opts *DockerProviderOptions)
)

func (f DockerProviderOptionFunc) ApplyDockerTo(opts *DockerProviderOptions) {
	_ = "STUB: not implemented"
	return
}

func Generic2DockerOptions(opts ...GenericProviderOption) []DockerProviderOption {
	_ = "STUB: not implemented"
	return nil
}

func WithDefaultBridgeNetwork(bridgeNetworkName string) DockerProviderOption {
	_ = "STUB: not implemented"
	return *new(DockerProviderOption)
}

func (f GenericProviderOptionFunc) ApplyGenericTo(opts *GenericProviderOptions) {
	_ = "STUB: not implemented"

	// ContainerProvider allows the creation of containers on an arbitrary system
	return
}

type ContainerProvider interface {
	Close() error                                                                // close the provider
	CreateContainer(context.Context, ContainerRequest) (Container, error)        // create a container without starting it
	ReuseOrCreateContainer(context.Context, ContainerRequest) (Container, error) // reuses a container if it exists or creates a container without starting
	RunContainer(context.Context, ContainerRequest) (Container, error)           // create a container and start it
	Health(context.Context) error
	Config() TestcontainersConfig
}

// GetProvider provides the provider implementation for a certain type
func (t ProviderType) GetProvider(opts ...GenericProviderOption) (GenericProvider, error) {
	_ = "STUB: not implemented"
	return *new(GenericProvider), nil
}

// NewDockerProvider creates a Docker provider with the EnvClient
func NewDockerProvider(provOpts ...DockerProviderOption) (*DockerProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
