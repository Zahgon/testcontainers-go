package openfga

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

// OpenFGAContainer represents the OpenFGA container type used in the module
type OpenFGAContainer struct {
	testcontainers.Container
}

// GrpcEndpoint returns the gRPC endpoint for the OpenFGA container,
// which uses the 8081/tcp port.
func (c *OpenFGAContainer) GrpcEndpoint(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// HttpEndpoint returns the HTTP endpoint for the OpenFGA container,
// which uses the 8080/tcp port.
//
//nolint:revive,staticcheck //FIXME
func (c *OpenFGAContainer) HttpEndpoint(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// PlaygroundEndpoint returns the playground endpoint for the OpenFGA container,
// which is the HTTP endpoint with the path /playground in the port 3000/tcp.
func (c *OpenFGAContainer) PlaygroundEndpoint(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Deprecated: use Run instead
// RunContainer creates an instance of the OpenFGA container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*OpenFGAContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the OpenFGA container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*OpenFGAContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
