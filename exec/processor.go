package exec

import (
	"bytes"
	"io"
	"sync"

	"github.com/moby/moby/client"
)

// ProcessOptions defines options applicable to the reader processor
type ProcessOptions struct {
	ExecConfig client.ExecCreateOptions
	Reader     io.Reader
}

// NewProcessOptions returns a new ProcessOptions instance
// with the given command and default options:
// - detach: false
// - attach stdout: true
// - attach stderr: true
func NewProcessOptions(cmd []string) *ProcessOptions { _ = "STUB: not implemented"; return nil }

// ProcessOption defines a common interface to modify the reader processor
// These options can be passed to the Exec function in a variadic way to customize the returned Reader instance
type ProcessOption interface {
	Apply(opts *ProcessOptions)
}

type ProcessOptionFunc func(opts *ProcessOptions)

func (fn ProcessOptionFunc) Apply(opts *ProcessOptions) { _ = "STUB: not implemented"; return }

func WithUser(user string) ProcessOption { _ = "STUB: not implemented"; return *new(ProcessOption) }

func WithWorkingDir(workingDir string) ProcessOption {
	_ = "STUB: not implemented"
	return *new(ProcessOption)
}

func WithEnv(env []string) ProcessOption { _ = "STUB: not implemented"; return *new(ProcessOption) }

// safeBuffer is a goroutine safe buffer.
type safeBuffer struct {
	mtx sync.Mutex
	buf bytes.Buffer
	err error
}

// Error sets an error for the next read.
func (sb *safeBuffer) Error(err error) { _ = "STUB: not implemented"; return }

// Write writes p to the buffer.
// It is safe for concurrent use by multiple goroutines.
func (sb *safeBuffer) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Read reads up to len(p) bytes into p from the buffer.
// It is safe for concurrent use by multiple goroutines.
func (sb *safeBuffer) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Multiplexed returns a [ProcessOption] that configures the command execution
// to combine stdout and stderr into a single stream without Docker's multiplexing headers.
func Multiplexed() ProcessOption { _ = "STUB: not implemented"; return *new(ProcessOption) }

// returning fast to bypass those options with a nil reader,
// which could be the case when other options are used
// to configure the exec creation.
