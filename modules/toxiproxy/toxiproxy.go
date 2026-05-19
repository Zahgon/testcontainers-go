package toxiproxy

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	// ControlPort is the port of the Toxiproxy control API
	ControlPort = "8474/tcp"

	// firstProxiedPort is the first port of the range of ports that will be proxied
	firstProxiedPort = 8666
)

// Container represents the Toxiproxy container type used in the module
type Container struct {
	testcontainers.Container

	// proxiedEndpoints is a map of the proxied endpoints of the Toxiproxy container
	proxiedEndpoints map[int]string
}

// ProxiedEndpoint returns the endpoint for the proxied port in the Toxiproxy container,
// an error in case the port has no proxied endpoint.
func (c *Container) ProxiedEndpoint(p int) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// URI returns the URI of the Toxiproxy container
func (c *Container) URI(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Run creates an instance of the Toxiproxy container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	// Process custom options first
	return nil, nil
}

// Expose the ports for the proxies, starting from the first proxied port

// Update the listen port of the proxy

// Render the config file

// Apply the config file to the container with the proxies.

// Map the ports of the proxies to the container, so that we can use them in the tests
