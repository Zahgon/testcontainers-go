package etcd

import (
	"github.com/testcontainers/testcontainers-go"
)

type options struct {
	currentNode    int
	clusterNetwork *testcontainers.DockerNetwork
	nodeNames      []string
	clusterToken   string
	additionalArgs []string
	mountDataDir   bool // flag needed to avoid extra calculations with the lifecycle hooks
}

func defaultOptions() options { _ = "STUB: not implemented"; return *new(options) }

// Compiler check to ensure that Option implements the testcontainers.ContainerCustomizer interface.
var _ testcontainers.ContainerCustomizer = (Option)(nil)

// Option is an option for the Etcd container.
type Option func(*options) error

// Customize is a NOOP. It's defined to satisfy the testcontainers.ContainerCustomizer interface.
func (o Option) Customize(*testcontainers.GenericContainerRequest) error {
	_ = "STUB: not implemented"
	// NOOP to satisfy interface.
	return nil
}

// WithAdditionalArgs is an option to pass additional arguments to the etcd container.
// They will be appended last to the command line.
func WithAdditionalArgs(args ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDataDir is an option to mount the data directory, which is located at /data.etcd.
// The option will add a lifecycle hook to the container to change the permissions of the data directory.
func WithDataDir() Option { _ = "STUB: not implemented"; return *new(Option) }

// Avoid extra calculations with the lifecycle hooks

// WithNodes is an option to set the nodes of the etcd cluster.
// It should be used to create a cluster with more than one node.
func WithNodes(node1 string, node2 string, nodes ...string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// withCurrentNode is an option to set the current node index.
// It's an internal option and should not be used by the user.
func withCurrentNode(i int) Option { _ = "STUB: not implemented"; return *new(Option) }

// withClusterNetwork is an option to set the cluster network.
// It's an internal option and should not be used by the user.
func withClusterNetwork(n *testcontainers.DockerNetwork) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithClusterToken is an option to set the cluster token.
func WithClusterToken(token string) Option { _ = "STUB: not implemented"; return *new(Option) }

func withClusterOptions(opts []Option) Option { _ = "STUB: not implemented"; return *new(Option) }
