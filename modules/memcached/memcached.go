package memcached

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	defaultPort = "11211/tcp"
)

// Container represents the Memcached container type used in the module
type Container struct {
	testcontainers.Container
}

// Run creates an instance of the Memcached container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HostPort returns the host and port of the Memcached container
func (c *Container) HostPort(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
