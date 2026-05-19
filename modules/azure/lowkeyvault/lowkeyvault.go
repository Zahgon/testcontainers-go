package lowkeyvault

import (
	"context"
	"net/http"

	"github.com/testcontainers/testcontainers-go"
)

// possible access modes
const (
	Local = iota
	Network
)

const (
	// defaultAPIPort is the default port used by for the Lowkey Vault Key Vault API endpoints
	defaultAPIPort = "8443/tcp"
	// defaultMetadataPort is the default port used for the Lowkey Vault Metadata endpoints
	defaultMetadataPort = "8080/tcp"
)

// Container represents the Lowkey Vault container type used in the module
type Container struct {
	testcontainers.Container
	localHostName  string
	remoteHostName string
}

// WithNetworkAlias sets the alias of the container for the provided network and adds the specified name as a key vault alias for the default, "localhost", vault.
func WithNetworkAlias(alias string, forNetwork *testcontainers.DockerNetwork) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// The <port> placeholder will be replaced by the container automatically just in time

// Run creates an instance of the Lowkey Vault container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	// Initialize with module defaults
	return nil, nil
}

// Add user-provided options

// Client prepares a client which will accept insecure (self-signed) certificates.
func (c *Container) Client(ctx context.Context) (http.Client, error) {
	_ = "STUB: not implemented"
	return *new(http.Client), nil
}

func (c *Container) fetchDefaultCertPassword(authority string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Container) fetchDefaultCertContent(authority string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Container) fetchContent(url string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConnectionURL returns the connection URL for the Lowkey Vault API based on the provided access mode.
func (c *Container) ConnectionURL(ctx context.Context, accessMode int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// IdentityEndpoint returns the URL value of the IDENTITY_ENDPOINT environment variable for the managed identity simulation. This will be used to obtain an access token for the Lowkey Vault API.
func (c *Container) IdentityEndpoint(ctx context.Context, accessMode int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// IdentityHeader returns the value of the IDENTITY_HEADER environment variable for the managed identity simulation.
func (c *Container) IdentityHeader() string { _ = "STUB: not implemented"; return "" }

func (c *Container) mappedHostAuthority(ctx context.Context, exposedPort string, accessMode int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Container) resolveLocalHostName(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Container) resolveNetworkHostName(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
