package middleware

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ydssx/morphix/common/rbac"
	"github.com/ydssx/morphix/pkg/jwt"
	"github.com/ydssx/morphix/pkg/util"
)

// Auth 是一个 Gin 中间件,用于验证 JWT 令牌并执行 RBAC 权限检查。
// 它从请求头中获取 Authorization 头,提取 Bearer 令牌,并使用 jwt.VerifyToken 验证令牌。
// 如果令牌有效,它会从令牌中提取角色信息,并使用 casbin 执行 RBAC 权限检查。
// 如果权限检查通过,它会将 JWT 声明添加到请求上下文中,并调用下一个中间件。
// 如果令牌无效或权限不足,它会中止请求并返回相应的错误。
func Auth() gin.HandlerFunc {
	e := rbac.NewCasbinEnforcer()
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")
		token := strings.ReplaceAll(auth, "Bearer ", "")
		claims, err := jwt.VerifyToken(token)
		if err != nil {
			ctx.Abort()
			util.FailWithMsg(ctx, "验证失败")
		}
		ok, err := e.Enforce(claims.Role, ctx.Request.URL.Path, ctx.Request.Method)
		if err != nil {
			ctx.Abort()
			util.FailWithError(ctx, err)
		}
		if !ok {
			ctx.Abort()
			util.FailWithMsg(ctx, "权限不足")
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
