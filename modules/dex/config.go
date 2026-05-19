package dex

import (
	"log/slog"
)

// testBcryptCost is the bcrypt work factor used when hashing test passwords.
// Dex v2.45+ enforces a minimum cost of 10; cost 10 is the minimum and still
// ~10x faster than the default cost 14 used in production.
const testBcryptCost = 10

// dexYAML mirrors Dex's config shape. Marshaled directly via yaml.v3 so no
// string field can inject structural characters.
type dexYAML struct {
	Issuer           string          `yaml:"issuer"`
	Storage          storageBlock    `yaml:"storage"`
	Web              endpointBlock   `yaml:"web"`
	GRPC             grpcBlock       `yaml:"grpc"`
	Logger           loggerBlock     `yaml:"logger"`
	OAuth2           oauth2Block     `yaml:"oauth2"`
	EnablePasswordDB bool            `yaml:"enablePasswordDB"`
	StaticClients    []yamlClient    `yaml:"staticClients,omitempty"`
	StaticPasswords  []yamlPassword  `yaml:"staticPasswords,omitempty"`
	Connectors       []yamlConnector `yaml:"connectors,omitempty"`
}

type loggerBlock struct {
	Level string `yaml:"level"`
}

type storageBlock struct {
	Type   string            `yaml:"type"`
	Config map[string]string `yaml:"config,omitempty"`
}

type endpointBlock struct {
	HTTP string `yaml:"http"`
}

type grpcBlock struct {
	Addr string `yaml:"addr"`
}

type oauth2Block struct {
	SkipApprovalScreen bool     `yaml:"skipApprovalScreen"`
	GrantTypes         []string `yaml:"grantTypes"`
	PasswordConnector  string   `yaml:"passwordConnector,omitempty"`
}

type yamlClient struct {
	ID           string   `yaml:"id"`
	Secret       string   `yaml:"secret"`
	Name         string   `yaml:"name"`
	Public       bool     `yaml:"public"`
	RedirectURIs []string `yaml:"redirectURIs,omitempty"`
	GrantTypes   []string `yaml:"grantTypes,omitempty"`
}

type yamlPassword struct {
	Email    string `yaml:"email"`
	Hash     string `yaml:"hash"`
	Username string `yaml:"username"`
	UserID   string `yaml:"userID"`
}

type yamlConnector struct {
	Type string `yaml:"type"`
	ID   string `yaml:"id"`
	Name string `yaml:"name"`
}

// baseGrantTypes is the server-level grantTypes list emitted into the
// Dex config by default. client_credentials is intentionally omitted — Dex
// ≥ v2.46.0 rejects it at startup unless
// DEX_CLIENT_CREDENTIAL_GRANT_ENABLED_BY_DEFAULT=true is set, and v2.45.x
// treats advertising an unenabled grant as a configuration error. render
// appends client_credentials when WithEnableClientCredentials() has set
// the matching env var on the container.
var baseGrantTypes = []string{
	"authorization_code",
	"refresh_token",
	"password",
}

// render serializes an options struct into a Dex YAML config payload. It
// validates that at least one auth source is configured and that the
// issuer has been populated.
func render(o options) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Dex requires oauth2.passwordConnector to name the connector ID used for
// the password grant (ROPC). When the built-in password DB is active its
// synthetic connector ID is "local".

// dexLogLevel maps a standard library slog.Level to the string vocabulary
// Dex recognises in its YAML `logger.level` field. Values between slog's
// fixed levels round up to the next defined level (e.g. slog.LevelInfo+1
// → "warn", slog.LevelWarn+1 → "error"); sub-debug values clamp to
// "debug".
func dexLogLevel(l slog.Level) string { _ = "STUB: not implemented"; return "" }

// newUUIDv4 generates an RFC 4122 v4 UUID without importing a third-party dep.
// Returns an error from crypto/rand.Read rather than panicking so callers in
// the container-startup and gRPC-admin paths can surface it up the chain.
func newUUIDv4() (string, error) { _ = "STUB: not implemented"; return "", nil }
