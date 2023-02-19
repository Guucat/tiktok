package service

import (
	"tiktok/dao/mysql"
	"tiktok/model"
)

func RelationAction(fromUserId, toUserId, actionType interface{}) (err error) {
	if actionType == "1" {
		return mysql.AddRelationCount(fromUserId, toUserId)
	}
	return mysql.SubRelationCount(fromUserId, toUserId)
}

func FollowList(userId interface{}) (followMessage []model.User, err error) {
	return mysql.GetFollowList(userId)
}

func FollowerList(userId interface{}) (followerMessage []model.User, err error) {
	return mysql.GetFollowerList(userId)
}
