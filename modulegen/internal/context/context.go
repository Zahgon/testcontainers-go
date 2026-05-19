package context

// Context is the context for the module generation.
// It provides the necessary information to generate the module or example,
// such as the different paths to the files and directories.
type Context struct {
	RootDir string
}

// DependabotConfigFile returns, from the root directory, the relative path
// to the dependabot config file, "/.github/dependabot.yml".
func (ctx Context) DependabotConfigFile() string { _ = "STUB: not implemented"; return "" }

// DocsDir returns, from the root directory, the relative path to the docs directory, "/docs".
func (ctx Context) DocsDir() string { _ = "STUB: not implemented"; return "" }

// ExamplesDir returns, from the root directory, the relative path to the examples directory, "/examples".
func (ctx Context) ExamplesDir() string { _ = "STUB: not implemented"; return "" }

// ExamplesDocsDir returns, from the docs directory, the relative path to the examples directory, "/docs/examples".
func (ctx Context) ExamplesDocsDir() string { _ = "STUB: not implemented"; return "" }

// ModulesDir returns, from the root directory, the relative path to the modules directory, "/modules".
func (ctx Context) ModulesDir() string { _ = "STUB: not implemented"; return "" }

// ModulesDocsDir returns, from the docs directory, the relative path to the modules directory, "/docs/modules".
func (ctx Context) ModulesDocsDir() string { _ = "STUB: not implemented"; return "" }

// GithubDir returns, from the root directory, the relative path to the github directory, "/.github".
func (ctx Context) GithubDir() string { _ = "STUB: not implemented"; return "" }

// GithubWorkflowsDir returns, from the github directory, the relative path to the workflows directory, "/.github/workflows".
func (ctx Context) GithubWorkflowsDir() string { _ = "STUB: not implemented"; return "" }

// GoModFile returns, from the root directory, the relative path to the go.mod file, "/go.mod".
func (ctx Context) GoModFile() string { _ = "STUB: not implemented"; return "" }

// getModulesByBaseDir returns, from the root directory, the relative paths to the modules or examples,
// depending on the base directory.
func (ctx Context) getModulesByBaseDir(baseDir string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getMarkdownsFromDir returns, from the docs directory, the relative paths to the markdown files,
// depending on the base directory.
func (ctx Context) getMarkdownsFromDir(baseDir string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetExamples returns, from the root directory, a slice with the names of the examples
// that are located in the "examples" directory.
func (ctx Context) GetExamples() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// GetModules returns, from the root directory, a slice with the names of the modules
// that are located in the "modules" directory.
func (ctx Context) GetModules() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// GetExamplesDocs returns, from the docs directory, a slice with the names of the markdown files
// that are located in the "docs/examples" directory.
func (ctx Context) GetExamplesDocs() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// GetModulesDocs returns, from the docs directory, a slice with the names of the markdown files
// that are located in the "docs/modules" directory.
func (ctx Context) GetModulesDocs() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// MkdocsConfigFile returns, from the root directory, the relative path to the mkdocs config file, "/mkdocs.yml".
func (ctx Context) MkdocsConfigFile() string { _ = "STUB: not implemented"; return "" }

// VSCodeWorkspaceFile returns, from the root directory, the relative path to the vscode workspace file, "/.vscode/.testcontainers-go.code-workspace".
func (ctx Context) VSCodeWorkspaceFile() string { _ = "STUB: not implemented"; return "" }

// New returns a new Context with the given root directory.
func New(dir string) Context { _ = "STUB: not implemented"; return *new(Context) }

// GetRootContext returns a new Context with the current working directory.
func GetRootContext() (Context, error) { _ = "STUB: not implemented"; return *new(Context), nil }
