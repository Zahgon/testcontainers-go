package vscode

// Config is a struct that represents the vscode workspace file.
type Config struct {
	Folders []Folder `json:"folders"`
}

// Folder is a struct that represents a folder in the vscode workspace.
type Folder struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func newConfig(examples []string, modules []string) *Config { _ = "STUB: not implemented"; return nil }
