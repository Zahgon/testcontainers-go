package compose

import (
	"context"
	"io"
	"sync"

	"github.com/compose-spec/compose-go/v2/cli"
	"github.com/compose-spec/compose-go/v2/types"
	"github.com/docker/cli/cli/command"
	"github.com/docker/compose/v5/pkg/api"
	"github.com/moby/moby/client"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/log"
	"github.com/testcontainers/testcontainers-go/wait"
)

type stackUpOptionFunc func(s *stackUpOptions)

func (f stackUpOptionFunc) applyToStackUp(o *stackUpOptions) {
	_ = "STUB: not implemented"

	// RunServices is comparable to 'docker compose run' as it only creates a subset of containers
	// instead of all services defined by the project
	return
}

func RunServices(serviceNames ...string) StackUpOption {
	_ = "STUB: not implemented"
	return *new(StackUpOption)
}

// Deprecated: will be removed in the next major release
// IgnoreOrphans - Ignore legacy containers for services that are not defined in the project
type IgnoreOrphans bool

// Deprecated: will be removed in the next major release
//
//nolint:unused
func (io IgnoreOrphans) applyToStackUp(co *api.CreateOptions, _ *api.StartOptions) {
	_ = "STUB: not implemented"
	return
}

// Recreate will recreate the containers that are already running
type Recreate string

func (r Recreate) applyToStackUp(o *stackUpOptions) { _ = "STUB: not implemented"; return }

// RecreateDependencies will recreate the dependencies of the services that are already running
type RecreateDependencies string

func (r RecreateDependencies) applyToStackUp(o *stackUpOptions) { _ = "STUB: not implemented"; return }

func validateRecreate(r string) string { _ = "STUB: not implemented"; return "" }

// RemoveOrphans will clean up containers that are not declared on the compose model but own the same labels
type RemoveOrphans bool

func (ro RemoveOrphans) applyToStackUp(o *stackUpOptions) { _ = "STUB: not implemented"; return }

func (ro RemoveOrphans) applyToStackDown(o *stackDownOptions) { _ = "STUB: not implemented"; return }

// Wait won't return until containers reached the running|healthy state
type Wait bool

func (w Wait) applyToStackUp(o *stackUpOptions) { _ = "STUB: not implemented"; return }

type RemoveVolumes bool

func (ro RemoveVolumes) applyToStackDown(o *stackDownOptions) { _ = "STUB: not implemented"; return }

// RemoveImages used by services
type RemoveImages uint8

func (ri RemoveImages) applyToStackDown(o *stackDownOptions) { _ = "STUB: not implemented"; return }

type ComposeStackReaders []io.Reader

func (r ComposeStackReaders) applyToComposeStack(o *composeStackOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// mark the file for removal as it was generated on the fly

type ComposeStackFiles []string

func (f ComposeStackFiles) applyToComposeStack(o *composeStackOptions) error {
	_ = "STUB: not implemented"
	return nil
}

type ComposeProfiles []string

func (p ComposeProfiles) applyToComposeStack(o *composeStackOptions) error {
	_ = "STUB: not implemented"
	return nil
}

type StackIdentifier string

func (f StackIdentifier) applyToComposeStack(o *composeStackOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f StackIdentifier) String() string { _ = "STUB: not implemented"; return "" }

const (
	// RemoveImagesAll - remove all images used by the stack
	RemoveImagesAll RemoveImages = iota
	// RemoveImagesLocal - remove only images that don't have a tag
	RemoveImagesLocal
)

type DockerCompose struct {
	// used to synchronize operations
	lock sync.RWMutex

	// dockerCli is the Docker CLI instance used internally by the compose service.
	// It is stored here so its HTTP transport connections can be closed after Down().
	dockerCli *command.DockerCli

	// name/identifier of the stack that will be started
	// by default a UUID will be used
	name string

	// paths to stack files that will be considered when compiling the final compose project
	configs []string

	// used to remove temporary files that were generated on the fly
	temporaryConfigs map[string]bool

	// used to set logger in DockerContainer
	logger log.Logger

	// wait strategies that are applied per service when starting the stack
	// only one strategy can be added to a service, to use multiple use wait.ForAll(...)
	waitStrategies map[string]wait.Strategy

	// Used to synchronise writes to the containers.
	containersLock sync.Mutex

	// cache for containers that are part of the stack
	// used in ServiceContainer(...) function to avoid calls to the Docker API
	containers map[string]*testcontainers.DockerContainer

	// cache for networks in the compose stack
	networks map[string]*testcontainers.DockerNetwork

	// docker/compose API service instance used to control the compose stack
	composeService api.Compose

	// Docker API client used to interact with single container instances and the Docker API e.g. to list containers
	dockerClient client.APIClient

	// options used to compile the compose project
	// e.g. environment settings, ...
	projectOptions []cli.ProjectOptionsFn

	// profiles applied to the compose project after compilation.
	projectProfiles []string

	// compiled compose project
	// can be nil if the stack wasn't started yet
	project *types.Project

	// sessionID is used to identify the reaper session
	sessionID string

	// provider is used to docker operations.
	provider *testcontainers.DockerProvider
}

func (d *DockerCompose) ServiceContainer(ctx context.Context, svcName string) (*testcontainers.DockerContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DockerCompose) Services() []string { _ = "STUB: not implemented"; return nil }

func (d *DockerCompose) Down(ctx context.Context, opts ...StackDownOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Close releases the HTTP transport connections held by the internal Docker CLI
// and the testcontainers Docker client, preventing net/http persistConn goroutine
// leaks. Call Close after Down when the compose stack will no longer be used.
func (d *DockerCompose) Close() error { _ = "STUB: not implemented"; return nil }

func (d *DockerCompose) Up(ctx context.Context, opts ...StackUpOption) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// NewReaper is deprecated: we need to find a way to create the reaper for compose
// bypassing the deprecation.
//nolint:staticcheck // intentional use of deprecated API for compose

// Cleanup on error, otherwise set termSignal to nil before successful return.

// Need to call Connect at least once to ensure the initial
// connection is cleaned up.

// No need to cleanup.

// Connect to the reaper and set the termination signal for each network.

// Lookup the containers for each service and connect them
// to the reaper if needed.

// wait here for the containers lookup to finish

// pinning the variables

func (d *DockerCompose) WaitForService(s string, strategy wait.Strategy) ComposeStack {
	_ = "STUB: not implemented"
	return *new(ComposeStack)
}

func (d *DockerCompose) WithEnv(m map[string]string) ComposeStack {
	_ = "STUB: not implemented"
	return *new(ComposeStack)
}

func (d *DockerCompose) WithOsEnv() ComposeStack {
	_ = "STUB: not implemented"
	return *new(ComposeStack)
}

// cachedContainer returns the cached container for svcName or nil if it doesn't exist.
func (d *DockerCompose) cachedContainer(svcName string) *testcontainers.DockerContainer {
	_ = "STUB: not implemented"
	return nil
}

// lookupContainer is used to retrieve the container instance from the cache or the Docker API.
//
// Safe for concurrent calls.
func (d *DockerCompose) lookupContainer(ctx context.Context, svcName string) (*testcontainers.DockerContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// lookupNetworks is used to retrieve the networks that are part of the compose stack.
//
// Safe for concurrent calls.
func (d *DockerCompose) lookupNetworks(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DockerCompose) compileProject(ctx context.Context) (*types.Project, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// default, will be overridden by `run` command

// add a label for each env file, indexed by its position

func withEnv(env map[string]string) func(*cli.ProjectOptions) error {
	_ = "STUB: not implemented"
	return nil
}
