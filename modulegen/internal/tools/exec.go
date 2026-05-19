package tools

// GoModTidy synchronizes the dependencies for a module or example,
// running the `go mod tidy` command in the given directory.
func GoModTidy(cmdDir string) error { _ = "STUB: not implemented"; return nil }

// GoVet checks the generated code for errors,
// running the `go vet ./...` command in the given directory.
func GoVet(cmdDir string) error { _ = "STUB: not implemented"; return nil }

// MakeLint runs the `make lint` command in the given directory.
func MakeLint(cmdDir string) error { _ = "STUB: not implemented"; return nil }

func runCommand(cmdDir string, command string, args ...string) error {
	_ = "STUB: not implemented"
	return nil
}
