package mongodb

// mongoCli is cli to interact with MongoDB. If username and password are provided
// it will use credentials to authenticate.
type mongoCli struct {
	mongoshBaseCmd string
	mongoBaseCmd   string
}

func newMongoCli(username string, password string) mongoCli {
	_ = "STUB: not implemented"
	return *new(mongoCli)
}

func (m mongoCli) eval(command string, args ...any) []string { _ = "STUB: not implemented"; return nil }
