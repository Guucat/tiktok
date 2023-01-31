package mysql

import (
	"log"
	"tiktok/model"
)

func InsertVideo(v *model.Video) (err error) {
	if err = DB.Create(v).Error; err != nil {
		log.Println("InsertVideo failed to insert", err)
	}
	return
}
