package api

import (
	"easypay/model"
	"easypay/tool"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

var MySecret = []byte("ddzyyds")

//JWTAuthMiddleware jwt鉴权
func JWTAuthMiddleware(ctx *gin.Context) {
	userName, _ := ctx.Cookie("username")

	claim := model.MyClaims{
		Username: userName,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "easy_pay",
			ExpiresAt: time.Now().Add(time.Second * 10).Unix(),
		},
	}
	tokenString, err := ctx.Cookie("jwt")
	if err != nil {
		tool.RespErrorWithData(ctx, "您还没有登入")
		ctx.Abort()
		return
	}

	//解析
	token, err := jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})

	if token.Valid {
		ctx.JSON(200, gin.H{
			"您好！": userName,
		})
		return
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			tool.RespErrorWithData(ctx, "token格式有误")
			fmt.Println(err)
			ctx.Abort()
			return
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			tool.RespErrorWithData(ctx, "验证已过期！")
			fmt.Println(err)
			ctx.Abort()
			return
		} else {
			tool.RespErrorWithData(ctx, "签证失败")
		}
	}
	tool.RespErrorWithData(ctx, "签证失败")
	fmt.Println(err)
}
