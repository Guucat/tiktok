package controller

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"strconv"
	"tiktok/dao/mysql"
	. "tiktok/mid"
	"tiktok/model"
	s "tiktok/service"
	"time"
)

func Upload(c *gin.Context) {
	videoId, err := s.GetStoreId() // 获取视频唯一id
	title := c.PostForm("title")
	authorId := c.GetInt64("id")
	if err != nil {
		Fail(c, "upload failed", nil)
		return
	}

	file, _ := c.FormFile("data") //获取视频数据
	if file == nil {
		log.Println("vedio is empty", err)
		Fail(c, "video is empty", nil)
		return
	}

	f, err := file.Open()
	if err != nil {
		log.Println("fail to load video data", err)
		Fail(c, "upload failed", nil)
		return
	}

	// upload asynchronization
	go func() {
		defer f.Close()
		//存储视频到本地
		videoWriter, path, imageId, err := s.CreateFile()
		if err != nil {
			log.Println("fail to save video data", err)
			Fail(c, "upload failed", nil)
			return
		}
		fReader := io.TeeReader(f, videoWriter)
		//上传视频,返回视频url
		videoUrl, err := s.StoreFileWithId(fReader, strconv.FormatInt(videoId, 10)+".mp4")
		if err != nil {
			Fail(c, "upload failed", nil)
			return
		}

		// 上传图片，返回图片url
		imageUrl, err := s.GetSnapshot(path, imageId)
		if err != nil {
			Fail(c, "upload failed", nil)
			return
		}

		vId, _ := strconv.ParseInt(strconv.FormatInt(videoId, 10), 10, 64)
		videoInfo := &model.Video{
			Id:       vId,
			AuthorId: authorId,
			PlayUrl:  videoUrl,
			CoverUrl: imageUrl,
			Title:    title,
		}
		err = s.SaveVideoInfo(videoInfo)
		if err != nil {
			Fail(c, "upload failed", nil)
			return
		}
	}()

	Ok(c, "success to upload file", nil)
}

func GetFeed(c *gin.Context) {
	data := gin.H{}
	data["next_time"] = time.Now().Unix()
	Ok(c, "test success", data)
}

func List(c *gin.Context) {
	userId := c.Query("user_id")
	videos, err := s.GetVideoList(userId)
	if err != nil {
		log.Println("fail to get video list", err)
		Fail(c, "user doesn't exist", gin.H{"video_list": nil})
		return
	}

	userDao, err := s.GetAuthorMessage(userId)
	if err != nil {
		log.Println("fail to get user info", err)
		Fail(c, "user doesn't exist", gin.H{"video_list": nil})
		return
	}
	isFollow := mysql.GetIsFollower(userId, userId)
	user := User{
		Id:              userDao.Id,
		Name:            userDao.Username,
		FollowCount:     userDao.FollowCount,
		FollowerCount:   userDao.FollowerCount,
		Avatar:          userDao.Avatar,
		BackgroundImage: userDao.BackgroundImage,
		Signature:       userDao.Signature,
		TotalFavorited:  userDao.TotalFavorited,
		WorkCount:       userDao.WorkCount,
		FavoriteCount:   userDao.FavoriteCount,
		IsFollow:        isFollow,
	}

	videoList := make([]Video, 0, 20)
	for _, v := range videos {
		isFavorite := mysql.IsFavorite(userId, v.Id)
		video := Video{
			Id:            v.Id,
			Title:         v.Title,
			Author:        user,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    isFavorite,
		}
		videoList = append(videoList, video)
	}
	Ok(c, "success", gin.H{"video_list": videoList})
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
		userDao, err := s.GetAuthorMessage(v.AuthorId)
		if err != nil {
			log.Println("fail to get user info", err)
			Fail(c, "user doesn't exist", gin.H{"" +
				"next_time": nil,
				"video_list": nil,
			})
			return
		}
		isFollow := mysql.GetIsFollower(id, v.AuthorId)
		user := User{
			Id:              userDao.Id,
			Name:            userDao.Username,
			FollowCount:     userDao.FollowCount,
			FollowerCount:   userDao.FollowerCount,
			Avatar:          userDao.Avatar,
			BackgroundImage: userDao.BackgroundImage,
			Signature:       userDao.Signature,
			TotalFavorited:  userDao.TotalFavorited,
			WorkCount:       userDao.WorkCount,
			FavoriteCount:   userDao.FavoriteCount,
			IsFollow:        isFollow,
		}

		list[i] = Video{
			Id:            v.Id,
			Title:         v.Title,
			Author:        user,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    mysql.IsFavorite(id, v.Id),
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
