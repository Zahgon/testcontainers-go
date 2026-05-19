package redpanda

import (
	"context"
	_ "embed"

	"github.com/testcontainers/testcontainers-go"
)

var (
	//go:embed mounts/redpanda.yaml.tpl
	nodeConfigTpl string

	//go:embed mounts/bootstrap.yaml.tpl
	bootstrapConfigTpl string

	//go:embed mounts/entrypoint-tc.sh
	entrypoint []byte
)

const (
	defaultKafkaAPIPort       = "9092/tcp"
	defaultAdminAPIPort       = "9644/tcp"
	defaultSchemaRegistryPort = "8081/tcp"
	defaultHTTPProxyPort      = "8082/tcp"

	redpandaDir         = "/etc/redpanda"
	entrypointFile      = "/entrypoint-tc.sh"
	bootstrapConfigFile = ".bootstrap.yaml"
	certFile            = "cert.pem"
	keyFile             = "key.pem"

	bootstrapAdminAPIUser     = "redpanda_bootstrap_admin_user"
	bootstrapAdminAPIPassword = "redpanda_bootstrap_admin_password"
)

// Container represents the Redpanda container type used in the module.
type Container struct {
	testcontainers.Container
	urlScheme string
}

// Deprecated: use Run instead
// RunContainer creates an instance of the Redpanda container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the Redpanda container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	// 1. Gather all config options (defaults and then apply provided options)
	return nil, nil
}

// 2. If the image is not at least v23.3, disable wasm transform

// 3. Build module options

// Wait for the ports to be mapped without accessing them,
// because container needs Redpanda configuration before Redpanda is started
// and the mapped ports are part of that configuration.

// 4. If enabled, bootstrap user account

// set the RP_BOOTSTRAP_USER env var

// add our internal bootstrap admin user to superusers

// enable admin_api_require_auth

// 5. Bootstrap config file contains cluster configurations which will only be considered
// the very first time you start a cluster.

// We need a custom entrypoint that waits until the actual Redpanda node config is mounted.
// Once the redpanda config is mounted we will call the original entrypoint with the same parameters.
// We have to do this kind of two-step process, because we need to know the mapped
// port, so that we can use this in Redpanda's advertised listeners configuration for
// the Kafka API.

// 7. Create certificate and key for TLS connections.

// 8. Append user-provided options

// 9. Add listener network aliases as final step (must be after user options that add networks)

// 9. Get mapped port for the Kafka API, so that we can render and then mount
// the Redpanda config with the advertised Kafka address.

// 6. Render redpanda.yaml config and mount it.

// 7. Wait until Redpanda is ready to serve requests.

// Redpanda's admin API returns 404 for requests to "/".

// 8. Create Redpanda Service Accounts if configured to do so.

// KafkaSeedBroker returns the seed broker that should be used for connecting
// to the Kafka API with your Kafka client. It'll be returned in the format:
// "host:port" - for example: "localhost:55687".
func (c *Container) KafkaSeedBroker(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// AdminAPIAddress returns the address to the Redpanda Admin API. This
// is an HTTP-based API and thus the returned format will be: http://host:port.
func (c *Container) AdminAPIAddress(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// SchemaRegistryAddress returns the address to the schema registry API. This
// is an HTTP-based API and thus the returned format will be: http://host:port.
func (c *Container) SchemaRegistryAddress(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// HTTPProxyAddress returns the address to the HTTP Proxy API (pandaproxy). This
// is an HTTP-based API and thus the returned format will be: http://host:port.
func (c *Container) HTTPProxyAddress(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// renderBootstrapConfig renders the config template for the .bootstrap.yaml config,
// which configures Redpanda's cluster properties.
// Reference: https://docs.redpanda.com/docs/reference/cluster-properties/
func renderBootstrapConfig(settings options) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// withListeners creates a CustomizeRequestOption that validates and sets network aliases for the provided listeners.
// The container must be attached to at least one network.
func withListeners(listeners []listener) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// renderNodeConfig renders the redpanda.yaml node config and returns it as
// byte array.
func renderNodeConfig(settings options, hostIP string, advertisedKafkaPort uint16) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type redpandaBootstrapConfigTplParams struct {
	Superusers                  []string
	KafkaAPIEnableAuthorization bool
	AutoCreateTopics            bool
	EnableWasmTransform         bool
	ExtraBootstrapConfig        map[string]any
}

type redpandaConfigTplParams struct {
	KafkaAPI         redpandaConfigTplParamsKafkaAPI
	SchemaRegistry   redpandaConfigTplParamsSchemaRegistry
	HTTPProxy        redpandaConfigTplParamsHTTPProxy
	AutoCreateTopics bool
	EnableTLS        bool
}

type redpandaConfigTplParamsKafkaAPI struct {
	AdvertisedHost       string
	AdvertisedPort       int
	AuthenticationMethod string
	EnableAuthorization  bool
	Listeners            []listener
}

type redpandaConfigTplParamsSchemaRegistry struct {
	AuthenticationMethod string
}

type redpandaConfigTplParamsHTTPProxy struct {
	AuthenticationMethod HTTPProxyAuthMethod
}

type listener struct {
	Address              string
	Port                 int
	AuthenticationMethod string
}

// isAtLeastVersion returns true if the base image (without tag) is in a version or above
func isAtLeastVersion(image, major string) bool { _ = "STUB: not implemented"; return false }

// version >= v8.x
