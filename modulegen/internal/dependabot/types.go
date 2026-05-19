package dependabot

// Config the configuration for the dependabot config file
type Config struct {
	Version int     `yaml:"version"`
	Updates Updates `yaml:"updates"`
}

// Schedule the schedule for the dependabot config file
type Schedule struct {
	Interval string `yaml:"interval"`
	Day      string `yaml:"day"`
}

// Update the update for the dependabot config file
type Update struct {
	PackageEcosystem      string   `yaml:"package-ecosystem"`
	Directories           []string `yaml:"directories"`
	Schedule              Schedule `yaml:"schedule"`
	OpenPullRequestsLimit int      `yaml:"open-pull-requests-limit"`
	RebaseStrategy        string   `yaml:"rebase-strategy"`
}

// Updates the updates for the dependabot config file
type Updates []Update

// addUpdate adds an update to the config
func (c *Config) addUpdate(modulePath string) error { _ = "STUB: not implemented"; return nil }

// look up the update in the gomodsUpdate
