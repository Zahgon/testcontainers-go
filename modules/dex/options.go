package dex

import (
	"log/slog"

	"github.com/testcontainers/testcontainers-go"
)

// connector is an internal record of a WithConnector call.
type connector struct {
	Type ConnectorType
	ID   string
	Name string
}

// options is the module-internal accumulator for Run().
type options struct {
	clients                 []Client
	users                   []User
	connectors              []connector
	issuer                  string
	skipApprovalScreen      bool
	storage                 Storage
	logLevel                slog.Level
	logger                  *slog.Logger
	enablePasswordDB        bool
	enableClientCredentials bool
}

func defaultOptions() options { _ = "STUB: not implemented"; return *new(options) }

// Option is a functional option for the Dex module. Options return an error
// so user-supplied values can be validated at Run time rather than failing
// silently in the rendered YAML.
//
// NOTE: Options must be passed directly to dex.Run. They satisfy the
// testcontainers.ContainerCustomizer interface only so the Run signature
// can accept them alongside generic tc-go customizers (e.g. network.With*)
// — Option.Customize is a no-op, so an Option forwarded through any
// wrapper that dispatches via Customize (instead of type-asserting to
// Option) is silently dropped.
type Option func(*options) error

// Compiler check: Option implements testcontainers.ContainerCustomizer.
var _ testcontainers.ContainerCustomizer = Option(nil)

// Customize is a no-op; real state mutation happens inside Run. See the
// Option type-level doc for why this is a no-op.
func (o Option) Customize(*testcontainers.GenericContainerRequest) error {
	_ = "STUB: not implemented"

	// WithClient registers a static client in Dex's YAML config. Unlike
	// gRPC-added clients, these may declare custom grant types.
	return nil
}

func WithClient(c Client) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUser registers a static password entry. The password DB connector is
// enabled by default, so no extra option is needed to consume the entry.
func WithUser(u User) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithConnector enables a Dex connector by type. For ConnectorPassword this
// is a no-op — the password DB is enabled by default and the template
// handles it separately; id and name are ignored and blank-field
// validation is skipped in that case. For other connectors
// (e.g. ConnectorMock) the entry is added to the rendered YAML, and blank
// id or name returns an error.
func WithConnector(t ConnectorType, id, name string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithIssuer overrides the default host:mappedPort-derived issuer. When set,
// Run uses the fast-path (direct YAML bind-mount). Callers are responsible
// for ensuring the URL is reachable from every client (tests and sibling
// containers).
func WithIssuer(url string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSkipApprovalScreen toggles Dex's oauth2.skipApprovalScreen. Default: true.
func WithSkipApprovalScreen(skip bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithStorage sets Dex's storage backend. Default: StorageSQLite.
func WithStorage(s Storage) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDisablePasswordDB disables Dex's built-in password connector. The
// caller must then configure at least one other connector via WithConnector,
// otherwise Run returns ErrNoAuthSource.
func WithDisablePasswordDB() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLogger routes Dex container logs through the supplied slog.Logger.
// When unset, Dex container logs are discarded. Calling WithLogger(nil)
// is a no-op; to discard logs again after setting a logger, drop the
// option rather than passing nil.
func WithLogger(logger *slog.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLogLevel sets Dex's own --log-level flag. Accepts a standard library
// slog.Level; values are mapped to Dex's level vocabulary (debug, info,
// warn, error). Default: slog.LevelInfo.
func WithLogLevel(level slog.Level) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnableClientCredentials enables Dex's OAuth2 client_credentials grant
// via the DEX_CLIENT_CREDENTIAL_GRANT_ENABLED_BY_DEFAULT=true environment
// variable.
//
// Requires Dex ≥ v2.46.0 or the dexidp/dex:master image tag. Earlier
// releases silently ignore the flag and token exchanges fail with
// unsupported_grant_type. This module does not validate the image tag —
// the caller must pin a compatible image.
func WithEnableClientCredentials() Option { _ = "STUB: not implemented"; return *new(Option) }
