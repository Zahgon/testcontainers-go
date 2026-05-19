package cosmosdb

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

// ContainerPolicy ensures that requests always target the CosmosDB emulator container endpoint.
// It overrides the CosmosDB client's globalEndpointManager, which would otherwise dynamically
// update [http.Request.Host] based on global endpoint discovery, pinning all requests to the container.
type ContainerPolicy struct {
	endpoint string
}

func NewContainerPolicy(ctx context.Context, c *Container) (*ContainerPolicy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *ContainerPolicy) Do(req *policy.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ClientOptions returns Azure CosmosDB client options that contain ContainerPolicy.
func (p *ContainerPolicy) ClientOptions() *azcosmos.ClientOptions {
	_ = "STUB: not implemented"
	return nil
}
