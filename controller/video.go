package controller

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"strconv"
	. "tiktok/mid"
	"tiktok/model"
	s "tiktok/service"
	"time"
)

func Upload(c *gin.Context) {
	videoId, err := s.GetStoreId() // 获取视频唯一id
	if err != nil {
		Fail(c, "upload failed", nil)
		return
	}

	file, _ := c.FormFile("data") //获取视频数据
	f, err := file.Open()
	if err != nil {
		log.Println("fail to load video data", err)
		Fail(c, "upload failed", nil)
		return
	}
	defer f.Close()

	videoWriter, path, imageId, err := s.CreateFile() //存储视频到本地
	if err != nil {
		log.Println("fail to save video data", err)
		Fail(c, "upload failed", nil)
		return
	}
	fReader := io.TeeReader(f, videoWriter)

	videoUrl, err := s.StoreFileWithId(fReader, strconv.FormatInt(videoId, 10)+".mp4") //上传视频
	if err != nil {
		Fail(c, "upload failed", nil)
		return
	}

	imageUrl, err := s.GetSnapshot(path, imageId) // 上传图片
	if err != nil {
		Fail(c, "upload failed", nil)
		return
	}

	title := c.PostForm("title")
	vId, _ := strconv.ParseInt(strconv.FormatInt(videoId, 10), 10, 64)
	videoInfo := &model.Video{
		Id:       vId,
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

func List(c *gin.Context) {
	id := c.Query("user_id")
	videos, err := s.GetVideoList(id)
	if err != nil {
		log.Println("fail to get video list", err)
		Fail(c, "user doesn't exist", gin.H{"video_list": nil})
		return
	}
	list := make([]Video, len(videos))
	for i, v := range videos {
		author, err := s.GetUserInfo(c.GetString("id"), id)
		if err != nil {
			log.Println("fail to get user info", err)
			Fail(c, "user doesn't exist", gin.H{"video_list": nil})
			return
		}
		list[i] = Video{
			Id:    v.Id,
			Title: v.Title,
			Author: User{
				Id:            author["id"].(int64),
				Name:          author["name"].(string),
				FollowCount:   author["follow_count"].(int64),
				FollowerCount: author["follower_count"].(int64),
				IsFollow:      author["is_follow"].(bool),
			},
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    s.IsFavorite(id, strconv.FormatInt(v.Id, 10)),
		}
	}
	Ok(c, "success", gin.H{"video_list": list})
}

func Feed(c *gin.Context) {
	stamp := c.Query("latest_time")
	start := time.Now().Format("2006-01-02 15:04:05")
	if stamp != "" {
		timeStamp, _ := strconv.ParseInt(stamp, 10, 64)
		start = time.Unix(timeStamp, 0).Format("2006-01-02 15:04:05")
	}
	videos, err := s.GetVideoFeed(start)
	if err != nil {
		log.Println("fail to get video list", err)
		Fail(c, "user doesn't exist", gin.H{
			"next_time":  nil,
			"video_list": nil,
		})
		return
	}

	id := c.GetString("id")
	if id == "" {
		id = "0"
	}
	list := make([]Video, len(videos))
	for i, v := range videos {
		author, err := s.GetUserInfo(id, strconv.FormatInt(v.AuthorId, 10))
		if err != nil {
			log.Println("fail to get user info", err)
			Fail(c, "user doesn't exist", gin.H{"" +
				"next_time": nil,
				"video_list": nil,
			})
			return
		}
		list[i] = Video{
			Id:    v.Id,
			Title: v.Title,
			Author: User{
				Id:            author["id"].(int64),
				Name:          author["name"].(string),
				FollowCount:   author["follow_count"].(int64),
				FollowerCount: author["follower_count"].(int64),
				IsFollow:      author["is_follow"].(bool),
			},
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    s.IsFavorite(id, strconv.FormatInt(v.Id, 10)),
		}
	}

	// 如果没有新视频, 循环获取
	nextTime := time.Now().Unix()
	if len(videos) > 0 {
		nextTime = videos[len(videos)-1].CreateTime.Unix()
	}
	Ok(c, "success", gin.H{
		"next_time":  nextTime,
		"video_list": list,
	})
}
