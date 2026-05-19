package dex

import (
	"log/slog"
	"strings"

	"github.com/testcontainers/testcontainers-go"
)

// slogConsumer adapts a *slog.Logger to testcontainers.LogConsumer. Dex
// emits logfmt-style lines (key=value); we parse level + msg and preserve
// remaining fields as slog attrs.
//
// Stderr lines are promoted to at least slog.LevelWarn because Dex writes
// runtime errors there.
type slogConsumer struct {
	logger *slog.Logger
}

// Compile check: *slogConsumer implements testcontainers.LogConsumer.
var _ testcontainers.LogConsumer = (*slogConsumer)(nil)

func newSlogConsumer(l *slog.Logger) *slogConsumer { _ = "STUB: not implemented"; return nil }

// Accept implements testcontainers.LogConsumer.
func (s *slogConsumer) Accept(l testcontainers.Log) { _ = "STUB: not implemented"; return }

// accept is the testable inner method. It takes the raw content string and
// the testcontainers log type (STDOUT/STDERR) so unit tests don't need to
// construct a real tc-go Log value.
func (s *slogConsumer) accept(content, logType string) { _ = "STUB: not implemented"; return }

// logfmtUnescaper unescapes the two backslash sequences logfmt allows
// inside quoted values: \" → " and \\ → \.
var logfmtUnescaper = strings.NewReplacer(`\\`, `\`, `\"`, `"`)

// parseLogfmt is a minimal logfmt parser — enough for Dex's default
// format (level=... msg=...). Unknown keys become slog attrs. Quoted
// values are unquoted and have their \" / \\ escapes expanded.
func parseLogfmt(line string) (slog.Level, string, []slog.Attr) {
	_ = "STUB: not implemented"
	return *new(slog.Level), "", nil
}

// Dex timestamps are redundant — slog adds its own.

type kv struct{ key, val string }

func tokenizeLogfmt(line string) []kv { _ = "STUB: not implemented"; return nil }

// skip '='

func mapLevel(s string) slog.Level { _ = "STUB: not implemented"; return *new(slog.Level) }
