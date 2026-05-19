package opensearch

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	defaultPassword = "admin"
	defaultUsername = "admin"
	defaultHTTPPort = "9200/tcp"
)

// OpenSearchContainer represents the OpenSearch container type used in the module
type OpenSearchContainer struct {
	testcontainers.Container
	User     string
	Password string
}

// Deprecated: use Run instead
// RunContainer creates an instance of the OpenSearch container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*OpenSearchContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the OpenSearch container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*OpenSearchContainer, error) {
	_ = "STUB: not implemented"
	// Gather all config options (defaults and then apply provided options)
	return nil, nil
}

// Set memlock to unlimited (no soft or hard limit)

// Maximum number of open files for the opensearch user - set to at least 65536

// the wait strategy does not support TLS at the moment,
// so we need to disable it in the strategy for now.

// Address retrieves the address of the OpenSearch container.
// It will use http as protocol, as TLS is not supported at the moment.
func (c *OpenSearchContainer) Address(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
