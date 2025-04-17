package config

import "github.com/spf13/cobra"

func Load(cmd *cobra.Command) *Config {
	cfg, ok := FromContext(cmd.Context())
	if !ok {
		panic("context missing config")
	}
	return cfg
}
