package router

import (
	"github.com/gin-gonic/gin"
	c "tiktok/controller"
	"tiktok/mid/jwt"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	dGroup := r.Group("/douyin")
	{
		dGroup.GET("/user/", jwt.Auth(), c.UserInfo)
	}

	userGroup := dGroup.Group("/user")
	{
		userGroup.POST("/login/", c.Login)
		userGroup.POST("/register/", c.Register)
	}

	publishGroup := dGroup.Group("/publish")
	{
		publishGroup.POST("/action/", jwt.Auth(), c.Upload)
	}

	return r
}
