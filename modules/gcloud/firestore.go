package gcloud

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

// Deprecated: use [firestore.Run] instead
// RunFirestoreContainer creates an instance of the GCloud container type for Firestore.
func RunFirestoreContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*GCloudContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deprecated: use [firestore.Run] instead
// RunFirestore creates an instance of the GCloud container type for Firestore.
func RunFirestore(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*GCloudContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
