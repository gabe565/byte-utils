package config

import "github.com/spf13/cobra"

const (
	FlagDecimal   = "decimal"
	FlagSpace     = "space"
	FlagPrecision = "precision"
	FlagInvert    = "invert"
)

func (c *Base) RegisterFlags(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVarP(&c.Decimal, FlagDecimal, "d", c.Decimal, "Use decimal instead of binary")
	fs.IntVarP(&c.Precision, FlagPrecision, "p", c.Precision, "Number of decimal places to use")
	fs.BoolVarP(&c.Space, FlagSpace, "s", c.Space, "Remove the space between the number and the unit")
}

func (c *Bytect) RegisterFlags(cmd *cobra.Command) {
	c.Base.RegisterFlags(cmd)
}

func (c *Bytefmt) RegisterFlags(cmd *cobra.Command) {
	c.Base.RegisterFlags(cmd)
	fs := cmd.Flags()
	fs.BoolVarP(&c.Invert, FlagInvert, "i", c.Invert, "Convert formatted bytes to raw")
}
