package neo4j

import (
	"github.com/testcontainers/testcontainers-go"
)

type LabsPlugin string

const (
	// labsPlugins {

	Apoc             LabsPlugin = "apoc"
	ApocCore         LabsPlugin = "apoc-core"
	Bloom            LabsPlugin = "bloom"
	GraphDataScience LabsPlugin = "graph-data-science"
	NeoSemantics     LabsPlugin = "n10s"
	Streams          LabsPlugin = "streams"

	// }
)

// WithoutAuthentication disables authentication.
func WithoutAuthentication() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithAdminPassword sets the admin password for the default account
// An empty string disables authentication.
// The default password is "password".
func WithAdminPassword(adminPassword string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithLabsPlugin registers one or more Neo4jLabsPlugin for download and server startup.
// There might be plugins not supported by your selected version of Neo4j.
func WithLabsPlugin(plugins ...LabsPlugin) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithNeo4jSetting adds Neo4j a single configuration setting to the container.
// The setting can be added as in the official Neo4j configuration, the function automatically translates the setting
// name (e.g. dbms.tx_log.rotation.size) into the format required by the Neo4j container.
// This function can be called multiple times. A warning is emitted if a key is overwritten.
// See WithNeo4jSettings to add multiple settings at once
// Note: credentials must be configured with WithAdminPassword
func WithNeo4jSetting(key, value string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithNeo4jSettings adds multiple Neo4j configuration settings to the container.
// The settings can be added as in the official Neo4j configuration, the function automatically translates each setting
// name (e.g. dbms.tx_log.rotation.size) into the format required by the Neo4j container.
// This function can be called multiple times. A warning is emitted if a key is overwritten.
// See WithNeo4jSetting to add a single setting
// Note: credentials must be configured with WithAdminPassword
func WithNeo4jSettings(settings map[string]string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// Deprecated: use testcontainers.WithLogger instead
//
// WithLogger sets a custom logger to be used by the container
// Consider calling this before other "With functions" as these may generate logs
var (
	WithLogger                                       = testcontainers.WithLogger
	_          testcontainers.CustomizeRequestOption = WithLogger(nil).Customize
)

func addSetting(req *testcontainers.GenericContainerRequest, key string, newVal string) error {
	_ = "STUB: not implemented"
	return nil
}

// make sure AUTH is not overwritten by a setting

func formatNeo4jConfig(name string) string { _ = "STUB: not implemented"; return "" }

// WithAcceptCommercialLicenseAgreement sets the environment variable
// NEO4J_ACCEPT_LICENSE_AGREEMENT to "yes", indicating that the user accepts
// the commercial licence agreement of Neo4j Enterprise Edition. The license
// agreement is available at https://neo4j.com/terms/licensing/.
func WithAcceptCommercialLicenseAgreement() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// WithAcceptEvaluationLicenseAgreement sets the environment variable
// NEO4J_ACCEPT_LICENSE_AGREEMENT to "eval", indicating that the user accepts
// the evaluation agreement of Neo4j Enterprise Edition. The evaluation
// agreement is available at https://neo4j.com/terms/enterprise_us/. Please
// read the terms of the evaluation agreement before you accept.
func WithAcceptEvaluationLicenseAgreement() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}
