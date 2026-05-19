package registry

import (
	"context"

	"github.com/cpuguy83/dockercfg"
	"github.com/moby/moby/api/types/registry"

	"github.com/testcontainers/testcontainers-go"
)

const (
	// registryPort is the default port used by the Registry container.
	registryPort = "5000/tcp"

	// DefaultImage is the default image used by the Registry container.
	DefaultImage = "registry:2.8.3"
)

// RegistryContainer represents the Registry container type used in the module
type RegistryContainer struct {
	testcontainers.Container
	RegistryName string
}

// Address returns the address of the Registry container, using the HTTP protocol
func (c *RegistryContainer) Address(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// HostAddress returns the host address including port of the Registry container.
func (c *RegistryContainer) HostAddress(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// This is a workaround for WSL, where localhost is not reachable from Docker.

// localAddress returns the local address of the machine
// which can be used to connect to the local registry.
// This avoids the issues with localhost on WSL.
func localAddress(ctx context.Context) (string, error) { _ = "STUB: not implemented"; return "", nil }

// getEndpointWithAuth returns the HTTP endpoint of the Registry container, along with the image auth
// for the image reference.
// E.g. imageRef = "localhost:5000/alpine:latest"
func getEndpointWithAuth(ctx context.Context, imageRef string) (string, string, registry.AuthConfig, error) {
	_ = "STUB: not implemented"
	return "", "", *new(registry.AuthConfig), nil
}

// DeleteImage deletes an image reference from the Registry container.
// It will use the HTTP endpoint of the Registry container to delete it,
// doing a HEAD request to get the image digest and then a DELETE request
// to actually delete the image.
// E.g. imageRef = "localhost:5000/alpine:latest"
func (c *RegistryContainer) DeleteImage(ctx context.Context, imageRef string) error {
	_ = "STUB: not implemented"
	return nil
}

// ImageExists checks if an image exists in the Registry container. It will use the v2 HTTP endpoint
// of the Registry container to check if the image reference exists.
// E.g. imageRef = "localhost:5000/alpine:latest"
func (c *RegistryContainer) ImageExists(ctx context.Context, imageRef string) error {
	_ = "STUB: not implemented"
	return nil
}

// PushImage pushes an image to the Registry container. It will use the internally stored RegistryName
// to push the image to the container, and it will finally wait for the image to be pushed.
func (c *RegistryContainer) PushImage(ctx context.Context, ref string) error {
	_ = "STUB: not implemented"
	return nil
}

// PullImage pulls an image from an external registry into the local Docker daemon.
// Differently from PushImage, which uploads an image to the testcontainers managed registry,
// this method downloads (copies) the specified image reference so it becomes
// available locally for further operations such as tagging or pushing.
// It uses the same platform as the registry container's image.
func (c *RegistryContainer) PullImage(ctx context.Context, ref string) error {
	_ = "STUB: not implemented"
	return nil
}

// Use the same platform as the registry container's image.

// TagImage tags an image from the local Registry.
// This function is helpful when you want to push an image to your Registry
// instance made by testcontainer.
func (c *RegistryContainer) TagImage(ctx context.Context, image, ref string) error {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: use Run instead
// RunContainer creates an instance of the Registry container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*RegistryContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the Registry container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*RegistryContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convenient for testing

// SetDockerAuthConfig sets the DOCKER_AUTH_CONFIG environment variable with
// authentication for the given host, username and password sets.
// It returns a function to reset the environment back to the previous state.
func SetDockerAuthConfig(host, username, password string, additional ...string) (func(), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DockerAuthConfig returns a map of AuthConfigs including base64 encoded Auth field
// for the provided details. It also accepts additional host, username and password
// triples to add more auth configurations.
func DockerAuthConfig(host, username, password string, additional ...string) (map[string]dockercfg.AuthConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
