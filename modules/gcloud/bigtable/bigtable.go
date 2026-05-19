package bigtable

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	// DefaultProjectID is the default project ID for the BigTable container.
	DefaultProjectID  = "test-project"
	defaultPortNumber = "9000"
	defaultPort       = defaultPortNumber + "/tcp"
)

// Container represents the BigTable container type used in the module
type Container struct {
	testcontainers.Container
	settings options
}

// ProjectID returns the project ID of the BigTable container.
func (c *Container) ProjectID() string { _ = "STUB: not implemented"; return "" }

// URI returns the URI of the BigTable container.
func (c *Container) URI() string { _ = "STUB: not implemented"; return "" }

// Run creates an instance of the BigTable GCloud container type.
// The URI uses the empty string as the protocol.
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
