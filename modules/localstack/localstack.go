package localstack

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	defaultPort            = 4566
	hostnameExternalEnvVar = "HOSTNAME_EXTERNAL"
	localstackHostEnvVar   = "LOCALSTACK_HOST"
)

var recentVersionTags = []string{
	"community-archive",
	"latest",
	"s3",
	"s3-latest",
	"stable",
}

func isMinimumVersion(image string, minVersion string) bool {
	_ = "STUB: not implemented"
	return false
}

// Re-check after stripping the arch suffix (e.g. "community-archive-amd64" -> "community-archive").

// WithNetwork creates a network with the given name and attaches the container to it, setting the network alias
// on that network to the given alias.
//
// Deprecated: use network.WithNetwork or network.WithNewNetwork instead
func WithNetwork(_ string, alias string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// Deprecated: use Run instead
// RunContainer creates an instance of the LocalStack container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*LocalStackContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the LocalStack container type
// - overrideReq: a function that can be used to override the default container request, usually used to set the image version, environment variables for localstack, etc.
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*LocalStackContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// configure the docker host after all the options have been applied

// Deprecated: use RunContainer instead
// StartContainer creates an instance of the LocalStack container type, being possible to pass a custom request and options:
// - overrideReq: a function that can be used to override the default container request, usually used to set the image version, environment variables for localstack, etc.
func StartContainer(ctx context.Context, overrideReq OverrideContainerRequestOption) (*LocalStackContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// configureDockerHost returns an option that configures the docker host environment variable
func configureDockerHost(ctx context.Context, envVar string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

func setDockerHost(ctx context.Context, req *testcontainers.GenericContainerRequest, envVar string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// if the container is not connected to the default network, use the last network alias in the first network
// for that we need to check if the container is connected to a network and if it has network aliases
