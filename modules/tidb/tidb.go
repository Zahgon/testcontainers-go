package tidb

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	defaultUser     = "root"
	defaultPassword = ""
	defaultDatabase = "test"
	defaultPort     = "4000/tcp"
	restAPIPort     = "10080/tcp"
)

// Container represents the TiDB container type used in the module
type Container struct {
	testcontainers.Container
	username string
	password string
	database string
}

// ConnectionString returns a DSN connection string for the TiDB container,
// using the MySQL driver format. It is possible to pass extra parameters
// to the connection string, e.g. "tls=skip-verify".
func (c *Container) ConnectionString(ctx context.Context, args ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// MustConnectionString panics if the connection string cannot be determined.
func (c *Container) MustConnectionString(ctx context.Context, args ...string) string {
	_ = "STUB: not implemented"
	return ""
}

// Run creates an instance of the TiDB container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
