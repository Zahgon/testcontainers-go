package mssql

import (
	"context"
	"io"

	"github.com/testcontainers/testcontainers-go"
)

const (
	defaultPort     = "1433/tcp"
	defaultUsername = "sa" // default microsoft system administrator
	defaultPassword = "Strong@Passw0rd"
)

// MSSQLServerContainer represents the MSSQLServer container type used in the module
type MSSQLServerContainer struct {
	testcontainers.Container
	password string
	username string
}

// Password returns the password for the MSSQLServer container
func (c *MSSQLServerContainer) Password() string {
	_ = "STUB: not implemented"

	// WithAcceptEULA sets the ACCEPT_EULA environment variable to "Y"
	return ""
}

func WithAcceptEULA() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithPassword sets the MSSQL_SA_PASSWORD environment variable to the provided password
func WithPassword(password string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithInitSQL adds SQL scripts to be executed after the container is ready.
// The scripts are executed in the order they are provided using sqlcmd tool.
func WithInitSQL(files ...io.Reader) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// targetPath is a dummy path to store the script in the container

// NOTE: we add both legacy and new mssql-tools paths to ensure compatibility

// Deprecated: use Run instead
// RunContainer creates an instance of the MSSQLServer container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*MSSQLServerContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the MSSQLServer container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*MSSQLServerContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate EULA acceptance after applying user options

// Retrieve password from container environment

// ConnectionString returns the connection string for the MSSQLServer container
func (c *MSSQLServerContainer) ConnectionString(ctx context.Context, args ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
