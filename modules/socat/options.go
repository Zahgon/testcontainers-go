package socat

import (
	"github.com/testcontainers/testcontainers-go"
)

type options struct {
	// targets is the map of targets of the socat container
	targets    map[int]Target
	targetsCmd string
}

func defaultOptions() options { _ = "STUB: not implemented"; return *new(options) }

// Compiler check to ensure that Option implements the testcontainers.ContainerCustomizer interface.
var _ testcontainers.ContainerCustomizer = (Option)(nil)

// Option is an option for the Socat container.
type Option func(*options) error

// Customize is a NOOP. It's defined to satisfy the testcontainers.ContainerCustomizer interface.
func (o Option) Customize(*testcontainers.GenericContainerRequest) error {
	_ = "STUB: not implemented"
	// NOOP to satisfy interface.
	return nil
}

// Target represents a target for the socat container.
// Create a new target with NewTarget or NewTargetWithInternalPort.
type Target struct {
	exposedPort  int
	internalPort int
	host         string
}

// ExposedPort returns the exposed port of the target.
func (t Target) ExposedPort() int { _ = "STUB: not implemented"; return 0 }

func (t Target) toCmd() string { _ = "STUB: not implemented"; return "" }

// NewTarget creates a new target for the socat container.
// The host of the target must be without the port,
// as it is internally mapped to the exposed port.
// The exposed port is exposed by the socat container.
func NewTarget(exposedPort int, host string) Target { _ = "STUB: not implemented"; return *new(Target) }

// NewTargetWithInternalPort creates a new target for the socat container.
// The host of the target must be without the port,
// as it is internally mapped to the exposed port.
// The exposed port is the port of the socat container, and
// the internal port is the port of the target container.
func NewTargetWithInternalPort(exposedPort int, internalPort int, host string) Target {
	_ = "STUB: not implemented"
	// If the internal port is not set, use the exposed port
	return *new(Target)
}

// WithTarget sets a single target for the socat container.
// The host of the target must be without the port, as it is internally mapped to the exposed port.
// Multiple calls to WithTarget will accumulate targets.
func WithTarget(target Target) Option { _ = "STUB: not implemented"; return *new(Option) }
