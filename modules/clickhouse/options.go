package clickhouse

import (
	_ "embed"

	"github.com/testcontainers/testcontainers-go"
)

//go:embed mounts/zk_config.xml.tpl
var zookeeperConfigTpl string

// ZookeeperOptions arguments for zookeeper in clickhouse
type ZookeeperOptions struct {
	Host, Port string
}

// renderZookeeperConfig generate default zookeeper configuration for clickhouse
func renderZookeeperConfig(settings ZookeeperOptions) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WithZookeeper pass a config to connect clickhouse with zookeeper and make clickhouse as cluster.
// It creates a temporary file in the host filesystem with the config and copies it to the container
// at /etc/clickhouse-server/config.d/zookeeper_config.xml. This file is not cleaned up automatically,
// and it's removed by the OS.
func WithZookeeper(host, port string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// write data to the temporary file

// WithInitScripts sets the init scripts to be run when the container starts
func WithInitScripts(scripts ...string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithConfigFile sets the XML config file to be used for the clickhouse container.
// The file is copied to the container at /etc/clickhouse-server/config.d/config.xml,
// which is the default location for ClickHouse config files.
func WithConfigFile(configFile string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithYamlConfigFile sets the YAML config file to be used for the clickhouse container
// The file is copied to the container at /etc/clickhouse-server/config.d/config.yaml,
// which is the default location for ClickHouse YAML config files.
func WithYamlConfigFile(configFile string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithDatabase sets the initial database to be created when the container starts
// It can be used to define a different name for the default database that is created when the image is first started.
// If it is not specified, then the default value("clickhouse") will be used.
func WithDatabase(dbName string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithPassword sets the initial password of the user to be created when the container starts
// It is required for you to use the ClickHouse image. It must not be empty or undefined.
// This environment variable sets the password for ClickHouse.
func WithPassword(password string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithUsername sets the initial username to be created when the container starts
// It is used in conjunction with WithPassword to set a user and its password.
// It will create the specified user with superuser power.
// If it is not specified, then the default user of clickhouse will be used.
func WithUsername(user string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}
