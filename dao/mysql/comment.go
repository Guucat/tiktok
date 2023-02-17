package mysql

import (
	"gorm.io/gorm"
	"log"
	"tiktok/model"
)

// AddComment 添加评论
func AddComment(videoId, commentId, userId, commentText interface{}) (err error) {
	//开启事务
	err = DB.Transaction(func(tx *gorm.DB) error {
		//存入评论数据
		if err = tx.Table("comment").Create(map[string]interface{}{
			"id":       commentId,
			"user_id":  userId,
			"video_id": videoId,
			"content":  commentText,
		}).Error; err != nil {
			log.Println("Fail to comment", err)
			return err
		}
		//增加视频评论数
		if err = tx.Table("videos").Where("id = ?", videoId).
			UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
			log.Println("Fail to comment", err)
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// DeleteComment 删除评论
func DeleteComment(commentId, videoId interface{}) (err error) {
	//开启事务
	err = DB.Transaction(func(tx *gorm.DB) error {
		//软删除评论数据
		if err = tx.Table("comment").Where("id = ?", commentId).
			Update("state = ?", 0).Error; err != nil {
			log.Println("Fail to comment", err)
			return err
		}
		//减少视频评论数
		if err = tx.Table("videos").Where("id = ?", videoId).
			UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
			log.Println("Fail to comment", err)
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func GetCommentListByVideoId(videoId interface{}) (commentMessage []model.Comment, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Table("comment").Where("video_id = ? AND state = ?", videoId, 1).
			Find(&commentMessage).Error; err != nil {
			log.Println("Fetch error", err)
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return commentMessage, nil
}
