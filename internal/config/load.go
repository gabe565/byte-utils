package config

import "github.com/spf13/cobra"

func Load[T Config](cmd *cobra.Command) *T {
	cfg, ok := FromContext[T](cmd.Context())
	if !ok {
		panic("context missing config")
	}
	return cfg
}
