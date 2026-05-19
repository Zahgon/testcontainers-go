package cockroachdb

import (
	"errors"

	"github.com/testcontainers/testcontainers-go"
)

// errInsecureWithPassword is returned when trying to use insecure mode with a password.
var errInsecureWithPassword = errors.New("insecure mode cannot be used with a password")

// WithDatabase sets the name of the database to create and use.
// This will be converted to lowercase as CockroachDB forces the database to be lowercase.
// The database creation will be skipped if data exists in the `/cockroach/cockroach-data` directory within the container.
func WithDatabase(database string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithUser sets the name of the user to create and connect as.
// This will be converted to lowercase as CockroachDB forces the user to be lowercase.
// The user creation will be skipped if data exists in the `/cockroach/cockroach-data` directory within the container.
func WithUser(user string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithPassword sets the password of the user to create and connect as.
// The user creation will be skipped if data exists in the `/cockroach/cockroach-data` directory within the container.
// This will error if insecure mode is enabled.
func WithPassword(password string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithStoreSize sets the amount of available [in-memory storage].
//
// [in-memory storage]: https://www.cockroachlabs.com/docs/stable/cockroach-start#store
func WithStoreSize(size string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// Wasn't found, add it.

// WithNoClusterDefaults disables the default cluster settings script.
//
// Without this option Cockroach containers run `data/cluster-defaults.sql` on startup
// which configures the settings recommended by Cockroach Labs for [local testing clusters]
// unless data exists in the `/cockroach/cockroach-data` directory within the container.
//
// [local testing clusters]: https://www.cockroachlabs.com/docs/stable/local-testing
func WithNoClusterDefaults() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithInitScripts adds the given scripts to those automatically run when the container starts.
// These will be ignored if data exists in the `/cockroach/cockroach-data` directory within the container.
func WithInitScripts(scripts ...string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithInsecure enables insecure mode which disables TLS.
func WithInsecure() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// configure sets the CockroachDBContainer options from the given request and updates the request
// wait strategies to match the options.
// This option must be called after all the options have been applied, in order to extract
// the credentials from the environment variables and the TLS strategy from the wait strategy.
func (c *CockroachDBContainer) configure() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// refresh the credentials from the environment variables

// Walk the wait strategies to find the TLS strategy and either remove it or
// update the client certificate files to match the user and configure the
// container to use the TLS strategy.

// If insecure mode is enabled, the certificate strategy is removed.

// Update the client certificate files to match the user which may have changed.

// Stop the walk as the certificate strategy has been found.
