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
	douyin.GET("/publish/list/", jwt.Auth(), c.List)
	douyin.GET("/feed/", c.Feed)

	// 互动接口 - I
	douyin.POST("/favorite/action/", jwt.Auth(), c.FavoriteAction)
	douyin.GET("/favorite/list/", jwt.Auth(), c.FavoriteList)
	douyin.POST("/comment/action/", jwt.Auth(), c.CommentAction)
	//douyin.GET("/comment/list/", jwt.Auth(), c.CommentList)

	// 社交接口
	douyin.POST("/relation/action/", jwt.Auth(), c.FollowAction)
	//douyin.GET("/relation/follow/list/", jwt.Auth(), c.FollowList)
	//douyin.GET("/relation/follower/list/", jwt.Auth(), c.FollowerList)

	return r
}
