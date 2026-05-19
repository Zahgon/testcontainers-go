package spanner

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	// DefaultProjectID is the default project ID for the Spanner container.
	DefaultProjectID  = "test-project"
	defaultPortNumber = "9010"
	defaultPort       = defaultPortNumber + "/tcp"
)

// Container represents the Spanner container type used in the module
type Container struct {
	testcontainers.Container
	settings options
}

// ProjectID returns the project ID of the Spanner container.
func (c *Container) ProjectID() string { _ = "STUB: not implemented"; return "" }

// URI returns the URI of the Spanner container.
func (c *Container) URI() string { _ = "STUB: not implemented"; return "" }

// Run creates an instance of the Spanner GCloud container type.
// The URI uses the empty string as the protocol.
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
