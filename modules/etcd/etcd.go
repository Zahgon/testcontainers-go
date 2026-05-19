package etcd

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	clientPort          = "2379"
	peerPort            = "2380"
	dataDir             = "/data.etcd"
	defaultClusterToken = "mys3cr3ttok3n"
	scheme              = "http"
)

// EtcdContainer represents the etcd container type used in the module. It can be used to create a single-node instance or a cluster.
// For the cluster, the first node creates the cluster and the other nodes join it as child nodes.
type EtcdContainer struct {
	testcontainers.Container
	// childNodes contains the child nodes of the current node, forming a cluster
	childNodes []*EtcdContainer
	opts       options
}

// Terminate terminates the etcd container, its child nodes, and the network in which the cluster is running
// to communicate between the nodes.
func (c *EtcdContainer) Terminate(ctx context.Context, opts ...testcontainers.TerminateOption) error {
	_ = "STUB: not implemented"

	// child nodes has no other children
	return nil
}

// remove the cluster network if it was created, but only for the first node
// we could check if the current node is the first one (index 0),
// and/or check that there are no child nodes

// Run creates an instance of the etcd container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*EtcdContainer, error) {
	_ = "STUB: not implemented"
	// Process custom options first to extract settings
	return nil, nil
}

// Build moduleOpts with defaults

// Append user options

// configure CMD with the nodes

// Initialise the etcd container with the current settings.
// The cluster network, if needed, is already part of the settings,
// so the following error handling returns a partially initialised container,
// allowing the caller to clean up the resources with the Terminate method.

// apply the network to the current node

// only the first node creates the cluster

// move to the next node

// return the parent cluster node and the error, so the caller can clean up.

// configureCluster configures the cluster settings, ensuring that the cluster is properly configured with the necessary network and options,
// avoiding duplicate application of options to be passed to the successive nodes.
func configureCluster(ctx context.Context, settings *options, opts []testcontainers.ContainerCustomizer) ([]testcontainers.ContainerCustomizer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// pass cluster options to each node

// if the option is of type Option, it won't be applied to the settings
// this prevents the same option from being applied multiple times (e.g. updating the current node)

// the first time the network is created

// set the network for the first node

// save the network for the next nodes

// we finally need to re-apply all the etcd-specific options

// configureCMD configures the etcd command line arguments, based on the settings provided,
// in order to create a cluster or a single-node instance.
func configureCMD(settings options) []string { _ = "STUB: not implemented"; return nil }

// ClientEndpoint returns the client endpoint for the etcd container, and an error if any.
// For a cluster, it returns the client endpoint of the first node.
func (c *EtcdContainer) ClientEndpoint(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ClientEndpoints returns the client endpoints for the etcd cluster.
func (c *EtcdContainer) ClientEndpoints(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PeerEndpoint returns the peer endpoint for the etcd container, and an error if any.
// For a cluster, it returns the peer endpoint of the first node.
func (c *EtcdContainer) PeerEndpoint(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// PeerEndpoints returns the peer endpoints for the etcd cluster.
func (c *EtcdContainer) PeerEndpoints(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
