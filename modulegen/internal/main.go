package internal

import (
	"github.com/testcontainers/testcontainers-go/modulegen/internal/context"
)

// Generate generates all the files for a module or example,
// running the `go mod tidy`, `go vet` and `make lint` commands
// in the given directory.
func Generate(moduleVar context.TestcontainersModuleVar, isModule bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Refresh refreshes the modules and examples, returning an error if something goes wrong.
func Refresh(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// update examples in mkdocs
// update examples in dependabot
// update vscode workspace

// ProjectGenerator is the interface for the project generators, which takes
// a context and for each module in the context, adds it to the project files,
// returning an error if something goes wrong.
type ProjectGenerator interface {
	Generate(ctx context.Context, examples []string, modules []string) error
}

// FileGenerator is the interface for the file generators, which takes
// a module and generate a file for it, returning an error if something goes wrong.
type FileGenerator interface {
	AddModule(context.Context, context.TestcontainersModule) error
}

func GenerateFiles(ctx context.Context, tcModule context.TestcontainersModule) error {
	_ = "STUB: not implemented"
	return nil
}

// creates Makefile for module
// creates go.mod for module
// update examples in mkdocs
// update examples in dependabot

// they are based on the content of the modules in the project workspace,
// not in the new module to be added, that's why they happen after the actual
// module generation

// update vscode workspace
