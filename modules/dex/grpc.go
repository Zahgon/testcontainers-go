package dex

import (
	"context"

	"google.golang.org/grpc"
)

// AddClient creates a client via Dex's gRPC admin API.
//
// Note: Dex's api.Client proto has no grant_types field — clients added this
// way inherit Dex's defaults (authorization_code + refresh_token). For
// custom grants, register the client pre-start via WithClient.
//
// Not safe for concurrent use.
func (c *Container) AddClient(ctx context.Context, cl Client) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveClient deletes a client by ID.
//
// Not safe for concurrent use.
func (c *Container) RemoveClient(ctx context.Context, id string) error {
	_ = "STUB: not implemented"
	return nil
}

// AddUser registers a user in Dex's password DB via gRPC.
//
// Not safe for concurrent use.
func (c *Container) AddUser(ctx context.Context, u User) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveUser deletes a user by email.
//
// Not safe for concurrent use.
func (c *Container) RemoveUser(ctx context.Context, email string) error {
	_ = "STUB: not implemented"
	return nil
}

func dial(target string) (*grpc.ClientConn, error) { _ = "STUB: not implemented"; return nil, nil }
