package testcontainers

import (
	"context"
	"io"
	"time"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/api/types/network"
	"github.com/moby/moby/api/types/registry"
	"github.com/moby/moby/client"

	tcexec "github.com/testcontainers/testcontainers-go/exec"
	"github.com/testcontainers/testcontainers-go/wait"
)

// Deprecated: Use [Container]
//
// DeprecatedContainer shows methods that were supported before, but are now deprecated
type DeprecatedContainer interface {
	GetHostEndpoint(ctx context.Context, port string) (string, string, error)
	GetIPAddress(ctx context.Context) (string, error)
	LivenessCheckPorts(ctx context.Context) (network.PortSet, error)
	Terminate(ctx context.Context) error
}

// Container allows getting info about and controlling a single container instance
type Container interface {
	GetContainerID() string                                                      // get the container id from the provider
	Endpoint(context.Context, string) (string, error)                            // get proto://ip:port string for the lowest exposed port
	PortEndpoint(ctx context.Context, port string, proto string) (string, error) // get proto://ip:port string for the given exposed port
	Host(context.Context) (string, error)                                        // get host where the container port is exposed
	Inspect(context.Context) (*container.InspectResponse, error)                 // get container info
	MappedPort(context.Context, string) (network.Port, error)                    // get externally mapped port for a container port
	Ports(context.Context) (network.PortMap, error)                              // Deprecated: Use c.Inspect(ctx).NetworkSettings.Ports instead
	SessionID() string                                                           // get session id
	IsRunning() bool                                                             // IsRunning returns true if the container is running, false otherwise.
	Start(context.Context) error                                                 // start the container
	Stop(context.Context, *time.Duration) error                                  // stop the container

	// Terminate stops and removes the container and its image if it was built and not flagged as kept.
	Terminate(ctx context.Context, opts ...TerminateOption) error

	Logs(context.Context) (io.ReadCloser, error)                    // Get logs of the container
	FollowOutput(LogConsumer)                                       // Deprecated: it will be removed in the next major release
	StartLogProducer(context.Context, ...LogProductionOption) error // Deprecated: Use the ContainerRequest instead
	StopLogProducer() error                                         // Deprecated: it will be removed in the next major release
	Name(context.Context) (string, error)                           // Deprecated: Use c.Inspect(ctx).Name instead
	State(context.Context) (*container.State, error)                // returns container's running state
	Networks(context.Context) ([]string, error)                     // get container networks
	NetworkAliases(context.Context) (map[string][]string, error)    // get container network aliases for a network
	Exec(ctx context.Context, cmd []string, options ...tcexec.ProcessOption) (int, io.Reader, error)
	ContainerIP(context.Context) (string, error)    // get container ip
	ContainerIPs(context.Context) ([]string, error) // get all container IPs
	CopyToContainer(ctx context.Context, fileContent []byte, containerFilePath string, fileMode int64) error
	CopyDirToContainer(ctx context.Context, hostDirPath string, containerParentPath string, fileMode int64) error
	CopyFileToContainer(ctx context.Context, hostFilePath string, containerFilePath string, fileMode int64) error
	CopyFileFromContainer(ctx context.Context, filePath string) (io.ReadCloser, error)
	GetLogProductionErrorChannel() <-chan error
}

// ImageBuildInfo defines what is needed to build an image
type ImageBuildInfo interface {
	BuildOptions() (client.ImageBuildOptions, error) // converts the ImageBuildInfo to a build.ImageBuildOptions
	GetContext() (io.Reader, error)                  // the path to the build context
	GetDockerfile() string                           // the relative path to the Dockerfile, including the file itself
	GetRepo() string                                 // get repo label for image
	GetTag() string                                  // get tag label for image
	BuildLogWriter() io.Writer                       // for output of build log, use io.Discard to disable the output
	ShouldBuildImage() bool                          // return true if the image needs to be built
	GetBuildArgs() map[string]*string                // return the environment args used to build the Dockerfile
	GetAuthConfigs() map[string]registry.AuthConfig  // Deprecated. Testcontainers will detect registry credentials automatically. Return the auth configs to be able to pull from an authenticated docker registry
}

