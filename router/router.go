package router

import (
	"github.com/gin-gonic/gin"
	c "tiktok/controller"
	"tiktok/mid/jwt"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	douyin := r.Group("/douyin")

	// 基础模块接口
	douyin.GET("/user/", jwt.Auth(), c.UserInfo)
	douyin.POST("/user/login/", c.Login)
	douyin.POST("/user/register/", c.Register)
	douyin.POST("/publish/action/", jwt.Auth(), c.Upload)

	return r
}
