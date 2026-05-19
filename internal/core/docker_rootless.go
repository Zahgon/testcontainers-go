package core

import (
	"context"
	"errors"
)

var (
	ErrRootlessDockerNotFound               = errors.New("rootless Docker not found")
	ErrRootlessDockerNotFoundHomeDesktopDir = errors.New("checked path: ~/.docker/desktop/docker.sock")
	ErrRootlessDockerNotFoundHomeRunDir     = errors.New("checked path: ~/.docker/run/docker.sock")
	ErrRootlessDockerNotFoundRunDir         = errors.New("checked path: /run/user/${uid}/docker.sock")
	ErrRootlessDockerNotFoundXDGRuntimeDir  = errors.New("checked path: $XDG_RUNTIME_DIR")
	ErrRootlessDockerNotSupportedWindows    = errors.New("rootless Docker is not supported on Windows")
	ErrXDGRuntimeDirNotSet                  = errors.New("XDG_RUNTIME_DIR is not set")
)

// baseRunDir is the base directory for the "/run/user/${uid}" directory.
// It is a variable so it can be modified for testing.
var baseRunDir = "/run"

// IsWindows returns if the current OS is Windows. For that it checks the GOOS environment variable or the runtime.GOOS constant.
func IsWindows() bool { _ = "STUB: not implemented"; return false }

// rootlessDockerSocketPath returns if the path to the rootless Docker socket exists.
// The rootless socket path is determined by the following order:
//
//  1. XDG_RUNTIME_DIR environment variable.
//  2. ~/.docker/run/docker.sock file.
//  3. ~/.docker/desktop/docker.sock file.
//  4. /run/user/${uid}/docker.sock file.
//  5. Else, return ErrRootlessDockerNotFound, wrapping specific errors for each of the above paths.
//
// It should include the Docker socket schema (unix://) in the returned path.
func rootlessDockerSocketPath(_ context.Context) (string, error) {
	_ = "STUB: not implemented"
	// adding a manner to test it on non-windows machines, setting the GOOS env var to windows
	// This is needed because runtime.GOOS is a constant that returns the OS of the machine running the test
	return "", nil
}

func fileExists(f string) bool { _ = "STUB: not implemented"; return false }

func parseURL(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// return the original URL, as it is a valid TCP URL

// rootlessSocketPathFromEnv returns the path to the rootless Docker socket from the XDG_RUNTIME_DIR environment variable.
// It should include the Docker socket schema (unix://) in the returned path.
func rootlessSocketPathFromEnv() (string, error) { _ = "STUB: not implemented"; return "", nil }

// rootlessSocketPathFromHomeRunDir returns the path to the rootless Docker socket from the ~/.docker/run/docker.sock file.
func rootlessSocketPathFromHomeRunDir() (string, error) { _ = "STUB: not implemented"; return "", nil }

// rootlessSocketPathFromHomeDesktopDir returns the path to the rootless Docker socket from the ~/.docker/desktop/docker.sock file.
func rootlessSocketPathFromHomeDesktopDir() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// rootlessSocketPathFromRunDir returns the path to the rootless Docker socket from the /run/user/<uid>/docker.sock file.
func rootlessSocketPathFromRunDir() (string, error) { _ = "STUB: not implemented"; return "", nil }
