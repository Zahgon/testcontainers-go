package core

const (
	// LabelBase is the base label for all testcontainers labels.
	LabelBase = "org.testcontainers"

	// LabelLang specifies the language which created the test container.
	LabelLang = LabelBase + ".lang"

	// LabelReaper identifies the container as a reaper.
	LabelReaper = LabelBase + ".reaper"

	// LabelRyuk identifies the container as a ryuk.
	LabelRyuk = LabelBase + ".ryuk"

	// LabelSessionID specifies the session ID of the container.
	LabelSessionID = LabelBase + ".sessionId"

	// LabelVersion specifies the version of testcontainers which created the container.
	LabelVersion = LabelBase + ".version"

	// LabelReap specifies the container should be reaped by the reaper.
	LabelReap = LabelBase + ".reap"
)

// DefaultLabels returns the standard set of labels which
// includes LabelSessionID if the reaper is enabled.
func DefaultLabels(sessionID string) map[string]string { _ = "STUB: not implemented"; return nil }

// AddDefaultLabels adds the default labels for sessionID to target.
func AddDefaultLabels(sessionID string, target map[string]string) {
	_ = "STUB: not implemented"
	return
}

// MergeCustomLabels sets labels from src to dst.
// If a key in src has [LabelBase] prefix returns an error.
// If dst is nil returns an error.
func MergeCustomLabels(dst, src map[string]string) error { _ = "STUB: not implemented"; return nil }
