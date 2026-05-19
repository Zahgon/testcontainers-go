package consul

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	defaultHTTPAPIPort = "8500"
	defaultBrokerPort  = "8600"
)

const (
	// Deprecated: it will be removed in the next major version.
	DefaultBaseImage = "hashicorp/consul:1.15"
)

// ConsulContainer represents the Consul container type used in the module.
type ConsulContainer struct {
	testcontainers.Container
}

// ApiEndpoint returns host:port for the HTTP API endpoint.
//
//nolint:revive,staticcheck //FIXME
func (c *ConsulContainer) ApiEndpoint(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// WithConfigString takes in a JSON string of keys and values to define a configuration to be used by the instance.
func WithConfigString(config string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithConfigFile takes in a path to a JSON file to define a configuration to be used by the instance.
func WithConfigFile(configPath string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// Deprecated: use Run instead
// RunContainer creates an instance of the Consul container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*ConsulContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the Consul container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*ConsulContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
