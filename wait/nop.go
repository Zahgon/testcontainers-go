package wait

import (
	"context"
	"io"
	"time"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/api/types/network"

	"github.com/testcontainers/testcontainers-go/exec"
)

var (
	_ Strategy        = (*NopStrategy)(nil)
	_ StrategyTimeout = (*NopStrategy)(nil)
)

type NopStrategy struct {
	timeout        *time.Duration
	waitUntilReady func(context.Context, StrategyTarget) error
}

func ForNop(
	waitUntilReady func(context.Context, StrategyTarget) error,
) *NopStrategy {
	_ = "STUB: not implemented"
	return nil
}

func (ws *NopStrategy) Timeout() *time.Duration {
	_ = "STUB: not implemented"

	// String returns a human-readable description of the wait strategy.
	return nil
}

func (ws *NopStrategy) String() string { _ = "STUB: not implemented"; return "" }

func (ws *NopStrategy) WithStartupTimeout(timeout time.Duration) *NopStrategy {
	_ = "STUB: not implemented"
	return nil
}

func (ws *NopStrategy) WaitUntilReady(ctx context.Context, target StrategyTarget) error {
	_ = "STUB: not implemented"
	return nil
}

type NopStrategyTarget struct {
	ReaderCloser   io.ReadCloser
	ContainerState container.State
}

func (st NopStrategyTarget) Host(_ context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (st NopStrategyTarget) Inspect(_ context.Context) (*container.InspectResponse, error) {
	_ = "STUB: not implemented"

	// Deprecated: use Inspect instead
	return nil, nil
}

func (st NopStrategyTarget) Ports(_ context.Context) (network.PortMap, error) {
	_ = "STUB: not implemented"
	return *new(network.PortMap), nil
}

func (st NopStrategyTarget) MappedPort(_ context.Context, n string) (network.Port, error) {
	_ = "STUB: not implemented"
	return *new(network.Port), nil
}

func (st NopStrategyTarget) Logs(_ context.Context) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (st NopStrategyTarget) Exec(_ context.Context, _ []string, _ ...exec.ProcessOption) (int, io.Reader, error) {
	_ = "STUB: not implemented"
	return 0, *new(io.Reader), nil
}

func (st NopStrategyTarget) State(_ context.Context) (*container.State, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (st NopStrategyTarget) CopyFileFromContainer(context.Context, string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}
