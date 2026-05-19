package testcontainers

import (
	"context"
	"fmt"
	"io"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/api/types/network"

	"github.com/testcontainers/testcontainers-go/log"
)

// ContainerRequestHook is a hook that will be called before a container is created.
// It can be used to modify container configuration before it is created,
// using the different lifecycle hooks that are available:
// - Creating
// For that, it will receive a ContainerRequest, modify it and return an error if needed.
type ContainerRequestHook func(ctx context.Context, req ContainerRequest) error

// ContainerHook is a hook that will be called after a container is created
// It can be used to modify the state of the container after it is created,
// using the different lifecycle hooks that are available:
// - Created
// - Starting
// - Started
// - Readied
// - Stopping
// - Stopped
// - Terminating
// - Terminated
// For that, it will receive a Container, modify it and return an error if needed.
type ContainerHook func(ctx context.Context, ctr Container) error

// ContainerLifecycleHooks is a struct that contains all the hooks that can be used
// to modify the container lifecycle. All the container lifecycle hooks except the PreCreates hooks
// will be passed to the container once it's created
type ContainerLifecycleHooks struct {
	PreBuilds      []ContainerRequestHook
	PostBuilds     []ContainerRequestHook
	PreCreates     []ContainerRequestHook
	PostCreates    []ContainerHook
	PreStarts      []ContainerHook
	PostStarts     []ContainerHook
	PostReadies    []ContainerHook
	PreStops       []ContainerHook
	PostStops      []ContainerHook
	PreTerminates  []ContainerHook
	PostTerminates []ContainerHook
}

// DefaultLoggingHook is a hook that will log the container lifecycle events
var DefaultLoggingHook = func(logger log.Logger) ContainerLifecycleHooks {
	shortContainerID := func(c Container) string {
		return c.GetContainerID()[:12]
	}

	return ContainerLifecycleHooks{
		PreBuilds: []ContainerRequestHook{
			func(_ context.Context, req ContainerRequest) error {
				logger.Printf("🐳 Building image %s:%s", req.GetRepo(), req.GetTag())
				return nil
			},
		},
		PostBuilds: []ContainerRequestHook{
			func(_ context.Context, req ContainerRequest) error {
				logger.Printf("✅ Built image %s", req.Image)
				return nil
			},
		},
		PreCreates: []ContainerRequestHook{
			func(_ context.Context, req ContainerRequest) error {
				logger.Printf("🐳 Creating container for image %s", req.Image)
				return nil
			},
		},
		PostCreates: []ContainerHook{
			func(_ context.Context, c Container) error {
				logger.Printf("✅ Container created: %s", shortContainerID(c))
				return nil
			},
		},
		PreStarts: []ContainerHook{
			func(_ context.Context, c Container) error {
				logger.Printf("🐳 Starting container: %s", shortContainerID(c))
				return nil
			},
		},
		PostStarts: []ContainerHook{
			func(_ context.Context, c Container) error {
				logger.Printf("✅ Container started: %s", shortContainerID(c))
				return nil
			},
		},
		PostReadies: []ContainerHook{
			func(_ context.Context, c Container) error {
				logger.Printf("🔔 Container is ready: %s", shortContainerID(c))
				return nil
			},
		},
		PreStops: []ContainerHook{
			func(_ context.Context, c Container) error {
				logger.Printf("🐳 Stopping container: %s", shortContainerID(c))
				return nil
			},
		},
		PostStops: []ContainerHook{
			func(_ context.Context, c Container) error {
				logger.Printf("✅ Container stopped: %s", shortContainerID(c))
				return nil
			},
		},
		PreTerminates: []ContainerHook{
			func(_ context.Context, c Container) error {
				logger.Printf("🐳 Terminating container: %s", shortContainerID(c))
				return nil
			},
		},
		PostTerminates: []ContainerHook{
			func(_ context.Context, c Container) error {
				logger.Printf("🚫 Container terminated: %s", shortContainerID(c))
				return nil
			},
		},
	}
}

// defaultPreCreateHook is a hook that will apply the default configuration to the container
var defaultPreCreateHook = func(p *DockerProvider, dockerInput *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig) ContainerLifecycleHooks {
	return ContainerLifecycleHooks{
		PreCreates: []ContainerRequestHook{
			func(ctx context.Context, req ContainerRequest) error {
				return p.preCreateContainerHook(ctx, req, dockerInput, hostConfig, networkingConfig)
			},
		},
	}
}

// defaultCopyFileToContainerHook is a hook that will copy files to the container after it's created
// but before it's started
var defaultCopyFileToContainerHook = func(files []ContainerFile) ContainerLifecycleHooks {
	return ContainerLifecycleHooks{
		PostCreates: []ContainerHook{
			// copy files to container after it's created
			func(ctx context.Context, c Container) error {
				for _, f := range files {
					if err := f.validate(); err != nil {
						return fmt.Errorf("invalid file: %w", err)
					}

					var err error
					// Bytes takes precedence over HostFilePath
					if f.Reader != nil {
						bs, ioerr := io.ReadAll(f.Reader)
						if ioerr != nil {
							return fmt.Errorf("can't read from reader: %w", ioerr)
						}

						err = c.CopyToContainer(ctx, bs, f.ContainerFilePath, f.FileMode)
					} else {
						err = c.CopyFileToContainer(ctx, f.HostFilePath, f.ContainerFilePath, f.FileMode)
					}

					if err != nil {
						return fmt.Errorf("can't copy %s to container: %w", f.HostFilePath, err)
					}
				}

				return nil
			},
		},
	}
}

