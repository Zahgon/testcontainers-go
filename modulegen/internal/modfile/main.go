package modfile

import (
	"golang.org/x/mod/modfile"
)

// GenerateModFile generates a go.mod file for a module or example.
func GenerateModFile(exampleDir string, rootGoModFilePath string, directory string, tcVersion string) error {
	_ = "STUB: not implemented"
	return nil
}

func newModFile(moduleStmt string, goStmt string, tcPath string, tcVersion string) (*modfile.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
