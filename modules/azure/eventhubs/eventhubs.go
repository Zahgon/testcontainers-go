package eventhubs

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/azure/azurite"
)

const (
	defaultAMPQPort        = "5672/tcp"
	defaultHTTPPort        = "5300/tcp"
	connectionStringFormat = "Endpoint=sb://%s;SharedAccessKeyName=%s;SharedAccessKey=%s;UseDevelopmentEmulator=true;"

	// aliasEventhubs is the alias for the eventhubs container in the network
	aliasEventhubs = "eventhubs"

	// aliasAzurite is the alias for the azurite container in the network
	aliasAzurite = "azurite"

	// containerConfigFile is the path to the eventhubs config file
	containerConfigFile = "/Eventhubs_Emulator/ConfigFiles/Config.json"
)

// Container represents the Azure Event Hubs container type used in the module
type Container struct {
	testcontainers.Container
	azuriteOptions *options
}

// AzuriteContainer returns the azurite container that is used by the eventhubs container
func (c *Container) AzuriteContainer() *azurite.Container { _ = "STUB: not implemented"; return nil }

// Terminate terminates the eventhubs container, the azurite container, and the network to communicate between them.
func (c *Container) Terminate(ctx context.Context, opts ...testcontainers.TerminateOption) error {
	_ = "STUB: not implemented"
	return nil
}

// terminate the eventhubs container

// terminate the azurite container if it was created

// remove the azurite network if it was created

// Run creates an instance of the Azure Event Hubs container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	// Process custom options first to extract settings
	return nil, nil
}

// Build moduleOpts with defaults

// start the azurite container first

// apply the network to the eventhubs container

// validate the EULA after all the options are applied

// ConnectionString returns the connection string for the eventhubs container,
// using the following format:
// Endpoint=sb://<hostname>:<port>;SharedAccessKeyName=<key-name>;SharedAccessKey=<key>;UseDevelopmentEmulator=true;
func (c *Container) ConnectionString(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	// we are passing an empty proto to get the host:port string
	return "", nil
}
