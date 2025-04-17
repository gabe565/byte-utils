package config

import "github.com/spf13/cobra"

const (
	FlagDecimal   = "decimal"
	FlagSpace     = "space"
	FlagPrecision = "precision"
	FlagInvert    = "invert"
)

func (c *Config) RegisterFlags(cmd *cobra.Command, withInvert bool) {
	fs := cmd.Flags()
	fs.BoolVarP(&c.Decimal, FlagDecimal, "d", c.Decimal, "Use decimal instead of binary")
	fs.IntVarP(&c.Precision, FlagPrecision, "p", c.Precision, "Number of decimal places to use")
	fs.BoolVarP(&c.Space, FlagSpace, "s", c.Space, "Remove the space between the number and the unit")
	if withInvert {
		fs.BoolVarP(&c.Invert, FlagInvert, "i", c.Invert, "Convert formatted bytes to raw")
	}
}
