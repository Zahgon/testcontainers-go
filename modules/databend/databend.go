package databend

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	defaultUser         = "databend"
	defaultPassword     = "databend"
	defaultDatabaseName = "default"
	defaultPort         = "8000/tcp"
)

// DatabendContainer represents the Databend container type used in the module
type DatabendContainer struct {
	testcontainers.Container
	username string
	password string
	database string
}

// Deprecated: use testcontainers.ContainerCustomizer instead
var _ testcontainers.ContainerCustomizer = (*DatabendOption)(nil)

// Deprecated: use testcontainers.ContainerCustomizer instead
// DatabendOption is an option for the Databend container.
type DatabendOption func(*DatabendContainer)

// Deprecated: use testcontainers.ContainerCustomizer instead
// Customize is a NOOP. It's defined to satisfy the testcontainers.ContainerCustomizer interface.
func (o DatabendOption) Customize(*testcontainers.GenericContainerRequest) error {
	_ = "STUB: not implemented"
	// NOOP to satisfy interface.
	return nil
}

// Run creates an instance of the Databend container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*DatabendContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// set default credentials

// refresh the credentials from the environment variables

// MustConnectionString panics if the address cannot be determined.
func (c *DatabendContainer) MustConnectionString(ctx context.Context, args ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *DatabendContainer) ConnectionString(ctx context.Context, args ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// databend://databend:databend@localhost:8000/default?sslmode=disable

// WithUsername sets the username for the Databend container.
// WithUsername is [Run] option that configures the default query user by setting
// the `QUERY_DEFAULT_USER` container environment variable.
func WithUsername(username string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithPassword sets the password for the Databend container.
func WithPassword(password string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}
