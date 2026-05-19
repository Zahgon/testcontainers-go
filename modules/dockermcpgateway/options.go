package dockermcpgateway

import (
	"github.com/testcontainers/testcontainers-go"
)

type options struct {
	tools   map[string][]string
	secrets map[string]string
}

func defaultOptions() options { _ = "STUB: not implemented"; return *new(options) }

// Compiler check to ensure that Option implements the testcontainers.ContainerCustomizer interface.
var _ testcontainers.ContainerCustomizer = (Option)(nil)

// Option is an option for the DockerMCPGateway container.
type Option func(*options) error

// Customize is a NOOP. It's defined to satisfy the testcontainers.ContainerCustomizer interface.
func (o Option) Customize(*testcontainers.GenericContainerRequest) error {
	_ = "STUB: not implemented"
	// NOOP to satisfy interface.
	return nil
}

// WithTools sets a server's tools to use in the DockerMCPGateway container.
// Multiple calls to this function with the same server will append to the existing tools for that server.
// No duplicate tools will be added for the same server.
func WithTools(server string, tools []string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Append only unique tools to avoid duplicates

// If the server does not exist, create a new entry

// WithServers sets the servers to use in the DockerMCPGateway container.
// Multiple calls to this function will append to the existing values.
func WithSecret(key, value string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSecrets sets the secrets to use in the DockerMCPGateway container.
// Multiple calls to this function will merge the secrets into the existing map.
func WithSecrets(secrets map[string]string) Option { _ = "STUB: not implemented"; return *new(Option) }
