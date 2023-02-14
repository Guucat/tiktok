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

func List(c *gin.Context) {
	id := c.GetString("id")
	videos, err := s.GetVideoList(id)
	if err != nil {
		log.Println("fail to get video list", err)
		Fail(c, "user doesn't exist", gin.H{"video_list": nil})
		return
	}
	list := make([]Video, len(videos))
	for i, v := range videos {
		author, err := s.GetUserInfo(id, id)
		if err != nil {
			log.Println("fail to get user info", err)
			Fail(c, "user doesn't exist", gin.H{"video_list": nil})
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
	start := time.Now()
	if stamp != "" {
		timeStamp, _ := strconv.ParseInt(stamp, 10, 64)
		start = time.Unix(timeStamp, 0)
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
	list := make([]Video, len(videos))
	for i, v := range videos {
		author, err := s.GetUserInfo(id, id)
		if err != nil {
			log.Println("fail to get user info", err)
			Fail(c, "user doesn't exist", gin.H{"" +
				"next_time": nil,
				"video_list": nil,
			})
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
	nextTime := time.Now()
	if len(videos) > 0 {
		nextTime = videos[len(videos)-1].CreateTime
	}
	Ok(c, "success", gin.H{
		"next_time":  nextTime,
		"video_list": list,
	})
}
