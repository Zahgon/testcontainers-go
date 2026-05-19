package eventhubs

import (
	"io"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/azure/azurite"
)

type options struct {
	azuriteImage     string
	azuriteOptions   []testcontainers.ContainerCustomizer
	azuriteContainer *azurite.Container
	network          *testcontainers.DockerNetwork
}

func defaultOptions() options { _ = "STUB: not implemented"; return *new(options) }

// Satisfy the testcontainers.ContainerCustomizer interface
var _ testcontainers.ContainerCustomizer = (Option)(nil)

// Option is an option for the EventHubs container.
type Option func(*options) error

// Customize is a NOOP. It's defined to satisfy the testcontainers.ContainerCustomizer interface.
func (o Option) Customize(*testcontainers.GenericContainerRequest) error {
	_ = "STUB: not implemented"
	// NOOP to satisfy interface.
	return nil
}

// WithAzurite sets the image and options for the Azurite container.
// By default, the image is "mcr.microsoft.com/azure-storage/azurite:3.33.0".
func WithAzurite(img string, opts ...testcontainers.ContainerCustomizer) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithAcceptEULA sets the ACCEPT_EULA environment variable to "Y" for the eventhubs container.
func WithAcceptEULA() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithConfig sets the eventhubs config file for the eventhubs container,
// copying the content of the reader to the container file at
// "/Eventhubs_Emulator/ConfigFiles/Config.json".
func WithConfig(r io.Reader) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// validateEula validates that the EULA is accepted for the eventhubs container.
func validateEula() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}
