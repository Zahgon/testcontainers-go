package forgejo

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	defaultHTTPPort = "3000/tcp"
	defaultSSHPort  = "22/tcp"
	defaultUser     = "forgejo-admin"
	defaultPassword = "forgejo-admin"
	defaultEmail    = "admin@forgejo.local"
)

// Container represents the Forgejo container type used in the module
type Container struct {
	testcontainers.Container
	adminUsername string
	adminPassword string
}

// AdminUsername returns the admin username for the Forgejo instance.
func (c *Container) AdminUsername() string { _ = "STUB: not implemented"; return "" }

// AdminPassword returns the admin password for the Forgejo instance.
func (c *Container) AdminPassword() string { _ = "STUB: not implemented"; return "" }

// extractAdminCredentials parses FORGEJO_ADMIN_* env vars from the container
// environment, falling back to the default values for any that are not set.
func extractAdminCredentials(env []string) (username, password, email string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

// Run creates an instance of the Forgejo container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	// Closure variables populated by the PostReadies hook so we can avoid
	// a second container.Inspect call after Run returns.
	return nil, nil
}

// Use SQLite for simplicity in tests (no external DB needed).
// INSTALL_LOCK skips the install wizard so the instance is ready to use.

// Add lifecycle hook to create admin user after container is ready.
// The hook reads credentials from container env vars so that user-provided
// options (which override the defaults above) are respected.
// The command runs as the "git" user because Forgejo refuses to run CLI
// commands as root.

// Store credentials in closure for Run to use later.

// Credentials were populated by the PostReadies hook above.

// ConnectionString returns the HTTP URL for the Forgejo instance
func (c *Container) ConnectionString(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// SSHConnectionString returns the SSH endpoint for Git operations
func (c *Container) SSHConnectionString(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// WithAdminCredentials sets the admin username, password, and email for the Forgejo instance.
// These credentials are used to create an admin user after the container is ready.
func WithAdminCredentials(username, password, email string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithConfig sets a Forgejo configuration value using the FORGEJO__section__key
// environment variable format.
// See https://forgejo.org/docs/latest/admin/config-cheat-sheet/ for available options.
func WithConfig(section, key, value string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}
