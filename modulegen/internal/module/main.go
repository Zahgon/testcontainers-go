package module

import (
	"text/template"

	"github.com/testcontainers/testcontainers-go/modulegen/internal/context"
)

// Generator is a struct that contains the logic to generate the module files.
type Generator struct{}

// AddModule creates the go.mod file for the module
func (g Generator) AddModule(ctx context.Context, tcModule context.TestcontainersModule) error {
	_ = "STUB: not implemented"
	return nil
}

func generateGoFiles(moduleDir string, tcModule context.TestcontainersModule) error {
	_ = "STUB: not implemented"
	return nil
}

func generateGoModFile(ctx context.Context, moduleDir string, tcModule context.TestcontainersModule) error {
	_ = "STUB: not implemented"
	return nil
}

// GenerateFiles generates the module files from the template files.
func GenerateFiles(moduleDir string, moduleName string, funcMap template.FuncMap, tcModule any) error {
	_ = "STUB: not implemented"
	return nil
}
