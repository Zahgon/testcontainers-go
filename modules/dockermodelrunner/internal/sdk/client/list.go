package client

import (
	"context"

	"github.com/testcontainers/testcontainers-go/modules/dockermodelrunner/internal/sdk/types"
)

// ListModels lists all models that are already pulled using the Docker Model Runner format.
func (c *Client) ListModels(ctx context.Context) ([]types.ModelResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
