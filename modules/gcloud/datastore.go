package gcloud

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

// Deprecated: use [datastore.Run] instead
// RunDatastoreContainer creates an instance of the GCloud container type for Datastore.
func RunDatastoreContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*GCloudContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deprecated: use [datastore.Run] instead
// RunDatastore creates an instance of the GCloud container type for Datastore.
func RunDatastore(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*GCloudContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
