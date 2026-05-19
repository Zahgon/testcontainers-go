package testcontainers

import (
	"time"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/api/types/network"

	tcexec "github.com/testcontainers/testcontainers-go/exec"
	"github.com/testcontainers/testcontainers-go/wait"
)

// ContainerCustomizer is an interface that can be used to configure the Testcontainers container
// request. The passed request will be merged with the default one.
type ContainerCustomizer interface {
	Customize(req *GenericContainerRequest) error
}

// CustomizeRequestOption is a type that can be used to configure the Testcontainers container request.
// The passed request will be merged with the default one.
type CustomizeRequestOption func(req *GenericContainerRequest) error

func (opt CustomizeRequestOption) Customize(req *GenericContainerRequest) error {
	_ = "STUB: not implemented"

	// CustomizeRequest returns a function that can be used to merge the passed container request with the one that is used by the container.
	// Slices and Maps will be appended.
	return nil
}

func CustomizeRequest(src GenericContainerRequest) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithDockerfile allows to build a container from a Dockerfile
func WithDockerfile(df FromDockerfile) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithConfigModifier allows to override the default container config
func WithConfigModifier(modifier func(config *container.Config)) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithEndpointSettingsModifier allows to override the default endpoint settings
func WithEndpointSettingsModifier(modifier func(settings map[string]*network.EndpointSettings)) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithEnv sets the environment variables for a container.
// If the environment variable already exists, it will be overridden.
func WithEnv(envs map[string]string) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithHostConfigModifier allows to override the default host config
func WithHostConfigModifier(modifier func(hostConfig *container.HostConfig)) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithHostPortAccess allows to expose the host ports to the container
func WithHostPortAccess(ports ...int) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithName will set the name of the container.
func WithName(containerName string) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithNoStart will prevent the container from being started after creation.
func WithNoStart() CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithReuseByName will mark a container to be reused if it exists or create a new one if it doesn't.
// A container name must be provided to identify the container to be reused.
func WithReuseByName(containerName string) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithImage sets the image for a container
func WithImage(image string) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// imageSubstitutor {

// ImageSubstitutor represents a way to substitute container image names
type ImageSubstitutor interface {
	// Description returns the name of the type and a short description of how it modifies the image.
	// Useful to be printed in logs
	Description() string
	Substitute(image string) (string, error)
}

// }

// CustomHubSubstitutor represents a way to substitute the hub of an image with a custom one,
// using provided value with respect to the HubImageNamePrefix configuration value.
type CustomHubSubstitutor struct {
	hub string
}

// NewCustomHubSubstitutor creates a new CustomHubSubstitutor
func NewCustomHubSubstitutor(hub string) CustomHubSubstitutor {
	_ = "STUB: not implemented"
	return *new(CustomHubSubstitutor)
}

// Description returns the name of the type and a short description of how it modifies the image.
func (c CustomHubSubstitutor) Description() string { _ = "STUB: not implemented"; return "" }

