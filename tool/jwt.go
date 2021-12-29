package tool

import (
	"easypay/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("ddzyyds")

func SetToken(ctx *gin.Context, username string) (string, error) {
	myClaim := model.MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "easy_pay",
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)

	s, err2 := token.SignedString(MySecret)
	if err2 != nil {
		RespErrorWithData(ctx, err2)
		return s, err2
	}
	return s, nil
}
