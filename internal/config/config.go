package config

import (
	"sync"
	"time"
)

const ReaperDefaultImage = "testcontainers/ryuk:0.13.0"

var (
	tcConfig     Config
	tcConfigOnce = new(sync.Once)
)

// testcontainersConfig {

// Config represents the configuration for Testcontainers.
// User values are read from ~/.testcontainers.properties file which can be overridden
// using the specified environment variables. For more information, see [Custom Configuration].
//
// The Ryuk prefixed fields controls the [Garbage Collector] feature, which ensures that
// resources are cleaned up after the test execution.
//
// [Garbage Collector]: https://golang.testcontainers.org/features/garbage_collector/
// [Custom Configuration]: https://golang.testcontainers.org/features/configuration/
type Config struct {
	// Host is the address of the Docker daemon.
	//
	// Environment variable: DOCKER_HOST
	Host string `properties:"docker.host,default="`

	// TLSVerify is a flag to enable or disable TLS verification when connecting to a Docker daemon.
	//
	// Environment variable: DOCKER_TLS_VERIFY
	TLSVerify int `properties:"docker.tls.verify,default=0"`

	// CertPath is the path to the directory containing the Docker certificates.
	// This is used when connecting to a Docker daemon over TLS.
	//
	// Environment variable: DOCKER_CERT_PATH
	CertPath string `properties:"docker.cert.path,default="`

	// HubImageNamePrefix is the prefix used for the images pulled from the Docker Hub.
	// This is useful when running tests in environments with restricted internet access.
	//
	// Environment variable: TESTCONTAINERS_HUB_IMAGE_NAME_PREFIX
	HubImageNamePrefix string `properties:"hub.image.name.prefix,default="`

	// RyukDisabled is a flag to enable or disable the Garbage Collector.
	// Setting this to true will prevent testcontainers from automatically cleaning up
	// resources, which is particularly important in tests which timeout as they
	// don't run test clean up.
	//
	// Environment variable: TESTCONTAINERS_RYUK_DISABLED
	RyukDisabled bool `properties:"ryuk.disabled,default=false"`

	// RyukPrivileged is a flag to enable or disable the privileged mode for the Garbage Collector container.
	// Setting this to true will run the Garbage Collector container in privileged mode.
	//
	// Environment variable: TESTCONTAINERS_RYUK_CONTAINER_PRIVILEGED
	RyukPrivileged bool `properties:"ryuk.container.privileged,default=false"`

	// RyukReconnectionTimeout is the time to wait before attempting to reconnect to the Garbage Collector container.
	//
	// Environment variable: RYUK_RECONNECTION_TIMEOUT
	RyukReconnectionTimeout time.Duration `properties:"ryuk.reconnection.timeout,default=10s"`

	// RyukConnectionTimeout is the time to wait before timing out when connecting to the Garbage Collector container.
	//
	// Environment variable: RYUK_CONNECTION_TIMEOUT
	RyukConnectionTimeout time.Duration `properties:"ryuk.connection.timeout,default=1m"`

	// RyukVerbose is a flag to enable or disable verbose logging for the Garbage Collector.
	//
	// Environment variable: RYUK_VERBOSE
	RyukVerbose bool `properties:"ryuk.verbose,default=false"`

	// TestcontainersHost is the address of the Testcontainers host.
	//
	// Environment variable: TESTCONTAINERS_DOCKER_SOCKET_OVERRIDE
	TestcontainersHost string `properties:"tc.host,default="`
}

// }

// Read reads from testcontainers properties file, if it exists
// it is possible that certain values get overridden when set as environment variables
func Read() Config { _ = "STUB: not implemented"; return *new(Config) }

// Reset resets the singleton instance of the Config struct,
// allowing to read the configuration again.
// Handy for testing, so do not use it in production code
// This function is not thread-safe
func Reset() { _ = "STUB: not implemented"; return }

func read() Config { _ = "STUB: not implemented"; return *new(Config) }

// init from a file

func parseBool(input string) bool { _ = "STUB: not implemented"; return false }

// readTestcontainersEnv reads the environment variable with the given name.
// It checks for the environment variable with the given name first, and then
// checks for the environment variable with the given name prefixed with "TESTCONTAINERS_".
func readTestcontainersEnv(envVar string) string { _ = "STUB: not implemented"; return "" }

// TODO: remove this prefix after the next major release
