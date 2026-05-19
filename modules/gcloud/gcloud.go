package gcloud

import (
	"context"
	"io"

	"github.com/testcontainers/testcontainers-go"
)

const defaultProjectID = "test-project"

// Deprecated: use the specialized containers instead:
// - [bigquery.Container]
// - [bigtable.Container]
// - [datastore.Container]
// - [firestore.Container]
// - [pubsub.Container]
// - [spanner.Container]
type GCloudContainer struct {
	testcontainers.Container
	Settings options
	URI      string
}

// newGCloudContainer creates a new GCloud container, obtaining the URL to access the container from the specified port.
func newGCloudContainer(ctx context.Context, img string, port int, settings options, proto string, opts ...testcontainers.ContainerCustomizer) (*GCloudContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type options struct {
	ProjectID        string
	bigQueryDataYaml io.Reader
}

func defaultOptions() options { _ = "STUB: not implemented"; return *new(options) }

// Compiler check to ensure that Option implements the testcontainers.ContainerCustomizer interface.
var _ testcontainers.ContainerCustomizer = (*Option)(nil)

// Option is an option for the GCloud container.
type Option func(*options) error

// Customize is a NOOP. It's defined to satisfy the testcontainers.ContainerCustomizer interface.
func (o Option) Customize(*testcontainers.GenericContainerRequest) error {
	_ = "STUB: not implemented"
	// NOOP to satisfy interface.
	return nil
}

// WithProjectID sets the project ID for the GCloud container.
func WithProjectID(projectID string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Deprecated: Use [bigquery.WithDataYAML] instead.
// WithDataYAML seeds the BigQuery project for the GCloud container with an [io.Reader] representing
// the data yaml file, which is used to copy the file to the container, and then processed to seed
// the BigQuery project.
//
// Other GCloud containers will ignore this option.
// If this option is passed multiple times, an error is returned.
func WithDataYAML(r io.Reader) Option { _ = "STUB: not implemented"; return *new(Option) }

// applyOptions applies the options to the container request and returns the settings.
func applyOptions(opts []testcontainers.ContainerCustomizer) (options, error) {
	_ = "STUB: not implemented"
	return *new(options), nil
}
