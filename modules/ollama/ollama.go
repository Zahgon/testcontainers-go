package ollama

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

// Deprecated: it will be removed in the next major version.
const DefaultOllamaImage = "ollama/ollama:0.5.7"

// OllamaContainer represents the Ollama container type used in the module
type OllamaContainer struct {
	testcontainers.Container
}

// ConnectionString returns the connection string for the Ollama container,
// using the default port 11434.
func (c *OllamaContainer) ConnectionString(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Commit it commits the current file system changes in the container into a new target image.
// The target image name should be unique, as this method will commit the current state
// of the container into a new image with the given name, so it doesn't override existing images.
// It should be used for creating an image that contains a loaded model.
func (c *OllamaContainer) Commit(ctx context.Context, targetImage string) error {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: use Run instead
// RunContainer creates an instance of the Ollama container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*OllamaContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the Ollama container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*OllamaContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if we need to use the local process

// Only request a GPU if NOT using local process and the host supports it.

// Now we have processed all the options, we can check if we need to use the local process.
