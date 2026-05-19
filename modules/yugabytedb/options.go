package yugabytedb

import (
	"github.com/testcontainers/testcontainers-go"
)

// WithDatabaseName sets the initial database name for the yugabyteDB container.
func WithDatabaseName(dbName string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithDatabaseUser sets the initial database user for the yugabyteDB container.
func WithDatabaseUser(dbUser string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithDatabasePassword sets the initial database password for the yugabyteDB container.
func WithDatabasePassword(dbPassword string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithKeyspace sets the initial keyspace for the yugabyteDB container.
func WithKeyspace(keyspace string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithUser sets the initial user for the yugabyteDB container.
func WithUser(user string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithPassword sets the initial password for the yugabyteDB container.
func WithPassword(password string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}
