package mongodb

import (
	"context"
	_ "embed"

	"github.com/testcontainers/testcontainers-go"
)

//go:embed mount/entrypoint-tc.sh
var entrypointContent []byte

const (
	defaultPort         = "27017/tcp"
	entrypointPath      = "/tmp/entrypoint-tc.sh"
	keyFilePath         = "/tmp/mongo_keyfile"
	replicaSetOptEnvKey = "testcontainers.mongodb.replicaset_name"
)

// MongoDBContainer represents the MongoDB container type used in the module
type MongoDBContainer struct {
	testcontainers.Container
	username   string
	password   string
	replicaSet string
}

// Deprecated: use Run instead
// RunContainer creates an instance of the MongoDB container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*MongoDBContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the MongoDB container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*MongoDBContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// configure the request for the replicaset after all the options have been applied

// Inspect the container to get environment variables

// refresh the credentials from the environment variables

// WithUsername sets the initial username to be created when the container starts
// It is used in conjunction with WithPassword to set a username and its password.
// It will create the specified user with superuser power.
func WithUsername(username string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithPassword sets the initial password of the user to be created when the container starts
// It is used in conjunction with WithUsername to set a username and its password.
// It will set the superuser password for MongoDB.
func WithPassword(password string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithReplicaSet sets the replica set name for Single node MongoDB replica set.
func WithReplicaSet(replSetName string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// ConnectionString returns the connection string for the MongoDB container.
// If you provide a username and a password, the connection string will also include them.
func (c *MongoDBContainer) ConnectionString(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func setupEntrypointForAuth() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

func configureReplicaset() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

func noAuthReplicaSet(replSetName string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

func initiateReplicaSet(cli mongoCli, replSetName string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

func withAuthReplicaset(
	replSetName string,
	username string,
	password string,
) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}
