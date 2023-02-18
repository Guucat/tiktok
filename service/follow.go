package service

import "tiktok/dao/mysql"

func FollowAtion(userID, interface{}) error {
	return mysql.AddFollow()
}

func CanFollow() {
	return mysql.CancelFollow()
}

func FollowList() {
	return mysql.GetFollowList()
}

func Followerist() {
	return mysql.GetFollowerList()
}
