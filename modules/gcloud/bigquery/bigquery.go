package bigquery

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	// DefaultProjectID is the default project ID for the BigQuery container.
	DefaultProjectID = "test-project"

	// bigQueryDataYamlPath is the path to the data yaml file in the container.
	bigQueryDataYamlPath = "/testcontainers-data.yaml"

	defaultPortNumber9050 = "9050"
	defaultPortNumber9060 = "9060"
	defaultPort9050       = defaultPortNumber9050 + "/tcp"
	defaultPort9060       = defaultPortNumber9060 + "/tcp"
)

// Container represents the BigQuery container type used in the module
type Container struct {
	testcontainers.Container
	settings options
}

// ProjectID returns the project ID of the BigQuery container.
func (c *Container) ProjectID() string { _ = "STUB: not implemented"; return "" }

// URI returns the URI of the BigQuery container.
func (c *Container) URI() string { _ = "STUB: not implemented"; return "" }

// Run creates an instance of the BigQuery GCloud container type.
// The URI uses http:// as the protocol.
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
