package mysql

import (
	"log"
	"testing"
	"tiktok/pkg/model"
)

func TestInsertVideo(t *testing.T) {
	// gorm 会默认零值为缺省值
	id := -1
	v := &model.Video{
		Id:       int64(id),
		AuthorId: int64(id),
		PlayUrl:  "test",
		CoverUrl: "test",
		Title:    "tt",
		//Time:     time.Now(),
	}
	if err := InsertVideo(v); err != nil {
		log.Fatal("failed", err)
	}
	DB.Delete(&model.Video{}, int64(id))
}
