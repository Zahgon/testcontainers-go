package testcontainers

import (
	"archive/tar"
	"context"
	"errors"
	"io"
	"regexp"
	"sync"
	"sync/atomic"
	"time"

	"github.com/containerd/errdefs"
	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/api/types/network"
	"github.com/moby/moby/client"
	specs "github.com/opencontainers/image-spec/specs-go/v1"

	tcexec "github.com/testcontainers/testcontainers-go/exec"
	"github.com/testcontainers/testcontainers-go/internal/config"
	"github.com/testcontainers/testcontainers-go/log"
	"github.com/testcontainers/testcontainers-go/wait"
)

// Implement interfaces
var _ Container = (*DockerContainer)(nil)

const (
	Bridge        = "bridge" // Bridge network name (as well as driver)
	Podman        = "podman"
	ReaperDefault = "reaper_default" // Default network name when bridge is not available
	packagePath   = "github.com/testcontainers/testcontainers-go"
)

var (
	// createContainerFailDueToNameConflictRegex is a regular expression that matches the container is already in use error.
	createContainerFailDueToNameConflictRegex = regexp.MustCompile("[Tt]he container name .* is already in use by .*")

	// minLogProductionTimeout is the minimum log production timeout.
	minLogProductionTimeout = time.Duration(5 * time.Second)

	// maxLogProductionTimeout is the maximum log production timeout.
	maxLogProductionTimeout = time.Duration(60 * time.Second)

	// errLogProductionStop is the cause for stopping log production.
	errLogProductionStop = errors.New("log production stopped")
)

// DockerContainer represents a container started using Docker
type DockerContainer struct {
	// Container ID from Docker
	ID           string
	WaitingFor   wait.Strategy
	Image        string
	exposedPorts []string // a reference to the container's requested exposed ports. It allows checking they are ready before any wait strategy

	isRunning     atomic.Bool
	imageWasBuilt bool
	// keepBuiltImage makes Terminate not remove the image if imageWasBuilt.
	keepBuiltImage    bool
	provider          *DockerProvider
	sessionID         string
	terminationSignal chan bool
	consumersMtx      sync.Mutex // protects consumers
	consumers         []LogConsumer

	// TODO: Remove locking and wait group once the deprecated StartLogProducer and
	// StopLogProducer have been removed and hence logging can only be started and
	// stopped once.

	// logProductionCancel is used to signal the log production to stop.
	logProductionCancel context.CancelCauseFunc
	logProductionCtx    context.Context
	// logProductionDone is closed when the log production goroutine exits.
	logProductionDone chan struct{}

	logProductionTimeout *time.Duration
	logger               log.Logger
	lifecycleHooks       []ContainerLifecycleHooks

	healthStatus container.HealthStatus // container health status, will default to healthStatusNone if no healthcheck is present
}

// SetLogger sets the logger for the container
func (c *DockerContainer) SetLogger(logger log.Logger) {
	_ = "STUB: not implemented"

	// SetProvider sets the provider for the container
	return
}

func (c *DockerContainer) SetProvider(provider *DockerProvider) { _ = "STUB: not implemented"; return }

// SetTerminationSignal sets the termination signal for the container
func (c *DockerContainer) SetTerminationSignal(signal chan bool) { _ = "STUB: not implemented"; return }

func (c *DockerContainer) GetContainerID() string { _ = "STUB: not implemented"; return "" }

func (c *DockerContainer) IsRunning() bool { _ = "STUB: not implemented"; return false }

