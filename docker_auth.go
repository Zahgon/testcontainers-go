package testcontainers

import (
	"context"
	"sync"

	"github.com/cpuguy83/dockercfg"
	"github.com/moby/moby/api/types/registry"
)

// defaultRegistryFn is variable overwritten in tests to check for behaviour with different default values.
var defaultRegistryFn = defaultRegistry

// getRegistryCredentials is a variable overwritten in tests to mock the dockercfg.GetRegistryCredentials function.
var getRegistryCredentials = dockercfg.GetRegistryCredentials

// DockerImageAuth returns the auth config for the given Docker image, extracting first its Docker registry.
// Finally, it will use the credential helpers to extract the information from the docker config file
// for that registry, if it exists.
func DockerImageAuth(ctx context.Context, image string) (string, registry.AuthConfig, error) {
	_ = "STUB: not implemented"
	return "", *new(registry.AuthConfig), nil
}

// dockerImageAuth returns the auth config for the given Docker image.
func dockerImageAuth(ctx context.Context, image string, configs map[string]registry.AuthConfig) (string, registry.AuthConfig, error) {
	_ = "STUB: not implemented"
	return "", *new(registry.AuthConfig), nil
}

// Normalize Docker Hub aliases for credential lookup

// This is https://index.docker.io/v1/

func getRegistryAuth(reg string, cfgs map[string]registry.AuthConfig) (registry.AuthConfig, bool) {
	_ = "STUB: not implemented"
	return *new(registry.AuthConfig), false
}

// fallback match using authentication key host

// url.Parse: The url may be relative (a path, without a host) [...]

// defaultRegistry returns the default registry to use when pulling images
// It will use the docker daemon to get the default registry, returning "https://index.docker.io/v1/" if
// it fails to get the information from the daemon
func defaultRegistry(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

// authConfigResult is a result looking up auth details for key.
type authConfigResult struct {
	key string
	cfg registry.AuthConfig
	err error
}

// credentialsCache is a cache for registry credentials.
type credentialsCache struct {
	entries map[string]credentials
	mtx     sync.RWMutex
}

// credentials represents the username and password for a registry.
type credentials struct {
	username string
	password string
}

var creds = &credentialsCache{entries: map[string]credentials{}}

// AuthConfig updates the details in authConfig for the given hostname
// as determined by the details in configKey.
func (c *credentialsCache) AuthConfig(hostname, configKey string, authConfig *registry.AuthConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// get returns the username and password for the given hostname
// as determined by the details in configPath.
// If the username is empty, the password is an identity token.
func (c *credentialsCache) get(hostname, configKey string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// No entry found, request and cache.

// configKey returns a key to use for caching credentials based on
// the contents of the currently active config.
func configKey(cfg *dockercfg.Config) (string, error) { _ = "STUB: not implemented"; return "", nil }

// getDockerAuthConfigs returns a map with the auth configs from the docker config file
// using the registry as the key
func getDockerAuthConfigs() (map[string]registry.AuthConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Look up credentials from the credential store.

// Create auth from the username and password encoding.

// In the case where the auth field in the .docker/conf.json is empty, and the user has
// credential helpers registered the auth comes from there.

// getDockerConfig returns the docker config file. It will internally check, in this particular order:
// 1. the DOCKER_AUTH_CONFIG environment variable, unmarshalling it into a dockercfg.Config
// 2. the DOCKER_CONFIG environment variable, as the path to the config file
// 3. else it will load the default config file, which is ~/.docker/config.json
func getDockerConfig() (*dockercfg.Config, error) { _ = "STUB: not implemented"; return nil, nil }
