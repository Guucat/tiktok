package router

import (
	"github.com/gin-gonic/gin"
	c "tiktok/controller"
	"tiktok/mid/jwt"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	douyin := r.Group("/douyin")
	{
		douyin.GET("/user/", jwt.Auth(), c.UserInfo)

		user := douyin.Group("/user")
		{
			user.POST("/login/", c.Login)
			user.POST("/register/", c.Register)
		}

		publish := douyin.Group("/publish")
		{
			publish.POST("/action/", jwt.Auth(), c.Upload)
		}
	}
	return r
}
