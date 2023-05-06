package main

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"google.golang.org/grpc"
	"log"
	"net"
	"tiktok/pkg/model"
	comment_proto "tiktok/server/comment/api"
	comment_srv "tiktok/server/comment/comment-srv"
	"tiktok/server/comment/dao"
)

func main() {
	// comment list cache aside【读多写多优化】 先写数据库【mq 分摊写压力】，再写缓存，设置较短缓存有效期 -> 最终一致性
	dao.Init()
	go Consumer("comment_action")

	lis, err := net.Listen("tcp", ":7010")
	if err != nil {
		log.Panicln("failed to listen: ", err)
	}
	s := grpc.NewServer()
	comment_proto.RegisterCommentServer(s, comment_srv.NewGrpcCommentServer())
	if err = s.Serve(lis); err != nil {
		log.Panicln("comment failed to serve: ", err)
	}
	log.Println("comment grpc server退出")

}

func Consumer(topic string) {
	consumer, err := sarama.NewConsumer([]string{"114.55.132.72:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to connect consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)

	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		//defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func() {
			for msg := range pc.Messages() {
				//fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				fmt.Println("消费一条消息", topic)
				ca := model.CommentAction{}
				_ = json.Unmarshal(msg.Value, &ca)
				// 增加评论
				if ca.ActionType == "1" {
					// 先写数据库
					if err := dao.AddComment(&ca); err != nil {
						fmt.Println("写数据库失败", err)
					}
					// 更新缓存
					if err := dao.PushCommentId(ca.VideoId); err != nil {
						fmt.Println("写comment id list缓存失败", err)
					}
					// 更新缓存
					if err := dao.IncrCommentCount(ca.VideoId); err != nil {
						fmt.Println("写comment count 缓存失败", err)
					}
				} else { // 删除评论
					// 先删数据库

					// 再检查缓存，再遍历缓存获取index，再ltrim 删缓存
				}

			}
			fmt.Println("done")
		}()
	}
}
