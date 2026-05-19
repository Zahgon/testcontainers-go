package artemis

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	defaultBrokerPort = "61616/tcp"
	defaultHTTPPort   = "8161/tcp"
	defaultUser       = "artemis"
	defaultPassword   = "artemis"
)

// Container represents the Artemis container type used in the module.
type Container struct {
	testcontainers.Container
	user     string
	password string
}

// User returns the administrator username.
func (c *Container) User() string {
	_ = "STUB: not implemented"

	// Password returns the administrator password.
	return ""
}

func (c *Container) Password() string {
	_ = "STUB: not implemented"

	// BrokerEndpoint returns the host:port for the combined protocols endpoint.
	// The endpoint accepts CORE, MQTT, AMQP, STOMP, HORNETQ and OPENWIRE protocols.
	return ""
}

func (c *Container) BrokerEndpoint(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ConsoleURL returns the URL for the management console.
func (c *Container) ConsoleURL(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// WithCredentials sets the administrator credentials. The default is artemis:artemis.
func WithCredentials(user, password string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithAnonymousLogin enables anonymous logins.
func WithAnonymousLogin() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// Additional arguments sent to the `artemis create` command.
// The default is `--http-host 0.0.0.0 --relax-jolokia`.
// Setting this value will override the default.
// See the documentation on `artemis create` for available options.
func WithExtraArgs(args string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// Deprecated: use Run instead.
// RunContainer creates an instance of the Artemis container type.
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the Artemis container type with a given image
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// initialize the credentials

// refresh the credentials from the environment variables
