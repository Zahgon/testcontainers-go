package testcontainers

import "github.com/testcontainers/testcontainers-go/log"

// Validate our types implement the required interfaces.
var (
	_ ContainerCustomizer   = LoggerOption{}
	_ GenericProviderOption = LoggerOption{}
	_ DockerProviderOption  = LoggerOption{}
)

// WithLogger returns a generic option that sets the logger to be used.
//
// Consider calling this before other "With functions" as these may generate logs.
//
// This can be given a TestLogger to collect the logs from testcontainers into a
// test case.
func WithLogger(logger log.Logger) LoggerOption {
	_ = "STUB: not implemented"
	return *new(LoggerOption)
}

// LoggerOption is a generic option that sets the logger to be used.
//
// It can be used to set the logger for providers and containers.
type LoggerOption struct {
	logger log.Logger
}

// ApplyGenericTo implements GenericProviderOption.
func (o LoggerOption) ApplyGenericTo(opts *GenericProviderOptions) {
	_ = "STUB: not implemented"
	return

	// ApplyDockerTo implements DockerProviderOption.
}

func (o LoggerOption) ApplyDockerTo(opts *DockerProviderOptions) { _ = "STUB: not implemented"; return }

// Customize implements ContainerCustomizer.
func (o LoggerOption) Customize(req *GenericContainerRequest) error {
	_ = "STUB: not implemented"
	return nil
}
