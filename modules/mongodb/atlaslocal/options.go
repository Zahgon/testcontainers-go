package atlaslocal

import (
	"github.com/testcontainers/testcontainers-go"
)

const (
	passwordContainerPath      = "/run/secrets/mongo-root-password"
	usernameContainerPath      = "/run/secrets/mongo-root-username"
	envMongotLogFile           = "MONGOT_LOG_FILE"
	envRunnerLogFile           = "RUNNER_LOG_FILE"
	envMongoDBInitDatabase     = "MONGODB_INITDB_DATABASE"
	envMongoDBInitUsername     = "MONGODB_INITDB_ROOT_USERNAME"
	envMongoDBInitPassword     = "MONGODB_INITDB_ROOT_PASSWORD"
	envMongoDBInitUsernameFile = "MONGODB_INITDB_ROOT_USERNAME_FILE"
	envMongoDBInitPasswordFile = "MONGODB_INITDB_ROOT_PASSWORD_FILE"
	envDoNotTrack              = "DO_NOT_TRACK"
)

// Compiler check to ensure that Option implements the testcontainers.ContainerCustomizer interface.
var _ testcontainers.ContainerCustomizer = (Option)(nil)

type options struct {
	username          string
	password          string
	localUsernameFile string
	localPasswordFile string
	noTelemetry       bool
	database          string
	mongotLogPath     string
	runnerLogPath     string

	files []testcontainers.ContainerFile
}

func (opts options) env() map[string]string { _ = "STUB: not implemented"; return nil }

func (opts options) validate() error { _ = "STUB: not implemented"; return nil }

// If username or password is specified, both must be provided.

// If username file or password file is specified, both must be provided.

// Setting credentials both inline and using files will result in a panic
// from the container, so we short circuit here.

// parseUsername will return either the username provided by WithUsername or
// from the local file specified by WithUsernameFile. If both are provided, this
// function will return an error. If neither is provided, an empty string is
// returned.
func (opts options) parseUsername() (string, error) { _ = "STUB: not implemented"; return "", nil }

// parsePassword will return either the password provided by WithPassword or
// from the local file specified by WithPasswordFile. If both are provided, this
// function will return an error. If neither is provided, an empty string is
// returned.
func (opts options) parsePassword() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Option is an option for the Atlas Local container.
type Option func(*options) error

// Customize is a NOOP. It's defined to satisfy the testcontainers.ContainerCustomizer interface.
func (o Option) Customize(*testcontainers.GenericContainerRequest) error {
	_ = "STUB: not implemented"
	// NOOP to satisfy interface.
	return nil
}

// WithUsername sets the MongoDB root username by setting the
// MONGODB_INITDB_ROOT_USERNAME environment variable.
func WithUsername(username string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPassword sets the MongoDB root password by setting the the
// MONGODB_INITDB_ROOT_PASSWORD environment variable.
func WithPassword(password string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUsernameFile mounts a local file as the MongoDB root username secret at
// /run/secrets/mongo-root-username and sets MONGODB_INITDB_ROOT_USERNAME_FILE.
// The path must be absolute and exist; no-op if empty.
func WithUsernameFile(usernameFile string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Must be an absolute path.

// Must exist and be a file.

// WithPasswordFile mounts a local file as the MongoDB root password secret at
// /run/secrets/mongo-root-password and sets MONGODB_INITDB_ROOT_PASSWORD_FILE.
// Path must be absolute and an existing file; no-op if empty.
func WithPasswordFile(passwordFile string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Must be an absolute path.

// Must exist and be a file.

// WithNoTelemetry opts out of telemetry for the MongoDB Atlas Local
// container by setting the DO_NOT_TRACK environment variable to 1.
func WithNoTelemetry() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithInitDatabase sets MONGODB_INITDB_DATABASE environment variable so the
// init scripts and the default connection string target the specified database
// instead of the default "test" database.
func WithInitDatabase(database string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithInitScripts mounts a directory containing .sh/.js init scripts into
// /docker-entrypoint-initdb.d so they run in alphabetical order on startup. If
// called multiple times, this function removes any prior init-scripts bind and
// uses only the latest on specified.
func WithInitScripts(scriptsDir string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Make shell scripts executable.

// WithMongotLogStdout writes to /dev/stdout inside the container. See
// (*Container).ReadMongotLogs to read the logs locally.
func WithMongotLogToStdout() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMongotLogToStderr writes to /dev/stderr inside the container. See
// (*Container).ReadMongotLogs to read the logs locally.
func WithMongotLogToStderr() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMongotLogFile writes the mongot logs to /tmp/mongot.log inside the
// container. See (*Container).ReadMongotLogs to read the logs locally.
func WithMongotLogFile() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRunnerLogToStdout writes to /dev/stdout inside the container. See
// (*Container).ReadRunnerLogs to read the logs locally.
func WithRunnerLogToStdout() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRunnerLogToStderr writes to /dev/stderr inside the container. See
// (*Container).ReadRunnerLogs to read the logs locally.
func WithRunnerLogToStderr() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRunnerLogFile writes the runner logs to /tmp/runner.log inside the
// container. See (*Container).ReadRunnerLogs to read the logs locally.
func WithRunnerLogFile() Option { _ = "STUB: not implemented"; return *new(Option) }
