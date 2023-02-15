package mysql

import (
	"gorm.io/gorm"
	"log"
	"tiktok/model"
)

func AddFavoriteCount(videoId string, userId interface{}) (err error) {
	//开启事务
	err = DB.Transaction(func(tx *gorm.DB) error {
		//增加点赞数
		if err = tx.Table("videos").Where("id = ?", videoId).
			UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			log.Println("Fail to like", err)
			return err
		}
		//存入点赞数据
		if err = tx.Table("user_favorite_video").Create(map[string]interface{}{
			"user_id":  userId,
			"video_id": videoId,
		}).Error; err != nil {
			log.Println("Fail to like", err)
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func SubFavoriteCount(videoId string, userId interface{}) (err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {

		if err = tx.Table("videos").Where("id = ?", videoId).
			UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
			log.Println("Fail to like", err)
			return err
		}
		if err = tx.Table("user_favorite_video").Where("user_id = ? AND video_id = ?", userId, videoId).
			Update("state", "0").Error; err != nil {
			log.Println("Fail to like", err)
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func GetFavoriteListByUserId(userId interface{}) (videoMessage []model.Video, err error) {
	videoIds := make([]int, 0)
	err = DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Table("user_favorite_video").Where("user_id = ? AND state = ?", userId, 1).
			Select("video_id").Find(&videoIds).Error; err != nil {
			log.Println("Fetch error", err)
			return err
		}
		for _, videoId := range videoIds {
			var video model.Video
			if err = tx.Table("videos").Where("id = ?", videoId).
				Find(&video).Error; err != nil {
				log.Println("Fetch error", err)
				return err
			}
			videoMessage = append(videoMessage, video)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return videoMessage, nil
}

func GetUserMessageById(userId interface{}) (user model.User, err error) {
	if err = DB.Table("users").Where("id = ?", userId).
		Find(&user).Error; err != nil {
		log.Println("No result", err)
		return
	}
	if err != nil {
		log.Println("No result", err)
		return
	}
	return
}

// IsFollowerAuthor 自己是否关注了他
func IsFollowerAuthor(userId, authorId interface{}) bool {
	n := 0
	DB.Table("followers").Select("count(*)").Where("`from` = ? and `to` = ?", userId, authorId).Find(&n)
	return n == 1
}
