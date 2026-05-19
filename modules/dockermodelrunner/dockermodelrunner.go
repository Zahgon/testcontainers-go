package dockermodelrunner

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/dockermodelrunner/internal/sdk/client"
	"github.com/testcontainers/testcontainers-go/modules/socat"
)

const (
	modelRunnerEntrypoint = "model-runner.docker.internal"
	modelRunnerPort       = 80
)

// Container represents the DockerModelRunner container type used in the module
type Container struct {
	*socat.Container
	*client.Client
	model   string
	baseURL string
}

// Run creates an instance of the DockerModelRunner container type.
func Run(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	return nil,

		// Process model runner options.
		nil
}

// Add socat options, which are applied to the socat container.
