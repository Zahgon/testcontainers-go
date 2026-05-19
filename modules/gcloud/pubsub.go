package gcloud

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

// Deprecated: use [pubsub.Run] instead
// RunPubsubContainer creates an instance of the GCloud container type for Pubsub.
func RunPubsubContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*GCloudContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deprecated: use [pubsub.Run] instead
// RunPubsub creates an instance of the GCloud container type for Pubsub.
func RunPubsub(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*GCloudContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
