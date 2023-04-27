package mysql

import (
	"gorm.io/gorm"
	"log"
	"strconv"
	"tiktok/pkg/model"
	"time"
)

func AddMessageInfo(fromUserId, toUserId, content interface{}) (err error) {
	//开启事务
	err = DB.Transaction(func(tx *gorm.DB) error {
		//存入数据
		if err = tx.Table("message").Create(map[string]interface{}{
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

func GetMessageList(toUserId, fromUserId interface{}, preMsgTime string) (messageChat []model.Message, err error) {
	// string转化为时间，layout必须为 "2006-01-02 15:04:05"
	timeLayout := "2006-01-02 15:04:05" //转化所需模板
	lastTime, _ := strconv.ParseInt(preMsgTime, 10, 64)
	lastMsgTime := time.Unix(lastTime, 10).Format(timeLayout)

	if err = DB.Table("message").Where("create_time > ?", lastMsgTime).
		Where(DB.Where("from_user_id = ? AND to_user_id = ? AND state = 1", fromUserId, toUserId).
			Or("from_user_id = ? AND to_user_id = ? AND state = 1", toUserId, fromUserId)).
		//Where("from_user_id = ? AND to_user_id = ? AND state = 1 AND create_time > ? ", fromUserId, toUserId, lastMsgTime).
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
