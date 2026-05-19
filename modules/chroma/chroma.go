package chroma

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

// ChromaContainer represents the Chroma container type used in the module
type ChromaContainer struct {
	testcontainers.Container
}

// Deprecated: use Run instead
// RunContainer creates an instance of the Chroma container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*ChromaContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the Chroma container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*ChromaContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RESTEndpoint returns the REST endpoint of the Chroma container
func (c *ChromaContainer) RESTEndpoint(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
