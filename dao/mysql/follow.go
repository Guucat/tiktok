package mysql

import (
	"gorm.io/gorm"
	"log"
)

// AddFollow 关注用户
func AddFollow(toUserId, userID interface{}) (err error) {
	//开启事务
	err = DB.Transaction(func(tx *gorm.DB) error {

		//增加关注记录
		if err = tx.Table("followers").Create(map[string]interface{}{
			"from": userID,
			"to":   toUserId,
		}).Error; err != nil {
			log.Println("Fail to Follow", err)
			return err
		}

		//增加用户的关注数
		if err = tx.Table("users").Where("id = ?", userID).
			UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
			log.Println("Fail to follow", err)
			return err
		}

		//增加用户的粉丝数
		if err = tx.Table("users").Where("id = ?", toUserId).
			UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
			log.Println("Fail to follow", err)
			return err
		}

		return nil
	})
	return err
}

// CancelFollow 取消关注用户
func CancelFollow(toUserId, userID interface{}) (err error) {
	//开启事务
	err = DB.Transaction(func(tx *gorm.DB) error {

		//软删除关注记录
		if err = tx.Table("followers").
			Where("`to` = ? and `from` = ?", toUserId, userID).
			Update("state", 0).Error; err != nil {
			log.Println("Fail to Cancel Follow", err)
			return err
		}

		//减少用户的关注数
		if err = tx.Table("users").Where("id = ?", userID).
			UpdateColumn("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
			log.Println("Fail to Cancel Follow", err)
			return err
		}

		//减少用户的粉丝数
		if err = tx.Table("users").Where("id = ?", toUserId).
			UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
			log.Println("Fail to Cancel Follow", err)
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

//// GetFollowList 获取关注列表
//func GetFollowList(){
//
//}
//
//// GetFollowerList 获取粉丝列表
//func GetFollowerList(){
//
//}
