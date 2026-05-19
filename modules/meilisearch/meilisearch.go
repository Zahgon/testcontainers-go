package meilisearch

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	defaultMasterKey = "just-a-master-key-for-test"
	defaultHTTPPort  = "7700/tcp"
	masterKeyEnvVar  = "MEILI_MASTER_KEY"
)

// MeilisearchContainer represents the Meilisearch container type used in the module
type MeilisearchContainer struct {
	testcontainers.Container
	masterKey string
}

// MasterKey retrieves the master key of the Meilisearch container
func (c *MeilisearchContainer) MasterKey() string {
	_ = "STUB: not implemented"

	// Run creates an instance of the Meilisearch container type
	return ""
}

func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*MeilisearchContainer, error) {
	_ = "STUB: not implemented"
	// Gather all config options (defaults and then apply provided options)
	return nil, nil
}

// the wait strategy does not support TLS at the moment,
// so we need to disable it in the strategy for now.

// Address retrieves the address of the Meilisearch container.
// It will use http as protocol, as TLS is not supported at the moment.
func (c *MeilisearchContainer) Address(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
