package shared

import (
	"github.com/testcontainers/testcontainers-go"
)

const (
	// DefaultProjectID is the default project ID for the Pubsub container.
	DefaultProjectID = "test-project"
)

// Options represents the options for the different GCloud containers.
// This type must contain all the options that are common to all the GCloud containers.
type Options struct {
	ProjectID string
	URI       string
}

// Compiler check to ensure that Option implements the testcontainers.ContainerCustomizer interface.
var _ testcontainers.ContainerCustomizer = (*Option)(nil)

// Option is an option for the GCloud container.
type Option func(*Options) error

// Customize is a NOOP. It's defined to satisfy the testcontainers.ContainerCustomizer interface.
func (o Option) Customize(*testcontainers.GenericContainerRequest) error {
	_ = "STUB: not implemented"
	// NOOP to satisfy interface.
	return nil
}

// DefaultOptions returns a new Options instance with the default project ID.
func DefaultOptions() Options { _ = "STUB: not implemented"; return *new(Options) }

// WithProjectID sets the project ID for the GCloud container.
func WithProjectID(projectID string) Option { _ = "STUB: not implemented"; return *new(Option) }
