package service

import "tiktok/dao/mysql"

func FollowAction(toUserId, userID interface{}) error {
	return mysql.AddFollow(toUserId, userID)
}

func CanFollowAction(toUserId, userID interface{}) error {
	return mysql.CancelFollow(toUserId, userID)
}

//func FollowList() {
//	return mysql.GetFollowList()
//}
//
//func Followerist() {
//	return mysql.GetFollowerList()
//}
