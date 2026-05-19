package socat

import (
	"context"
	"net/url"

	"github.com/testcontainers/testcontainers-go"
)

const (
	// defaultImage {
	// DefaultImage is the default image used by the Socat container.
	DefaultImage = "alpine/socat:1.8.0.1"
	// }
)

// Container represents the Socat container type used in the module.
// A socat container is used as a TCP proxy, enabling any TCP port
// of another container to be exposed publicly, even if that container
// does not make the port public itself.
type Container struct {
	testcontainers.Container
	targetURLs map[int]*url.URL
}

// Run creates an instance of the Socat container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	// Gather all config options (defaults and then apply provided options)
	return nil, nil
}

// Only check if the socat binary is available if there are targets to expose.
// This is because the socat container exits otherwise.

// TargetURL returns the URL for the exposed port of a target, nil if the port is not mapped
func (c *Container) TargetURL(exposedPort int) *url.URL { _ = "STUB: not implemented"; return nil }
