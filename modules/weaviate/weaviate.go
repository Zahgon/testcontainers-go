package weaviate

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	httpPort = "8080/tcp"
	grpcPort = "50051/tcp"
)

// WeaviateContainer represents the Weaviate container type used in the module
type WeaviateContainer struct {
	testcontainers.Container
}

// Deprecated: use Run instead
// RunContainer creates an instance of the Weaviate container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*WeaviateContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the Weaviate container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*WeaviateContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HttpHostAddress returns the schema and host of the Weaviate container.
// At the moment, it only supports the http scheme.
//
//nolint:revive,staticcheck //FIXME
func (c *WeaviateContainer) HttpHostAddress(ctx context.Context) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// GrpcHostAddress returns the gRPC host of the Weaviate container.
// At the moment, it only supports unsecured gRPC connection.
func (c *WeaviateContainer) GrpcHostAddress(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
