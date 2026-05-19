package arangodb

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	defaultPort = "8529/tcp"

	// DefaultUser is the default username for the ArangoDB container.
	// This is the username to be used when connecting to the ArangoDB instance.
	DefaultUser = "root"

	defaultPassword = "root"
)

// Container represents the ArangoDB container type used in the module
type Container struct {
	testcontainers.Container
	password string
}

// Credentials returns the credentials for the ArangoDB container:
// first return value is the username, second is the password.
func (c *Container) Credentials() (string, string) { _ = "STUB: not implemented"; return "", "" }

// HTTPEndpoint returns the HTTP endpoint of the ArangoDB container, using the following format: `http://$host:$port`.
func (c *Container) HTTPEndpoint(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Run creates an instance of the ArangoDB container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// configure the wait strategy after all the options have been applied
