package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tiktok/dao/mysql"
	"tiktok/service"
	"time"
)

type CommentActionResponse struct {
	StatusCode int32   `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg,omitempty"` // 返回状态描述
	Comment    Comment `json:"comment,omitempty"`
}

func CommentAction(c *gin.Context) {
	actionType := c.Query("action_type")
	videoId := c.Query("video_id")
	commentText := c.Query("comment_text")
	commentId, _ := service.GetStoreId() // 获取唯一id
	userId, _ := c.Get("id")

	if actionType != "1" {
		commentId := c.Query("comment_id")
		err := service.DelCommentAction(commentId, videoId)
		if err != nil {
			log.Println("Delete error", err)
			return
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"StatusCode": 0,
			"StatusMsg":  "success to delete",
			"comment":    nil,
		})
	}

	err := service.CommentAction(videoId, commentId, userId, commentText)
	if err != nil {
		log.Println("Comment error", err)
		return
	}
	userMessage, err := service.GetAuthorMessage(userId)
	if err != nil {
		log.Println("Comment error", err)
		return
	}
	comment := Comment{
		Id: commentId,
		User: User{
			Id:              userMessage.Id,
			Name:            userMessage.Username,
			FollowCount:     userMessage.FollowCount,
			FollowerCount:   userMessage.FollowerCount,
			Avatar:          userMessage.Avatar,
			BackgroundImage: userMessage.BackgroundImage,
			Signature:       userMessage.Signature,
			TotalFavorited:  userMessage.TotalFavorited,
			WorkCount:       userMessage.WorkCount,
			FavoriteCount:   userMessage.FavoriteCount,
			IsFollow:        userMessage.IsFollow,
		},
		Content:    commentText,
		CreateDate: time.Now().Format("01-02"),
	}
	c.JSON(http.StatusOK, CommentActionResponse{
		StatusCode: 0,
		StatusMsg:  "success to Comment",
		Comment:    comment,
	})
}

type CommentListResponse struct {
	StatusCode  int32     `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg   string    `json:"status_msg,omitempty"` // 返回状态描述
	CommentList []Comment `json:"comment_list,omitempty"`
}

func CommentList(c *gin.Context) {
	videoId := c.Query("video_id")
	commentMessage, err := service.CommentList(videoId)
	userId, _ := c.Get("id")
	if err != nil {
		log.Println("Fetch error", err)
		return
	}

	var commentList = make([]Comment, 0, 10)
	for _, commentDao := range commentMessage {
		userMessage, err := service.GetAuthorMessage(commentDao.UserId)
		if err != nil {
			log.Println("Fetch error", err)
			return
		}
		isFollower := mysql.GetIsFollower(userId, commentDao.UserId)
		user := User{
			Id:              userMessage.Id,
			Name:            userMessage.Username,
			FollowCount:     userMessage.FollowCount,
			FollowerCount:   userMessage.FollowerCount,
			Avatar:          userMessage.Avatar,
			BackgroundImage: userMessage.BackgroundImage,
			Signature:       userMessage.Signature,
			TotalFavorited:  userMessage.TotalFavorited,
			WorkCount:       userMessage.WorkCount,
			FavoriteCount:   userMessage.FavoriteCount,
			IsFollow:        isFollower,
		}
		comment := Comment{
			Id:         commentDao.Id,
			User:       user,
			Content:    commentDao.Content,
			CreateDate: commentDao.CreateDate,
		}
		commentList = append(commentList, comment)
	}

	c.JSON(http.StatusOK, CommentListResponse{
		StatusCode:  0,
		StatusMsg:   "success",
		CommentList: commentList,
	})
}
