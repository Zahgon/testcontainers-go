package dependabot

import (
	"github.com/testcontainers/testcontainers-go/modulegen/internal/context"
)

// Generator is a struct that contains the logic to generate the dependabot config file.
type Generator struct{}

// AddModule update dependabot with the new module
func (g Generator) AddModule(ctx context.Context, tcModule context.TestcontainersModule) error {
	_ = "STUB: not implemented"
	return nil
}

// Generate generates dependabot config file from source
func (g Generator) Generate(ctx context.Context, examples []string, modules []string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetUpdates returns the updates from the dependabot config file
func GetUpdates(configFile string) (Updates, error) {
	_ = "STUB: not implemented"
	return *new(Updates), nil
}

// CopyConfig helper function to copy the dependabot config file to a another file
// in the tests.
func CopyConfig(configFile string, tmpFile string) error { _ = "STUB: not implemented"; return nil }
