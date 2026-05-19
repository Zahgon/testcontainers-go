package servicebus

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mssql"
)

const (
	defaultPort                = "5672/tcp"
	defaultHTTPPort            = "5300/tcp"
	defaultSharedAccessKeyName = "RootManageSharedAccessKey"
	defaultSharedAccessKey     = "SAS_KEY_VALUE"
	connectionStringFormat     = "Endpoint=sb://%s;SharedAccessKeyName=%s;SharedAccessKey=%s;UseDevelopmentEmulator=true;"

	// aliasServiceBus is the alias for the servicebus container in the network
	aliasServiceBus = "servicebus"

	// aliasMSSQL is the alias for the mssql network
	aliasMSSQL = "mssql"

	// defaultMSSQLImage is the default image for the mssql container
	defaultMSSQLImage = "mcr.microsoft.com/mssql/server:2022-CU14-ubuntu-22.04"

	// containerConfigFile is the path to the config file for the servicebus container
	containerConfigFile = "/ServiceBus_Emulator/ConfigFiles/Config.json"
)

// Container represents the Azure ServiceBus container type used in the module
type Container struct {
	testcontainers.Container
	mssqlOptions *options
}

// MSSQLContainer returns the mssql container that is used by the servicebus container
func (c *Container) MSSQLContainer() *mssql.MSSQLServerContainer {
	_ = "STUB: not implemented"
	return nil
}

// Terminate terminates the servicebus container, the mssql container, and the network to communicate between them.
func (c *Container) Terminate(ctx context.Context, opts ...testcontainers.TerminateOption) error {
	_ = "STUB: not implemented"
	return nil
}

// terminate the servicebus container

// terminate the mssql container if it was created

// remove the mssql network if it was created

// Run creates an instance of the Azure ServiceBus container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	// Process custom options first to extract settings
	return nil, nil
}

// Build moduleOpts with defaults

// default is zero because the MSSQL container is started first

// Start the mssql container first. The EULA is accepted by default, as it is required by the servicebus emulator.

// apply the network to the servicebus container

// validate the EULA after all the options are applied

// ConnectionString returns the connection string for the servicebus container,
// using the following format:
// Endpoint=sb://<hostname>:<port>;SharedAccessKeyName=<key-name>;SharedAccessKey=<key>;UseDevelopmentEmulator=true;
func (c *Container) ConnectionString(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	// we are passing an empty proto to get the host:port string
	return "", nil
}
