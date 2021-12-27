package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	engine := gin.Default()
	engine.GET("help", help)

	engine.POST("/register", register)
	engine.POST("/login", login)
	userGroup := engine.Group("/user")
	{
		userGroup.Use(auth)
		userGroup.POST("/topup", topUp)
		userGroup.POST("/transfer", transfer)
		userGroup.PUT("password", changePassword)
		userGroup.GET("/getrecord", getRecord)
	}
	engine.Run(":8020")
}
