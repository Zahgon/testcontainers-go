package k3s

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	// containerPorts {
	defaultKubeSecurePort     = "6443/tcp"
	defaultRancherWebhookPort = "8443/tcp"
	// }
	defaultKubeConfigK3sPath = "/etc/rancher/k3s/k3s.yaml"
)

// K3sContainer represents the K3s container type used in the module
type K3sContainer struct {
	testcontainers.Container
}

// path to the k3s manifests directory
const k3sManifests = "/var/lib/rancher/k3s/server/manifests/"

// WithManifest loads the manifest into the cluster. K3s applies it automatically during the startup process
func WithManifest(manifestPath string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// Deprecated: use Run instead
// RunContainer creates an instance of the K3s container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*K3sContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the K3s container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*K3sContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Host which will be used to access the Kubernetes server from tests.

func getContainerHost(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (string, error) {
	_ = "STUB: not implemented"
	// Use a dummy request to get the provider from options.
	return "", nil
}

// Fall back to localhost.

// GetKubeConfig returns the modified kubeconfig with server url
func (c *K3sContainer) GetKubeConfig(ctx context.Context) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func kubeConfigWithServerURL(kubeConfigYaml, server string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func marshal(config *KubeConfigValue) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func unmarshal(bytes []byte) (*KubeConfigValue, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *K3sContainer) LoadImages(ctx context.Context, images ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *K3sContainer) LoadImagesWithOpts(ctx context.Context, images []string, opts ...testcontainers.SaveImageOption) error {
	_ = "STUB: not implemented"
	return nil
}

// save image

// Close the file handle immediately: SaveImages and CopyFileToContainer
// open the file by name.
