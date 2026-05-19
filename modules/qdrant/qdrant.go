package qdrant

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

// QdrantContainer represents the Qdrant container type used in the module
type QdrantContainer struct {
	testcontainers.Container
}

// Deprecated: use Run instead
// RunContainer creates an instance of the Qdrant container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*QdrantContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the Qdrant container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*QdrantContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RESTEndpoint returns the REST endpoint of the Qdrant container
func (c *QdrantContainer) RESTEndpoint(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GRPCEndpoint returns the gRPC endpoint of the Qdrant container
func (c *QdrantContainer) GRPCEndpoint(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// WebUI returns the web UI endpoint of the Qdrant container
func (c *QdrantContainer) WebUI(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
