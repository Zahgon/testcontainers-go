package dynamodb

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	port          = "8000/tcp"
	containerName = "tc_dynamodb_local"
)

// DynamoDBContainer represents the DynamoDB container type used in the module
type DynamoDBContainer struct {
	testcontainers.Container
}

// Run creates an instance of the DynamoDB container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*DynamoDBContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConnectionString returns DynamoDB local endpoint host and port in <host>:<port> format
func (c *DynamoDBContainer) ConnectionString(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// WithSharedDB allows container reuse between successive runs. Data will be persisted
func WithSharedDB() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithDisableTelemetry - DynamoDB local will not send any telemetry
func WithDisableTelemetry() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// if other flags (e.g. -sharedDb) exist, append to them
