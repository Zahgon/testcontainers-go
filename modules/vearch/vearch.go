package vearch

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

// VearchContainer represents the Vearch container type used in the module
type VearchContainer struct {
	testcontainers.Container
}

// Deprecated: use Run instead
// RunContainer creates an instance of the Vearch container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*VearchContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the Vearch container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*VearchContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RESTEndpoint returns the REST endpoint of the Vearch container
func (c *VearchContainer) RESTEndpoint(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
