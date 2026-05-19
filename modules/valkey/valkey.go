package valkey

import (
	"context"
	"crypto/tls"

	"github.com/testcontainers/testcontainers-go"
)

// ValkeyContainer represents the Valkey container type used in the module
type ValkeyContainer struct {
	testcontainers.Container
	settings options
}

// valkeyServerProcess is the name of the valkey server process
const valkeyServerProcess = "valkey-server"

type LogLevel string

const (
	// valkeyPort is the port for the Valkey connection
	valkeyPort = "6379/tcp"

	// LogLevelDebug is the debug log level
	LogLevelDebug LogLevel = "debug"
	// LogLevelVerbose is the verbose log level
	LogLevelVerbose LogLevel = "verbose"
	// LogLevelNotice is the notice log level
	LogLevelNotice LogLevel = "notice"
	// LogLevelWarning is the warning log level
	LogLevelWarning LogLevel = "warning"
)

// ConnectionString returns the connection string for the Valkey container
func (c *ValkeyContainer) ConnectionString(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// TLSConfig returns the TLS configuration for the Valkey container, nil if TLS is not enabled.
func (c *ValkeyContainer) TLSConfig() *tls.Config { _ = "STUB: not implemented"; return nil }

// Deprecated: use Run instead
// RunContainer creates an instance of the Valkey container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*ValkeyContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the Valkey container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*ValkeyContainer, error) {
	_ = "STUB: not implemented"
	// Process custom options first
	return nil, nil
}

// wait for the TLS port to be available

// Generate TLS certificates in the fly and add them to the container before it starts.
// Update the CMD to use the TLS certificates.

// Update the CMD to use the TLS certificates.

// Disable the default port, as described in https://redis.io/docs/latest/operate/oss_and_stack/management/security/encryption/#running-manually

// Match the server cert's common name

// WithConfigFile sets the config file to be used for the valkey container.
// The config file must be the first argument to valkey-server.
func WithConfigFile(configFile string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// Prepend the config file as the first argument

// WithLogLevel sets the log level for the valkey server process
// See https://redis.io/docs/reference/modules/modules-api-ref/#redismodule_log for more information.
func WithLogLevel(level LogLevel) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithSnapshotting sets the snapshotting configuration for the valkey server process. You can configure Valkey to have it
// save the dataset every N seconds if there are at least M changes in the dataset.
// This method allows Valkey to benefit from copy-on-write semantics.
// See https://redis.io/docs/management/persistence/#snapshotting for more information.
func WithSnapshotting(seconds int, changedKeys int) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}
