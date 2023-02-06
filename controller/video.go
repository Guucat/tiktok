package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	. "tiktok/mid"
	"tiktok/model"
	s "tiktok/service"
	"time"
)

func Upload(c *gin.Context) {
	// 获取视频全局id
	videoId, err := s.GetStoreId()
	if err != nil {
		Fail(c, "upload failed", nil)
		return
	}

	// 获取视频数据
	file, _ := c.FormFile("data")
	f, err := file.Open()
	if err != nil {
		log.Println("fail to load video data", err)
		Fail(c, "upload failed", nil)
		return
	}
	defer f.Close()

	// 视频上传oss并获取视频url
	videoUrl, err := s.StoreFileWithId(f, strconv.FormatInt(videoId, 10)+".mp4")
	if err != nil {
		Fail(c, "upload failed", nil)
		return
	}

	// 获取封面url
	imageUrl, err := s.GetSnapshot(f)
	if err != nil {
		Fail(c, "upload failed", nil)
		return
	}

	// 初始化视频信息，存放mysql
	title := c.PostForm("title")
	videoInfo := &model.Video{
		Id:       videoId,
		AuthorId: c.GetInt64("id"),
		PlayUrl:  videoUrl,
		CoverUrl: imageUrl,
		Title:    title,
	}
	err = s.SaveVideoInfo(videoInfo)
	if err != nil {
		Fail(c, "upload failed", nil)
		return
	}
	Ok(c, "success to upload file", nil)
}

func GetFeed(c *gin.Context) {
	data := gin.H{}
	data["next_time"] = time.Now().Unix()
	Ok(c, "test success", data)
}
