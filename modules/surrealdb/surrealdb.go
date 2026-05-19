package surrealdb

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

// SurrealDBContainer represents the SurrealDB container type used in the module
type SurrealDBContainer struct {
	testcontainers.Container
}

// URL returns the connection string for the OpenLDAP container
func (c *SurrealDBContainer) URL(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// WithUsername sets the initial username to be created when the container starts
// It is used in conjunction with WithPassword to set a username and its password.
// It will create the specified user with superuser power.
func WithUsername(username string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithPassword sets the initial password of the user to be created when the container starts
// It is used in conjunction with WithUsername to set a username and its password.
// It will set the superuser password for SurrealDB.
func WithPassword(password string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithAuthentication enables authentication for the SurrealDB instance
func WithAuthentication() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithStrictMode enables strict mode for the SurrealDB instance
func WithStrictMode() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithAllowAllCaps enables all caps for the SurrealDB instance
func WithAllowAllCaps() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// Deprecated: use Run instead
// RunContainer creates an instance of the SurrealDB container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*SurrealDBContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the SurrealDB container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*SurrealDBContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