// FromDockerfile represents the parameters needed to build an image from a Dockerfile
// rather than using a pre-built one
type FromDockerfile struct {
	Context        string                         // the path to the context of the docker build
	ContextArchive io.ReadSeeker                  // the tar archive file to send to docker that contains the build context
	Dockerfile     string                         // the path from the context to the Dockerfile for the image, defaults to "Dockerfile"
	Repo           string                         // the repo label for image, defaults to UUID
	Tag            string                         // the tag label for image, defaults to UUID
	BuildArgs      map[string]*string             // enable user to pass build args to docker daemon
	PrintBuildLog  bool                           // Deprecated: Use BuildLogWriter instead
	BuildLogWriter io.Writer                      // for output of build log, defaults to io.Discard
	AuthConfigs    map[string]registry.AuthConfig // Deprecated. Testcontainers will detect registry credentials automatically. Enable auth configs to be able to pull from an authenticated docker registry
	// KeepImage describes whether DockerContainer.Terminate should not delete the
	// container image. Useful for images that are built from a Dockerfile and take a
	// long time to build. Keeping the image also Docker to reuse it.
	KeepImage bool
	// BuildOptionsModifier Modifier for the build options before image build. Use it for
	// advanced configurations while building the image. Please consider that the modifier
	// is called after the default build options are set.
	BuildOptionsModifier func(*client.ImageBuildOptions)
}

type ContainerFile struct {
	HostFilePath      string    // If Reader is present, HostFilePath is ignored
	Reader            io.Reader // If Reader is present, HostFilePath is ignored
	ContainerFilePath string
	FileMode          int64
}

// validate validates the ContainerFile
func (c *ContainerFile) validate() error { _ = "STUB: not implemented"; return nil }

// ContainerRequest represents the parameters used to get a running container
type ContainerRequest struct {
	FromDockerfile
	HostAccessPorts          []int
	Image                    string
	ImageSubstitutors        []ImageSubstitutor
	Entrypoint               []string
	Env                      map[string]string
	ExposedPorts             []string // allow specifying protocol info
	Cmd                      []string
	Labels                   map[string]string
	Mounts                   ContainerMounts
	Tmpfs                    map[string]string
	RegistryCred             string // Deprecated: Testcontainers will detect registry credentials automatically
	WaitingFor               wait.Strategy
	Name                     string                                     // for specifying container name
	Hostname                 string                                     // Deprecated: Use [ConfigModifier] instead. S
	WorkingDir               string                                     // Deprecated: Use [ConfigModifier] instead. Specify the working directory of the container
	ExtraHosts               []string                                   // Deprecated: Use HostConfigModifier instead
	Privileged               bool                                       // Deprecated: Use [HostConfigModifier] instead. For starting privileged container
	Networks                 []string                                   // for specifying network names
	NetworkAliases           map[string][]string                        // for specifying network aliases
	NetworkMode              container.NetworkMode                      // Deprecated: Use HostConfigModifier instead
	Resources                container.Resources                        // Deprecated: Use HostConfigModifier instead
	Files                    []ContainerFile                            // files which will be copied when container starts
	User                     string                                     // Deprecated: Use [ConfigModifier] instead. For specifying uid:gid
	SkipReaper               bool                                       // Deprecated: The reaper is globally controlled by the .testcontainers.properties file or the TESTCONTAINERS_RYUK_DISABLED environment variable
	ReaperImage              string                                     // Deprecated: use WithImageName ContainerOption instead. Alternative reaper image
	ReaperOptions            []ContainerOption                          // Deprecated: the reaper is configured at the properties level, for an entire test session
	AutoRemove               bool                                       // Deprecated: Use HostConfigModifier instead. If set to true, the container will be removed from the host when stopped
	AlwaysPullImage          bool                                       // Always pull image
	ImagePlatform            string                                     // ImagePlatform describes the platform which the image runs on.
	Binds                    []string                                   // Deprecated: Use HostConfigModifier instead
	ShmSize                  int64                                      // Deprecated: Use [HostConfigModifier] instead. Amount of memory shared with the host (in bytes)
	CapAdd                   []string                                   // Deprecated: Use HostConfigModifier instead. Add Linux capabilities
	CapDrop                  []string                                   // Deprecated: Use HostConfigModifier instead. Drop Linux capabilities
	ConfigModifier           func(*container.Config)                    // Modifier for the config before container creation
	HostConfigModifier       func(*container.HostConfig)                // Modifier for the host config before container creation
	EndpointSettingsModifier func(map[string]*network.EndpointSettings) // Modifier for the network settings before container creation
	LifecycleHooks           []ContainerLifecycleHooks                  // define hooks to be executed during container lifecycle
	LogConsumerCfg           *LogConsumerConfig                         // define the configuration for the log producer and its log consumers to follow the logs
}

// sessionID returns the session ID for the container request.
func (c *ContainerRequest) sessionID() string { _ = "STUB: not implemented"; return "" }

// containerOptions functional options for a container
type containerOptions struct {
	ImageName           string
	RegistryCredentials string // Deprecated: Testcontainers will detect registry credentials automatically
}

// Deprecated: it will be removed in the next major release
// functional option for setting the reaper image
type ContainerOption func(*containerOptions)

// Deprecated: it will be removed in the next major release
// WithImageName sets the reaper image name
func WithImageName(imageName string) ContainerOption {
	_ = "STUB: not implemented"
	return *new(ContainerOption)
}

