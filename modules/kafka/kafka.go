package kafka

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const publicPort = "9093/tcp"
const (
	starterScript = "/usr/sbin/testcontainers_start.sh"

	// starterScript {
	starterScriptContent = `#!/bin/bash
source /etc/confluent/docker/bash-config
export KAFKA_ADVERTISED_LISTENERS=%s,BROKER://%s:9092
echo Starting Kafka KRaft mode
sed -i '/KAFKA_ZOOKEEPER_CONNECT/d' /etc/confluent/docker/configure
echo 'kafka-storage format --ignore-formatted -t "$(kafka-storage random-uuid)" -c /etc/kafka/kafka.properties' >> /etc/confluent/docker/configure
echo '' > /etc/confluent/docker/ensure
/etc/confluent/docker/configure
/etc/confluent/docker/launch`
	// }
)

// KafkaContainer represents the Kafka container type used in the module
type KafkaContainer struct {
	testcontainers.Container
	ClusterID string
}

// Deprecated: use Run instead
// RunContainer creates an instance of the Kafka container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*KafkaContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run creates an instance of the Kafka container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*KafkaContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// envVars {

// }

// this CMD will wait for the starter script to be copied into the container and then execute it

// Use a single hook to copy the starter script and wait for
// the Kafka server to be ready. This prevents the wait running
// if the starter script fails to copy.

// 1. copy the starter script into the container

// 2. wait for the Kafka server to be ready

// configure the controller quorum voters after all the options have been applied

// Inspect the container to get the CLUSTER_ID environment variable

// copyStarterScript copies the starter script into the container.
func copyStarterScript(ctx context.Context, c testcontainers.Container) error {
	_ = "STUB: not implemented"
	return nil
}

func WithClusterID(clusterID string) testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// Brokers retrieves the broker connection strings from Kafka with only one entry,
// defined by the exposed public port.
func (kc *KafkaContainer) Brokers(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// configureControllerQuorumVoters returns an option that sets the quorum voters for the controller.
// For that, it will check if there are any network aliases defined for the container and use the
// first alias in the first network. Else, it will use localhost.
func configureControllerQuorumVoters() testcontainers.CustomizeRequestOption {
	_ = "STUB: not implemented"
	return *new(testcontainers.CustomizeRequestOption)
}

// }

// validateKRaftVersion validates if the image version is compatible with KRaft mode,
// which is available since version 7.0.0.
func validateKRaftVersion(fqName string) error { _ = "STUB: not implemented"; return nil }

// do not validate if the image is not the official one.
// not raising an error here, letting the image start and
// eventually evaluate an error if it exists.

// semver requires the version to start with a "v"

// remove the architecture suffix

// version < v7.4.0
