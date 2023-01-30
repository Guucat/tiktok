package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
)

var Oss *oss.Bucket

const (
	endpoint        = "http://oss-cn-hangzhou.aliyuncs.com"
	accessKeyId     = "LTAI5tArBU9y1AEGeJVV8YjP"
	accessKeySecret = "KS7csu53v4jK7f6BUHKsMhTCd2v6FS"
)

func Init() {

	log.Println("==============初始化oss连接=============")
	oss.Timeout(10, 120)
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		log.Println("fail to connect oss, timeout")
		panic(err)
	}

	bucket, err := client.Bucket("tiktok-syam")
	if err != nil {
		log.Println("fail to get bucket instance")
		panic(err)
	}
	Oss = bucket
	log.Printf("==============bucket name %s=============\n", bucket.BucketName)
}
