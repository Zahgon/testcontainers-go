package openldap

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	defaultUser     = "admin"
	defaultPassword = "adminpassword"
	defaultRoot     = "dc=example,dc=org"
	defaultAdminDn  = "cn=admin,dc=example,dc=org"
)

// OpenLDAPContainer represents the OpenLDAP container type used in the module
type OpenLDAPContainer struct {
	testcontainers.Container
	adminUsername string
	adminPassword string
	rootDn        string
}

// ConnectionString returns the connection string for the OpenLDAP container
func (c *OpenLDAPContainer) ConnectionString(ctx context.Context, _ ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// LoadLdif loads an ldif file into the OpenLDAP container
func (c *OpenLDAPContainer) LoadLdif(ctx context.Context, ldif []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// WithAdminUsername sets the initial admin username to be created when the container starts
// It is used in conjunction with WithAdminPassword to set a username and its password.
// It will create the specified user with admin power.
func WithAdminUsername(username string) testcontainers.ContainerCustomizer {
	_ = "STUB: not implemented"
	return *new(testcontainers.ContainerCustomizer)
}

// WithAdminPassword sets the initial admin password of the user to be created when the container starts
// It is used in conjunction with WithAdminUsername to set a username and its password.
// It will set the admin password for OpenLDAP.
func WithAdminPassword(password string) testcontainers.ContainerCustomizer {
	_ = "STUB: not implemented"
	return *new(testcontainers.ContainerCustomizer)
}

// WithRoot sets the root of the OpenLDAP instance
func WithRoot(root string) testcontainers.ContainerCustomizer {
	_ = "STUB: not implemented"
	return *new(testcontainers.ContainerCustomizer)
}

// WithInitialLdif sets the initial ldif file to be loaded into the OpenLDAP container
func WithInitialLdif(ldif string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// Deprecated: use Run instead
// RunContainer creates an instance of the OpenLDAP container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*OpenLDAPContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the OpenLDAP container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*OpenLDAPContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve the actual env vars set on the container
