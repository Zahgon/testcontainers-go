package clickhouse

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	defaultUser         = "default"
	defaultDatabaseName = "clickhouse"
)

const (
	// containerPorts {
	httpPort   = "8123/tcp"
	nativePort = "9000/tcp"
	// }
)

// ClickHouseContainer represents the ClickHouse container type used in the module
type ClickHouseContainer struct {
	testcontainers.Container
	DbName   string //nolint:staticcheck //FIXME
	User     string
	Password string
}

// ConnectionHost returns the host and port of the clickhouse container, using the default, native 9000 port, and
// obtaining the host and exposed port from the container
func (c *ClickHouseContainer) ConnectionHost(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ConnectionString returns the dsn string for the clickhouse container, using the default, native 9000 port, and
// obtaining the host and exposed port from the container. It also accepts a variadic list of extra arguments
// which will be appended to the dsn string. The format of the extra arguments is the same as the
// connection string format, e.g. "dial_timeout=300ms" or "skip_verify=false"
func (c *ClickHouseContainer) ConnectionString(ctx context.Context, args ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Deprecated: use Run instead
// RunContainer creates an instance of the ClickHouse container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*ClickHouseContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the ClickHouse container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*ClickHouseContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// initialize the credentials

// refresh the credentials from the environment variables
