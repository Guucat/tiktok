package service

import (
	"tiktok/dao/mysql"
	"tiktok/model"
)

func MessageAction(toUserId, fromUserId, content interface{}) error {
	return mysql.AddMessageInfo(toUserId, fromUserId, content)
}

func MessageChat(fromUserId, toUserId interface{}, preMsgTime string) ([]model.Message, error) {
	return mysql.GetMessageList(fromUserId, toUserId, preMsgTime)
}
