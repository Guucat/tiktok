package mysql

import (
	"gorm.io/gorm"
	"log"
	"tiktok/model"
)

func InsertVideo(v *model.Video) (err error) {
	//开启事务
	err = DB.Transaction(func(tx *gorm.DB) error {
		//存入视频数据
		if err = tx.Create(v).Error; err != nil {
			log.Println("InsertVideo failed to insert", err)
			return err
		}
		//增加用户作品数
		if err = tx.Table("users").Where("id = ?", v.AuthorId).
			UpdateColumn("work_count", gorm.Expr("work_count + ?", 1)).Error; err != nil {
			log.Println("InsertVideo failed to insert", err)
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func QueryVideoList(id, start interface{}) ([]model.Video, error) {
	var list []model.Video
	var tx *gorm.DB
	if id == "" {
		tx = DB.Select("id, author_id, play_url, cover_url, favorite_count, comment_count, title, create_time").
			Where("create_time < ?", start.(string)).
			Order("create_time desc").Limit(15).
			Find(&list)
	} else {
		tx = DB.Select("id, author_id, play_url, cover_url, favorite_count, comment_count, title").
			Where("author_id = ?", id).
			Order("create_time desc").
			Find(&list)

	}
	return list, tx.Error
}

func IsFavorite(userId, videioId interface{}) bool {
	n := -1
	DB.Table("user_favorite_video").
		Select("id").
		Where("user_id = ? and video_id = ? and state = 1", userId, videioId).
		Find(&n)
	if n == -1 {
		return false
	}
	return true
}
