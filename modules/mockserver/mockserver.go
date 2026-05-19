package mockserver

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

// MockServerContainer represents the MockServer container type used in the module
type MockServerContainer struct {
	testcontainers.Container
}

// Deprecated: use Run instead
// RunContainer creates an instance of the MockServer container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*MockServerContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the MockServer container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*MockServerContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// URL returns the URL of the MockServer container
func (c *MockServerContainer) URL(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
