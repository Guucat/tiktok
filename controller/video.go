package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	. "tiktok/mid"
	. "tiktok/mid/oss"
	s "tiktok/service"
	"time"
)

func Upload(c *gin.Context) {
	// 获取视频全局id
	videoId, err := s.GetStoreId()
	log.Println(videoId)
	if err != nil {
		Fail(c, "fail to generate id", nil)
		return
	}
	// 将视频以id为名存放在本地

	// 获取封面全局id
	// 将封面以id为名存放在本地
	// 调用oss存储视频，返回 video url
	// 调用oss存放封面，返回 image url
	file, _ := c.FormFile("data")
	f, err := file.Open()
	if err != nil {
		Fail(c, "fail to load video data", nil)
		return
	}
	err = Oss.PutObject(strconv.FormatInt(videoId, 10)+".mp4", f)
	if err != nil {
		Fail(c, "fail to store file", nil)
		return
	}

	Ok(c, "success to upload file", nil)
	// 初始化视频信息，存放mysql
	//title := c.PostForm("title")

}

func GetFeed(c *gin.Context) {
	data := gin.H{}
	data["next_time"] = time.Now().Unix()
	Ok(c, "test success", data)
}
