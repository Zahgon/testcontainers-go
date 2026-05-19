package atlaslocal

import (
	"context"
	"io"

	"github.com/testcontainers/testcontainers-go"
)

const defaultPort = "27017/tcp"

// Container represents the MongoDBAtlasLocal container type used in the module.
type Container struct {
	testcontainers.Container
	userOpts options
}

// Run creates an instance of the MongoDBAtlasLocal container type.
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set the defaults

// ConnectionString returns the connection string for the MongoDB Atlas Local
// container. If you provide a username and a password, the connection string
// will also include them.
func (ctr *Container) ConnectionString(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// If MONGODB_INITDB_DATABASE is set, use it as the default database in the
// connection string.

// ReadMongotLogs returns a reader for mongot logs in the container. Reads from
// stdout/stderr or /tmp/mongot.log if configured.
//
// This method return the os.ErrNotExist sentinel error if it is called with
// no log file configured.
func (ctr *Container) ReadMongotLogs(ctx context.Context) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// ReadRunnerLogs() returns a reader for runner logs in the container. Reads
// from stdout/stderr or /tmp/runner.log if configured.
//
// This method return the os.ErrNotExist sentinel error if it is called with
// no log file configured.
func (ctr *Container) ReadRunnerLogs(ctx context.Context) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}
