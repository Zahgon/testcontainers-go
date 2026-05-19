package vault

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	defaultPort = "8200"
)

// VaultContainer represents the vault container type used in the module
type VaultContainer struct {
	testcontainers.Container
}

// Deprecated: use Run instead
// RunContainer creates an instance of the Vault container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*VaultContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the Vault container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*VaultContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WithToken is a container option function that sets the root token for the Vault
func WithToken(token string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithInitCommand is an option function that adds a set of initialization commands to the Vault's configuration
func WithInitCommand(commands ...string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// HttpHostAddress returns the http host address of Vault.
// It returns a string with the format http://<host>:<port>
//
//nolint:revive,staticcheck //FIXME
func (v *VaultContainer) HttpHostAddress(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
