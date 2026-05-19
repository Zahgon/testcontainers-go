package testcontainers

import (
	"context"
	"sync"

	"github.com/moby/moby/client"
)

// DockerClient is a wrapper around the docker client that is used by testcontainers-go.
// It implements the SystemAPIClient interface in order to cache the docker info and reuse it.
type DockerClient struct {
	*client.Client // client is embedded into our own client
}

var (
	// dockerInfo stores the docker info to be reused in the Info method
	dockerInfo     client.SystemInfoResult
	dockerInfoSet  bool
	dockerInfoLock sync.Mutex
)

// implements SystemAPIClient interface
var _ client.SystemAPIClient = &DockerClient{}

// Events returns a channel to listen to events that happen to the docker daemon.
func (c *DockerClient) Events(ctx context.Context, options client.EventsListOptions) client.EventsResult {
	_ = "STUB: not implemented"
	return *new(client.EventsResult)
}

// Info returns information about the docker server. The result of Info is cached
// and reused every time Info is called.
// It will also print out the docker server info, and the resolved Docker paths, to the default logger.
func (c *DockerClient) Info(ctx context.Context, options client.InfoOptions) (client.SystemInfoResult, error) {
	_ = "STUB: not implemented"
	return *new(client.SystemInfoResult), nil
}

// RegistryLogin logs into a Docker registry.
func (c *DockerClient) RegistryLogin(ctx context.Context, options client.RegistryLoginOptions) (client.RegistryLoginResult, error) {
	_ = "STUB: not implemented"
	return *new(client.RegistryLoginResult), nil
}

// DiskUsage returns the disk usage of all images.
func (c *DockerClient) DiskUsage(ctx context.Context, options client.DiskUsageOptions) (client.DiskUsageResult, error) {
	_ = "STUB: not implemented"
	return *new(client.DiskUsageResult), nil
}

// Ping pings the docker server.
func (c *DockerClient) Ping(ctx context.Context, options client.PingOptions) (client.PingResult, error) {
	_ = "STUB: not implemented"
	return *new(client.PingResult), nil
}

// Deprecated: Use NewDockerClientWithOpts instead.
func NewDockerClient() (*client.Client, error) { _ = "STUB: not implemented"; return nil, nil }

func NewDockerClientWithOpts(ctx context.Context, opt ...client.Opt) (*DockerClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Fallback to environment, including the original options
