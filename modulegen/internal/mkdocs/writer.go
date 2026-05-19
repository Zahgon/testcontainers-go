package mkdocs

func writeConfig(configFile string, config *Config) error { _ = "STUB: not implemented"; return nil }

// simple solution to replace the empty strings, as mapping those fields
// into the MkDocs config is not supported yet
func overrideData(data []byte) []byte { _ = "STUB: not implemented"; return nil }

func setEmoji(content string, key string, value string) string {
	_ = "STUB: not implemented"
	return ""
}
