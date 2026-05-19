package neo4j

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	// containerPorts {
	defaultBoltPort  = "7687/tcp"
	defaultHTTPPort  = "7474/tcp"
	defaultHTTPSPort = "7473/tcp"
	// }
)

// Neo4jContainer represents the Neo4j container type used in the module
type Neo4jContainer struct {
	testcontainers.Container
}

// BoltUrl returns the bolt url for the Neo4j container, using the bolt port, in the format of neo4j://host:port
//
//nolint:revive,staticcheck //FIXME
func (c Neo4jContainer) BoltUrl(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Deprecated: use Run instead
// RunContainer creates an instance of the Neo4j container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*Neo4jContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the Neo4j container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Neo4jContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isHTTPOk() func(status int) bool { _ = "STUB: not implemented"; return nil }
