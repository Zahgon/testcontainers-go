// Package dex provides a testcontainers module for the Dex OIDC provider.
//
// Supported grants: authorization_code, refresh_token, password. The
// client_credentials grant requires Dex ≥ v2.46.0 (or dexidp/dex:master)
// together with WithEnableClientCredentials() — this sets the
// DEX_CLIENT_CREDENTIAL_GRANT_ENABLED_BY_DEFAULT=true env var that gates
// the feature. Earlier releases return unsupported_grant_type.
//
// Example:
//
//	ctx := context.Background()
//	app, err := dex.NewClient("my-app",
//	    dex.WithClientSecret("s3cr3t"),
//	    dex.WithClientRedirectURIs("http://localhost/callback"),
//	)
//	if err != nil { log.Fatal(err) }
//	user, err := dex.NewUser("u@example.com", "u", "p")
//	if err != nil { log.Fatal(err) }
//
//	c, err := dex.Run(ctx, "dexidp/dex:v2.45.1",
//	    dex.WithClient(app),
//	    dex.WithUser(user),
//	)
//	defer testcontainers.TerminateContainer(c)
package dex

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	httpPort = "5556/tcp"
	grpcPort = "5557/tcp"

	configPath = "/etc/dex/dex.yml"
)

// Container is a running Dex OIDC provider.
type Container struct {
	testcontainers.Container
	issuer string
}

// Run starts Dex. The image is required (tc-go convention). Module options
// (WithClient, WithUser, WithIssuer, ...) and generic tc-go customizers may
// be mixed in the opts slice.
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Wait for Dex to serve discovery + gRPC port before returning. The
// discovery path is derived from the issuer's path component so
// WithIssuer("http://host/idp") probes "/idp/.well-known/...", not
// the default "/dex/...".

// No wait strategy on the request — Dex would fail readiness before the
// PostStart hook copies the YAML. The hook performs its own wait after
// copying.

// IssuerURL returns Dex's issuer URL. Empty if Run has not started.
func (c *Container) IssuerURL() string {
	_ = "STUB: not implemented"

	// ConfigEndpoint returns the OIDC discovery document URL.
	return ""
}

func (c *Container) ConfigEndpoint() string { _ = "STUB: not implemented"; return "" }

// JWKSEndpoint returns the JSON Web Key Set URL.
func (c *Container) JWKSEndpoint() string { _ = "STUB: not implemented"; return "" }

// TokenEndpoint returns the OAuth2 token URL.
func (c *Container) TokenEndpoint() string { _ = "STUB: not implemented"; return "" }

// AuthEndpoint returns the OAuth2 authorization URL.
func (c *Container) AuthEndpoint() string { _ = "STUB: not implemented"; return "" }

// GRPCEndpoint returns host:mappedPort for Dex's gRPC admin API. Errors
// propagate from the Docker API, or report that the container has not been
// started.
func (c *Container) GRPCEndpoint(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
