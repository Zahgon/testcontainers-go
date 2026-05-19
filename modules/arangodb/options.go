package arangodb

import (
	"github.com/testcontainers/testcontainers-go"
)

// WithRootPassword sets the password for the ArangoDB root user
func WithRootPassword(password string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// withWaitStrategy sets the wait strategy for the ArangoDB container
// once we know the credentials.
func withWaitStrategy() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}
