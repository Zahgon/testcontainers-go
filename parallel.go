package testcontainers

import (
	"context"
	"sync"
)

const (
	defaultWorkersCount = 8
)

type ParallelContainerRequest []GenericContainerRequest

// ParallelContainersOptions represents additional options for parallel running
type ParallelContainersOptions struct {
	WorkersCount int // count of parallel workers. If field empty(zero), default value will be 'defaultWorkersCount'
}

// ParallelContainersRequestError represents error from parallel request
type ParallelContainersRequestError struct {
	Request GenericContainerRequest
	Error   error
}

type ParallelContainersError struct {
	Errors []ParallelContainersRequestError
}

func (gpe ParallelContainersError) Error() string { _ = "STUB: not implemented"; return "" }

// parallelContainersResult represents result.
type parallelContainersResult struct {
	ParallelContainersRequestError
	Container Container
}

func parallelContainersRunner(
	ctx context.Context,
	requests <-chan GenericContainerRequest,
	results chan<- parallelContainersResult,
	wg *sync.WaitGroup,
) {
	_ = "STUB: not implemented"
	return
}

// ParallelContainers creates a generic containers with parameters and run it in parallel mode
func ParallelContainers(ctx context.Context, reqs ParallelContainerRequest, opt ParallelContainersOptions) ([]Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// run workers
