package scylladb

import (
	"context"
	"io"

	"github.com/testcontainers/testcontainers-go"
)

const (
	port           = "9042/tcp"
	shardAwarePort = "19042/tcp"
	alternatorPort = "8000/tcp"
)

// Container represents a ScyllaDB container type used in the module
type Container struct {
	testcontainers.Container
}

// WithConfig sets the YAML config file as an [io.Reader] to be used for the ScyllaDB container
func WithConfig(r io.Reader) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithShardAwareness enable shard-awareness in the ScyllaDB container so you can use the `19042` port.
func WithShardAwareness() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithAlternator enables the Alternator (DynamoDB Compatible API) service in the ScyllaDB container,
// using the default HTTP port 8000.
// It will set the "alternator-port" parameter to the specified port.
// It will also set the "alternator-write-isolation" parameter to "always" as a command line argument to the container.
func WithAlternator() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithCustomCommands sets custom commands with values for the ScyllaDB container.
// Pass the command and the values as a list of strings in the following format: "--flag1=value", "--flag2", etc.
// In case of an invalid flag (not starting with "--" or "-"), this option returns an error,
// not applying any changes to the command line. Else, flags that exist in the command line overwrite the default commands.
// See more in the [ScyllaDB docs].
//
// [ScyllaDB docs]: https://opensource.docs.scylladb.com/stable/operating-scylla/procedures/tips/best-practices-scylla-on-docker.html
func WithCustomCommands(flags ...string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// ShardAwareConnectionHost returns the host and port of the ScyllaDB container with the shard-aware port
func (c Container) ShardAwareConnectionHost(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// NonShardAwareConnectionHost returns the host and port of the ScyllaDB container with the non-shard-aware port
func (c Container) NonShardAwareConnectionHost(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// AlternatorConnectionHost returns the host and port of the ScyllaDB container with the alternator port
func (c Container) AlternatorConnectionHost(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Run starts a ScyllaDB container with the specified image and options
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// setCommandFlags sets the flags in the command line.
// It takes the container request and a map of flags,
// and checks if the flag is present in the command line, overriding the value if it is.
// If the flag is not present, it's added to the end of the command line.
func setCommandFlags(req *testcontainers.GenericContainerRequest, flagsMap map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// The flag is present in the command line, so it's removed from the flagsMap
// to avoid adding it to the end of the command line.

// The extra flags not present in the command line are added to the end of the command line,
// and this could be in any order.
