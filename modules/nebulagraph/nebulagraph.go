package nebulagraph

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

// Cluster represents a running NebulaGraph cluster for testing
type Cluster struct {
	graphd   testcontainers.Container
	metad    testcontainers.Container
	storaged testcontainers.Container
	network  *testcontainers.DockerNetwork
}

// RunCluster starts a NebulaGraph cluster (metad, storaged, graphd and activator) containers within a Docker network
func RunCluster(ctx context.Context,
	graphdImg string, graphdCustomizers []testcontainers.ContainerCustomizer,
	storagedImg string, storagedCustomizers []testcontainers.ContainerCustomizer,
	metadImg string, metadCustomizers []testcontainers.ContainerCustomizer,
) (*Cluster, error) {
	_ = "STUB: not implemented"
	// 1. Create a custom network
	return nil, nil
}

// 2. Start metad

// 3. Start graphd (needed for storage registration)

// 4. Start storaged

// 5. Run storage registration command with retry logic

// ConnectionString returns the host:port for connecting to NebulaGraph graphd
func (c *Cluster) ConnectionString(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Terminate stops all NebulaGraph containers
func (c *Cluster) Terminate(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func terminateContainersAndRemoveNetwork(ctx context.Context, netRes *testcontainers.DockerNetwork, containers ...testcontainers.Container) []error {
	_ = "STUB: not implemented"
	return nil
}

// Retry network removal: after container termination, Docker may not
// have fully disconnected endpoints yet, causing "has active endpoints".
