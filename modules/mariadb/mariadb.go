package mariadb

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	rootUser            = "root"
	defaultUser         = "test"
	defaultPassword     = "test"
	defaultDatabaseName = "test"
)

// MariaDBContainer represents the MariaDB container type used in the module
type MariaDBContainer struct {
	testcontainers.Container
	username string
	password string
	database string
}

// WithDefaultCredentials applies the default credentials to the container request.
// It will look up for MARIADB environment variables.
func WithDefaultCredentials() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// https://github.com/docker-library/docs/tree/master/mariadb#environment-variables
// From tag 10.2.38, 10.3.29, 10.4.19, 10.5.10 onwards, and all 10.6 and later tags,
// the MARIADB_* equivalent variables are provided. MARIADB_* variants will always be
// used in preference to MYSQL_* variants.
func withMySQLEnvVars() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// look up for MARIADB environment variables and apply the same to MYSQL

// apply the same value to the MYSQL environment variables

func WithUsername(username string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

func WithPassword(password string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

func WithDatabase(database string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

func WithConfigFile(configFile string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

func WithScripts(scripts ...string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// Deprecated: use Run instead
// RunContainer creates an instance of the MariaDB container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*MariaDBContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the MariaDB container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*MariaDBContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Apply MySQL environment variables after user customization
// In future releases of MariaDB, they could remove the MYSQL_* environment variables
// at all. Then we can remove this customization.

// validate credentials before running the container

// Inspect the container to get environment variables

// MustConnectionString panics if the address cannot be determined.
func (c *MariaDBContainer) MustConnectionString(ctx context.Context, args ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *MariaDBContainer) ConnectionString(ctx context.Context, args ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// validateCredentials validates the credentials to ensure that an empty password
// is only used with the root user.
func validateCredentials() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}
