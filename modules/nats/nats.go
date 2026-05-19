package nats

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	defaultClientPort     = "4222/tcp"
	defaultRoutingPort    = "6222/tcp"
	defaultMonitoringPort = "8222/tcp"
)

// NATSContainer represents the NATS container type used in the module
type NATSContainer struct {
	testcontainers.Container
	User     string
	Password string
}

// Deprecated: use Run instead
// RunContainer creates an instance of the NATS container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*NATSContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the NATS container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*NATSContainer, error) {
	_ = "STUB: not implemented"
	// Gather all config options (defaults and then apply provided options)
	return nil, nil
}

// Include the command line arguments

// always prepend the dash because it was removed in the options

func (c *NATSContainer) MustConnectionString(ctx context.Context) string {
	_ = "STUB: not implemented"
	return ""
}

// ConnectionString returns a connection string for the NATS container
func (c *NATSContainer) ConnectionString(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
