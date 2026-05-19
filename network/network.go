package network

import (
	"context"

	"github.com/moby/moby/api/types/network"
	"github.com/moby/moby/client"

	"github.com/testcontainers/testcontainers-go"
)

// New creates a new network with a random UUID name, calling the already existing GenericNetwork APIs.
// Those existing APIs are deprecated and will be removed in the future, so this function will
// implement the new network APIs when they will be available.
// By default, the network is created with the following options:
// - Driver: bridge
// - Labels: the Testcontainers for Go generic labels, to be managed by Ryuk. Please see the GenericLabels() function
// And those options can be modified by the user, using the CreateModifier function field.
func New(ctx context.Context, opts ...NetworkCustomizer) (*testcontainers.DockerNetwork, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:staticcheck

//nolint:staticcheck

// Return a DockerNetwork struct instead of the Network interface,
// following the "accept interface, return struct" pattern.

// NetworkCustomizer is an interface that can be used to configure the network create request.
type NetworkCustomizer interface {
	Customize(req *client.NetworkCreateOptions) error
}

// CustomizeNetworkOption is a type that can be used to configure the network create request.
type CustomizeNetworkOption func(req *client.NetworkCreateOptions) error

// Customize implements the NetworkCustomizer interface,
// applying the option to the network create request.
func (opt CustomizeNetworkOption) Customize(req *client.NetworkCreateOptions) error {
	_ = "STUB: not implemented"

	// WithAttachable allows to set the network as attachable.
	return nil
}

func WithAttachable() CustomizeNetworkOption {
	_ = "STUB: not implemented"
	return *new(CustomizeNetworkOption)
}

// WithCheckDuplicate allows to check if a network with the same name already exists.
//
// Deprecated: CheckDuplicate is deprecated since API v1.44, but it defaults to true when sent by the client package to older daemons.
func WithCheckDuplicate() CustomizeNetworkOption {
	_ = "STUB: not implemented"
	return *new(CustomizeNetworkOption)
}

// WithDriver allows to override the default network driver, which is "bridge".
func WithDriver(driver string) CustomizeNetworkOption {
	_ = "STUB: not implemented"
	return *new(CustomizeNetworkOption)
}

// WithEnableIPv6 allows to set the network as IPv6 enabled.
// Please use this option if and only if IPv6 is enabled on the Docker daemon.
func WithEnableIPv6() CustomizeNetworkOption {
	_ = "STUB: not implemented"
	return *new(CustomizeNetworkOption)
}

// WithInternal allows to set the network as internal.
func WithInternal() CustomizeNetworkOption {
	_ = "STUB: not implemented"
	return *new(CustomizeNetworkOption)
}

// WithLabels allows to set the network labels, adding the new ones
// to the default Testcontainers for Go labels.
func WithLabels(labels map[string]string) CustomizeNetworkOption {
	_ = "STUB: not implemented"
	return *new(CustomizeNetworkOption)
}

// WithIPAM allows to change the default IPAM configuration.
func WithIPAM(ipam *network.IPAM) CustomizeNetworkOption {
	_ = "STUB: not implemented"
	return *new(CustomizeNetworkOption)
}

// WithNetwork reuses an already existing network, attaching the container to it.
// Finally it sets the network alias on that network to the given alias.
func WithNetwork(aliases []string, nw *testcontainers.DockerNetwork) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithNetworkName attachs a container to an already existing network, by its name.
// If the network is not "bridge", it sets the network alias on that network
// to the given alias, else, it returns an error. This is because network-scoped alias
// is supported only for containers in user defined networks.
func WithNetworkName(aliases []string, networkName string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// attaching to the network because it was created with success or it already existed.

// WithBridgeNetwork attachs a container to the "bridge" network.
// There is no need to set the network alias, as it is not supported for the "bridge" network.
func WithBridgeNetwork() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithNewNetwork creates a new network with random name and customizers, and attaches the container to it.
// Finally it sets the network alias on that network to the given alias.
func WithNewNetwork(ctx context.Context, aliases []string, opts ...NetworkCustomizer) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// attaching to the network because it was created with success or it already existed.
