package cassandra

import (
	"context"
	"crypto/tls"
	_ "embed"

	"github.com/testcontainers/testcontainers-go"
)

const (
	port    = "9042/tcp"
	sslPort = "9142/tcp"
)

//go:embed testdata/cassandra-ssl.yaml
var sslConfigYAML []byte

// CassandraContainer represents the Cassandra container type used in the module
type CassandraContainer struct {
	testcontainers.Container
	settings options
}

// ConnectionHost returns the host and port of the cassandra container, using the default, native 9042 port, and
// obtaining the host and exposed port from the container
func (c *CassandraContainer) ConnectionHost(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// TLSConfig returns the TLS configuration for secure client connections.
// Returns an error if TLS is not enabled on the container.
func (c *CassandraContainer) TLSConfig() (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WithConfigFile sets the YAML config file to be used for the cassandra container
// It will also set the "configFile" parameter to the path of the config file
// as a command line argument to the container.
func WithConfigFile(configFile string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithInitScripts sets the init cassandra queries to be run when the container starts
func WithInitScripts(scripts ...string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// Deprecated: use Run instead
// RunContainer creates an instance of the Cassandra container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*CassandraContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the Cassandra container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*CassandraContainer, error) {
	_ = "STUB: not implemented"
	// Process custom options to extract settings
	return nil, nil
}

// Configure TLS if enabled

// Store the TLS config for client connections

// Add SSL port and configure networking for SSL
// We need CASSANDRA_BROADCAST_RPC_ADDRESS when CASSANDRA_RPC_ADDRESS is 0.0.0.0

// Mount the SSL config and keystore

// Update wait strategy to also wait for SSL port
