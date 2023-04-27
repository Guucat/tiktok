package service

import (
	"tiktok/dao/mysql"
	model2 "tiktok/pkg/model"
)

func FavoriteAction(videoId, actionType string, userId interface{}) (err error) {
	if actionType == "1" {
		return mysql.AddFavoriteCount(videoId, userId)
	}
	return mysql.SubFavoriteCount(videoId, userId)
}

func FavoriteList(userId interface{}) (videoMessage []model2.Video, err error) {
	return mysql.GetFavoriteListByUserId(userId)
}

func GetAuthorMessage(authorId interface{}) (model2.User, error) {
	return mysql.GetUserMessageById(authorId)
}
