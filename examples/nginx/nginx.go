package nginx

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

type nginxContainer struct {
	testcontainers.Container
	URI string
}

func startContainer(ctx context.Context) (*nginxContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
