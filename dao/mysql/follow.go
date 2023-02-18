package mysql

import "gorm.io/gorm"

// 关注用户
func addFollowUser(toUserId string, userID string)(err error) {
	//开启事务
	err = DB.Transaction(func(tx *gorm.DB) error {

	}
	return nil

}
