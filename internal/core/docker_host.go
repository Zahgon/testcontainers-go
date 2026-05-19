package core

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/moby/moby/client"
)

type dockerHostContext string

var DockerHostContextKey = dockerHostContext("docker_host")

var (
	ErrDockerHostNotSet               = errors.New("DOCKER_HOST is not set")
	ErrDockerSocketOverrideNotSet     = errors.New("TESTCONTAINERS_DOCKER_SOCKET_OVERRIDE is not set")
	ErrDockerSocketNotSetInContext    = errors.New("socket not set in context")
	ErrDockerSocketNotSetInProperties = errors.New("socket not set in ~/.testcontainers.properties")
	ErrNoUnixSchema                   = errors.New("URL schema is not unix")
	ErrSocketNotFound                 = errors.New("socket not found")
	ErrSocketNotFoundInPath           = errors.New("docker socket not found in " + DockerSocketPath)
	// ErrTestcontainersHostNotSetInProperties this error is specific to Testcontainers
	ErrTestcontainersHostNotSetInProperties = errors.New("tc.host not set in ~/.testcontainers.properties")
)

var (
	dockerHostCache    string
	dockerHostErrCache error
	dockerHostOnce     sync.Once
)

var (
	dockerSocketPathCache string
	dockerSocketPathOnce  sync.Once
)

// deprecated
// see https://github.com/testcontainers/testcontainers-java/blob/main/core/src/main/java/org/testcontainers/dockerclient/DockerClientConfigUtils.java#L46
func DefaultGatewayIP() (string, error) {
	_ = "STUB: not implemented"
	// see https://github.com/testcontainers/testcontainers-java/blob/3ad8d80e2484864e554744a4800a81f6b7982168/core/src/main/java/org/testcontainers/dockerclient/DockerClientConfigUtils.java#L27
	return "", nil
}

// dockerHostCheck Use a vanilla Docker client to check if the Docker host is reachable.
// It will avoid recursive calls to this function.
var dockerHostCheck = func(ctx context.Context, host string) error {
	cli, err := client.New(client.FromEnv, client.WithHost(host))
	if err != nil {
		return fmt.Errorf("new client: %w", err)
	}
	defer cli.Close()

	_, err = cli.Info(ctx, client.InfoOptions{})
	if err != nil {
		return fmt.Errorf("docker info: %w", err)
	}

	return nil
}

// MustExtractDockerHost Extracts the docker host from the different alternatives, caching the result to avoid unnecessary
// calculations. Use this function to get the actual Docker host. This function does not consider Windows containers at the moment.
// The possible alternatives are:
//
//  1. Docker host from the "tc.host" property in the ~/.testcontainers.properties file.
//  2. DOCKER_HOST environment variable.
//  3. Docker host from context.
//  4. Docker host from the default docker socket path, without the unix schema.
//  5. Docker host from the "docker.host" property in the ~/.testcontainers.properties file.
//  6. Rootless docker socket path.
//  7. Else, because the Docker host is not set, it panics.
func MustExtractDockerHost(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func ExtractDockerHost(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// MustExtractDockerSocket Extracts the docker socket from the different alternatives, removing the socket schema and
// caching the result to avoid unnecessary calculations. Use this function to get the docker socket path,
// not the host (e.g. mounting the socket in a container). This function does not consider Windows containers at the moment.
// The possible alternatives are:
//
//  1. Docker host from the "tc.host" property in the ~/.testcontainers.properties file.
//  2. The TESTCONTAINERS_DOCKER_SOCKET_OVERRIDE environment variable.
//  3. Using a Docker client, check if the Info().OperatingSystem is "Docker Desktop" and return the default docker socket path for rootless docker.
//  4. Else, Get the current Docker Host from the existing strategies: see MustExtractDockerHost.
//  5. If the socket contains the unix schema, the schema is removed (e.g. unix:///var/run/docker.sock -> /var/run/docker.sock)
//  6. Else, the default location of the docker socket is used (/var/run/docker.sock)
//
// It panics if a Docker client cannot be created, or the Docker host cannot be discovered.
func MustExtractDockerSocket(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

// extractDockerHost Extracts the docker host from the different alternatives, without caching the result.
// This internal method is handy for testing purposes.
func extractDockerHost(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// extractDockerSocket Extracts the docker socket from the different alternatives, without caching the result.
// It will internally use the default Docker client, calling the internal method extractDockerSocketFromClient with it.
// This internal method is handy for testing purposes.
// It panics if a Docker client cannot be created, or the Docker host is not discovered.
func extractDockerSocket(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

// a Docker client is required to get the Docker info

// extractDockerSocketFromClient Extracts the docker socket from the different alternatives, without caching the result,
// and receiving an instance of the Docker API client interface.
// This internal method is handy for testing purposes, passing a mock type simulating the desired behaviour.
// It panics if the Docker Info call errors, or the Docker host is not discovered.
func extractDockerSocketFromClient(ctx context.Context, cli client.APIClient) string {
	_ = "STUB: not implemented"
	// check that the socket is not a tcp or unix socket
	return ""
}

// this use case will cover the case when the docker host is a tcp socket

// Docker Info is required to get the Operating System

// Because Docker Desktop runs in a VM, we need to use the default docker path for rootless docker

// Docker host is required to get the Docker socket

// isHostNotSet returns true if the error is related to the Docker host
// not being set, false otherwise.
func isHostNotSet(err error) bool { _ = "STUB: not implemented"; return false }

// dockerHostFromEnv returns the docker host from the DOCKER_HOST environment variable, if it's not empty
func dockerHostFromEnv(_ context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// dockerHostFromContext returns the docker host from the Go context, if it's not empty
func dockerHostFromContext(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// dockerHostFromProperties returns the docker host from the ~/.testcontainers.properties file, if it's not empty
func dockerHostFromProperties(_ context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// dockerSocketOverridePath returns the docker socket from the TESTCONTAINERS_DOCKER_SOCKET_OVERRIDE environment variable,
// if it's not empty
func dockerSocketOverridePath() (string, error) { _ = "STUB: not implemented"; return "", nil }

// dockerSocketPath returns the docker socket from the default docker socket path, if it's not empty
// and the socket exists
func dockerSocketPath(_ context.Context) (string, error) { _ = "STUB: not implemented"; return "", nil }

// testcontainersHostFromProperties returns the testcontainers host from the ~/.testcontainers.properties file, if it's not empty
func testcontainersHostFromProperties(_ context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Validate the URL format

// Return the original URL to preserve schema for Docker client

// DockerEnvFile is the file that is created when running inside a container.
// It's a variable to allow testing.
// TODO: Remove this once context rework is done, which eliminates need for the default network creation.
var DockerEnvFile = "/.dockerenv"

// InAContainer returns true if the code is running inside a container
// See https://github.com/docker/docker/blob/a9fa38b1edf30b23cae3eade0be48b3d4b1de14b/daemon/initlayer/setup_unix.go#L25
func InAContainer() bool { _ = "STUB: not implemented"; return false }

func inAContainer(path string) bool {
	_ = "STUB: not implemented"
	// see https://github.com/testcontainers/testcontainers-java/blob/3ad8d80e2484864e554744a4800a81f6b7982168/core/src/main/java/org/testcontainers/dockerclient/DockerClientConfigUtils.java#L15
	return false
}
