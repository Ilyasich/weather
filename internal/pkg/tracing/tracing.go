package tracing

import "context"

var trace struct{}

func Add(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, trace, token)
}

func Get(ctx context.Context) string {
	value := ctx.Value(trace)
	if token, ok := value.(string); ok {
		return token
	}

	return ""
}