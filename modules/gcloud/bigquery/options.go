package bigquery

import (
	"io"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/gcloud/internal/shared"
)

// Options aliases the common GCloud options
type options = shared.Options

// Option aliases the common GCloud option type
type Option = shared.Option

// defaultOptions returns a new Options instance with the default project ID.
func defaultOptions() options { _ = "STUB: not implemented"; return *new(options) }

// WithProjectID re-exports the common GCloud WithProjectID option
var WithProjectID = shared.WithProjectID

// WithDataYAML seeds the BigQuery project for the GCloud container with an [io.Reader] representing
// the data yaml file, which is used to copy the file to the container, and then processed to seed
// the BigQuery project.
//
// Other GCloud containers will ignore this option.
func WithDataYAML(r io.Reader) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}
