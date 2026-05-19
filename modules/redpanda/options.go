package redpanda

import (
	"github.com/testcontainers/testcontainers-go"
)

// HTTPProxyAuthMethod defines the authentication method for HTTP Proxy.
type HTTPProxyAuthMethod string

const (
	HTTPProxyAuthMethodNone      HTTPProxyAuthMethod = "none"
	HTTPProxyAuthMethodHTTPBasic HTTPProxyAuthMethod = "http_basic"
	HTTPProxyAuthMethodOIDC      HTTPProxyAuthMethod = "oidc"
)

type options struct {
	// Superusers is a list of service account names.
	Superusers []string

	// KafkaEnableAuthorization is a flag to require authorization for Kafka connections.
	KafkaEnableAuthorization bool

	// KafkaAuthenticationMethod is either "none" for plaintext or "sasl"
	// for SASL (scram sha 256) authentication.
	KafkaAuthenticationMethod string

	// SchemaRegistryAuthenticationMethod is either "none" for no authentication
	// or "http_basic" for HTTP basic authentication.
	SchemaRegistryAuthenticationMethod string

	// HTTPProxyAuthenticationMethod is the authentication method for HTTP Proxy (pandaproxy).
	// Valid values are "none", "http_basic", or "oidc".
	HTTPProxyAuthenticationMethod HTTPProxyAuthMethod

	// EnableWasmTransform is a flag to enable wasm transform.
	EnableWasmTransform bool

	// ServiceAccounts is a map of username (key) to password (value) of users
	// that shall be created, so that you can use these to authenticate against
	// Redpanda (either for the Kafka API or Schema Registry HTTP access).
	// You must use SCRAM-SHA-256 as algorithm when authenticating on the
	// Kafka API.
	ServiceAccounts map[string]string

	// AutoCreateTopics is a flag to allow topic auto creation.
	AutoCreateTopics bool

	// EnableTLS is a flag to enable TLS.
	EnableTLS bool

	cert, key []byte

	// Listeners is a list of custom listeners that can be provided to access the
	// containers form within docker networks
	Listeners []listener

	// ExtraBootstrapConfig is a map of configs to be interpolated into the
	// container's bootstrap.yml
	ExtraBootstrapConfig map[string]any

	// enableAdminAPIAuthentication enables Admin API authentication
	enableAdminAPIAuthentication bool
}

func defaultOptions() options { _ = "STUB: not implemented"; return *new(options) }

// Compiler check to ensure that Option implements the testcontainers.ContainerCustomizer interface.
var _ testcontainers.ContainerCustomizer = (Option)(nil)

// Option is an option for the Redpanda container.
type Option func(*options) error

// Customize is a NOOP. It's defined to satisfy the testcontainers.ContainerCustomizer interface.
func (o Option) Customize(*testcontainers.GenericContainerRequest) error {
	_ = "STUB: not implemented"
	// NOOP to satisfy interface.
	return nil
}

// WithNewServiceAccount includes a new user with username (key) and password (value)
// that shall be created, so that you can use these to authenticate against
// Redpanda (either for the Kafka API or Schema Registry HTTP access).
func WithNewServiceAccount(username, password string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSuperusers defines the superusers added to the redpanda config.
// By default, there are no superusers.
func WithSuperusers(superusers ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnableSASL enables SASL scram sha 256 authentication.
// By default, no authentication (plaintext) is used.
// When setting an authentication method, make sure to add users
// as well as authorize them using the WithSuperusers() option.
func WithEnableSASL() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnableKafkaAuthorization enables authorization for connections on the Kafka API.
func WithEnableKafkaAuthorization() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnableWasmTransform enables wasm transform.
// Should not be used with RP versions before 23.3
func WithEnableWasmTransform() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnableSchemaRegistryHTTPBasicAuth enables HTTP basic authentication for
// Schema Registry.
func WithEnableSchemaRegistryHTTPBasicAuth() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithHTTPProxyAuthMethod sets the authentication method for HTTP Proxy.
// If an invalid method is provided, it defaults to "none".
func WithHTTPProxyAuthMethod(method HTTPProxyAuthMethod) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Invalid method, default to "none"

// WithAutoCreateTopics enables topic auto creation.
func WithAutoCreateTopics() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTLS(cert, key []byte) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithListener adds a custom listener to the Redpanda containers. Listener
// will be aliases to all networks, so they can be accessed from within docker
// networks. At least one network must be attached to the container, if not an
// error will be thrown when starting the container.
func WithListener(lis string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBootstrapConfig adds an arbitrary config kvp to the Redpanda container.
// Per the name, this config will be interpolated into the generated bootstrap
// config file, which is particularly useful for configs requiring a restart
// when otherwise applied to a running Redpanda instance.
func WithBootstrapConfig(cfg string, val any) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithAdminAPIAuthentication enables Admin API Authentication.
// It sets `admin_api_require_auth` configuration to true and configures a bootstrap user account.
// See https://docs.redpanda.com/current/deploy/deployment-option/self-hosted/manual/production/production-deployment/#bootstrap-a-user-account
func WithAdminAPIAuthentication() Option { _ = "STUB: not implemented"; return *new(Option) }
