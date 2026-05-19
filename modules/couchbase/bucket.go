package couchbase

type bucket struct {
	name              string
	flushEnabled      bool
	queryPrimaryIndex bool
	quota             int
	numReplicas       int
}

// NewBucket creates a new bucket with the given name, using default values for all other fields.
func NewBucket(name string) bucket { _ = "STUB: not implemented"; return *new(bucket) }

// WithReplicas sets the number of replicas for this bucket. The minimum value is 0 and the maximum value is 3.
func (b bucket) WithReplicas(numReplicas int) bucket {
	_ = "STUB: not implemented"
	return *new(bucket)
}

// WithFlushEnabled sets whether the bucket should be flushed when the container is stopped.
func (b bucket) WithFlushEnabled(flushEnabled bool) bucket {
	_ = "STUB: not implemented"
	return *new(bucket)
}

// WithQuota sets the bucket quota in megabytes. The minimum value is 100 MB.
func (b bucket) WithQuota(quota int) bucket {
	_ = "STUB: not implemented"
	// Couchbase Server 6.5.0 has a minimum of 100 MB
	return *new(bucket)
}

// WithPrimaryIndex sets whether the primary index should be created for this bucket.
func (b bucket) WithPrimaryIndex(primaryIndex bool) bucket {
	_ = "STUB: not implemented"
	return *new(bucket)
}
