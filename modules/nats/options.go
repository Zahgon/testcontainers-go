package nats

import (
	"io"

	"github.com/testcontainers/testcontainers-go"
)

type options struct {
	CmdArgs map[string]string
}

func defaultOptions() options { _ = "STUB: not implemented"; return *new(options) }

// Compiler check to ensure that CmdOption implements the testcontainers.ContainerCustomizer interface.
var _ testcontainers.ContainerCustomizer = (*CmdOption)(nil)

// CmdOption is an option for the NATS container.
type CmdOption func(opts *options)

// Customize is a NOOP. It's defined to satisfy the testcontainers.ContainerCustomizer interface.
func (o CmdOption) Customize(_ *testcontainers.GenericContainerRequest) error {
	_ = "STUB: not implemented"
	// NOOP to satisfy interface.
	return nil
}

func WithUsername(username string) CmdOption { _ = "STUB: not implemented"; return *new(CmdOption) }

func WithPassword(password string) CmdOption { _ = "STUB: not implemented"; return *new(CmdOption) }

// WithArgument adds an argument and its value to the NATS container.
// The argument flag does not need to include the dashes.
func WithArgument(flag string, value string) CmdOption {
	_ = "STUB: not implemented"
	return *new(CmdOption)
}

// remove all dashes to make it easier to use

// WithConfigFile pass a content of io.Reader to the NATS container as /etc/nats.conf
// Changing the connectivity (listen address or ports) can break the container setup.
func WithConfigFile(config io.Reader) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}
