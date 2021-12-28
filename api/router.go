package api

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	engine := gin.Default()
	engine.GET("/help", help)
	engine.POST("/login_with_jwt", loginJWT)
	engine.POST("/register", register)
	engine.POST("/login", login)
	engine.GET("/jwt_text", JWTAuthMiddleware(), help)
	userGroup := engine.Group("/user")
	{
		userGroup.Use(auth)
		userGroup.POST("/top_up", topUp)
		userGroup.POST("/transfer", transfer)
		userGroup.PUT("password", changePassword)
		userGroup.GET("/get_all_record", getRecord)
		userGroup.GET("/get_money", getMoney)
		userGroup.GET("/get_exact_record", getRecordWithPeople)

	}
	engine.Run(":8020")
}
