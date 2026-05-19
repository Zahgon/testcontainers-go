package influxdb

import (
	"context"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// InfluxDbContainer represents the InfluxDB container type used in the module
//
//nolint:staticcheck //FIXME
type InfluxDbContainer struct {
	testcontainers.Container
}

// Deprecated: use Run instead
// RunContainer creates an instance of the InfluxDB container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*InfluxDbContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the InfluxDB container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*InfluxDbContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:revive,staticcheck //FIXME
func (c *InfluxDbContainer) MustConnectionUrl(ctx context.Context) string {
	_ = "STUB: not implemented"
	return ""
}

//nolint:revive,staticcheck //FIXME
func (c *InfluxDbContainer) ConnectionUrl(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func WithUsername(username string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

func WithPassword(password string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

func WithDatabase(database string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

func WithConfigFile(configFile string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// withV2 configures the influxdb container to be compatible with InfluxDB v2
func withV2(req *testcontainers.GenericContainerRequest, org, bucket string) error {
	_ = "STUB: not implemented"
	return nil
}

// Always setup, we wont be migrating from v1 to v2

// WithV2 configures the influxdb container to be compatible with InfluxDB v2
func WithV2(org, bucket string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

const dockerSecretPath = "/run/secrets"

func secretsPath(path string) string { _ = "STUB: not implemented"; return "" }

// WithV2Auth configures the influxdb container to be compatible with InfluxDB v2 and sets the username and password
// for the initial user.
func WithV2Auth(org, bucket, username, password string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithV2SecretsAuth configures the container to be compatible with InfluxDB v2 and sets the username and password file path
func WithV2SecretsAuth(org, bucket, usernameFile, passwordFile string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithV2Retention configures the default bucket's retention
func WithV2Retention(retention time.Duration) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithV2AdminToken sets the admin token for the influxdb container
func WithV2AdminToken(token string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithV2SecretsAdminToken sets the admin token for the influxdb container using a file
func WithV2SecretsAdminToken(tokenFile string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithInitDb returns a request customizer that initialises the database using the file `docker-entrypoint-initdb.d`
// located in `srcPath` directory.
//
//nolint:staticcheck //FIXME
func WithInitDb(srcPath string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

func waitForHTTPHealth() *wait.HTTPStrategy { _ = "STUB: not implemented"; return nil }
