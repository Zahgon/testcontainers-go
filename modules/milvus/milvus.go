package milvus

import (
	"context"
	_ "embed"
	"io"

	"github.com/testcontainers/testcontainers-go"
)

//go:embed mounts/embedEtcd.yaml.tpl
var embedEtcdConfigTpl string

const (
	embedEtcdContainerPath = "/milvus/configs/embedEtcd.yaml"
	defaultClientPort      = 2379
	etcdPort               = "2379/tcp"
	httpPort               = "9091/tcp"
	grpcPort               = "19530/tcp"
)

// MilvusContainer represents the Milvus container type used in the module
type MilvusContainer struct {
	testcontainers.Container
}

// ConnectionString returns the connection string for the milvus container, using the default 19530 port, and
// obtaining the host and exposed port from the container.
func (c *MilvusContainer) ConnectionString(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Deprecated: use Run instead
// RunContainer creates an instance of the Milvus container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*MilvusContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the Milvus container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*MilvusContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Adapted from https://github.com/milvus-io/milvus/blob/v2.6.3/scripts/standalone_embed.sh

type embedEtcdConfigTplParams struct {
	Port int
}

// renderEmbedEtcdConfig renders the embed etcd config template with the given port
// and returns it as an io.Reader.
func renderEmbedEtcdConfig(port int) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}
