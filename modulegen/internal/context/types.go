package context

// TestcontainersModule represents a module or an example of the Testcontainers project.
// It defines the common properties for both modules and examples,
// such as the image name, if it's a module or an example, the name and the title of the module.
// It provides methods to be used during the code generation of the module or example.
type TestcontainersModule struct {
	// Image fully qualified name of the Docker image
	Image string

	// IsModule if true, the module will be generated as a Go module, otherwise an example
	IsModule bool

	// Name name of the module
	Name string

	// TitleName title of the name: m.g. "mongodb" -> "MongoDB"
	TitleName string

	// TCVersion Testcontainers for Go version
	TCVersion string
}

// ContainerName returns the name of the container, which is the lower-cased title of the example
// If the title is set, it will be used instead of the name
func (m *TestcontainersModule) ContainerName() string {
	_ = "STUB: not implemented"

	// Entrypoint returns the name of the entrypoint function, which is the lower-cased title of the example
	// If the example is a module, the entrypoint will be "Run"
	return ""
}

func (m *TestcontainersModule) Entrypoint() string { _ = "STUB: not implemented"; return "" }

// Lower returns the lower-cased name of the module
func (m *TestcontainersModule) Lower() string { _ = "STUB: not implemented"; return "" }

// ParentDir returns the parent directory of the module: "modules" or "examples"
func (m *TestcontainersModule) ParentDir() string { _ = "STUB: not implemented"; return "" }

// Title returns the title of the module. If it's not set,
// it will be the title of the lower-cased name.
func (m *TestcontainersModule) Title() string { _ = "STUB: not implemented"; return "" }

// Type returns the type of the module: "module" or "example"
func (m *TestcontainersModule) Type() string { _ = "STUB: not implemented"; return "" }

// Validate validates the module name and title
func (m *TestcontainersModule) Validate() error { _ = "STUB: not implemented"; return nil }
