package make

import (
	"github.com/testcontainers/testcontainers-go/modulegen/internal/context"
)

// Generator is a struct that contains the logic to generate the Makefile.
type Generator struct{}

// AddModule update Makefile with the new module
func (g Generator) AddModule(ctx context.Context, tcModule context.TestcontainersModule) error {
	_ = "STUB: not implemented"
	return nil
}

// GenerateMakefile creates Makefile for example
func GenerateMakefile(ctx context.Context, tcModule context.TestcontainersModule) error {
	_ = "STUB: not implemented"
	return nil
}
