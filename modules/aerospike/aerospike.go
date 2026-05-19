package aerospike

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	// port is the port used for client connections
	port = "3000/tcp"
	// fabricPort is the port used for Intra-cluster communication port.
	// Replica writes, migrations, and other node-to-node communications use the Fabric port.
	fabricPort = "3001/tcp"
	// heartbeatPort is the port used for heartbeat communication
	// between nodes in the Aerospike cluster
	heartbeatPort = "3002/tcp"
	// infoPort is the port used for info commands
	infoPort = "3003/tcp"
)

// Container is the Aerospike container type used in the module
type Container struct {
	testcontainers.Container
}

// Run creates an instance of the Aerospike container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
