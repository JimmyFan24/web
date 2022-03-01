package config

import "web/internal/logagent/options"

// Config is the running configuration structure of the LogAgent
type Config struct {
	*options.Options
}

// CreateConfigFromOptions creates a running configuration instance based
// on a given LogAgentServer command line or configuration file option.
func CreateConfigFromOptions(opt *options.Options) (*Config,error) {
	return &Config{opt},nil
}


