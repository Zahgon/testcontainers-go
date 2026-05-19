package valkey

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

// Option is an option for the Valkey container.
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

// createTLSCerts creates a CA certificate, a client certificate and a Valkey certificate.
func createTLSCerts() (caCert *tlscert.Certificate, clientCert *tlscert.Certificate, valkeyCert *tlscert.Certificate, err error) {
	_ = "STUB: not implemented"
	// ips is the extra list of IPs to include in the certificates.
	// It's used to allow the client and Valkey certificates to be used in the same host
	// when the tests are run using a remote docker daemon.
	return nil, nil, nil, nil
}

// Generate CA certificate

// Generate client certificate

// Generate Valkey certificate
