package network

import (
	"context"

	"github.com/moby/moby/api/types/network"
)

const (
	// FilterByID uses to filter network by identifier.
	FilterByID = "id"

	// FilterByName uses to filter network by name.
	FilterByName = "name"
)

// Get returns a network by its ID.
func Get(ctx context.Context, id string) (network.Summary, error) {
	_ = "STUB: not implemented"
	return *new(network.Summary), nil
}

// GetByName returns a network by its name.
func GetByName(ctx context.Context, name string) (network.Summary, error) {
	_ = "STUB: not implemented"
	return *new(network.Summary), nil
}

func get(ctx context.Context, filter string, value string) (network.Summary, error) {
	_ = "STUB: not implemented"
	return *
	// initialize to the zero value
	new(network.Summary), nil
}
