package redis

import (
	"context"
	"crypto/tls"

	"github.com/testcontainers/testcontainers-go"
)

type LogLevel string

const (
	// redisPort is the port for the Redis connection
	redisPort = "6379/tcp"

	// LogLevelDebug is the debug log level
	LogLevelDebug LogLevel = "debug"
	// LogLevelVerbose is the verbose log level
	LogLevelVerbose LogLevel = "verbose"
	// LogLevelNotice is the notice log level
	LogLevelNotice LogLevel = "notice"
	// LogLevelWarning is the warning log level
	LogLevelWarning LogLevel = "warning"
)

type RedisContainer struct {
	testcontainers.Container
	settings options
}

// ConnectionString returns the connection string for the Redis container.
// It uses the default 6379 port.
func (c *RedisContainer) ConnectionString(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// TLSConfig returns the TLS configuration for the Redis container, nil if TLS is not enabled.
func (c *RedisContainer) TLSConfig() *tls.Config { _ = "STUB: not implemented"; return nil }

// Deprecated: use Run instead
// RunContainer creates an instance of the Redis container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*RedisContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the Redis container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*RedisContainer, error) {
	_ = "STUB: not implemented"
	// Process custom options to extract settings
	return nil, nil
}

// Generate TLS certificates in the fly and add them to the container before it starts.
// Update the CMD to use the TLS certificates.

// Update the CMD to use the TLS certificates.

// Disable the default port, as described in https://redis.io/docs/latest/operate/oss_and_stack/management/security/encryption/#running-manually

// Match the server cert's common name

// Append the customizers passed to the Run function.
