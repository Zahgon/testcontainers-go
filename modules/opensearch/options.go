package opensearch

import "github.com/testcontainers/testcontainers-go"

// Options is a struct for specifying options for the OpenSearch container.
type Options struct {
	Password string
	Username string
}

func defaultOptions() *Options { _ = "STUB: not implemented"; return nil }

// Compiler check to ensure that Option implements the testcontainers.ContainerCustomizer interface.
var _ testcontainers.ContainerCustomizer = (*Option)(nil)

// Option is an option for the OpenSearch container.
type Option func(*Options) error

// Customize is a NOOP. It's defined to satisfy the testcontainers.ContainerCustomizer interface.
func (o Option) Customize(*testcontainers.GenericContainerRequest) error {
	_ = "STUB: not implemented"
	// NOOP to satisfy interface.
	return nil
}

// WithPassword sets the password for the OpenSearch container.
func WithPassword(password string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUsername sets the username for the OpenSearch container.
func WithUsername(username string) Option { _ = "STUB: not implemented"; return *new(Option) }
