package mysql

import (
	"gorm.io/gorm"
	"log"
	"tiktok/pkg/model"
)

// AddRelationCount 关注
func AddRelationCount(fromUserId, toUserId interface{}) (err error) {
	//开启事务
	err = DB.Transaction(func(tx *gorm.DB) error {
		//判断记录是否存在或被软删除
		id := -1
		if err = tx.Table("followers").Where("from_user_id = ? and to_user_id = ?", fromUserId, toUserId).
			Select("id").Find(&id).Error; err != nil {
			log.Println("Fail", err)
			return err
		}
		if id == -1 {
			//数据库还不存在该数据，存入数据
			if err = tx.Table("followers").Create(map[string]interface{}{
				"from_user_id": fromUserId,
				"to_user_id":   toUserId,
			}).Error; err != nil {
				log.Println("Fail", err)
				return err
			}
			//增加关注数
			if err = tx.Table("users").Where("id = ?", fromUserId).
				UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
				log.Println("Fail", err)
				return err
			}
			//增加对方粉丝数
			if err = tx.Table("users").Where("id = ?", toUserId).
				UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
				log.Println("Fail", err)
				return err
			}
			return nil
		}

		//数据库已存在该数据，只是软删除了，把state=1
		if err = tx.Table("followers").Where("from_user_id = ? and to_user_id = ?", fromUserId, toUserId).
			Update("state", 1).Error; err != nil {
			log.Println("Fail", err)
			return err
		}

		//增加关注数
		if err = tx.Table("users").Where("id = ?", fromUserId).
			UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
			log.Println("Fail", err)
			return err
		}
		//增加对方粉丝数
		if err = tx.Table("users").Where("id = ?", toUserId).
			UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
			log.Println("Fail", err)
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// SubRelationCount 取消关注
func SubRelationCount(fromUserId, toUserId interface{}) (err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		id := -1
		if err = tx.Table("followers").
			Where("from_user_id = ? and to_user_id = ? and state = 1", fromUserId, toUserId).
			Select("id").Find(&id).Error; err != nil {
			log.Println("Fail", err)
			return err
		}
		if id != -1 {
			if err = tx.Table("followers").Where("from_user_id = ? and to_user_id = ?", fromUserId, toUserId).
				Update("state", "0").Error; err != nil {
				log.Println("Fail", err)
				return err
			}
			if err = tx.Table("users").Where("id = ?", fromUserId).
				UpdateColumn("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
				log.Println("Fail", err)
				return err
			}
			if err = tx.Table("users").Where("id = ?", toUserId).
				UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
				log.Println("Fail", err)
				return err
			}
			return nil
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// GetFollowList 获取关注用户
func GetFollowList(userId interface{}) (followMessage []model.User, err error) {
	userIds := make([]int, 0)
	err = DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Table("followers").Where("from_user_id = ? AND state = ?", userId, 1).
			Select("to_user_id").Find(&userIds).Error; err != nil {
			log.Println("Fetch error", err)
			return err
		}
		for _, userId := range userIds {
			var user model.User
			user, err = GetUserMessageById(userId)
			followMessage = append(followMessage, user)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return followMessage, nil
}

// GetFollowerList 获取粉丝用户
func GetFollowerList(userId interface{}) (followMessage []model.User, err error) {
	userIds := make([]int, 0)
	err = DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Table("followers").Where("to_user_id = ? AND state = ?", userId, 1).
			Select("from_user_id").Find(&userIds).Error; err != nil {
			log.Println("Fetch error", err)
			return err
		}
		for _, userId := range userIds {
			var user model.User
			user, err = GetUserMessageById(userId)
			followMessage = append(followMessage, user)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return followMessage, nil
}
