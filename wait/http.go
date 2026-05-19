package wait

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/moby/moby/api/types/network"
)

// Implement interface
var (
	_ Strategy        = (*HTTPStrategy)(nil)
	_ StrategyTimeout = (*HTTPStrategy)(nil)
)

type HTTPStrategy struct {
	// all Strategies should have a startupTimeout to avoid waiting infinitely
	timeout *time.Duration

	// additional properties
	Port                   network.Port
	Path                   string
	StatusCodeMatcher      func(status int) bool
	ResponseMatcher        func(body io.Reader) bool
	UseTLS                 bool
	AllowInsecure          bool
	TLSConfig              *tls.Config // TLS config for HTTPS
	Method                 string      // http method
	Body                   io.Reader   // http request body
	Headers                map[string]string
	ResponseHeadersMatcher func(headers http.Header) bool
	PollInterval           time.Duration
	UserInfo               *url.Userinfo
	ForceIPv4LocalHost     bool
}

// NewHTTPStrategy constructs an HTTP strategy waiting on port 80 and status code 200
func NewHTTPStrategy(path string) *HTTPStrategy { _ = "STUB: not implemented"; return nil }

func defaultStatusCodeMatcher(status int) bool { _ = "STUB: not implemented"; return false }

// fluent builders for each property
// since go has neither covariance nor generics, the return type must be the type of the concrete implementation
// this is true for all properties, even the "shared" ones like startupTimeout

// WithStartupTimeout can be used to change the default startup timeout
func (ws *HTTPStrategy) WithStartupTimeout(timeout time.Duration) *HTTPStrategy {
	_ = "STUB: not implemented"
	return nil
}

// WithPort set the port to wait for.
// Default is the lowest numbered port.
func (ws *HTTPStrategy) WithPort(port string) *HTTPStrategy { _ = "STUB: not implemented"; return nil }

func (ws *HTTPStrategy) WithStatusCodeMatcher(statusCodeMatcher func(status int) bool) *HTTPStrategy {
	_ = "STUB: not implemented"
	return nil
}

func (ws *HTTPStrategy) WithResponseMatcher(matcher func(body io.Reader) bool) *HTTPStrategy {
	_ = "STUB: not implemented"
	return nil
}

func (ws *HTTPStrategy) WithTLS(useTLS bool, tlsconf ...*tls.Config) *HTTPStrategy {
	_ = "STUB: not implemented"
	return nil
}

func (ws *HTTPStrategy) WithAllowInsecure(allowInsecure bool) *HTTPStrategy {
	_ = "STUB: not implemented"
	return nil
}

func (ws *HTTPStrategy) WithMethod(method string) *HTTPStrategy {
	_ = "STUB: not implemented"
	return nil
}

func (ws *HTTPStrategy) WithBody(reqdata io.Reader) *HTTPStrategy {
	_ = "STUB: not implemented"
	return nil
}

func (ws *HTTPStrategy) WithHeaders(headers map[string]string) *HTTPStrategy {
	_ = "STUB: not implemented"
	return nil
}

func (ws *HTTPStrategy) WithResponseHeadersMatcher(matcher func(http.Header) bool) *HTTPStrategy {
	_ = "STUB: not implemented"
	return nil
}

func (ws *HTTPStrategy) WithBasicAuth(username, password string) *HTTPStrategy {
	_ = "STUB: not implemented"
	return nil
}

// WithPollInterval can be used to override the default polling interval of 100 milliseconds
func (ws *HTTPStrategy) WithPollInterval(pollInterval time.Duration) *HTTPStrategy {
	_ = "STUB: not implemented"
	return nil
}

// WithForcedIPv4LocalHost forces usage of localhost to be ipv4 127.0.0.1
// to avoid ipv6 docker bugs https://github.com/moby/moby/issues/42442 https://github.com/moby/moby/issues/42375
func (ws *HTTPStrategy) WithForcedIPv4LocalHost() *HTTPStrategy {
	_ = "STUB: not implemented"
	return nil
}

// ForHTTP is a convenience method similar to Wait.java
// https://github.com/testcontainers/testcontainers-java/blob/1d85a3834bd937f80aad3a4cec249c027f31aeb4/core/src/main/java/org/testcontainers/containers/wait/strategy/Wait.java
func ForHTTP(path string) *HTTPStrategy { _ = "STUB: not implemented"; return nil }

func (ws *HTTPStrategy) Timeout() *time.Duration {
	_ = "STUB: not implemented"

	// String returns a human-readable description of the wait strategy.
	return nil
}

func (ws *HTTPStrategy) String() string { _ = "STUB: not implemented"; return "" }

// WaitUntilReady implements Strategy.WaitUntilReady
func (ws *HTTPStrategy) WaitUntilReady(ctx context.Context, target StrategyTarget) error {
	_ = "STUB: not implemented"
	return nil
}

// to avoid ipv6 docker bugs https://github.com/moby/moby/issues/42442 https://github.com/moby/moby/issues/42375

// No specific port requested; inspect container to find lowest exposed TCP port.
// We wait one polling interval before we grab the ports
// otherwise they might not be bound yet on startup.

// Port should now be bound so just continue.

// Find the lowest numbered exposed tcp port.

// Specific port requested; use MappedPort to resolve it.

// cache the body into a byte-slice so that it can be iterated over multiple times
