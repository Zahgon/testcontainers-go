package redpanda

import (
	"context"
	"net/http"
)

// AdminAPIClient is a client for the Redpanda Admin API.
type AdminAPIClient struct {
	BaseURL  string
	username string
	password string
	client   *http.Client
}

// NewAdminAPIClient creates a new AdminAPIClient.
func NewAdminAPIClient(baseURL string) *AdminAPIClient { _ = "STUB: not implemented"; return nil }

// WithHTTPClient sets the HTTP client for the AdminAPIClient.
func (cl *AdminAPIClient) WithHTTPClient(c *http.Client) *AdminAPIClient {
	_ = "STUB: not implemented"
	return nil

	// WithAuthentication sets the username and password for the AdminAPIClient.
}

func (cl *AdminAPIClient) WithAuthentication(username, password string) *AdminAPIClient {
	_ = "STUB: not implemented"
	return nil
}

// Username returns the username of the AdminAPIClient.
func (cl *AdminAPIClient) Username() string {
	_ = "STUB: not implemented"

	// Password returns the password of the AdminAPIClient.
	return ""
}

func (cl *AdminAPIClient) Password() string { _ = "STUB: not implemented"; return "" }

type createUserRequest struct {
	User      string `json:"username,omitempty"`
	Password  string `json:"password"`
	Algorithm string `json:"algorithm"`
}

// CreateUser creates a new user in Redpanda using Admin API.
func (cl *AdminAPIClient) CreateUser(ctx context.Context, username, password string) error {
	_ = "STUB: not implemented"
	return nil
}
