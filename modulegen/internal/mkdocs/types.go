package mkdocs

type Config struct {
	SiteName string   `yaml:"site_name"`
	SiteURL  string   `yaml:"site_url"`
	Plugins  []string `yaml:"plugins"`
	Theme    struct {
		Name      string `yaml:"name"`
		CustomDir string `yaml:"custom_dir"`
		Palette   struct {
			Scheme string `yaml:"scheme"`
		} `yaml:"palette"`
		Font struct {
			Text string `yaml:"text"`
			Code string `yaml:"code"`
		} `yaml:"font"`
		Logo    string `yaml:"logo"`
		Favicon string `yaml:"favicon"`
	} `yaml:"theme"`
	ExtraCSS           []string `yaml:"extra_css"`
	ExtraJavascript    []string `yaml:"extra_javascript"`
	RepoName           string   `yaml:"repo_name"`
	RepoURL            string   `yaml:"repo_url"`
	MarkdownExtensions []any    `yaml:"markdown_extensions"`
	Nav                []struct {
		Home               string   `yaml:"Home,omitempty"`
		Quickstart         string   `yaml:"Quickstart,omitempty"`
		Features           []any    `yaml:"Features,omitempty"`
		Examples           []string `yaml:"Examples,omitempty"`
		Modules            []string `yaml:"Modules,omitempty"`
		SystemRequirements []any    `yaml:"System Requirements,omitempty"`
		Dependabot         string   `yaml:"Dependabot,omitempty"`
		UsageMetrics       string   `yaml:"Usage Metrics,omitempty"`
		Contributing       string   `yaml:"Contributing,omitempty"`
		GettingHelp        string   `yaml:"Getting help,omitempty"`
	} `yaml:"nav"`
	EditURI string `yaml:"edit_uri"`
	Extra   struct {
		LatestVersion string `yaml:"latest_version"`
	} `yaml:"extra"`
}

func (c *Config) addModule(isModule bool, moduleMd string, indexMd string) {
	_ = "STUB: not implemented"
	return
}

// make sure the index.md is the first element in the list of examples in the nav

// filter out the index.md file

// prepend the index.md file
