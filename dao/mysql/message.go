package mysql

import (
	"gorm.io/gorm"
	"log"
	"tiktok/model"
)

func AddMessageInfo(fromUserId, toUserId, content interface{}) (err error) {
	//开启事务
	err = DB.Transaction(func(tx *gorm.DB) error {
		//存入数据
		if err = tx.Table("user_favorite_video").Create(map[string]interface{}{
			"from_user_id": fromUserId,
			"to_user_id":   toUserId,
			"content":      content,
		}).Error; err != nil {
			log.Println("Fail to message", err)
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func GetMessageList(toUserId, fromUserId interface{}) (messageChat []model.Message, err error) {
	if err = DB.Table("message").
		Where("from_user_id = ? AND to_user_id = ? AND state = 1", fromUserId, toUserId).
		Find(&messageChat).Error; err != nil {
		log.Println("No result", err)
		return
	}
	if err != nil {
		log.Println("No result", err)
		return
	}
	return
}
