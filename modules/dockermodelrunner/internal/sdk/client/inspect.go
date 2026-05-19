package client

import (
	"context"

	"github.com/testcontainers/testcontainers-go/modules/dockermodelrunner/internal/sdk/types"
)

// InspectModel returns a model that is already pulled using the Docker Model Runner format.
// The name of the model is in the format of <name>:<tag>.
// The namespace and name defines Models as OCI Artifacts in Docker Hub, therefore the namespace is the organization and the name is the repository.
// E.g. "ai/smollm2:360M-Q4_K_M". See [Models_as_OCI_Artifacts] for more information.
//
// [Models_as_OCI_Artifacts]: https://hub.docker.com/u/ai
func (c *Client) InspectModel(ctx context.Context, namespace string, name string) (*types.ModelResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
