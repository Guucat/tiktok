package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tiktok/dao/mysql"
	"tiktok/service"
)

func RelationAction(c *gin.Context) {
	userId, _ := c.Get("id")
	toUserId := c.Query("to_user_id")
	actionType := c.Query("action_type")
	err := service.RelationAction(userId, toUserId, actionType)
	if err != nil {
		log.Println("Fail", err)
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "success",
	})
}

type FollowListResponse struct {
	StatusCode int32  `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` // 返回状态描述
	UserList   []User `json:"user_list,omitempty"`
}

func FollowList(c *gin.Context) {
	userId, _ := c.Get("id")
	followMessage, err := service.FollowList(userId)
	if err != nil {
		log.Println("fail", err)
		return
	}

	var userList = make([]User, 0, 10)
	for _, follow := range followMessage {
		user := User{
			Id:              follow.Id,
			Name:            follow.Username,
			FollowCount:     follow.FollowCount,
			FollowerCount:   follow.FollowerCount,
			Avatar:          follow.Avatar,
			BackgroundImage: follow.BackgroundImage,
			Signature:       follow.Signature,
			TotalFavorited:  follow.TotalFavorited,
			WorkCount:       follow.WorkCount,
			FavoriteCount:   follow.FavoriteCount,
			IsFollow:        true,
		}
		userList = append(userList, user)
	}
	c.JSON(http.StatusOK, FollowListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserList:   userList,
	})
}

func FollowerList(c *gin.Context) {
	userId, _ := c.Get("id")
	followerMessage, err := service.FollowerList(userId)
	if err != nil {
		log.Println("fail", err)
		return
	}

	var userList = make([]User, 0, 10)
	for _, follower := range followerMessage {
		isFollower := mysql.GetIsFollower(userId, follower.Id)
		user := User{
			Id:              follower.Id,
			Name:            follower.Username,
			FollowCount:     follower.FollowCount,
			FollowerCount:   follower.FollowerCount,
			Avatar:          follower.Avatar,
			BackgroundImage: follower.BackgroundImage,
			Signature:       follower.Signature,
			TotalFavorited:  follower.TotalFavorited,
			WorkCount:       follower.WorkCount,
			FavoriteCount:   follower.FavoriteCount,
			IsFollow:        isFollower,
		}
		userList = append(userList, user)
	}
	c.JSON(http.StatusOK, FollowListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserList:   userList,
	})
}

func FriendList(c *gin.Context) {
	userId, _ := c.Get("id")
	followerMessage, err := service.FollowerList(userId)
	if err != nil {
		log.Println("fail", err)
		return
	}

	var userList = make([]User, 0, 10)
	for _, follower := range followerMessage {
		isFollower := mysql.GetIsFollower(userId, follower.Id)
		if isFollower {
			user := User{
				Id:              follower.Id,
				Name:            follower.Username,
				FollowCount:     follower.FollowCount,
				FollowerCount:   follower.FollowerCount,
				Avatar:          follower.Avatar,
				BackgroundImage: follower.BackgroundImage,
				Signature:       follower.Signature,
				TotalFavorited:  follower.TotalFavorited,
				WorkCount:       follower.WorkCount,
				FavoriteCount:   follower.FavoriteCount,
				IsFollow:        isFollower,
			}
			userList = append(userList, user)
		}
	}
	c.JSON(http.StatusOK, FollowListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserList:   userList,
	})
}
