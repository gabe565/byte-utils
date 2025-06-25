package config

import "gabe565.com/utils/bytefmt"

type Config interface {
	Base | Bytect | Bytefmt
}

func NewShared() *Base {
	return &Base{Precision: 2}
}

type Base struct {
	Decimal   bool
	Precision int
	Space     bool
}

func (c *Base) NewEncoder() *bytefmt.Encoder {
	return bytefmt.NewEncoder().
		SetPrecision(c.Precision).
		SetUseSpace(c.Space)
}

func (c *Base) NewEncodeFunc() func(int64) string {
	if c.Decimal {
		return c.NewEncoder().EncodeDecimal
	}
	return c.NewEncoder().EncodeBinary
}

func NewBytect() *Bytect {
	return &Bytect{Base: NewShared()}
}

type Bytect struct {
	*Base
	Stdout bool
}

func NewBytefmt() *Bytefmt {
	return &Bytefmt{Base: NewShared()}
}

type Bytefmt struct {
	*Base
}
