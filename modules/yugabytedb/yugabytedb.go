package yugabytedb

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	ycqlPort = "9042/tcp"

	ycqlKeyspaceEnv = "YCQL_KEYSPACE"
	ycqlUserNameEnv = "YCQL_USER"
	ycqlPasswordEnv = "YCQL_PASSWORD"

	ycqlKeyspace = "yugabyte"
	ycqlUserName = "yugabyte"
	ycqlPassword = "yugabyte"
)

const (
	ysqlPort = "5433/tcp"

	ysqlDatabaseNameEnv     = "YSQL_DB"
	ysqlDatabaseUserEnv     = "YSQL_USER"
	ysqlDatabasePasswordEnv = "YSQL_PASSWORD"

	ysqlDatabaseName     = "yugabyte"
	ysqlDatabaseUser     = "yugabyte"
	ysqlDatabasePassword = "yugabyte"
)

// Container represents the yugabyteDB container type used in the module
type Container struct {
	testcontainers.Container

	ysqlDatabaseName     string
	ysqlDatabaseUser     string
	ysqlDatabasePassword string
}

// Run creates an instance of the yugabyteDB container type and automatically starts it.
// A default configuration is used for the container, but it can be customized using the
// provided options.
// When using default configuration values it is recommended to use the provided
// [*Container.YSQLConnectionString] and [*Container.YCQLConfigureClusterConfig]
// methods to use the container in their respective clients.
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Inspect the container to get the actual env var values after user customizations

// YSQLConnectionString returns a connection string for the yugabyteDB container
// using the configured database name, user, password, port, host and additional
// arguments.
// Additional arguments are appended to the connection string as query parameters
// in the form of key=value pairs separated by "&".
func (y *Container) YSQLConnectionString(ctx context.Context, args ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
