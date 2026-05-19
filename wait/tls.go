package wait

import (
	"context"
	"crypto/tls"
	"time"
)

// Validate we implement interface.
var _ Strategy = (*TLSStrategy)(nil)

// TLSStrategy is a strategy for handling TLS.
type TLSStrategy struct {
	// General Settings.
	timeout      *time.Duration
	pollInterval time.Duration

	// Custom Settings.
	certFiles *x509KeyPair
	rootFiles []string

	// State.
	tlsConfig *tls.Config
}

// x509KeyPair is a pair of certificate and key files.
type x509KeyPair struct {
	certPEMFile string
	keyPEMFile  string
}

// ForTLSCert returns a CertStrategy that will add a Certificate to the [tls.Config]
// constructed from PEM formatted certificate key file pair in the container.
func ForTLSCert(certPEMFile, keyPEMFile string) *TLSStrategy { _ = "STUB: not implemented"; return nil }

// ForTLSRootCAs returns a CertStrategy that sets the root CAs for the [tls.Config]
// using the given PEM formatted files from the container.
func ForTLSRootCAs(pemFiles ...string) *TLSStrategy { _ = "STUB: not implemented"; return nil }

// WithRootCAs sets the root CAs for the [tls.Config] using the given files from
// the container.
func (ws *TLSStrategy) WithRootCAs(files ...string) *TLSStrategy {
	_ = "STUB: not implemented"
	return nil
}

// WithCert sets the [tls.Config] Certificates using the given files from the container.
func (ws *TLSStrategy) WithCert(certPEMFile, keyPEMFile string) *TLSStrategy {
	_ = "STUB: not implemented"
	return nil
}

// WithServerName sets the server for the [tls.Config].
func (ws *TLSStrategy) WithServerName(serverName string) *TLSStrategy {
	_ = "STUB: not implemented"
	return nil
}

// WithStartupTimeout can be used to change the default startup timeout.
func (ws *TLSStrategy) WithStartupTimeout(startupTimeout time.Duration) *TLSStrategy {
	_ = "STUB: not implemented"
	return nil
}

// WithPollInterval can be used to override the default polling interval of 100 milliseconds.
func (ws *TLSStrategy) WithPollInterval(pollInterval time.Duration) *TLSStrategy {
	_ = "STUB: not implemented"
	return nil
}

// TLSConfig returns the TLS config once the strategy is ready.
// If the strategy is nil, it returns nil.
func (ws *TLSStrategy) TLSConfig() *tls.Config { _ = "STUB: not implemented"; return nil }

// String returns a human-readable description of the wait strategy.
func (ws *TLSStrategy) String() string { _ = "STUB: not implemented"; return "" }

// WaitUntilReady implements the [Strategy] interface.
// It waits for the CA, client cert and key files to be available in the container and
// uses them to setup the TLS config.
func (ws *TLSStrategy) WaitUntilReady(ctx context.Context, target StrategyTarget) error {
	_ = "STUB: not implemented"
	return nil
}
