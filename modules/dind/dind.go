package dind

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	defaultDockerDaemonPortNumber = "2375"
	defaultDockerDaemonPort       = defaultDockerDaemonPortNumber + "/tcp"
)

// Container represents the Docker in Docker container type used in the module
type Container struct {
	testcontainers.Container
}

// Run creates an instance of the Docker in Docker container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Host returns the endpoint to connect to the Docker daemon running inside the DinD container.
func (c *Container) Host(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// LoadImage loads an image into the DinD container.
// It creates a temporary file to save the image and then copies it to the container.
// This temporary file is deleted after the function returns.
func (c *Container) LoadImage(ctx context.Context, image string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// save image

// Close the file handle immediately: SaveImages and CopyFileToContainer
// open the file by name.
