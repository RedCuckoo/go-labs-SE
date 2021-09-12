package service

import (
	"context"
	"github.com/deliveroo/jsonrpc-go"
)

func CtxMiddleware(extenders ...func(context.Context) context.Context) jsonrpc.Middleware {
	return func(next jsonrpc.Next) jsonrpc.Next {
		return func(ctx context.Context, params interface{}) (interface{}, error) {
			for _, extender := range extenders {
				ctx = extender(ctx)
			}

			return next(ctx, params)
		}
	}
}
