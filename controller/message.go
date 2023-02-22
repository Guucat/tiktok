package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tiktok/service"
)

func MessageAction(c *gin.Context) {
	userId, _ := c.Get("id")
	toUserId := c.Query("to_user_id")
	//actionType := c.Query("action_type")
	content := c.Query("content")
	err := service.MessageAction(userId, toUserId, content)
	if err != nil {
		log.Println("Fail to message", err)
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "success to message",
	})
}

type MessageListResponse struct {
	StatusCode  int32     `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg   string    `json:"status_msg,omitempty"` // 返回状态描述
	MessageList []Message `json:"message_list"`
}

func MessageChat(c *gin.Context) {
	userId, _ := c.Get("id")
	toUserId := c.Query("to_user_id")
	preMsgTime := c.Query("pre_msg_time")
	fmt.Printf("c: preMsgTime is %v", preMsgTime)
	fmt.Println()
	messageChat, err := service.MessageChat(userId, toUserId, preMsgTime)
	if err != nil {
		log.Println("Fetch error", err)
		return
	}

	var messageList = make([]Message, 0, 10)
	for _, mesDao := range messageChat {
		timeUnix := mesDao.CreateTime.Unix()
		message := Message{
			Id:         mesDao.Id,
			FromUserId: mesDao.FromUserId,
			ToUserId:   mesDao.ToUserId,
			Content:    mesDao.Content,
			CreateTime: timeUnix,
		}

		messageList = append(messageList, message)
	}
	c.JSON(http.StatusOK, MessageListResponse{
		StatusCode:  0,
		StatusMsg:   "success",
		MessageList: messageList,
	})
}
