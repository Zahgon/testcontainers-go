package postgres

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/testcontainers/testcontainers-go"
)

const (
	defaultUser         = "postgres"
	defaultPassword     = "postgres"
	defaultSnapshotName = "migrated_template"
)

//go:embed resources/customEntrypoint.sh
var embeddedCustomEntrypoint string

// PostgresContainer represents the postgres container type used in the module
type PostgresContainer struct {
	testcontainers.Container
	dbName       string
	user         string
	password     string
	snapshotName string
	// sqlDriverName is passed to sql.Open() to connect to the database when making or restoring snapshots.
	// This can be set if your app imports a different postgres driver, f.ex. "pgx"
	sqlDriverName string
}

// MustConnectionString panics if the address cannot be determined.
func (c *PostgresContainer) MustConnectionString(ctx context.Context, args ...string) string {
	_ = "STUB: not implemented"
	return ""
}

// ConnectionString returns the connection string for the postgres container, using the default 5432 port, and
// obtaining the host and exposed port from the container. It also accepts a variadic list of extra arguments
// which will be appended to the connection string. The format of the extra arguments is the same as the
// connection string format, e.g. "connect_timeout=10" or "application_name=myapp"
func (c *PostgresContainer) ConnectionString(ctx context.Context, args ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// WithConfigFile sets the config file to be used for the postgres container
// It will also set the "config_file" parameter to the path of the config file
// as a command line argument to the container
func WithConfigFile(cfg string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithDatabase sets the initial database to be created when the container starts
// It can be used to define a different name for the default database that is created when the image is first started.
// If it is not specified, then the value of WithUser will be used.
func WithDatabase(dbName string) testcontainers.ContainerCustomizer {
	_ = "STUB: not implemented"
	return *new(testcontainers.ContainerCustomizer)
}

// WithInitScripts sets the init scripts to be run when the container starts.
// These init scripts will be executed in sorted name order as defined by the container's current locale, which defaults to en_US.utf8.
// If you need to run your scripts in a specific order, consider using `WithOrderedInitScripts` instead.
func WithInitScripts(scripts ...string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithOrderedInitScripts sets the init scripts to be run when the container starts.
// The scripts will be run in the order that they are provided in this function.
func WithOrderedInitScripts(scripts ...string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithPassword sets the initial password of the user to be created when the container starts
// It is required for you to use the PostgreSQL image. It must not be empty or undefined.
// This environment variable sets the superuser password for PostgreSQL.
func WithPassword(password string) testcontainers.ContainerCustomizer {
	_ = "STUB: not implemented"
	return *new(testcontainers.ContainerCustomizer)
}

// WithUsername sets the initial username to be created when the container starts
// It is used in conjunction with WithPassword to set a user and its password.
// It will create the specified user with superuser power and a database with the same name.
// If it is not specified, then the default user of postgres will be used.
func WithUsername(user string) testcontainers.ContainerCustomizer {
	_ = "STUB: not implemented"
	return *new(testcontainers.ContainerCustomizer)
}

// Deprecated: use Run instead
// RunContainer creates an instance of the Postgres container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*PostgresContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the Postgres container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*PostgresContainer, error) {
	_ = "STUB: not implemented"
	// Gather all config options (defaults and then apply provided options)
	return nil, nil
}

// defaults to the user name

// Retrieve the actual env vars set on the container

type snapshotConfig struct {
	snapshotName string
}

// SnapshotOption is the type for passing options to the snapshot function of the database
type SnapshotOption func(container *snapshotConfig) *snapshotConfig

// WithSnapshotName adds a specific name to the snapshot database created from the main database defined on the
// container. The snapshot must not have the same name as your main database, otherwise it will be overwritten
func WithSnapshotName(name string) SnapshotOption {
	_ = "STUB: not implemented"
	return *new(SnapshotOption)
}

// WithSSLSettings configures the Postgres server to run with the provided CA Chain
// This will not function if the corresponding postgres conf is not correctly configured.
// Namely the paths below must match what is set in the conf file
func WithSSLCert(caCertFile string, certFile string, keyFile string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// Snapshot takes a snapshot of the current state of the database as a template, which can then be restored using
// the Restore method. By default, the snapshot will be created under a database called migrated_template, you can
// customize the snapshot name with the options.
// If a snapshot already exists under the given/default name, it will be overwritten with the new snapshot.
func (c *PostgresContainer) Snapshot(ctx context.Context, opts ...SnapshotOption) error {
	_ = "STUB: not implemented"
	return nil
}

// execute the commands to create the snapshot, in order

// Update pg_database to remove the template flag, then drop the database if it exists.
// This is needed because dropping a template database will fail.
// https://www.postgresql.org/docs/current/manage-ag-templatedbs.html

// Create a copy of the database to another database to use as a template now that it was fully migrated

// Snapshot the template database so we can restore it onto our original database going forward

// Restore will restore the database to a specific snapshot. By default, it will restore the last snapshot taken on the
// database by the Snapshot method. If a snapshot name is provided, it will instead try to restore the snapshot by name.
func (c *PostgresContainer) Restore(ctx context.Context, opts ...SnapshotOption) error {
	_ = "STUB: not implemented"
	return nil
}

// execute the commands to restore the snapshot, in order

// Terminate all connections to the template database explicitly as the forced drop below will sometimes
// not terminate them and then fail to drop the database.

// Drop the database if it exists

// Then restore the previous snapshot

func (c *PostgresContainer) checkSnapshotConfig(opts []SnapshotOption) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *PostgresContainer) execCommandsSQL(ctx context.Context, cmds ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// snapshotConnection connects to the actual database using the "postgres" sql.DB driver, if it exists.
// The returned function should be called as a defer() to close the pool.
// No need to close the individual connection, that is done as part of the pool close.
// Also, no need to cache the connection pool, since it is a single connection which is very fast to establish.
func (c *PostgresContainer) snapshotConnection(ctx context.Context) (*sql.Conn, func(), error) {
	_ = "STUB: not implemented"
	// Connect to the database "postgres" instead of the app one
	return nil, nil, nil
}

// Try to use an actual postgres connection, if the driver is loaded

func (c *PostgresContainer) execCommandsFallback(ctx context.Context, cmds []string) error {
	_ = "STUB: not implemented"
	return nil
}
