package solace

import (
	"github.com/testcontainers/testcontainers-go"
)

type options struct {
	vpn      string
	username string
	password string
	services []Service           // enabled services
	queues   map[string][]string // queueName -> topics
	shmSize  int64
}

func defaultOptions() options { _ = "STUB: not implemented"; return *new(options) }

// 1 GiB

// Compiler check to ensure that Option implements the testcontainers.ContainerCustomizer interface.
var _ testcontainers.ContainerCustomizer = (Option)(nil)

// Option is an option for the Solace container.
type Option func(*options) error

// Customize is a NOOP. It's defined to satisfy the testcontainers.ContainerCustomizer interface.
func (o Option) Customize(*testcontainers.GenericContainerRequest) error {
	_ = "STUB: not implemented"
	// NOOP to satisfy interface.
	return nil
}

// WithServices configures the services to be exposed with their wait strategies
func WithServices(srv ...Service) Option { _ = "STUB: not implemented"; return *new(Option) }

// Clear existing services and use only the specified ones

// WithCredentials sets the client credentials (username, password)
func WithCredentials(username, password string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithVPN sets the VPN name
func WithVPN(vpn string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithQueue subscribes a given topic to a queue (for SMF/AMQP testing)
func WithQueue(queueName, topic string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithShmSize sets the size of the /dev/shm volume
func WithShmSize(size int64) Option { _ = "STUB: not implemented"; return *new(Option) }
