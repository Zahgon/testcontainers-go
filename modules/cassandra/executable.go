package cassandra

import (
	"github.com/testcontainers/testcontainers-go"
)

type initScript struct {
	testcontainers.ExecOptions
	File string
}

func (i initScript) AsCommand() []string { _ = "STUB: not implemented"; return nil }
