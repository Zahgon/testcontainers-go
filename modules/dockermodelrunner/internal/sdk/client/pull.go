package client

import (
	"context"
)

// PullModel creates a model in the Docker Model Runner, by pulling the model from Docker Hub.
func (c *Client) PullModel(ctx context.Context, fullyQualifiedModelName string) error {
	_ = "STUB: not implemented"
	return nil
}

// The Docker Model Runner returns a 200 status code for a successful pulls

// TODO: use a progressbar instead of multiple line output.
