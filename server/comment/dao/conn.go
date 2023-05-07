package dao

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"tiktok/pkg/model"
	"time"
)

var mysqlCon *gorm.DB
var redisCon *redis.Client
var kafkaCon sarama.SyncProducer

func GetMysqlCon() *gorm.DB {
	return mysqlCon
}

func GetRedisCon() *redis.Client {
	return redisCon
}

func GetKafkaCon() sarama.SyncProducer {
	return kafkaCon
}
func Init() {
	// mysql
	dsn := "shengyi:123456@tsy@tcp(rm-2vc34w5spf5nm2992eo.mysql.cn-chengdu.rds.aliyuncs.com)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(1000)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(2000)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	mysqlCon = db

	// redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "114.55.132.72:6379",
		Password: "",
		DB:       7,
		PoolSize: 100,
	})
	redisCon = rdb

	//
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	kafkaCon, err = sarama.NewSyncProducer([]string{"114.55.132.72:9092"}, config)
	if err != nil {
		panic("fail to connect kafka " + err.Error())
	}

}

func AddComment(c *model.CommentAction) (err error) {
	//开启事务
	err = mysqlCon.Transaction(func(tx *gorm.DB) error {
		//存入评论数据
		if err = tx.Table("comment").Create(map[string]interface{}{
			"id":          c.ContentId,
			"user_id":     c.MeId,
			"video_id":    c.VideoId,
			"content":     c.Content,
			"create_time": c.TimeDate,
			"update_time": c.TimeDate,
		}).Error; err != nil {
			log.Println("Fail to comment", err)
			return err
		}
		//增加视频评论数
		if err = tx.Table("videos").Where("id = ?", c.VideoId).
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

// IncrCommentCount 缓存不存在只写数据库
func IncrCommentCount(videoId string) error {
	prefix := "comment_count_"
	if redisCon.Exists(context.Background(), prefix+videoId).Val() != 1 {
		//n := 0
		//mysqlCon.Table("videos").Select("comment_count").Where("id = ?", videoId).Find(&n)
		//return redisCon.Set(context.Background(), prefix+videoId, n+1, cache.RandExpiredTimeSec(1, 1)).Err()

		return nil
	}
	return redisCon.Incr(context.Background(), prefix+videoId).Err()
}

// PushCommentId 缓存不存在只写数据库v
func PushCommentId(videoId string) error {
	prefix := "comment_list_"
	//commentIds := make([]int64, 0)
	if redisCon.Exists(context.Background(), prefix+videoId).Val() != 1 {
		//if err := mysqlCon.Table("comment").Select("id").Where("video_id = ? AND state = ?", videoId, 1).Order("create_time DESC").
		//	Find(&commentIds).Error; err != nil {
		//	log.Println(err)
		//	return err
		//}

		// 评论数为0
		//if len(commentIds) == 0 {
		//	return nil
		//}
		//redisCon.RPush(context.Background(), prefix+videoId, commentIds, cache.RandExpiredTimeSec(1, 1))
		return nil
	}
	return redisCon.RPush(context.Background(), prefix+videoId, videoId).Err()
}
