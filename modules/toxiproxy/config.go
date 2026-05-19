package toxiproxy

// proxy represents a single proxy configuration in the toxiproxy config file
type proxy struct {
	Name     string `json:"name"`
	Listen   string `json:"listen"`
	Upstream string `json:"upstream"`
	Enabled  bool   `json:"enabled"`

	listenIP     string
	upstreamIP   string
	listenPort   int
	upstreamPort int
}

// sanitize is a helper function that returns another proxy
// in which all the fields have been extracted from the
// string representation of the upstream and listen fields.
func (p *proxy) sanitize() error { _ = "STUB: not implemented"; return nil }

// newProxy creates a new proxy configuration with default values
func newProxy(name string, upstream string) (*proxy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
