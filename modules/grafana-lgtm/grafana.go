package grafanalgtm

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

const (
	GrafanaPort    = "3000/tcp"
	LokiPort       = "3100/tcp"
	TempoPort      = "3200/tcp"
	OtlpGrpcPort   = "4317/tcp"
	OtlpHttpPort   = "4318/tcp" //nolint:revive,staticcheck //FIXME
	PrometheusPort = "9090/tcp"
)

// GrafanaLGTMContainer represents the Grafana LGTM container type used in the module
type GrafanaLGTMContainer struct {
	testcontainers.Container
}

// Run creates an instance of the Grafana LGTM container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*GrafanaLGTMContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// return the container instance to allow the caller to clean up

// WithAdminCredentials sets the admin credentials for the Grafana LGTM container
func WithAdminCredentials(user, password string) testcontainers.ContainerCustomizer {
	_ = "STUB: not implemented"
	return *new(testcontainers.ContainerCustomizer)
}

// LokiEndpoint returns the Loki endpoint
func (c *GrafanaLGTMContainer) LokiEndpoint(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// MustLokiEndpoint returns the Loki endpoint or panics if an error occurs
func (c *GrafanaLGTMContainer) MustLokiEndpoint(ctx context.Context) string {
	_ = "STUB: not implemented"
	return ""
}

// TempoEndpoint returns the Tempo endpoint
func (c *GrafanaLGTMContainer) TempoEndpoint(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// MustTempoEndpoint returns the Tempo endpoint or panics if an error occurs
func (c *GrafanaLGTMContainer) MustTempoEndpoint(ctx context.Context) string {
	_ = "STUB: not implemented"
	return ""
}

// HttpEndpoint returns the HTTP URL
//
//nolint:revive,staticcheck //FIXME
func (c *GrafanaLGTMContainer) HttpEndpoint(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// MustHttpEndpoint returns the HTTP endpoint or panics if an error occurs
//
//nolint:revive,staticcheck //FIXME
func (c *GrafanaLGTMContainer) MustHttpEndpoint(ctx context.Context) string {
	_ = "STUB: not implemented"
	return ""
}

// OtlpHttpEndpoint returns the OTLP HTTP endpoint
//
//nolint:revive,staticcheck //FIXME
func (c *GrafanaLGTMContainer) OtlpHttpEndpoint(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// MustOtlpHttpEndpoint returns the OTLP HTTP endpoint or panics if an error occurs
//
//nolint:revive,staticcheck //FIXME
func (c *GrafanaLGTMContainer) MustOtlpHttpEndpoint(ctx context.Context) string {
	_ = "STUB: not implemented"
	return ""
}

// OtlpGrpcEndpoint returns the OTLP gRPC endpoint
func (c *GrafanaLGTMContainer) OtlpGrpcEndpoint(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// MustOtlpGrpcEndpoint returns the OTLP gRPC endpoint or panics if an error occurs
func (c *GrafanaLGTMContainer) MustOtlpGrpcEndpoint(ctx context.Context) string {
	_ = "STUB: not implemented"
	return ""
}

// PrometheusHttpEndpoint returns the Prometheus HTTP endpoint
//
//nolint:revive,staticcheck //FIXME
func (c *GrafanaLGTMContainer) PrometheusHttpEndpoint(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// MustPrometheusHttpEndpoint returns the Prometheus HTTP endpoint or panics if an error occurs
//
//nolint:revive,staticcheck //FIXME
func (c *GrafanaLGTMContainer) MustPrometheusHttpEndpoint(ctx context.Context) string {
	_ = "STUB: not implemented"
	return ""
}
