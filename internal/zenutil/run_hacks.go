package zenutil

import (
	"context"
)

type unixUIDContextKey struct{}

func WithUnixUID(parent context.Context, uid uint32) context.Context {
	return context.WithValue(parent, unixUIDContextKey{}, uid)
}

func ContextUnixUID(ctx context.Context) *uint32 {
	uid, ok := ctx.Value(unixUIDContextKey{}).(uint32)
	if ok {
		return &uid
	}
	return nil
}
