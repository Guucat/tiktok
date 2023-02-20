package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tiktok/dao/mysql"
	"tiktok/service"
)

type FavoriteListResponse struct {
	StatusCode int32   `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg,omitempty"` // 返回状态描述
	VideoList  []Video `json:"video_list,omitempty"`
}

func FavoriteAction(c *gin.Context) {
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")
	userId, _ := c.Get("id")
	err := service.FavoriteAction(videoId, actionType, userId)
	if err != nil {
		log.Println("Fail to like", err)
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "success to like",
	})
}

func FavoriteList(c *gin.Context) {
	userId, _ := c.Get("id")
	videoMessage, err := service.FavoriteList(userId)
	if err != nil {
		log.Println("Fetch error", err)
		return
	}
	var videoList = make([]Video, 0, 10)
	for _, videoDao := range videoMessage {
		authorMessage, err := service.GetAuthorMessage(videoDao.AuthorId)
		if err != nil {
			log.Println("Fetch error", err)
			return
		}
		isFollower := mysql.GetIsFollower(userId, authorMessage.Id)
		author := User{
			Id:              authorMessage.Id,
			Name:            authorMessage.Username,
			FollowCount:     authorMessage.FollowCount,
			FollowerCount:   authorMessage.FollowerCount,
			Avatar:          authorMessage.Avatar,
			BackgroundImage: authorMessage.BackgroundImage,
			Signature:       authorMessage.Signature,
			TotalFavorited:  authorMessage.TotalFavorited,
			WorkCount:       authorMessage.WorkCount,
			FavoriteCount:   authorMessage.FavoriteCount,
			IsFollow:        isFollower,
		}
		video := Video{
			Id:            videoDao.Id,
			Title:         videoDao.Title,
			Author:        author,
			PlayUrl:       videoDao.PlayUrl,
			CoverUrl:      videoDao.CoverUrl,
			FavoriteCount: videoDao.FavoriteCount,
			CommentCount:  videoDao.CommentCount,
			IsFavorite:    true,
		}
		videoList = append(videoList, video)
	}
	c.JSON(http.StatusOK, FavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		VideoList:  videoList,
	})
}
