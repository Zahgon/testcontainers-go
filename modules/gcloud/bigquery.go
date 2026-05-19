package gcloud

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

// Deprecated: use [bigquery.Run] instead.
// RunBigQueryContainer creates an instance of the GCloud container type for BigQuery.
func RunBigQueryContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*GCloudContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deprecated: use [bigquery.Run] instead.
// RunBigQuery creates an instance of the GCloud container type for BigQuery.
// The URI uses http:// as the protocol.
func RunBigQuery(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*GCloudContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Process data yaml file only for the BigQuery container.
