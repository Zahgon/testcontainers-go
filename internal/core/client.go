package core

import (
	"context"

	"github.com/moby/moby/client"
)

// NewClient returns a new docker client extracting the docker host from the different alternatives
func NewClient(ctx context.Context, ops ...client.Opt) (*client.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// for further information, read https://docs.docker.com/engine/security/protect-access/

// passed options have priority over the default ones
