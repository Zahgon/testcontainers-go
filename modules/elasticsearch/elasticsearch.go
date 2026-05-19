package elasticsearch

import (
	"context"
	"crypto/x509"
	"io"

	"github.com/testcontainers/testcontainers-go"
)

const (
	defaultHTTPPort   = "9200"
	defaultTCPPort    = "9300"
	defaultPassword   = "changeme"
	defaultUsername   = "elastic"
	defaultCaCertPath = "/usr/share/elasticsearch/config/certs/http_ca.crt"
	envPassword       = "ELASTIC_PASSWORD"
)

const (
	// Deprecated: it will be removed in the next major version
	DefaultBaseImage = "docker.elastic.co/elasticsearch/elasticsearch"
	// Deprecated: it will be removed in the next major version
	DefaultBaseImageOSS = "docker.elastic.co/elasticsearch/elasticsearch-oss"
)

// ElasticsearchContainer represents the Elasticsearch container type used in the module
type ElasticsearchContainer struct {
	testcontainers.Container
	Settings Options
}

// Deprecated: use Run instead
// RunContainer creates an instance of the Elasticsearch container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*ElasticsearchContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the Elasticsearch container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*ElasticsearchContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Gather all config options (defaults and then apply provided options)

// Transfer the password settings to the container request

// Set the default waiting strategy if not already set.

// certWriter is a helper that writes the details of a CA cert to options.
type certWriter struct {
	options  *Options
	certPool *x509.CertPool
}

// Read reads the CA cert from the reader and appends it to the options.
func (w *certWriter) Read(r io.Reader) error { _ = "STUB: not implemented"; return nil }

// configureWaitFor sets the req.WaitingFor strategy based on settings.
func configureWaitFor(options *Options) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// configureAddress sets the address of the Elasticsearch container.
// If the certificate is set, it will use https as protocol, otherwise http.
func (c *ElasticsearchContainer) configureAddress(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// sslRequired returns true if the SSL is required, otherwise false.
func sslRequired(req *testcontainers.GenericContainerRequest) bool {
	_ = "STUB: not implemented"
	return false
}

// These configuration keys explicitly disable CA generation.
// If any are set we skip the file retrieval.

// configurePassword transfers the password settings to the container request.
// If the password is not set, it will be set to "changeme" for Elasticsearch 8
func configurePassword(settings *Options) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// set "changeme" as default password for Elasticsearch 8

// major version 8 is secure by default and does not need this to enable authentication

// configureJvmOpts sets the default memory of the Elasticsearch instance to 2GB.
// This functions, which is only available since version 7, is called as a post create hook
// for the container request.
func configureJvmOpts() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// Sets default memory of elasticsearch instance to 2GB

// The temp file is closed to not leak a file descriptor.

// clean up

// Spaces are deliberate to allow user to define additional jvm options as elasticsearch resolves option files lexicographically
