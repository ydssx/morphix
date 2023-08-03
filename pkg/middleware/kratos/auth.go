package kratos

import (
	"context"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/ydssx/morphix/pkg/jwt"
)

func AuthServer() middleware.Middleware {
	return func(h middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			ts, ok := transport.FromServerContext(ctx)
			if ok {
				header := ts.RequestHeader()
				token := strings.ReplaceAll(header.Get("Authorization"), "Bearer ", "")
				claims, err := jwt.VerifyToken(token)
				if err != nil {
					return nil, errors.Unauthorized(err.Error(), "请登录")
				}
				ctx = NewContext(ctx, claims)
				return h(ctx, req)
			}
			return nil, errors.Forbidden("forbidden", "禁止访问")
		}
	}
}

type authKey struct{}

// NewContext put currentUser into context
func NewContext(ctx context.Context, c *jwt.Claims) context.Context {
	return context.WithValue(ctx, authKey{}, c)
}
