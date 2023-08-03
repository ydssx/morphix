package middleware

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ydssx/morphix/pkg/jwt"
	"github.com/ydssx/morphix/pkg/util"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")
		token := strings.ReplaceAll(auth, "Bearer ", "")
		claims, err := jwt.VerifyToken(token)
		if err != nil {
			ctx.Abort()
			util.FailWithMsg(ctx, "验证失败")
		}
		ctx.Request = ctx.Request.WithContext(NewContext(ctx.Request.Context(), claims))
		ctx.Next()
	}
}

type authKey struct{}

// NewContext put currentUser into context
func NewContext(ctx context.Context, c *jwt.Claims) context.Context {
	return context.WithValue(ctx, authKey{}, c)
}

func AuthFromGinContext(c *gin.Context) *jwt.Claims {
	return c.Value(authKey{}).(*jwt.Claims)
}
