package rabbitmq

import (
	"context"
	_ "embed"

	"github.com/testcontainers/testcontainers-go"
)

const (
	DefaultAMQPSPort      = "5671/tcp"
	DefaultAMQPPort       = "5672/tcp"
	DefaultHTTPSPort      = "15671/tcp"
	DefaultHTTPPort       = "15672/tcp"
	defaultPassword       = "guest"
	defaultUser           = "guest"
	defaultCustomConfPath = "/etc/rabbitmq/rabbitmq-testcontainers.conf"
)

//go:embed mounts/rabbitmq-testcontainers.conf.tpl
var customConfigTpl string

// RabbitMQContainer represents the RabbitMQ container type used in the module
type RabbitMQContainer struct {
	testcontainers.Container
	AdminPassword string
	AdminUsername string
}

// AmqpURL returns the URL for AMQP clients.
//
//nolint:staticcheck //FIXME
func (c *RabbitMQContainer) AmqpURL(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// AmqpURL returns the URL for AMQPS clients.
func (c *RabbitMQContainer) AmqpsURL(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// HttpURL returns the URL for HTTP management.
//
//nolint:revive,staticcheck //FIXME
func (c *RabbitMQContainer) HttpURL(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// HttpsURL returns the URL for HTTPS management.
//
//nolint:revive,staticcheck //FIXME
func (c *RabbitMQContainer) HttpsURL(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Deprecated: use Run instead
// RunContainer creates an instance of the RabbitMQ container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*RabbitMQContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the RabbitMQ container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*RabbitMQContainer, error) {
	_ = "STUB: not implemented"
	// Gather all config options (defaults and then apply provided options)
	return nil, nil
}

func withConfig(hostPath string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// applySSLSettings transfers the SSL settings to the container request.
func applySSLSettings(sslSettings *SSLSettings) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// To verify that TLS has been enabled on the node, container logs should contain an entry about a TLS listener being enabled
// See https://www.rabbitmq.com/ssl.html#enabling-tls-verify-configuration

func renderRabbitMQConfig(opts options) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
