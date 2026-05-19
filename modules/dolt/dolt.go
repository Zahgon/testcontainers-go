package dolt

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

// DoltContainer represents the Dolt container type used in the module
type DoltContainer struct {
	testcontainers.Container
	username string
	password string
	database string
}

// Deprecated: this function will be removed in the next major release.
func WithDefaultCredentials() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// withDefaultCredentials is the function that applies the default credentials to the container request.
// In case the provided username is the root user, the credentials will be removed.
func withDefaultCredentials() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// Deprecated: use Run instead
// RunContainer creates an instance of the Couchbase container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*DoltContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the Dolt container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*DoltContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// refresh the credentials from the environment variables

// withCredentials found the root user

// dolthub/dolt-sql-server does not create user or database, so we do so here

func (c *DoltContainer) initialize(ctx context.Context, createUser bool) error {
	_ = "STUB: not implemented"
	return nil
}

// create database

// create user

// grant user privileges

func (c *DoltContainer) initialConnectionString(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *DoltContainer) MustConnectionString(ctx context.Context, args ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *DoltContainer) ConnectionString(ctx context.Context, args ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func WithUsername(username string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

func WithPassword(password string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

func WithDoltCredsPublicKey(key string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

//nolint:revive,staticcheck //FIXME
func WithDoltCloneRemoteUrl(url string) testcontainers.CustomizeRequestOption {
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

func WithCredsFile(credsFile string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

func WithScripts(scripts ...string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}
