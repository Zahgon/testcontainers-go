package couchbase

import "github.com/testcontainers/testcontainers-go"

// Option is a function that configures the Couchbase container.
//
// Deprecated: Use the With* functions instead.
type Option func(*Config)

// Config is the configuration for the Couchbase container, that will be stored in the container itself.
type Config struct {
	enabledServices  []Service
	username         string
	password         string
	isEnterprise     bool
	buckets          []bucket
	imageName        string
	indexStorageMode indexStorageMode
}

// WithEnterpriseService enables the eventing service in the container.
// Only available in the Enterprise Edition of Couchbase Server.
//
// Deprecated: Use WithServiceEventing instead.
func WithEventingService() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAnalyticsService enables the analytics service in the container.
// Only available in the Enterprise Edition of Couchbase Server.
//
// Deprecated: Use [WithServiceAnalytics] instead.
func WithAnalyticsService() Option { _ = "STUB: not implemented"; return *new(Option) }

type credentialsCustomizer struct {
	username string
	password string
}

func (c credentialsCustomizer) Customize(_ *testcontainers.GenericContainerRequest) error {
	_ = "STUB: not implemented"
	// NOOP, we want to simply transfer the credentials to the container
	return nil
}

// WithAdminCredentials sets the username and password for the administrator user.
func WithAdminCredentials(username, password string) credentialsCustomizer {
	_ = "STUB: not implemented"
	return *new(credentialsCustomizer)
}

// WithCredentials sets the username and password for the administrator user.
//
// Deprecated: Use [WithAdminCredentials] instead.
func WithCredentials(username, password string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithBucket adds a bucket to the container.
//
// Deprecated: Use [WithBuckets] instead.
func WithBucket(bucket bucket) Option { _ = "STUB: not implemented"; return *new(Option) }

type bucketCustomizer struct {
	buckets []bucket
}

func (c bucketCustomizer) Customize(_ *testcontainers.GenericContainerRequest) error {
	_ = "STUB: not implemented"
	// NOOP, we want to simply transfer the buckets to the container
	return nil
}

// WithBucket adds buckets to the couchbase container
func WithBuckets(bucket ...bucket) bucketCustomizer {
	_ = "STUB: not implemented"
	return *new(bucketCustomizer)
}

// WithImageName allows to override the default image name.
//
// Deprecated: Use testcontainers.WithImage instead.
func WithImageName(imageName string) Option { _ = "STUB: not implemented"; return *new(Option) }

type indexStorageCustomizer struct {
	mode indexStorageMode
}

func (c indexStorageCustomizer) Customize(_ *testcontainers.GenericContainerRequest) error {
	_ = "STUB: not implemented"
	// NOOP, we want to simply transfer the index storage mode to the container
	return nil
}

// WithBucket adds buckets to the couchbase container
func WithIndexStorage(indexStorageMode indexStorageMode) indexStorageCustomizer {
	_ = "STUB: not implemented"
	return *new(indexStorageCustomizer)
}

// WithIndexStorageMode sets the storage mode to be used in the cluster.
//
// Deprecated: Use [WithIndexStorage] instead.
func WithIndexStorageMode(indexStorageMode indexStorageMode) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
