package config

import "gabe565.com/utils/bytefmt"

func New() *Config {
	return &Config{
		Precision: 2,
	}
}

type Config struct {
	Decimal   bool
	Precision int
	Space     bool
	Invert    bool
}

func (c *Config) NewEncoder() *bytefmt.Encoder {
	return bytefmt.NewEncoder().
		SetPrecision(c.Precision).
		SetUseSpace(c.Space)
}

func (c *Config) NewEncodeFunc() func(int64) string {
	if c.Decimal {
		return c.NewEncoder().EncodeDecimal
	}
	return c.NewEncoder().EncodeBinary
}