// Substitute replaces the hub of the image with the provided one, with certain conditions:
//   - if the hub is empty, the image is returned as is.
//   - if the image already contains a registry, the image is returned as is.
//   - if the HubImageNamePrefix configuration value is set, the image is returned as is.
func (c CustomHubSubstitutor) Substitute(image string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// prependHubRegistry represents a way to prepend a custom Hub registry to the image name,
// using the HubImageNamePrefix configuration value
type prependHubRegistry struct {
	prefix string
}

// newPrependHubRegistry creates a new prependHubRegistry
func newPrependHubRegistry(hubPrefix string) prependHubRegistry {
	_ = "STUB: not implemented"
	return *new(prependHubRegistry)
}

// Description returns the name of the type and a short description of how it modifies the image.
func (p prependHubRegistry) Description() string { _ = "STUB: not implemented"; return "" }

// Substitute prepends the Hub prefix to the image name, with certain conditions:
//   - if the prefix is empty, the image is returned as is.
//   - if the image is a non-hub image (e.g. where another registry is set), the image is returned as is.
//   - if the image is a Docker Hub image where the hub registry is explicitly part of the name
//     (i.e. anything with a registry.hub.docker.com host part), the image is returned as is.
func (p prependHubRegistry) Substitute(image string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// add the exclusions in the right order

// no prefix set at the configuration level
// non-hub image
// explicitly including docker.io
// explicitly including registry.hub.docker.com

// WithImageSubstitutors sets the image substitutors for a container
func WithImageSubstitutors(fn ...ImageSubstitutor) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithLogConsumers sets the log consumers for a container
func WithLogConsumers(consumer ...LogConsumer) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithLogConsumerConfig sets the log consumer config for a container.
// Beware that this option completely replaces the existing log consumer config,
// including the log consumers and the log production options,
// so it should be used with care.
func WithLogConsumerConfig(config *LogConsumerConfig) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// Executable represents an executable command to be sent to a container, including options,
// as part of the different lifecycle hooks.
type Executable interface {
	AsCommand() []string
	// Options can container two different types of options:
	// - Docker's ExecConfigs (WithUser, WithWorkingDir, WithEnv, etc.)
	// - testcontainers' ProcessOptions (i.e. Multiplexed response)
	Options() []tcexec.ProcessOption
}

// ExecOptions is a struct that provides a default implementation for the Options method
// of the Executable interface.
type ExecOptions struct {
	opts []tcexec.ProcessOption
}

func (ce ExecOptions) Options() []tcexec.ProcessOption {
	_ = "STUB: not implemented"

	// RawCommand is a type that implements Executable and represents a command to be sent to a container
	return nil
}

type RawCommand struct {
	ExecOptions
	cmds []string
}

func NewRawCommand(cmds []string, opts ...tcexec.ProcessOption) RawCommand {
	_ = "STUB: not implemented"
	return *new(RawCommand)
}

// AsCommand returns the command as a slice of strings
func (r RawCommand) AsCommand() []string {
	_ = "STUB: not implemented"

	// WithStartupCommand will execute the command representation of each Executable into the container.
	// It will leverage the container lifecycle hooks to call the command right after the container
	// is started.
	return nil
}

func WithStartupCommand(execs ...Executable) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithAfterReadyCommand will execute the command representation of each Executable into the container.
// It will leverage the container lifecycle hooks to call the command right after the container
// is ready.
func WithAfterReadyCommand(execs ...Executable) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithWaitStrategy replaces the wait strategy for a container, using 60 seconds as deadline
func WithWaitStrategy(strategies ...wait.Strategy) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithAdditionalWaitStrategy appends the wait strategy for a container, using 60 seconds as deadline
func WithAdditionalWaitStrategy(strategies ...wait.Strategy) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithWaitStrategyAndDeadline replaces the wait strategy for a container, including deadline
func WithWaitStrategyAndDeadline(deadline time.Duration, strategies ...wait.Strategy) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithAdditionalWaitStrategyAndDeadline appends the wait strategy for a container, including deadline
func WithAdditionalWaitStrategyAndDeadline(deadline time.Duration, strategies ...wait.Strategy) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithImageMount mounts an image to a container, passing the source image name,
// the relative subpath to mount in that image, and the mount point in the target container.
// This option validates that the subpath is a relative path, raising an error otherwise.
func WithImageMount(source string, subpath string, target ContainerMountTarget) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithAlwaysPull will pull the image before starting the container
func WithAlwaysPull() CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithImagePlatform sets the platform for a container
func WithImagePlatform(platform string) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithEntrypoint completely replaces the entrypoint of a container
func WithEntrypoint(entrypoint ...string) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithEntrypointArgs appends the entrypoint arguments to the entrypoint of a container
func WithEntrypointArgs(entrypointArgs ...string) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithExposedPorts appends the ports to the exposed ports for a container
func WithExposedPorts(ports ...string) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithCmd completely replaces the command for a container
func WithCmd(cmd ...string) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithCmdArgs appends the command arguments to the command for a container
func WithCmdArgs(cmdArgs ...string) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithLabels appends the labels to the labels for a container
func WithLabels(labels map[string]string) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithLifecycleHooks completely replaces the lifecycle hooks for a container
func WithLifecycleHooks(hooks ...ContainerLifecycleHooks) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithAdditionalLifecycleHooks appends lifecycle hooks to the existing ones for a container
func WithAdditionalLifecycleHooks(hooks ...ContainerLifecycleHooks) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithMounts appends the mounts to the mounts for a container
func WithMounts(mounts ...ContainerMount) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithTmpfs appends the tmpfs mounts to the tmpfs mounts for a container
func WithTmpfs(tmpfs map[string]string) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithFiles appends the files to the files for a container
func WithFiles(files ...ContainerFile) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}

// WithProvider sets the provider type for a container
func WithProvider(provider ProviderType) CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(CustomizeRequestOption)
}