// Endpoint gets proto://host:port string for the lowest numbered exposed port
// Will returns just host:port if proto is ""
func (c *DockerContainer) Endpoint(ctx context.Context, proto string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Get lowest numbered bound port.

// PortEndpoint gets proto://host:port string for the given exposed port
// It returns proto://host:port or proto://[IPv6host]:port string for the given exposed port.
// It returns just host:port or [IPv6host]:port if proto is blank.
func (c *DockerContainer) PortEndpoint(ctx context.Context, port string, proto string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Host gets host (ip or name) of the docker daemon where the container port is exposed
// Warning: this is based on your Docker host setting. Will fail if using an SSH tunnel
// You can use the "TESTCONTAINERS_HOST_OVERRIDE" env variable to set this yourself
func (c *DockerContainer) Host(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Inspect gets the raw container info
func (c *DockerContainer) Inspect(ctx context.Context) (*container.InspectResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MappedPort gets externally mapped port for a container port
func (c *DockerContainer) MappedPort(ctx context.Context, port string) (network.Port, error) {
	_ = "STUB: not implemented"
	return *new(network.Port), nil
}

// The old nat.Port type (a plain string) accepted empty strings:
// nat.SplitProtoPort("") returns ("", ""), so Port() == "" and
// no container port matches, yielding "not found".
// See https://github.com/docker/go-connections/blob/v0.6.0/nat/nat.go#L101-L110
// Skip parsing here to preserve that behavior and avoid a
// ParsePort error on empty input.

// Deprecated: use c.Inspect(ctx).NetworkSettings.Ports instead.
// Ports gets the exposed ports for the container.
func (c *DockerContainer) Ports(ctx context.Context) (network.PortMap, error) {
	_ = "STUB: not implemented"
	return *new(network.PortMap), nil
}

// SessionID gets the current session id
func (c *DockerContainer) SessionID() string {
	_ = "STUB: not implemented"

	// Start will start an already created container
	return ""
}

func (c *DockerContainer) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Stop stops the container.
//
// In case the container fails to stop gracefully within a time frame specified
// by the timeout argument, it is forcefully terminated (killed).
//
// If the timeout is nil, the container's StopTimeout value is used, if set,
// otherwise the engine default. A negative timeout value can be specified,
// meaning no timeout, i.e. no forceful termination is performed.
//
// All hooks are called in the following order:
//   - [ContainerLifecycleHooks.PreStops]
//   - [ContainerLifecycleHooks.PostStops]
//
// If the container is already stopped, the method is a no-op.
func (c *DockerContainer) Stop(ctx context.Context, timeout *time.Duration) error {
	_ = "STUB: not implemented"
	// Note we can't check isRunning here because we allow external creation
	// without exposing the ability to fully initialize the container state.
	// See: https://github.com/testcontainers/testcontainers-go/issues/2667
	// TODO: Add a check for isRunning when the above issue is resolved.
	return nil
}

// Terminate calls stops and then removes the container including its volumes.
// If its image was built it and all child images are also removed unless
// the [FromDockerfile.KeepImage] on the [ContainerRequest] was set to true.
//
// The following hooks are called in order:
//   - [ContainerLifecycleHooks.PreTerminates]
//   - [ContainerLifecycleHooks.PostTerminates]
//
// Default: timeout is 10 seconds.
func (c *DockerContainer) Terminate(ctx context.Context, opts ...TerminateOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Close reaper connection if it was attached.

// TODO: Handle errors from ContainerRemove more correctly, e.g. should we
// run the terminated hook?

// update container raw info
func (c *DockerContainer) inspectRawContainer(ctx context.Context) (*client.ContainerInspectResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Logs will fetch both STDOUT and STDERR from the current container. Returns a
// ReadCloser and leaves it up to the caller to extract what it wants.
func (c *DockerContainer) Logs(ctx context.Context) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// parseMultiplexedLogs handles the multiplexed log format used when TTY is disabled
func (c *DockerContainer) parseMultiplexedLogs(rc io.ReadCloser) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

// Deprecated: use the ContainerRequest.LogConsumerConfig field instead.
func (c *DockerContainer) FollowOutput(consumer LogConsumer) { _ = "STUB: not implemented"; return }

// followOutput adds a LogConsumer to be sent logs from the container's
// STDOUT and STDERR
func (c *DockerContainer) followOutput(consumer LogConsumer) { _ = "STUB: not implemented"; return }

// consumersCopy returns a copy of the current consumers.
func (c *DockerContainer) consumersCopy() []LogConsumer { _ = "STUB: not implemented"; return nil }

// resetConsumers resets the current consumers to the provided ones.
func (c *DockerContainer) resetConsumers(consumers []LogConsumer) {
	_ = "STUB: not implemented"
	return
}

// Deprecated: use c.Inspect(ctx).Name instead.
// Name gets the name of the container.
func (c *DockerContainer) Name(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// State returns container's running state.
func (c *DockerContainer) State(ctx context.Context) (*container.State, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Networks gets the names of the networks the container is attached to.
func (c *DockerContainer) Networks(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ContainerIP gets the IP address of the primary network within the container.
func (c *DockerContainer) ContainerIP(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

//  IPAddress is deprecated; use IP from "Networks" if only single network defined

// ContainerIPs gets the IP addresses of all the networks within the container.
func (c *DockerContainer) ContainerIPs(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NetworkAliases gets the aliases of the container for the networks it is attached to.
func (c *DockerContainer) NetworkAliases(ctx context.Context) (map[string][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Exec executes a command in the current container.
// It returns the exit status of the executed command, an [io.Reader] containing the combined
// stdout and stderr, and any encountered error. Note that reading directly from the [io.Reader]
// may result in unexpected bytes due to custom stream multiplexing headers.
// Use [tcexec.Multiplexed] option to read the combined output without the multiplexing headers.
// Alternatively, to separate the stdout and stderr from [io.Reader] and interpret these headers properly,
// [github.com/docker/docker/pkg/stdcopy.StdCopy] from the Docker API should be used.
func (c *DockerContainer) Exec(ctx context.Context, cmd []string, options ...tcexec.ProcessOption) (int, io.Reader, error) {
	_ = "STUB: not implemented"
	return 0, *new(io.Reader), nil
}

// processing all the options in a first loop because for the multiplexed option
// we first need to have a containerExecCreateResponse

// second loop to process the multiplexed option, as now we have a reader
// from the created exec response.

type FileFromContainer struct {
	underlying *io.ReadCloser
	tarreader  *tar.Reader
}

func (fc *FileFromContainer) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (fc *FileFromContainer) Close() error { _ = "STUB: not implemented"; return nil }

func (c *DockerContainer) CopyFileFromContainer(ctx context.Context, filePath string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// if we got here we have exactly one file in the TAR-stream
// so we advance the index by one so the next call to Read will start reading it

// CopyDirToContainer copies the contents of a directory to a parent path in the container. This parent path must exist in the container first
// as we cannot create it
func (c *DockerContainer) CopyDirToContainer(ctx context.Context, hostDirPath string, containerParentPath string, fileMode int64) error {
	_ = "STUB: not implemented"
	return nil
}

// it's not a dir: let the consumer handle the error

// create the directory under its parent

func (c *DockerContainer) CopyFileToContainer(ctx context.Context, hostFilePath string, containerFilePath string, fileMode int64) error {
	_ = "STUB: not implemented"
	return nil
}

// In Go 1.22 os.File is always an io.WriterTo. However, testcontainers
// currently allows Go 1.21, so we need to trick the compiler a little.

// Attempt optimized writeTo, implemented in linux

// CopyToContainer copies fileContent data to a file in container
func (c *DockerContainer) CopyToContainer(ctx context.Context, fileContent []byte, containerFilePath string, fileMode int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *DockerContainer) copyToContainer(ctx context.Context, fileContent func(tw io.Writer) error, fileContentSize int64, containerFilePath string, fileMode int64) error {
	_ = "STUB: not implemented"
	return nil
}

// logConsumerWriter is a writer that writes to a LogConsumer.
type logConsumerWriter struct {
	log       Log
	consumers []LogConsumer
}

// newLogConsumerWriter creates a new logConsumerWriter for logType that sends messages to all consumers.
func newLogConsumerWriter(logType string, consumers []LogConsumer) *logConsumerWriter {
	_ = "STUB: not implemented"
	return nil
}

// Write writes the p content to all consumers.
func (lw logConsumerWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

type LogProductionOption func(*DockerContainer)

// WithLogProductionTimeout is a functional option that sets the timeout for the log production.
// If the timeout is lower than 5s or greater than 60s it will be set to 5s or 60s respectively.
func WithLogProductionTimeout(timeout time.Duration) LogProductionOption {
	_ = "STUB: not implemented"
	return *new(LogProductionOption)
}

// Deprecated: use the ContainerRequest.LogConsumerConfig field instead.
func (c *DockerContainer) StartLogProducer(ctx context.Context, opts ...LogProductionOption) error {
	_ = "STUB: not implemented"
	return nil
}

// startLogProduction will start a concurrent process that will continuously read logs
// from the container and will send them to each added LogConsumer.
//
// Default log production timeout is 5s. It is used to set the context timeout
// which means that each log-reading loop will last at up to the specified timeout.
//
// Use functional option WithLogProductionTimeout() to override default timeout. If it's
// lower than 5s and greater than 60s it will be set to 5s or 60s respectively.
func (c *DockerContainer) startLogProduction(ctx context.Context, opts ...LogProductionOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate the log production timeout.

// Setup the log writers.

// Setup the log production context which will be used to stop the log production.

// We capture context cancel function to avoid data race with multiple
// calls to startLogProduction.

// Ensure the context is cancelled when log productions completes
// so that GetLogProductionErrorChannel functions correctly.

// Signal that the goroutine has exited so stopLogProduction can drain.

// logProducer read logs from the container and writes them to stdout, stderr until either:
//   - logProductionCtx is done
//   - A fatal error occurs
//   - No more logs are available
func (c *DockerContainer) logProducer(stdout, stderr io.Writer) {
	_ = "STUB: not implemented"
	// Clean up idle client connections.
	return
}

// Setup the log options, start from the beginning.

// Use a separate method so that timeout cancel function is
// called correctly.

// copyLogsTimeout copies logs from the container to stdout and stderr with a timeout.
// It returns true if the log production should be retried, false otherwise.
func (c *DockerContainer) copyLogsTimeout(stdout, stderr io.Writer, options *client.ContainerLogsOptions) bool {
	_ = "STUB: not implemented"
	return false
}

// No more logs available.

// Log production was stopped or caller context is done.

// Timeout or client connection closed, retry.

// Unexpected error, retry.

// Retry from the last log received.

// copyLogs copies logs from the container to stdout and stderr.
func (c *DockerContainer) copyLogs(ctx context.Context, stdout, stderr io.Writer, options client.ContainerLogsOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: it will be removed in the next major release.
func (c *DockerContainer) StopLogProducer() error { _ = "STUB: not implemented"; return nil }

// stopLogProduction will stop the concurrent process that is reading logs
// and sending them to each added LogConsumer
func (c *DockerContainer) stopLogProduction() error { _ = "STUB: not implemented"; return nil }

// Wait for the log production goroutine to finish draining any buffered
// logs before cancelling. When the container has already exited, the
// goroutine will reach EOF naturally and close logProductionDone on its
// own. The bounded timeout prevents blocking indefinitely when the
// container is still actively streaming (e.g. Stop() called on a running
// container).

// Goroutine already finished naturally; nothing more to do.

// Timed out waiting for natural exit; force-cancel now.

// Signal the log production to stop (for still-running containers).

// Wait for the goroutine to acknowledge the cancellation. Context
// cancellation propagates into the Docker transport and should unblock
// stdcopy.StdCopy promptly, but we bound the wait to match
// minLogProductionTimeout to guard against stuck kernel socket reads or
// daemon transport failures that might not honour context cancellation.

// Log production was stopped.

// Parent context is done.

// GetLogProductionErrorChannel exposes the only way for the consumer
// to be able to listen to errors and react to them.
func (c *DockerContainer) GetLogProductionErrorChannel() <-chan error {
	_ = "STUB: not implemented"
	return nil
}

// connectReaper connects the reaper to the container if it is needed.
func (c *DockerContainer) connectReaper(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Reaper is disabled or we are the reaper container.

// cleanupTermSignal triggers the termination signal if it was created and an error occurred.
func (c *DockerContainer) cleanupTermSignal(err error) { _ = "STUB: not implemented"; return }

// DockerNetwork represents a network started using Docker
type DockerNetwork struct {
	ID                string // Network ID from Docker
	Driver            string
	Name              string
	provider          *DockerProvider
	terminationSignal chan bool
}

// Remove is used to remove the network. It is usually triggered by as defer function.
func (n *DockerNetwork) Remove(ctx context.Context) error {
	_ = "STUB: not implemented"

	// close reaper if it was created
	return nil
}

func (n *DockerNetwork) SetTerminationSignal(signal chan bool) { _ = "STUB: not implemented"; return }

// DockerProvider implements the ContainerProvider interface
type DockerProvider struct {
	*DockerProviderOptions
	client    client.APIClient
	host      string
	hostCache string
	config    config.Config
	mtx       sync.Mutex
}

// Client gets the docker client used by the provider
func (p *DockerProvider) Client() client.APIClient {
	_ = "STUB: not implemented"

	// Close closes the docker client used by the provider
	return *new(client.APIClient)
}

func (p *DockerProvider) Close() error { _ = "STUB: not implemented"; return nil }

// SetClient sets the docker client to be used by the provider
func (p *DockerProvider) SetClient(c client.APIClient) { _ = "STUB: not implemented"; return }

var _ ContainerProvider = (*DockerProvider)(nil)

// BuildImage will build and image from context and Dockerfile, then return the tag
func (p *DockerProvider) BuildImage(ctx context.Context, img ImageBuildInfo) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// release resources in any case

// Error is already wrapped.

// Always process the output, even if it is not printed
// to ensure that errors during the build process are
// correctly handled.

// the first tag is the one we want

// CreateContainer fulfils a request for a container without starting it
func (p *DockerProvider) CreateContainer(ctx context.Context, req ContainerRequest) (con Container, err error) {
	_ = "STUB: not implemented"
	// defer the close of the Docker client connection the soonest
	return *new(Container), nil
}

// If default network is not bridge make sure it is attached to the request
// as container won't be attached to it automatically
// in case of Podman the bridge network is called 'podman' as 'bridge' would conflict

// always append the hub substitutor after the user-defined ones

// If requested always attempt to pull image

// Add the labels that identify this as a testcontainers container and
// allow the reaper to terminate it if requested.

// default hooks include logger hook and pre-create hook

// in the case the container needs to access a local port
// we need to forward the local port to the container

// a container lifecycle hook will be added, which will expose the host ports to the container
// using a SSHD server running in a container. The SSHD server will be started and will
// forward the host ports to the container ports.

// Container setup failed so ensure we clean up the sshd container too.

// Combine with the original LifecycleHooks to avoid duplicate logging hooks.

// #248: If there is more than one network specified in the request attach newly created container to them one by one

// This should match the fields set in ContainerFromDockerResponse.

// No wrap as it would stutter.

// Wrapped so the returned error is passed to the cleanup function.

// Return the container to allow caller to clean up.

func (p *DockerProvider) findContainerByName(ctx context.Context, name string) (*container.Summary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note that, 'name' filter will use regex to find the containers

func (p *DockerProvider) waitContainerCreation(ctx context.Context, name string) (*container.Summary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *DockerProvider) ReuseOrCreateContainer(ctx context.Context, req ContainerRequest) (con Container, err error) {
	_ = "STUB: not implemented"
	return *new(Container), nil
}

// Cleanup on error.

// default hooks include logger hook and pre-create hook

// Workaround for https://github.com/moby/moby/issues/50133.
// /containers/{id}/json API endpoint of Docker Engine takes data about container from master (not replica) database
// which is synchronized with container state after call of /containers/{id}/stop API endpoint.

// If a container was stopped programmatically, we want to ensure the container
// is running again, but only if it is not paused, as it's not possible to start
// a paused container. The Docker Engine returns the "cannot start a paused container,
// try unpause instead" error.

// cannot re-start a running container, but we still need
// to call the startup hooks.

// TODO: we should unpause the container here.

// attemptToPullImage tries to pull the image while respecting the ctx cancellations.
// Besides, if the image cannot be pulled due to ErrorNotFound then no need to retry but terminate immediately.
func (p *DockerProvider) attemptToPullImage(ctx context.Context, tag string, pullOpt client.ImagePullOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// see https://github.com/docker/docs/blob/e8e1204f914767128814dca0ea008644709c117f/engine/api/sdk/examples.md?plain=1#L649-L657

// download of docker image finishes at EOF of the pull request

// Health measure the healthiness of the provider. Right now we leverage the
// docker-client Info endpoint to see if the daemon is reachable.
func (p *DockerProvider) Health(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// RunContainer takes a RequestContainer as input and it runs a container via the docker sdk
func (p *DockerProvider) RunContainer(ctx context.Context, req ContainerRequest) (Container, error) {
	_ = "STUB: not implemented"
	return *new(Container), nil
}

// Config provides the TestcontainersConfig read from $HOME/.testcontainers.properties or
// the environment variables
func (p *DockerProvider) Config() TestcontainersConfig {
	_ = "STUB: not implemented"
	return *new(TestcontainersConfig)
}

// DaemonHost gets the host or ip of the Docker daemon where ports are exposed on
// Warning: this is based on your Docker host setting. Will fail if using an SSH tunnel
// You can use the "TESTCONTAINERS_HOST_OVERRIDE" env variable to set this yourself
func (p *DockerProvider) DaemonHost(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *DockerProvider) daemonHostLocked(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// infer from Docker host

// Deprecated: use network.New instead
// CreateNetwork returns the object representing a new network identified by its name
func (p *DockerProvider) CreateNetwork(ctx context.Context, req NetworkRequest) (net Network, err error) {
	_ = "STUB: not implemented"
	// defer the close of the Docker client connection the soonest
	return *new(Network), nil
}

// Cleanup on error.

// add the labels that the reaper will use to terminate the network to the request

// GetNetwork returns the object representing the network identified by its name
func (p *DockerProvider) GetNetwork(ctx context.Context, req NetworkRequest) (network.Inspect, error) {
	_ = "STUB: not implemented"
	return *new(network.Inspect), nil
}

func (p *DockerProvider) GetGatewayIP(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	// Use a default network as defined in the DockerProvider
	return "", nil
}

func (p *DockerProvider) getGatewayIP(ctx context.Context, defaultNetwork string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ensureDefaultNetwork ensures that defaultNetwork is set and creates
// it if it does not exist, returning its value.
// It is safe to call this method concurrently.
func (p *DockerProvider) ensureDefaultNetwork(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *DockerProvider) ensureDefaultNetworkLocked(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "",

		// Already set.
		nil
}

// TODO: remove once we have docker context support via #2810
// Prefer the default bridge network if it exists.
// This makes the results stable as network list order is not guaranteed.

// Create a bridge network for the container communications.

// If the network already exists, we can ignore the error as that can
// happen if we are running multiple tests in parallel and we only
// need to ensure that the network exists.

// ContainerFromType builds a Docker container struct from the response of the Docker API
func (p *DockerProvider) ContainerFromType(ctx context.Context, response container.Summary) (ctr *DockerContainer, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This should match the fields set in CreateContainer.

// Wrapped so the returned error is passed to the cleanup function.

// populate the raw representation of the container

// Return the container to allow caller to clean up.

// the health status of the container, if any

// ListImages list images from the provider. If an image has multiple Tags, each tag is reported
// individually with the same ID and same labels
func (p *DockerProvider) ListImages(ctx context.Context) ([]ImageInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SaveImages exports a list of images as an uncompressed tar
func (p *DockerProvider) SaveImages(ctx context.Context, output string, images ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// SaveImagesWithOpts exports a list of images as an uncompressed tar, passing options to the provider
func (p *DockerProvider) SaveImagesWithOpts(ctx context.Context, output string, images []string, opts ...SaveImageOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Attempt optimized readFrom, implemented in linux

func SaveDockerImageWithPlatforms(platforms ...specs.Platform) SaveImageOption {
	_ = "STUB: not implemented"
	return *new(SaveImageOption)
}

// PullImage pulls image from registry
func (p *DockerProvider) PullImage(ctx context.Context, img string) error {
	_ = "STUB: not implemented"
	return nil
}

var permanentClientErrors = []func(error) bool{
	errdefs.IsNotFound,
	errdefs.IsInvalidArgument,
	errdefs.IsUnauthorized,
	errdefs.IsPermissionDenied,
	errdefs.IsNotImplemented,
	errdefs.IsInternal,
}

func isPermanentClientError(err error) bool { _ = "STUB: not implemented"; return false }

func tryClose(r io.Reader) { _ = "STUB: not implemented"; return }
