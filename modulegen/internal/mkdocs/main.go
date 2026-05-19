package mkdocs

import (
	"github.com/testcontainers/testcontainers-go/modulegen/internal/context"
)

// Generator is a struct that contains the logic to generate the mkdocs config file.
type Generator struct{}

// AddModule update modules in mkdocs
func (g Generator) AddModule(ctx context.Context, tcModule context.TestcontainersModule) error {
	_ = "STUB: not implemented"
	return nil
}

// Generate refresh the mkdocs config file for all the modules,
// excluding compose as it has its own page in the docs.
func (g Generator) Generate(ctx context.Context, examples []string, modules []string) error {
	_ = "STUB: not implemented"
	return nil
}

// The compose module has its own page in the docs.

// CopyConfig helper function to copy the mkdocs config file to a another file
// in the tests.
func CopyConfig(configFile string, tmpFile string) error { _ = "STUB: not implemented"; return nil }
