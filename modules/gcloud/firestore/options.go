package firestore

import (
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/gcloud/internal/shared"
)

// options embeds the common GCloud options
type options struct {
	shared.Options
	datastoreMode bool
}

type Option func(o *options) error

// Customize is a NOOP. It's defined to satisfy the testcontainers.ContainerCustomizer interface.
func (o Option) Customize(*testcontainers.GenericContainerRequest) error {
	_ = "STUB: not implemented"
	// NOOP to satisfy interface.
	return nil
}

// defaultOptions returns a new Options instance with the default project ID.
func defaultOptions() options { _ = "STUB: not implemented"; return *new(options) }

// WithDatastoreMode sets the firestore emulator to run in datastore mode.
// Requires a cloud-sdk image with version 465.0.0 or higher
func WithDatastoreMode() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithProjectID re-exports the common GCloud WithProjectID option
var WithProjectID = shared.WithProjectID
