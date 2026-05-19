package gcloud

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

// Deprecated: use [spanner.Run] instead
// RunSpannerContainer creates an instance of the GCloud container type for Spanner.
func RunSpannerContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*GCloudContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deprecated: use [spanner.Run] instead
// RunSpanner creates an instance of the GCloud container type for Spanner.
func RunSpanner(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*GCloudContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
