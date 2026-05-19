package testcontainers

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/ssh"
)

const (
	// hubSshdImage {
	sshdImage string = "testcontainers/sshd:1.3.0"
	// }

	// HostInternal is the internal hostname used to reach the host from the container,
	// using the SSHD container as a bridge.
	HostInternal string = "host.testcontainers.internal"
	user         string = "root"
	sshPort             = "22/tcp"
)

// sshPassword is a random password generated for the SSHD container.
var sshPassword = uuid.NewString()

// exposeHostPorts performs all the necessary steps to expose the host ports to the container, leveraging
// the SSHD container to create the tunnel, and the container lifecycle hooks to manage the tunnel lifecycle.
// At least one port must be provided to expose.
// The steps are:
// 1. Create a new SSHD container.
// 2. Expose the host ports to the container after the container is ready.
// 3. Close the SSH sessions before killing the container.
func exposeHostPorts(ctx context.Context, req *ContainerRequest, ports ...int) (sshdConnectHook ContainerLifecycleHooks, err error) {
	_ = "STUB: not implemented"
	return *new(ContainerLifecycleHooks), nil
}

// Use the first network of the container to connect to the SSHD container.

// get the first network of the container to connect the SSHD container to it.

// WithNetwork reuses an already existing network, attaching the container to it.
// Finally it sets the network alias on that network to the given alias.
// TODO: Using an anonymous function to avoid cyclic dependencies with the network package.

// attaching to the network because it was created with success or it already existed.

// start the SSHD container with the provided options

// Ensure the SSHD container is stopped and removed in case of error.

// IP in the first network of the container.

// do not override the original HostConfigModifier

// adding the host internal alias to the container as an extra host
// to allow the container to reach the SSHD container.

// if the container is not in one of the modes, attach it to the first network of the SSHD container

// invoke the original HostConfigModifier with the updated hostConfig

// Context already canceled, need to create a new one to ensure
// the SSH session is closed.

// after the container is ready, create the SSH tunnel
// for each exposed port from the host.

// newSshdContainer creates a new SSHD container with the provided options.
func newSshdContainer(ctx context.Context, opts ...ContainerCustomizer) (*sshdContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Return the container and the error to the caller to handle it.

// sshdContainer represents the SSHD container type used for the port forwarding container.
// It's an internal type that extends the DockerContainer type, to add the SSH tunnelling capabilities.
type sshdContainer struct {
	Container
	port           string
	sshConfig      *ssh.ClientConfig
	portForwarders []*portForwarder
}

// Terminate stops the container and closes the SSH session
func (sshdC *sshdContainer) Terminate(ctx context.Context, opts ...TerminateOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Stop stops the container and closes the SSH session
func (sshdC *sshdContainer) Stop(ctx context.Context, timeout *time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// closePorts closes all port forwarders.
func (sshdC *sshdContainer) closePorts() error { _ = "STUB: not implemented"; return nil }

// Ensure the port forwarders are not used after closing.

// clientConfig sets up the SSHD client configuration.
func (sshdC *sshdContainer) clientConfig(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// exposeHostPort exposes the host ports to the container.
func (sshdC *sshdContainer) exposeHostPort(ctx context.Context, ports ...int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// portForwarder forwards a port from the container to the host.
type portForwarder struct {
	client      *ssh.Client
	listener    net.Listener
	dialTimeout time.Duration
	localAddr   string
	ctx         context.Context
	cancel      context.CancelFunc

	// closeMtx protects the close operation
	closeMtx sync.Mutex
	closeErr error
}

// newPortForwarder creates a new running portForwarder for the given port.
// The context is only used for the initial SSH connection.
func newPortForwarder(ctx context.Context, sshDAddr string, sshConfig *ssh.ClientConfig, port int) (pf *portForwarder, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ensure the connection is closed in case of error.

// Close closes the port forwarder.
func (pf *portForwarder) Close() error { _ = "STUB: not implemented"; return nil }

// Already closed.

// run forwards the port from the remote connection to the local connection.
func (pf *portForwarder) run() { _ = "STUB: not implemented"; return }

// The listener has been closed.

// Ignore errors as they are transient and we want requests to
// continue to be accepted.

// tunnel runs a tunnel between two connections; as soon as the forwarder
// context is cancelled or one connection copies returns, irrespective of
// the error, both connections are closed.
func (pf *portForwarder) tunnel(remote net.Conn) { _ = "STUB: not implemented"; return }

// Nothing we can do with the error.

//nolint:errcheck // Nothing useful we can do with the error.

//nolint:errcheck // Nothing useful we can do with the error.

// Wait for the context to be done before returning which triggers
// both connections to close. This is done to prevent the copies
// blocking forever on unused connections.
