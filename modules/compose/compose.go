package compose

import (
	"context"
	"errors"
	"io"

	"github.com/compose-spec/compose-go/v2/types"
	"github.com/docker/compose/v5/pkg/api"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/log"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	envProjectName = "COMPOSE_PROJECT_NAME"
	envComposeFile = "COMPOSE_FILE"
)

var ErrNoStackConfigured = errors.New("no stack files configured")

type composeStackOptions struct {
	Identifier     string
	Paths          []string
	temporaryPaths map[string]bool
	Logger         log.Logger
	Profiles       []string
}

type ComposeStackOption interface {
	applyToComposeStack(o *composeStackOptions) error
}

type stackUpOptions struct {
	// Services defines the services user interacts with
	Services []string
	// Remove legacy containers for services that are not defined in the project
	RemoveOrphans bool
	// Wait won't return until containers reached the running|healthy state
	Wait bool
	// Recreate define the strategy to apply on existing containers
	Recreate string
	// RecreateDependencies define the strategy to apply on dependencies services
	RecreateDependencies string
	// Project is the compose project used to define this app. Might be nil if user ran command just with project name
	Project *types.Project
}

type StackUpOption interface {
	applyToStackUp(o *stackUpOptions)
}

type stackDownOptions struct {
	api.DownOptions
}

type StackDownOption interface {
	applyToStackDown(do *stackDownOptions)
}

// ComposeStack defines operations that can be applied to a parsed compose stack
type ComposeStack interface {
	Up(ctx context.Context, opts ...StackUpOption) error
	Down(ctx context.Context, opts ...StackDownOption) error
	Services() []string
	WaitForService(s string, strategy wait.Strategy) ComposeStack
	WithEnv(m map[string]string) ComposeStack
	WithOsEnv() ComposeStack
	ServiceContainer(ctx context.Context, svcName string) (*testcontainers.DockerContainer, error)
}

// Deprecated: DockerComposer is the old shell escape based API
// use ComposeStack instead
// DockerComposer defines the contract for running Docker Compose
type DockerComposer interface {
	Down() ExecError
	Invoke() ExecError
	WaitForService(string, wait.Strategy) DockerComposer
	WithCommand([]string) DockerComposer
	WithEnv(map[string]string) DockerComposer
	WithExposedService(string, int, wait.Strategy) DockerComposer
}

type waitService struct {
	service       string
	publishedPort int
}

// WithRecreate defines the strategy to apply on existing containers. If any other value than
// api.RecreateNever, api.RecreateForce or api.RecreateDiverged is provided, the default value
// api.RecreateForce will be used.
func WithRecreate(recreate string) StackUpOption {
	_ = "STUB: not implemented"
	return *new(StackUpOption)
}

// WithRecreateDependencies defines the strategy to apply on container dependencies. If any other value than
// api.RecreateNever, api.RecreateForce or api.RecreateDiverged is provided, the default value
// api.RecreateForce will be used.
func WithRecreateDependencies(recreate string) StackUpOption {
	_ = "STUB: not implemented"
	return *new(StackUpOption)
}

func WithStackFiles(filePaths ...string) ComposeStackOption {
	_ = "STUB: not implemented"
	return *new(ComposeStackOption)
}

// WithStackReaders supports reading the compose file/s from a reader.
func WithStackReaders(readers ...io.Reader) ComposeStackOption {
	_ = "STUB: not implemented"
	return *new(ComposeStackOption)
}

// WithProfiles allows to enable/disable services based on the profiles defined in the compose file.
func WithProfiles(profiles ...string) ComposeStackOption {
	_ = "STUB: not implemented"
	return *new(ComposeStackOption)
}

func NewDockerCompose(filePaths ...string) (*DockerCompose, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewDockerComposeWith(opts ...ComposeStackOption) (*DockerCompose, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create Docker CLI for compose service (uses moby/moby client internally)

// Create a separate testcontainers Docker client for provider and direct API calls.
// Compose v5 uses moby/moby/client internally, which is not type-compatible with
// docker/docker/client used by testcontainers, so we cannot share the CLI client.

// Deprecated: NewLocalDockerCompose returns a DockerComposer compatible instance which is superseded
// by ComposeStack use NewDockerCompose instead to get a ComposeStack compatible instance
//
// NewLocalDockerCompose returns an instance of the local Docker Compose, using an
// array of Docker Compose file paths and an identifier for the Compose execution.
//
// It will iterate through the array adding '-f compose-file-path' flags to the local
// Docker Compose execution. The identifier represents the name of the execution,
// which will define the name of the underlying Docker network and the name of the
// running Compose services.
func NewLocalDockerCompose(filePaths []string, identifier string, opts ...LocalDockerComposeOption) *LocalDockerCompose {
	_ = "STUB: not implemented"
	return nil
}
