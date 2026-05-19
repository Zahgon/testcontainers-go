package pinecone

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

// Container represents the Pinecone container type used in the module
type Container struct {
	testcontainers.Container
}

// Run creates an instance of the Pinecone container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HttpEndpoint returns the http endpoint for the pinecone container
//
//nolint:revive,staticcheck //FIXME
func (c *Container) HttpEndpoint() (string, error) { _ = "STUB: not implemented"; return "", nil }
