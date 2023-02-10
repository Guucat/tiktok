package service

import (
	"tiktok/dao/mysql"
	"tiktok/model"
)

func FavoriteAction(videoId, actionType string, userId interface{}) (err error) {
	if actionType == "1" {
		return mysql.AddFavoriteCount(videoId, userId)
	}
	return mysql.SubFavoriteCount(videoId, userId)
}

func FavoriteList(userId interface{}) (videoMessage []model.Video, err error) {
	return mysql.GetFavoriteListByUserId(userId)
}

func GetAuthorMessage(authorId interface{}) (model.User, error) {
	return mysql.GetUserMessageById(authorId)
}