// Deprecated: Testcontainers will detect registry credentials automatically, and it will be removed in the next major release
// WithRegistryCredentials sets the reaper registry credentials
func WithRegistryCredentials(registryCredentials string) ContainerOption {
	_ = "STUB: not implemented"
	return *new(ContainerOption)
}

// Validate ensures that the ContainerRequest does not have invalid parameters configured to it
// ex. make sure you are not specifying both an image as well as a context
func (c *ContainerRequest) Validate() error { _ = "STUB: not implemented"; return nil }

// GetContext retrieve the build context for the request
// Must be closed when no longer needed.
func (c *ContainerRequest) GetContext() (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

// always pass context as absolute path

// only add .dockerignore if it exists

// parseDockerIgnore returns if the file exists, the excluded files and an error if any
func parseDockerIgnore(targetDir string) (bool, []string, error) {
	_ = "STUB: not implemented"
	// based on https://github.com/docker/cli/blob/master/cli/command/image/build/dockerignore.go#L14
	return false, nil, nil
}

// GetBuildArgs returns the env args to be used when creating from Dockerfile
func (c *ContainerRequest) GetBuildArgs() map[string]*string {
	_ = "STUB: not implemented"

	// GetDockerfile returns the Dockerfile from the ContainerRequest, defaults to "Dockerfile".
	// Sets FromDockerfile.Dockerfile to the default if blank.
	return nil
}

func (c *ContainerRequest) GetDockerfile() string { _ = "STUB: not implemented"; return "" }

// GetRepo returns the Repo label for image from the ContainerRequest, defaults to UUID.
// Sets FromDockerfile.Repo to the default value if blank.
func (c *ContainerRequest) GetRepo() string { _ = "STUB: not implemented"; return "" }

// GetTag returns the Tag label for image from the ContainerRequest, defaults to UUID.
// Sets FromDockerfile.Tag to the default value if blank.
func (c *ContainerRequest) GetTag() string { _ = "STUB: not implemented"; return "" }

// Deprecated: Testcontainers will detect registry credentials automatically, and it will be removed in the next major release.
// GetAuthConfigs returns the auth configs to be able to pull from an authenticated docker registry.
// Panics if an error occurs.
func (c *ContainerRequest) GetAuthConfigs() map[string]registry.AuthConfig {
	_ = "STUB: not implemented"
	return nil
}

// dockerFileImages returns the images from the request Dockerfile.
func (c *ContainerRequest) dockerFileImages() ([]string, error) {
	_ = "STUB: not implemented"
	return nil,

		// Source is a directory, we can read the Dockerfile directly.
		nil
}

// Source is an archive, we need to read it to get the Dockerfile.

// Reset the archive to the beginning.

// getAuthConfigsFromDockerfile returns the auth configs to be able to pull from an authenticated docker registry
func getAuthConfigsFromDockerfile(c *ContainerRequest) (map[string]registry.AuthConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get the auth configs once for all images as it can be a time-consuming operation.

// Credentials not found no config to add.

func (c *ContainerRequest) ShouldBuildImage() bool { _ = "STUB: not implemented"; return false }

func (c *ContainerRequest) ShouldKeepBuiltImage() bool {
	_ = "STUB: not implemented"

	// BuildLogWriter returns the io.Writer for output of log when building a Docker image from
	// a Dockerfile. It returns the BuildLogWriter from the ContainerRequest, defaults to io.Discard.
	// For backward compatibility, if BuildLogWriter is default and PrintBuildLog is true,
	// the function returns os.Stderr.
	//
	//nolint:staticcheck //FIXME
	return false
}

func (c *ContainerRequest) BuildLogWriter() io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

// BuildOptions returns the image build options when building a Docker image from a Dockerfile.
// It will apply some defaults and finally call the BuildOptionsModifier from the FromDockerfile struct,
// if set.
func (c *ContainerRequest) BuildOptions() (client.ImageBuildOptions, error) {
	_ = "STUB: not implemented"
	return *new(client.ImageBuildOptions), nil
}

// apply mandatory values after the modifier

// Make sure the auth configs from the Dockerfile are set right after the user-defined build options.

// make sure the first tag is the one defined in the ContainerRequest

// apply substitutors to the built image

// prepend the tag

// Do this as late as possible to ensure we don't leak the context on error/panic.

func (c *ContainerRequest) validateContextAndImage() error { _ = "STUB: not implemented"; return nil }

func (c *ContainerRequest) validateContextOrImageIsSpecified() error {
	_ = "STUB: not implemented"
	return nil
}

// validateMounts ensures that the mounts do not have duplicate targets.
// It will check the Mounts and HostConfigModifier.Binds fields.
func (c *ContainerRequest) validateMounts() error { _ = "STUB: not implemented"; return nil }
