package api

import (
	"easypay/tool"
	"github.com/gin-gonic/gin"
	"strings"
)

func auth(ctx *gin.Context) {
	username, err := ctx.Cookie("username")
	if err != nil {
		tool.RespErrorWithData(ctx, "请登陆后进行操作")
		ctx.Abort()
	}
	ctx.Set("username", username)
	ctx.Next()
}

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			tool.RespErrorWithData(c, "请求头中auth为空")
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			tool.RespErrorWithData(c, "请求头中auth格式有误")
			c.Abort()
			return
		}

		user, err := tool.ParseToken(parts[1])
		if err != nil {
			tool.RespErrorWithData(c, "无效的Token")
			c.Abort()
			return
		}

		c.Set("username", user.Username)
		c.Next()
	}
}
