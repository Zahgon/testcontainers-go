package azurite

import (
	"github.com/testcontainers/testcontainers-go"
)

type options struct {
	// EnabledServices is a list of services that should be enabled
	EnabledServices []Service
}

func defaultOptions() options { _ = "STUB: not implemented"; return *new(options) }

// Satisfy the testcontainers.ContainerCustomizer interface
var _ testcontainers.ContainerCustomizer = (Option)(nil)

// Option is an option for the Azurite container.
type Option func(*options) error

// Customize is a NOOP. It's defined to satisfy the testcontainers.ContainerCustomizer interface.
func (o Option) Customize(*testcontainers.GenericContainerRequest) error {
	_ = "STUB: not implemented"
	// NOOP to satisfy interface.
	return nil
}

// WithEnabledServices is a custom option to specify which services should be enabled.
func WithEnabledServices(services ...Service) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// valid service, continue

// WithInMemoryPersistence is a custom option to enable in-memory persistence for Azurite.
// This option is only available for Azurite v3.28.0 and later.
func WithInMemoryPersistence(megabytes float64) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}
