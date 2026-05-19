package firestore

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	// DefaultProjectID is the default project ID for the Firestore container.
	DefaultProjectID  = "test-project"
	defaultPortNumber = "8080"
	defaultPort       = defaultPortNumber + "/tcp"
)

// Container represents the Firestore container type used in the module
type Container struct {
	testcontainers.Container
	settings options
}

// ProjectID returns the project ID of the Firestore container.
func (c *Container) ProjectID() string { _ = "STUB: not implemented"; return "" }

// URI returns the URI of the Firestore container.
func (c *Container) URI() string { _ = "STUB: not implemented"; return "" }

// Run creates an instance of the Firestore GCloud container type.
// The URI uses the empty string as the protocol.
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	// Process custom options first to extract settings
	return nil, nil
}

// Build moduleOpts with defaults

// Append user options
