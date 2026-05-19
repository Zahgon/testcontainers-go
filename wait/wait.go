package wait

import (
	"context"
	"io"
	"time"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/api/types/network"

	"github.com/testcontainers/testcontainers-go/exec"
)

// Strategy defines the basic interface for a Wait Strategy
type Strategy interface {
	WaitUntilReady(context.Context, StrategyTarget) error
}

// StrategyTimeout allows MultiStrategy to configure a Strategy's Timeout
type StrategyTimeout interface {
	Timeout() *time.Duration
}

type StrategyTarget interface {
	Host(context.Context) (string, error)
	Inspect(context.Context) (*container.InspectResponse, error)
	Ports(ctx context.Context) (network.PortMap, error) // Deprecated: use Inspect instead
	MappedPort(context.Context, string) (network.Port, error)
	Logs(context.Context) (io.ReadCloser, error)
	Exec(context.Context, []string, ...exec.ProcessOption) (int, io.Reader, error)
	State(context.Context) (*container.State, error)
	CopyFileFromContainer(ctx context.Context, filePath string) (io.ReadCloser, error)
}

func checkTarget(ctx context.Context, target StrategyTarget) error {
	_ = "STUB: not implemented"
	return nil
}

func checkState(state *container.State) error { _ = "STUB: not implemented"; return nil }

func defaultStartupTimeout() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func defaultPollInterval() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }
