package couchbase

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

//nolint:staticcheck //FIXME
const (
	// containerPorts {

	MGMT_PORT     = "8091"
	MGMT_SSL_PORT = "18091"

	VIEW_PORT     = "8092"
	VIEW_SSL_PORT = "18092"

	QUERY_PORT     = "8093"
	QUERY_SSL_PORT = "18093"

	SEARCH_PORT     = "8094"
	SEARCH_SSL_PORT = "18094"

	ANALYTICS_PORT     = "8095"
	ANALYTICS_SSL_PORT = "18095"

	EVENTING_PORT     = "8096"
	EVENTING_SSL_PORT = "18096"

	KV_PORT     = "11210"
	KV_SSL_PORT = "11207"

	// }
)

// initialServices is the list of services that are enabled by default
var initialServices = []Service{kv, query, search, index}

type clusterInit func(context.Context) error

// CouchbaseContainer represents the Couchbase container type used in the module
type CouchbaseContainer struct {
	testcontainers.Container
	config *Config
}

// Deprecated: use Run instead
// RunContainer creates an instance of the Couchbase container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*CouchbaseContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the Couchbase container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*CouchbaseContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Process custom options first to extract config
// Start with initial services

// Transfer custom options to the config

// If the option is a bucketCustomizer, we need to add the buckets to the request

// If the option is a serviceCustomizer, we need to append the service

// If the option is a indexStorageCustomizer, we need to set the index storage mode

// If the option is a credentialsCustomizer, we need to set the credentials

// Build moduleOpts with defaults

// Append all customizers (initial services + user opts)

// StartContainer creates an instance of the Couchbase container type
//
// Deprecated: use RunContainer instead
func StartContainer(ctx context.Context, opts ...Option) (*CouchbaseContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConnectionString returns the connection string to connect to the Couchbase container instance.
// It returns a string with the format couchbase://<host>:<port>
func (c *CouchbaseContainer) ConnectionString(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Username returns the username of the Couchbase administrator.
func (c *CouchbaseContainer) Username() string { _ = "STUB: not implemented"; return "" }

// Password returns the password of the Couchbase administrator.
func (c *CouchbaseContainer) Password() string { _ = "STUB: not implemented"; return "" }

func (c *CouchbaseContainer) initCluster(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CouchbaseContainer) waitUntilNodeIsOnline(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CouchbaseContainer) initializeIsEnterprise(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CouchbaseContainer) renameNode(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CouchbaseContainer) initializeServices(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CouchbaseContainer) setMemoryQuotas(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CouchbaseContainer) configureAdminUser(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CouchbaseContainer) configureExternalPorts(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CouchbaseContainer) configureIndexer(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CouchbaseContainer) waitUntilAllNodesAreHealthy(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CouchbaseContainer) createBuckets(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CouchbaseContainer) isPrimaryIndexOnline(ctx context.Context, bucket bucket) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CouchbaseContainer) createPrimaryIndex(ctx context.Context, bucket bucket) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CouchbaseContainer) isQueryKeyspacePresent(ctx context.Context, bucket bucket) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CouchbaseContainer) waitForAllServicesEnabled(ctx context.Context, bucket bucket) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CouchbaseContainer) createBucket(ctx context.Context, bucket bucket) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CouchbaseContainer) doHTTPRequest(ctx context.Context, port, path, method string, body map[string]string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// retry with backoff

func (c *CouchbaseContainer) getURL(ctx context.Context, port, path string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *CouchbaseContainer) getInternalIPAddress(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *CouchbaseContainer) getEnabledServices() string { _ = "STUB: not implemented"; return "" }

func (c *CouchbaseContainer) checkAllServicesEnabled(rawConfig []byte) bool {
	_ = "STUB: not implemented"
	return false
}

type serviceCustomizer struct {
	enabledService Service
}

func (c serviceCustomizer) Customize(req *testcontainers.GenericContainerRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// withService creates a serviceCustomizer for the given service.
// It's private to prevent users from creating other services than the Analytics and Eventing services.
func withService(service Service) serviceCustomizer {
	_ = "STUB: not implemented"
	return *new(serviceCustomizer)
}

// WithServiceAnalytics enables the Analytics service.
func WithServiceAnalytics() serviceCustomizer {
	_ = "STUB: not implemented"
	return *new(serviceCustomizer)
}

// WithServiceEventing enables the Eventing service.
func WithServiceEventing() serviceCustomizer {
	_ = "STUB: not implemented"
	return *new(serviceCustomizer)
}

func contains(services []Service, service Service) bool { _ = "STUB: not implemented"; return false }
