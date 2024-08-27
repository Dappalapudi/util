package auth

import (
	"context"
	"fmt"
)

type ctxTokenKey struct{}

// TokenFromCtx returns token from ctx.
func TokenFromCtx(ctx context.Context) (string, error) {
	key := ctxTokenKey{}
	token, ok := ctx.Value(key).(string)
	if !ok {
		return "", fmt.Errorf("empty token")
	}
	return token, nil
}

// WithCtxToken sets token value inside context.
func WithCtxToken(ctx context.Context, token string) context.Context {
	key := ctxTokenKey{}
	return context.WithValue(ctx, key, token)
}
