package cosmosdb

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	defaultPort     = "8081/tcp"
	defaultProtocol = "http"

	// Well-known, publicly documented account key for the Azure CosmosDB Emulator.
	// See: https://learn.microsoft.com/en-us/azure/cosmos-db/how-to-develop-emulator
	testAccKey = "C2y6yDjf5/R+ob0N8A7Cgv30VRDJIWEHLM+4QDU5DE2nQ9nDuVTqobD4b8mGGyPMbIZnqyMsEcaGQy67XIw/Jw=="
)

// Container represents the CosmosDB container type used in the module
type Container struct {
	testcontainers.Container
}

// Run creates an instance of the CosmosDB container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	// Initialize with module defaults
	return nil, nil
}

// Add user-provided options

// ConnectionString returns a connection string that can be used to connect to the CosmosDB emulator.
// The connection string includes the account endpoint (host:port) and the default test account key.
// It returns an error if the port endpoint cannot be determined.
//
// Format: "AccountEndpoint=<host>:<port>;AccountKey=<accountKey>"
func (c *Container) ConnectionString(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
