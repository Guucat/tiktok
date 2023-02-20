package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tiktok/service"
	"time"
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

	messageChat, err := service.MessageChat(userId, toUserId)
	if err != nil {
		log.Println("Fetch error", err)
		return
	}

	var messageList = make([]Message, 0, 10)
	for _, mesDao := range messageChat {
		//获取本地location
		toBeCharge := mesDao.Model.CreateTime.String()                  //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
		timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
		loc, _ := time.LoadLocation("Local")                            //重要：获取时区
		theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
		sr := theTime.Unix()                                            //转化为时间戳 类型是int64
		fmt.Println(theTime)                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
		fmt.Println(sr)
		message := Message{
			Id:         mesDao.Id,
			FromUserId: mesDao.FromUserId,
			ToUserId:   mesDao.ToUserId,
			Content:    mesDao.Content,
			CreateTime: sr,
		}

		messageList = append(messageList, message)
	}
	c.JSON(http.StatusOK, MessageListResponse{
		StatusCode:  0,
		StatusMsg:   "success",
		MessageList: messageList,
	})
}
