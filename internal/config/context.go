package config

import "context"

type ctxKey uint8

const configKey ctxKey = iota

func NewContext[T Config](ctx context.Context, conf *T) context.Context {
	return context.WithValue(ctx, configKey, conf)
}

func FromContext[T Config](ctx context.Context) (*T, bool) {
	conf, ok := ctx.Value(configKey).(*T)
	return conf, ok
}
