package redis

import (
	"crypto/tls"

	"github.com/mdelapenya/tlscert"

	"github.com/testcontainers/testcontainers-go"
)

type options struct {
	tlsEnabled bool
	tlsConfig  *tls.Config
}

// Compiler check to ensure that Option implements the testcontainers.ContainerCustomizer interface.
var _ testcontainers.ContainerCustomizer = (Option)(nil)

// Option is an option for the Redis container.
type Option func(*options) error

// Customize is a NOOP. It's defined to satisfy the testcontainers.ContainerCustomizer interface.
func (o Option) Customize(*testcontainers.GenericContainerRequest) error {
	_ = "STUB: not implemented"
	// NOOP to satisfy interface.
	return nil
}

// WithTLS sets the TLS configuration for the redis container, setting
// the 6380/tcp port to listen on for TLS connections and using a secure URL (rediss://).
func WithTLS() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithConfigFile sets the config file to be used for the redis container, and sets the command to run the redis server
// using the passed config file
func WithConfigFile(configFile string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// prepend the command to run the redis server with the config file, which must be the first argument of the redis server process

// WithLogLevel sets the log level for the redis server process
// See "[RedisModule_Log]" for more information.
//
// [RedisModule_Log]: https://redis.io/docs/reference/modules/modules-api-ref/#redismodule_log
func WithLogLevel(level LogLevel) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithSnapshotting sets the snapshotting configuration for the redis server process. You can configure Redis to have it
// save the dataset every N seconds if there are at least M changes in the dataset.
// This method allows Redis to benefit from copy-on-write semantics.
// See [Snapshotting] for more information.
//
// [Snapshotting]: https://redis.io/docs/management/persistence/#snapshotting
func WithSnapshotting(seconds int, changedKeys int) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// createTLSCerts creates a CA certificate, a client certificate and a Redis certificate.
func createTLSCerts() (caCert *tlscert.Certificate, clientCert *tlscert.Certificate, redisCert *tlscert.Certificate, err error) {
	_ = "STUB: not implemented"
	// ips is the extra list of IPs to include in the certificates.
	// It's used to allow the client and Redis certificates to be used in the same host
	// when the tests are run using a remote docker daemon.
	return nil, nil, nil, nil
}

// Generate CA certificate

// Generate client certificate

// Generate Redis certificate
