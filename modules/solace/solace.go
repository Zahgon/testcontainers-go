package solace

import (
	"context"
	_ "embed"

	"github.com/testcontainers/testcontainers-go"
)

//go:embed mounts/solace.script.tpl
var customScriptTpl string

// Container represents a Solace container with additional settings
type Container struct {
	testcontainers.Container
	settings options
}

// Run starts a Solace container with the provided image and options
//
// The container requires specific ulimits to start successfully:
//   - nofile: 1048576 (number of open files)
//   - core: -1 (for core dumps)
//   - memlock: -1 (for memory locking)
//
// See https://docs.solace.com for more information.
// - https://docs.solace.com/Software-Broker/Managing-Core-Files.htm
// - https://docs.solace.com/Software-Broker/Container-Tasks/Config-Container-Storage.htm?Highlight=ulimit
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*Container, error) {
	_ = "STUB: not implemented"
	// Override default options with provided ones
	return nil, nil
}

// Build wait strategies

// Primary wait strategy for Solace to be ready

// Add port-based wait strategies for each service

// Where disk space permits, Solace docs recommends the core file size "rlimit" for event brokers to be "unlimited".

// Render CLI script for queue/topic configuration

// Copy the CLI script directly to the container

// Execute the script

// BrokerURLFor returns the origin URL for a given service
func (c *Container) BrokerURLFor(ctx context.Context, service Service) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Username returns the username configured for the Solace container
func (c *Container) Username() string { _ = "STUB: not implemented"; return "" }

// Password returns the password configured for the Solace container
func (c *Container) Password() string { _ = "STUB: not implemented"; return "" }

// VPN returns the VPN name configured for the Solace container
func (c *Container) VPN() string { _ = "STUB: not implemented"; return "" }

// renderSolaceScript renders the Solace CLI configuration script based on the provided settings.
// Reference: https://docs.solace.com/Admin-Ref/CLI-Reference/VMR_CLI_Commands.html
func renderSolaceScript(opts options) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Create template data structure
