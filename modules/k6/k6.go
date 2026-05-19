package k6

import (
	"context"
	"io"
	"net/url"

	"github.com/testcontainers/testcontainers-go"
)

// cacheTarget is the path to the cache volume in the container.
const cacheTarget = "/cache"

// K6Container represents the K6 container type used in the module
type K6Container struct {
	testcontainers.Container
}

type DownloadableFile struct {
	Uri         url.URL //nolint:revive,staticcheck //FIXME
	DownloadDir string
	User        string
	Password    string
}

func (d *DownloadableFile) getDownloadPath() string { _ = "STUB: not implemented"; return "" }

func downloadFileFromDescription(d DownloadableFile) error { _ = "STUB: not implemented"; return nil }

// Set up HTTPS request with basic authorization.

// WithTestScript mounts the given script into the ./test directory in the container
// and passes it to k6 as the test to run.
// The path to the script must be an absolute path
func WithTestScript(scriptPath string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithTestScriptReader copies files into the Container using the Reader API
// The script base name is not a path, neither absolute nor relative and should
// be just the file name of the script
func WithTestScriptReader(reader io.Reader, scriptBaseName string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// add script to the k6 run command

// WithRemoteTestScript takes a RemoteTestFileDescription and copies to container
func WithRemoteTestScript(d DownloadableFile) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// Deprecated: use [testcontainers.WithCmdArgs] instead
// WithCmdOptions pass the given options to the k6 run command
func WithCmdOptions(options ...string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// SetEnvVar adds a '--env' command-line flag to the k6 command in the container for setting an environment variable for the test script.
func SetEnvVar(variable string, value string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithCache sets a volume as a cache for building the k6 binary
// If a volume name is provided in the TC_K6_BUILD_CACHE, this volume is used and it will
// persist across test sessions.
// If no value is provided, a volume is created and automatically deleted when the test session ends.
func WithCache() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// if no volume is provided, create one and ensure add labels for garbage collection

// Deprecated: use Run instead
// RunContainer creates an instance of the K6 container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*K6Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the K6 container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*K6Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CacheMount returns the name of volume used as a cache or an empty string
// if no cache was found.
func (k *K6Container) CacheMount(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