// defaultLogConsumersHook is a hook that will start log consumers after the container is started
var defaultLogConsumersHook = func(cfg *LogConsumerConfig) ContainerLifecycleHooks {
	return ContainerLifecycleHooks{
		PostStarts: []ContainerHook{
			// Produce logs sending details to the log consumers.
			// See combineContainerHooks for the order of execution.
			func(ctx context.Context, c Container) error {
				if cfg == nil || len(cfg.Consumers) == 0 {
					return nil
				}

				dockerContainer := c.(*DockerContainer)
				dockerContainer.resetConsumers(cfg.Consumers)

				return dockerContainer.startLogProduction(ctx, cfg.Opts...)
			},
		},
		PostStops: []ContainerHook{
			// Stop the log production.
			// See combineContainerHooks for the order of execution.
			func(_ context.Context, c Container) error {
				if cfg == nil || len(cfg.Consumers) == 0 {
					return nil
				}

				dockerContainer := c.(*DockerContainer)
				return dockerContainer.stopLogProduction()
			},
		},
	}
}

// defaultReadinessHook is a hook that will wait for the container to be ready
var defaultReadinessHook = func() ContainerLifecycleHooks {
	return ContainerLifecycleHooks{
		PostStarts: []ContainerHook{
			// wait for the container to be ready
			func(ctx context.Context, c Container) error {
				dockerContainer := c.(*DockerContainer)

				// if a Wait Strategy has been specified, wait before returning
				if dockerContainer.WaitingFor != nil {
					strategy := dockerContainer.WaitingFor
					strategyDesc := "unknown strategy"
					if s, ok := strategy.(fmt.Stringer); ok {
						strategyDesc = s.String()
					}
					dockerContainer.logger.Printf(
						"⏳ Waiting for container id %s image: %s. Waiting for: %+v",
						dockerContainer.ID[:12], dockerContainer.Image, strategyDesc,
					)
					if err := strategy.WaitUntilReady(ctx, dockerContainer); err != nil {
						return fmt.Errorf("wait until ready: %w", err)
					}
				}

				dockerContainer.isRunning.Store(true)

				return nil
			},
		},
	}
}

