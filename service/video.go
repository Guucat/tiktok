package service

import (
	"bytes"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"log"
	"os"
	"strconv"
	"tiktok/dao/mysql"
	. "tiktok/mid/oss"
	"tiktok/model"
	"time"
)

const path = "./"

func GetStoreId() (int64, error) {
	// 单机版机器id固定
	node, err := snowflake.NewNode(0)
	if err != nil {
		log.Println("fail to generate uuid", err)
		return -1, err
	}
	return node.Generate().Int64(), err
}

// StoreFileWithId 上传文件至oss,返回访问url
func StoreFileWithId(f io.Reader, id string) (url string, err error) {
	err = Oss.PutObject(id, f)
	if err != nil {
		log.Println("oss upload failed", err)
	}
	return Root + id, nil
}

// GetSnapshot 截取视频第一帧作为封面，上传至oss并返回访问地址
func GetSnapshot(videoPath, id string) (url string, err error) {
	// 截取视频第一帧作为封面
	buf := bytes.NewBuffer(nil)
	frameNum := 1
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n, %d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Println("ffmpeg fail to input video: ", err)
		return
	}
	img, err := imaging.Decode(buf)
	if err != nil {
		log.Println("fail to decode video stream data: ", err)
		return
	}

	// 存储封面图片到本地
	imagePath := path + id + ".png"
	err = imaging.Save(img, imagePath)
	if err != nil {
		log.Println("fail to save image: ", err)
		return
	}
	defer os.Remove(imagePath)
	defer os.Remove(videoPath)

	// 上传图片至oss
	imageFH, _ := os.Open(imagePath)
	return StoreFileWithId(imageFH, id+".png")
}

func SaveVideoInfo(v *model.Video) error {
	return mysql.InsertVideo(v)
}

func GetVideoList(id string) ([]model.Video, error) {
	return mysql.QueryVideoList(id, nil)
}

func IsFavorite(userId, videoId string) bool {
	return mysql.QueryFavorite(userId, videoId)
}

func GetVideoFeed(start time.Time) ([]model.Video, error) {
	return mysql.QueryVideoList("", start)
}

func CreateFile() (io.Writer, string, string, error) {
	id, err := GetStoreId()
	strId := strconv.FormatInt(id, 10)
	if err != nil {
		return nil, "", "", err
	}
	filePath := path + strId + ".mp4"
	localFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("fail to create file", err)
		return nil, "", "", err
	}
	return localFile, filePath, strId, nil
}
