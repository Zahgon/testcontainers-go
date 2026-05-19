package ollama

import (
	"context"
	"io"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/api/types/network"
	"github.com/moby/moby/client"

	"github.com/testcontainers/testcontainers-go"
	tcexec "github.com/testcontainers/testcontainers-go/exec"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	localPort       = uint16(11434)
	localBinary     = "ollama"
	localServeArg   = "serve"
	localLogRegex   = `Listening on (.*:\d+) \(version\s(.*)\)`
	localNamePrefix = "local-ollama"
	localHostVar    = "OLLAMA_HOST"
	localLogVar     = "OLLAMA_LOGFILE"
)

var (
	// Ensure localProcess implements the required interfaces.
	_ testcontainers.Container           = (*localProcess)(nil)
	_ testcontainers.ContainerCustomizer = (*localProcess)(nil)

	// zeroTime is the zero time value.
	zeroTime time.Time
)

// stdWriter wraps an io.Writer to produce Docker-multiplexed output frames.
// This replaces stdcopy.NewStdWriter which is no longer exported from moby.
// The frame format is: 1 byte stream type + 3 bytes padding + 4 bytes big-endian length + payload.
type stdWriter struct {
	w      io.Writer
	prefix byte // 1 = stdout, 2 = stderr
}

func (s *stdWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// localProcess emulates the Ollama container using a local process to improve performance.
type localProcess struct {
	sessionID string

	// env is the combined environment variables passed to the Ollama binary.
	env []string

	// cmd is the command that runs the Ollama binary, not valid externally if nil.
	cmd *exec.Cmd

	// logName and logFile are the file where the Ollama logs are written.
	logName string
	logFile *os.File

	// host, port and version are extracted from log on startup.
	host    string
	port    uint16
	version string

	// waitFor is the strategy to wait for the process to be ready.
	waitFor wait.Strategy

	// done is closed when the process is finished.
	done chan struct{}

	// wg is used to wait for the process to finish.
	wg sync.WaitGroup

	// startedAt is the time when the process started.
	startedAt time.Time

	// mtx is used to synchronize access to the process state fields below.
	mtx sync.Mutex

	// finishedAt is the time when the process finished.
	finishedAt time.Time

	// exitErr is the error returned by the process.
	exitErr error

	// binary is the name of the Ollama binary.
	binary string
}

// run returns an OllamaContainer that uses the local Ollama binary instead of using a Docker container.
func (c *localProcess) run(ctx context.Context, img string, opts []testcontainers.ContainerCustomizer) (*OllamaContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Apply the updated details from the request.

// validateRequest checks that req is valid for the local Ollama binary.
func (c *localProcess) validateRequest(req testcontainers.GenericContainerRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate the image and extract the binary name.
// The image must be in the format "[<path>/]<binary>[:latest]".

// Check if the version is "latest" or not specified.

// Trim the path if present.

// Reset fields we support to their zero values.

// We don't need the logger.

// Only check the leaf fields.

// Start implements testcontainers.Container interface for the local Ollama binary.
func (c *localProcess) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Multiplex stdout and stderr to the log file matching the Docker API.

// Run the ollama serve command in background.

// Past this point, the process was started successfully.

// Reset the details to allow multiple start / stop cycles.

// Wait for the process to finish in a goroutine.

// waitStrategy waits until the Ollama process is ready.
func (c *localProcess) waitStrategy(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// extractLogDetails extracts the listening address and version from the log.
func (c *localProcess) extractLogDetails(pattern string, submatches [][][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Set OLLAMA_HOST variable to the extracted host so Exec can use it.

// Return the last error encountered.

// ContainerIP implements testcontainers.Container interface for the local Ollama binary.
func (c *localProcess) ContainerIP(_ context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "",

		// ContainerIPs returns a slice with the IP address of the local Ollama binary.
		nil
}

func (c *localProcess) ContainerIPs(_ context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// CopyToContainer implements testcontainers.Container interface for the local Ollama binary.
	// Returns [errors.ErrUnsupported].
}

func (c *localProcess) CopyToContainer(_ context.Context, _ []byte, _ string, _ int64) error {
	_ = "STUB: not implemented"
	return nil
}

// CopyDirToContainer implements testcontainers.Container interface for the local Ollama binary.
// Returns [errors.ErrUnsupported].
func (c *localProcess) CopyDirToContainer(_ context.Context, _ string, _ string, _ int64) error {
	_ = "STUB: not implemented"
	return nil
}

// CopyFileToContainer implements testcontainers.Container interface for the local Ollama binary.
// Returns [errors.ErrUnsupported].
func (c *localProcess) CopyFileToContainer(_ context.Context, _ string, _ string, _ int64) error {
	_ = "STUB: not implemented"
	return nil
}

// CopyFileFromContainer implements testcontainers.Container interface for the local Ollama binary.
// Returns [errors.ErrUnsupported].
func (c *localProcess) CopyFileFromContainer(_ context.Context, _ string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// GetLogProductionErrorChannel implements testcontainers.Container interface for the local Ollama binary.
// It returns a nil channel because the local Ollama binary doesn't have a production error channel.
func (c *localProcess) GetLogProductionErrorChannel() <-chan error {
	_ = "STUB: not implemented"

	// Exec implements testcontainers.Container interface for the local Ollama binary.
	// It executes a command using the local Ollama binary and returns the exit status
	// of the executed command, an [io.Reader] containing the combined stdout and stderr,
	// and any encountered error.
	//
	// Reading directly from the [io.Reader] may result in unexpected bytes due to custom
	// stream multiplexing headers. Use [tcexec.Multiplexed] option to read the combined output
	// without the multiplexing headers.
	// Alternatively, to separate the stdout and stderr from [io.Reader] and interpret these
	// headers properly, [stdcopy.StdCopy] from the Docker API should be used.
	return nil
}

func (c *localProcess) Exec(ctx context.Context, cmd []string, options ...tcexec.ProcessOption) (int, io.Reader, error) {
	_ = "STUB: not implemented"
	return 0, *new(io.Reader), nil
}

// Multiplex stdout and stderr to the buffer so they can be read separately later.

// Use process options to customize the command execution
// emulating the Docker API behaviour.

// validateExecOptions checks if the given exec options are supported by the local Ollama binary.
func (c *localProcess) validateExecOptions(options client.ExecCreateOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Inspect implements testcontainers.Container interface for the local Ollama binary.
// It returns a ContainerJSON with the state of the local Ollama binary.
func (c *localProcess) Inspect(ctx context.Context) (*container.InspectResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IsRunning implements testcontainers.Container interface for the local Ollama binary.
// It returns true if the local Ollama process is running, false otherwise.
func (c *localProcess) IsRunning() bool { _ = "STUB: not implemented"; return false }

// The process hasn't started yet.

// The process exited.

// The process is still running.

// Logs implements testcontainers.Container interface for the local Ollama binary.
// It returns the logs from the local Ollama binary.
func (c *localProcess) Logs(_ context.Context) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// State implements testcontainers.Container interface for the local Ollama binary.
// It returns the current state of the Ollama process, simulating a container state.
func (c *localProcess) State(_ context.Context) (*container.State, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Setting the Running field because it's required by the wait strategy
// to check if the given log message is present.

// Stop implements testcontainers.Container interface for the local Ollama binary.
// It gracefully stops the local Ollama process.
func (c *localProcess) Stop(ctx context.Context, d *time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// The process exited.

// Terminate implements testcontainers.Container interface for the local Ollama binary.
// It stops the local Ollama process, removing the log file.
func (c *localProcess) Terminate(ctx context.Context, opts ...testcontainers.TerminateOption) error {
	_ = "STUB: not implemented"
	return nil
}

// First try to stop gracefully.

// Still running, force kill.

// Best effort so we can continue with the cleanup.

// Wait for the process to exit so we can capture any error.

// cleanup performs all clean up, closing and removing the log file if set.
func (c *localProcess) cleanup() error { _ = "STUB: not implemented"; return nil }

// Prevent double cleanup.

// Endpoint implements testcontainers.Container interface for the local Ollama binary.
// It returns proto://host:port string for the Ollama port.
// It returns just host:port if proto is blank.
func (c *localProcess) Endpoint(ctx context.Context, proto string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetContainerID implements testcontainers.Container interface for the local Ollama binary.
func (c *localProcess) GetContainerID() string { _ = "STUB: not implemented"; return "" }

// Host implements testcontainers.Container interface for the local Ollama binary.
func (c *localProcess) Host(_ context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "",

		// MappedPort implements testcontainers.Container interface for the local Ollama binary.
		nil
}

func (c *localProcess) MappedPort(_ context.Context, port string) (network.Port, error) {
	_ = "STUB: not implemented"
	return *new(network.Port), nil
}

// Networks implements testcontainers.Container interface for the local Ollama binary.
// It returns a nil slice.
func (c *localProcess) Networks(_ context.Context) ([]string, error) {
	_ = "STUB: not implemented"

	// NetworkAliases implements testcontainers.Container interface for the local Ollama binary.
	// It returns a nil map.
	return nil, nil
}

func (c *localProcess) NetworkAliases(_ context.Context) (map[string][]string, error) {
	_ = "STUB: not implemented"

	// PortEndpoint implements testcontainers.Container interface for the local Ollama binary.
	// It returns proto://host:port or proto://[IPv6host]:port string for the given exposed port.
	// It returns just host:port or [IPv6host]:port if proto is blank.
	return nil, nil
}

func (c *localProcess) PortEndpoint(ctx context.Context, port string, proto string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// SessionID implements testcontainers.Container interface for the local Ollama binary.
func (c *localProcess) SessionID() string {
	_ = "STUB: not implemented"

	// Deprecated: it will be removed in the next major release.
	// FollowOutput is not implemented for the local Ollama binary.
	// It panics if called.
	return ""
}

func (c *localProcess) FollowOutput(_ testcontainers.LogConsumer) {
	_ = "STUB: not implemented"
	return
}

// Deprecated: use c.Inspect(ctx).NetworkSettings.Ports instead.
// Ports gets the exposed ports for the container.
func (c *localProcess) Ports(ctx context.Context) (network.PortMap, error) {
	_ = "STUB: not implemented"
	return *new(network.PortMap), nil
}

// Deprecated: it will be removed in the next major release.
// StartLogProducer implements testcontainers.Container interface for the local Ollama binary.
// It returns an error because the local Ollama binary doesn't have a log producer.
func (c *localProcess) StartLogProducer(context.Context, ...testcontainers.LogProductionOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: it will be removed in the next major release.
// StopLogProducer implements testcontainers.Container interface for the local Ollama binary.
// It returns an error because the local Ollama binary doesn't have a log producer.
func (c *localProcess) StopLogProducer() error { _ = "STUB: not implemented"; return nil }

// Deprecated: Use c.Inspect(ctx).Name instead.
// Name returns the name for the local Ollama binary.
func (c *localProcess) Name(context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Customize implements the [testcontainers.ContainerCustomizer] interface.
// It configures the environment variables set by [WithUseLocal] and sets up
// the wait strategy to extract the host, port and version from the log.
func (c *localProcess) Customize(req *testcontainers.GenericContainerRequest) error {
	_ = "STUB: not implemented"
	// Replace the default host port strategy with one that waits for a log entry
	// and extracts the host, port and version from it.
	return nil
}

// Setup the environment variables using a random port by default
// to avoid conflicts.

// isCleanupSafe reports whether all errors in err's tree are one of the
// following, so can safely be ignored:
//   - nil
//   - os: process already finished
//   - context deadline exceeded
func (c *localProcess) isCleanupSafe(err error) bool { _ = "STUB: not implemented"; return false }
