package dockermcpgateway

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	defaultPort = "8811/tcp"
	secretsPath = "/testcontainers/app/secrets"
)

// Container represents the DockerMCPGateway container type used in the module
type Container struct {
	testcontainers.Container
	tools map[string][]string
}

// Run creates an instance of the DockerMCPGateway container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Process custom options first to extract settings

// Build moduleOpts with defaults

// Append user options

// GatewayEndpoint returns the endpoint for the DockerMCPGateway container.
// It uses the mapped port for the default port (8811/tcp) and the "http" protocol.
func (c *Container) GatewayEndpoint(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Tools returns the tools configured for the DockerMCPGateway container,
// indexed by server name.
// The keys are the server names and the values are slices of tool names.
func (c *Container) Tools() map[string][]string { _ = "STUB: not implemented"; return nil }
