package azurite

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	// BlobPort is the default port used by Azurite
	BlobPort = "10000/tcp"
	// QueuePort is the default port used by Azurite
	QueuePort = "10001/tcp"
	// TablePort is the default port used by Azurite
	TablePort = "10002/tcp"

	// defaultCredentials {
	// AccountName is the default testing account name used by Azurite
	AccountName string = "devstoreaccount1"

	// AccountKey is the default testing account key used by Azurite
	AccountKey string = "Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw=="
	// }
)

// Container represents the Azurite container type used in the module
type Container struct {
	testcontainers.Container
	opts options
}

// ServiceURL returns the URL of the given service
func (c *Container) ServiceURL(ctx context.Context, srv Service) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// BlobServiceURL returns the URL of the Blob service
func (c *Container) BlobServiceURL(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// QueueServiceURL returns the URL of the Queue service
func (c *Container) QueueServiceURL(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// TableServiceURL returns the URL of the Table service
func (c *Container) TableServiceURL(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func servicePort(srv Service) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Run creates an instance of the Azurite container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	// 1. Gather all config options (defaults and then apply provided options)
	return nil, nil
}

// Use azurite-table in future once it matures. Graceful shutdown is currently very slow.

// 2. evaluate the enabled services to apply the right wait strategy and Cmd options