// buildingHook is a hook that will be called before a container image is built.
func (req ContainerRequest) buildingHook(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// builtHook is a hook that will be called after a container image is built.
func (req ContainerRequest) builtHook(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// creatingHook is a hook that will be called before a container is created.
func (req ContainerRequest) creatingHook(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// applyLifecycleHooks calls hook on all LifecycleHooks.
func (req ContainerRequest) applyLifecycleHooks(hook func(lifecycleHooks ContainerLifecycleHooks) error) error {
	_ = "STUB: not implemented"
	return nil
}

// createdHook is a hook that will be called after a container is created.
func (c *DockerContainer) createdHook(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// startingHook is a hook that will be called before a container is started.
func (c *DockerContainer) startingHook(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// startedHook is a hook that will be called after a container is started.
func (c *DockerContainer) startedHook(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// readiedHook is a hook that will be called after a container is ready.
func (c *DockerContainer) readiedHook(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// printLogs is a helper function that will print the logs of a Docker container
// We are going to use this helper function to inform the user of the logs when an error occurs
func (c *DockerContainer) printLogs(ctx context.Context, cause error) {
	_ = "STUB: not implemented"
	return
}

// stoppingHook is a hook that will be called before a container is stopped.
func (c *DockerContainer) stoppingHook(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// stoppedHook is a hook that will be called after a container is stopped.
func (c *DockerContainer) stoppedHook(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// terminatingHook is a hook that will be called before a container is terminated.
func (c *DockerContainer) terminatingHook(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// terminatedHook is a hook that will be called after a container is terminated.
func (c *DockerContainer) terminatedHook(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// applyLifecycleHooks applies all lifecycle hooks reporting the container logs on error if logError is true.
func (c *DockerContainer) applyLifecycleHooks(ctx context.Context, logError bool, hooks func(lifecycleHooks ContainerLifecycleHooks) []ContainerHook) error {
	_ = "STUB: not implemented"
	return nil
}

// Context has timed out so need a new context to get logs.

// Building is a hook that will be called before a container image is built.
func (c ContainerLifecycleHooks) Building(ctx context.Context) func(req ContainerRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// Building is a hook that will be called before a container image is built.
func (c ContainerLifecycleHooks) Built(ctx context.Context) func(req ContainerRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// Creating is a hook that will be called before a container is created.
func (c ContainerLifecycleHooks) Creating(ctx context.Context) func(req ContainerRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// containerRequestHook returns a function that will iterate over all
// the hooks and call them one by one until there is an error.
func containerRequestHook(ctx context.Context, hooks []ContainerRequestHook) func(req ContainerRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// containerHookFn is a helper function that will create a function to be returned by all the different
// container lifecycle hooks. The created function will iterate over all the hooks and call them one by one.
func containerHookFn(ctx context.Context, containerHook []ContainerHook) func(container Container) error {
	_ = "STUB: not implemented"
	return nil
}

// Created is a hook that will be called after a container is created
func (c ContainerLifecycleHooks) Created(ctx context.Context) func(container Container) error {
	_ = "STUB: not implemented"
	return nil
}

// Starting is a hook that will be called before a container is started
func (c ContainerLifecycleHooks) Starting(ctx context.Context) func(container Container) error {
	_ = "STUB: not implemented"
	return nil
}

// Started is a hook that will be called after a container is started
func (c ContainerLifecycleHooks) Started(ctx context.Context) func(container Container) error {
	_ = "STUB: not implemented"
	return nil
}

// Readied is a hook that will be called after a container is ready
func (c ContainerLifecycleHooks) Readied(ctx context.Context) func(container Container) error {
	_ = "STUB: not implemented"
	return nil
}

// Stopping is a hook that will be called before a container is stopped
func (c ContainerLifecycleHooks) Stopping(ctx context.Context) func(container Container) error {
	_ = "STUB: not implemented"
	return nil
}

// Stopped is a hook that will be called after a container is stopped
func (c ContainerLifecycleHooks) Stopped(ctx context.Context) func(container Container) error {
	_ = "STUB: not implemented"
	return nil
}

// Terminating is a hook that will be called before a container is terminated
func (c ContainerLifecycleHooks) Terminating(ctx context.Context) func(container Container) error {
	_ = "STUB: not implemented"
	return nil
}

// Terminated is a hook that will be called after a container is terminated
func (c ContainerLifecycleHooks) Terminated(ctx context.Context) func(container Container) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *DockerProvider) preCreateContainerHook(ctx context.Context, req ContainerRequest, dockerInput *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// validate only the mount sources that implement the Validator interface

// prepare mounts

// #248: Docker allows only one network to be specified during container creation
// If there is more than one network specified in the request container should be attached to them
// once it is created. We will take a first network if any specified in the request and use it to create container

// Expose ports automatically if the container request exposes zero ports and the container
// does not run in a container network. The NetworkMode check must be done after the pre-creation
// Modifiers are called, so the network mode is already set.

// combineContainerHooks returns a ContainerLifecycle hook as the result
// of combining the default hooks with the user-defined hooks.
//
// The order of hooks is the following:
// - Pre-hooks run the default hooks first then the user-defined hooks
// - Post-hooks run the user-defined hooks first then the default hooks
func combineContainerHooks(defaultHooks, userDefinedHooks []ContainerLifecycleHooks) ContainerLifecycleHooks {
	_ = "STUB: not implemented"
	// We use reflection here to ensure that any new hooks are handled.
	return *new(ContainerLifecycleHooks)
}

// Append the user-defined hooks after the default pre-hooks
// and because the post hooks are still empty, the user-defined
// post-hooks will be the first ones to be executed.

// Finally, append the default post-hooks.

func parseExposedPorts(specs []string) (network.PortSet, error) {
	_ = "STUB: not implemented"
	return *new(network.PortSet), nil
}

// mergePortBindings returns a PortMap for the given exposedPortSet.
//
// For each port in exposedPortSet, a binding is ensured:
//   - If configPortMap contains bindings for that port, those bindings are used.
//   - Otherwise, a default binding with HostPort "0" (ephemeral allocation)
//     is assigned.
//
// Bindings for ports not present in exposedPortSet are not preserved.
// Any binding with an empty HostPort is normalized to "0".
//
// TODO(thaJeztah): this logic seems the reverse of the docker CLI, which
// exposes ports if the user requests a port-mapping (i.e., if a port-mapping
// is requested, but not exposed, we map the port *and* add an entry to
// ExposedPorts). The logic here is the reverse; any port "mapped" in
// HostConfig.PortBindings is dropped if is not exposed.
func mergePortBindings(configPortMap network.PortMap, exposedPortSet network.PortSet) network.PortMap {
	_ = "STUB: not implemented"
	return *new(network.PortMap)
}

// Fix: Ensure that ports with empty HostPort get "0" for automatic allocation
// This fixes the UDP port binding issue where ports were getting HostPort:0 instead of being allocated

// Tell Docker to allocate a random port

// defaultHostConfigModifier provides a default modifier including the deprecated fields
func defaultConfigModifier(req ContainerRequest) func(config *container.Config) {
	_ = "STUB: not implemented"
	return nil
}

// defaultHostConfigModifier provides a default modifier including the deprecated fields
func defaultHostConfigModifier(req ContainerRequest) func(hostConfig *container.HostConfig) {
	_ = "STUB: not implemented"
	return nil
}
