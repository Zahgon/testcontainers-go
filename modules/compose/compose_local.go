package compose

import (
	"bytes"
	"io"

	"github.com/testcontainers/testcontainers-go/log"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	_ ComposeVersion = (*composeVersion1)(nil)
	_ ComposeVersion = (*composeVersion2)(nil)
)

type ComposeVersion interface {
	Format(parts ...string) string
}

type composeVersion1 struct{}

func (c composeVersion1) Format(parts ...string) string { _ = "STUB: not implemented"; return "" }

type composeVersion2 struct{}

func (c composeVersion2) Format(parts ...string) string { _ = "STUB: not implemented"; return "" }

// Deprecated: use ComposeStack instead
// LocalDockerCompose represents a Docker Compose execution using local binary
// docker compose or docker.exe compose, depending on the underlying platform
type LocalDockerCompose struct {
	ComposeVersion
	*LocalDockerComposeOptions
	Executable           string
	composeSubcommand    string
	ComposeFilePaths     []string
	absComposeFilePaths  []string
	Identifier           string
	Cmd                  []string
	Env                  map[string]string
	Services             map[string]any
	waitStrategySupplied bool
	WaitStrategyMap      map[waitService]wait.Strategy
}

type (
	// Deprecated: it will be removed in the next major release
	// LocalDockerComposeOptions defines options applicable to LocalDockerCompose
	LocalDockerComposeOptions struct {
		Logger log.Logger
	}

	// Deprecated: it will be removed in the next major release
	// LocalDockerComposeOption defines a common interface to modify LocalDockerComposeOptions
	// These options can be passed to NewLocalDockerCompose in a variadic way to customize the returned LocalDockerCompose instance
	LocalDockerComposeOption interface {
		ApplyToLocalCompose(opts *LocalDockerComposeOptions)
	}

	// Deprecated: it will be removed in the next major release
	// LocalDockerComposeOptionsFunc is a shorthand to implement the LocalDockerComposeOption interface
	LocalDockerComposeOptionsFunc func(opts *LocalDockerComposeOptions)
)

type ComposeLoggerOption struct {
	logger log.Logger
}

// WithLogger is a generic option that implements LocalDockerComposeOption
// It replaces the global Logging implementation with a user defined one e.g. to aggregate logs from testcontainers
// with the logs of specific test case
func WithLogger(logger log.Logger) ComposeLoggerOption {
	_ = "STUB: not implemented"
	return *new(ComposeLoggerOption)
}

// Deprecated: it will be removed in the next major release
func (o ComposeLoggerOption) ApplyToLocalCompose(opts *LocalDockerComposeOptions) {
	_ = "STUB: not implemented"
	return
}

func (o ComposeLoggerOption) applyToComposeStack(opts *composeStackOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: it will be removed in the next major release
func (f LocalDockerComposeOptionsFunc) ApplyToLocalCompose(opts *LocalDockerComposeOptions) {
	_ = "STUB: not implemented"

	// Deprecated: it will be removed in the next major release
	// Down executes docker compose down
	return
}

func (dc *LocalDockerCompose) Down() ExecError { _ = "STUB: not implemented"; return *new(ExecError) }

// Deprecated: it will be removed in the next major release
func (dc *LocalDockerCompose) getDockerComposeEnvironment() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: it will be removed in the next major release
func (dc *LocalDockerCompose) containerNameFromServiceName(service, separator string) string {
	_ = "STUB: not implemented"
	return ""
}

// Deprecated: it will be removed in the next major release
func (dc *LocalDockerCompose) applyStrategyToRunningContainer() error {
	_ = "STUB: not implemented"
	return nil
}

// The length should always be a list of 1, since we are matching one service name at a time

// Deprecated: it will be removed in the next major release
// Invoke invokes the docker compose
func (dc *LocalDockerCompose) Invoke() ExecError { _ = "STUB: not implemented"; return *new(ExecError) }

// Deprecated: it will be removed in the next major release
// WaitForService sets the strategy for the service that is to be waited on
func (dc *LocalDockerCompose) WaitForService(service string, strategy wait.Strategy) DockerComposer {
	_ = "STUB: not implemented"
	return *new(DockerComposer)
}

// Deprecated: it will be removed in the next major release
// WithCommand assigns the command
func (dc *LocalDockerCompose) WithCommand(cmd []string) DockerComposer {
	_ = "STUB: not implemented"
	return *

	// Deprecated: it will be removed in the next major release
	// WithEnv assigns the environment
	new(DockerComposer)
}

func (dc *LocalDockerCompose) WithEnv(env map[string]string) DockerComposer {
	_ = "STUB: not implemented"
	return *

	// Deprecated: it will be removed in the next major release
	// WithExposedService sets the strategy for the service that is to be waited on. If multiple strategies
	// are given for a single service running on different ports, both strategies will be applied on the same container
	new(DockerComposer)
}

func (dc *LocalDockerCompose) WithExposedService(service string, port int, strategy wait.Strategy) DockerComposer {
	_ = "STUB: not implemented"
	return *new(DockerComposer)
}

// Deprecated: it will be removed in the next major release
// determineVersion checks which version of docker compose is installed
// depending on the version services names are composed in a different way
func (dc *LocalDockerCompose) determineVersion() error { _ = "STUB: not implemented"; return nil }

// Deprecated: it will be removed in the next major release
// validate checks if the files to be run in the compose are valid YAML files, setting up
// references to all services in them
func (dc *LocalDockerCompose) validate() error { _ = "STUB: not implemented"; return nil }

// ExecError is super struct that holds any information about an execution error, so the client code
// can handle the result
type ExecError struct {
	Command      []string
	StdoutOutput []byte
	StderrOutput []byte
	Error        error
	Stdout       error
	Stderr       error
}

// execute executes a program with arguments and environment variables inside a specific directory
func execute(
	dirContext string, environment map[string]string, binary string, args []string,
) ExecError {
	_ = "STUB: not implemented"
	return *new(ExecError)
}

// add information about the CMD and arguments used

// Deprecated: it will be removed in the next major release
func executeCompose(dc *LocalDockerCompose, args []string) ExecError {
	_ = "STUB: not implemented"
	return *new(ExecError)
}

// initialise the command with the compose subcommand

// If the wait strategy has been executed once for all services during startup , disable it so that it is not invoked while tearing down

// capturingPassThroughWriter is a writer that remembers
// data written to it and passes it to w
type capturingPassThroughWriter struct {
	buf bytes.Buffer
	w   io.Writer
}

// newCapturingPassThroughWriter creates new capturingPassThroughWriter
func newCapturingPassThroughWriter(w io.Writer) *capturingPassThroughWriter {
	_ = "STUB: not implemented"
	return nil
}

func (w *capturingPassThroughWriter) Write(d []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Bytes returns bytes written to the writer
func (w *capturingPassThroughWriter) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// Which checks if a binary is present in PATH
func which(binary string) error { _ = "STUB: not implemented"; return nil }
