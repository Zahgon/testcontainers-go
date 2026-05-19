package gcloud

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

// Deprecated: use [bigtable.Run] instead
// RunBigTableContainer creates an instance of the GCloud container type for BigTable.
func RunBigTableContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*GCloudContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deprecated: use [bigtable.Run] instead
// RunBigTable creates an instance of the GCloud container type for BigTable.
func RunBigTable(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*GCloudContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
