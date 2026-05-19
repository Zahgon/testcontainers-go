package pulsar

import (
	"context"
	"io"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	defaultPulsarPort                      = "6650/tcp"
	defaultPulsarAdminPort                 = "8080/tcp"
	defaultPulsarCmd                       = "/pulsar/bin/apply-config-from-env.py /pulsar/conf/standalone.conf && bin/pulsar standalone"
	defaultPulsarCmdWithoutFunctionsWorker = "--no-functions-worker -nss"
	transactionTopicEndpoint               = "/admin/v2/persistent/pulsar/system/transaction_coordinator_assign/partitions"
)

var defaultWaitStrategies = []wait.Strategy{
	wait.ForHTTP("/admin/v2/clusters").WithPort(defaultPulsarAdminPort).WithResponseMatcher(func(r io.Reader) bool {
		respBytes, _ := io.ReadAll(r)
		resp := string(respBytes)
		return resp == `["standalone"]`
	}),
	wait.ForListeningPort(defaultPulsarPort),
}

type Container struct {
	testcontainers.Container
	LogConsumers []testcontainers.LogConsumer // Deprecated. Use the ContainerRequest instead. Needs to be exported to control the stop from the caller
}

func (c *Container) BrokerURL(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Container) HTTPServiceURL(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Container) resolveURL(ctx context.Context, port string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// WithFunctionsWorker enables the functions worker, which will override the default pulsar command
// and add a waiting strategy for the functions worker
func WithFunctionsWorker() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// Deprecated: use the testcontainers.WithLogConsumers functional option instead
// WithLogConsumers allows to add log consumers to the container.
// They will be automatically started and they will follow the container logs,
// but it's a responsibility of the caller to stop them calling StopLogProducer
func (c *Container) WithLogConsumers(ctx context.Context, _ ...testcontainers.LogConsumer) {
	_ = "STUB: not implemented"
	return
}

// not handling the error because it will return an error if and only if the producer is already started

// WithPulsarEnv allows to use the native APIs and set each variable with PULSAR_PREFIX_ as prefix.
func WithPulsarEnv(configVar string, configValue string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

func WithTransactions() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// clone defaultWaitStrategies

// Deprecated: use Run instead
// RunContainer creates an instance of the Pulsar container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the Pulsar container type, being possible to pass a custom request and options
// The created container will use the following defaults:
// - image: apachepulsar/pulsar:4.0.9
// - exposed ports: 6650/tcp, 8080/tcp
// - waiting strategy: wait for all the following strategies:
//   - the Pulsar admin API ("/admin/v2/clusters") to be ready on port 8080/tcp and return the response `["standalone"]`
//   - the Pulsar port (6650/tcp) to be listening
//
// - command: "/bin/bash -c /pulsar/bin/apply-config-from-env.py /pulsar/conf/standalone.conf && bin/pulsar standalone --no-functions-worker -nss"
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
