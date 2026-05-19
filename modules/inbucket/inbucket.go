package inbucket

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

// InbucketContainer represents the Inbucket container type used in the module
type InbucketContainer struct {
	testcontainers.Container
}

// SmtpConnection returns the connection string for the smtp server, using the default
// 2500 port, and obtaining the host and exposed port from the container.
//
//nolint:revive,staticcheck //FIXME
func (c *InbucketContainer) SmtpConnection(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// WebInterface returns the connection string for the web interface server,
// using the default 9000 port, and obtaining the host and exposed port from
// the container.
func (c *InbucketContainer) WebInterface(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Deprecated: use Run instead
// RunContainer creates an instance of the Inbucket container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*InbucketContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the Inbucket container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*InbucketContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
