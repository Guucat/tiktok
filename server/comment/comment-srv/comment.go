package comment_srv

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"tiktok/pkg/model"
	comment_proto "tiktok/server/comment/api"
	"tiktok/server/comment/dao"
	snowflake_proto "tiktok/server/snowflake/api"
	user_proto "tiktok/server/user/api"
	"time"
)

type GrpcCommentServer struct {
	comment_proto.UnimplementedCommentServer
	mysql           *gorm.DB
	redis           *redis.Client
	kafka           sarama.SyncProducer
	userClient      user_proto.UserClient
	snowflakeClient snowflake_proto.SnowflakeClient
}

func NewGrpcCommentServer() *GrpcCommentServer {
	conn1, err := grpc.Dial("localhost:7030", grpc.WithInsecure())
	if err != nil {
		log.Fatal("fail to dial user src: ", err)
	}

	conn2, err := grpc.Dial("localhost:7020", grpc.WithInsecure())
	if err != nil {
		log.Fatal("fail to dial snowflake.yaml srv: ", err)
	}
	return &GrpcCommentServer{
		mysql:           dao.GetMysqlCon(),
		redis:           dao.GetRedisCon(),
		kafka:           dao.GetKafkaCon(),
		userClient:      user_proto.NewUserClient(conn1),
		snowflakeClient: snowflake_proto.NewSnowflakeClient(conn2),
	}
}

func (g *GrpcCommentServer) SendMassage(topic string, data string) error {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	// 发送消息
	_, _, err := g.kafka.SendMessage(msg)
	if err == nil {
		fmt.Println("消息投递成功", topic)
	}
	return err
}

func (g *GrpcCommentServer) CommentAction(c context.Context, r *comment_proto.CommentActionRequest) (*comment_proto.CommentActionResponse, error) {
	host, _ := os.Hostname()
	log.Println("from comment srv: ", host)

	ca := model.CommentAction{
		VideoId:    r.VideoId,
		MeId:       r.MeId,
		ActionType: r.ActionType,
		Content:    r.CommentText,
		ContentId:  r.CommentId,
		TimeDate:   time.Now().Format("2006-01-02 15:04:05"),
	}
	if r.ActionType == "1" {
		res, err := g.snowflakeClient.GetId(c, &empty.Empty{})
		if err != nil {
			log.Println("fail to dial id srv", err)
			return nil, err
		}
		ca.ContentId = strconv.FormatInt(res.Id, 10)
	}
	b, _ := json.Marshal(ca)
	commentMes := string(b)
	err := g.SendMassage("comment_action", commentMes)
	if err != nil {
		log.Println("comment_action 投递消息失败", err)
	}

	id, _ := strconv.ParseInt(ca.ContentId, 10, 64)
	resp := comment_proto.CommentActionResponse{
		Id:         id,
		Content:    r.CommentText,
		CreateDate: ca.TimeDate,
	}
	resp.Id = id
	resp.Content = r.CommentText
	resp.CreateDate = ca.TimeDate

	return &resp, nil

}
func (g *GrpcCommentServer) CommentList(c context.Context, r *comment_proto.CommentListRequest) (*comment_proto.CommentListResponse, error) {
	host, _ := os.Hostname()
	log.Println("from comment srv: ", host)

	keyId := "comment_list_" + r.VideoId
	commentIds := make([]int64, 0)
	var results []*comment_proto.CommentActionResponse

	// id list 缓存
	if g.redis.Exists(c, keyId).Val() != 1 {
		fmt.Println("id list 缓存miss")
		if err := g.mysql.Table("comment").Select("id").Where("video_id = ? AND state = ?", r.VideoId, 1).Order("create_time DESC").
			Find(&commentIds).Error; err != nil {
			log.Println(err)
			return nil, err
		}
		// 评论数为0
		if len(commentIds) == 0 {
			return &comment_proto.CommentListResponse{}, nil
		}
		pipe := g.redis.Pipeline()
		for _, v := range commentIds {
			pipe.RPush(c, keyId, v)
		}
		//err := g.redis.RPush(c, keyId, commentIds).Err()
		_, err := pipe.Exec(c)
		if err != nil {
			fmt.Println("id list 缓存设置失败", err)
		}
	} else {
		fmt.Println("id list 缓存 hit")
		res := g.redis.LRange(c, keyId, 0, -1).Val()
		for _, id := range res {
			i, _ := strconv.ParseInt(id, 10, 64)
			commentIds = append(commentIds, i)
		}
	}

	for _, id := range commentIds {
		keyInfo := "comment_info_" + strconv.FormatInt(int64(id), 10)
		info := model.Comment{}
		// comment info 缓存
		if g.redis.Exists(c, keyInfo).Val() != 1 {
			fmt.Println("single comment info 缓存miss")
			if err := g.mysql.Table("comment").Select("user_id, create_time, content").Where("id = ?", id).
				Find(&info).Error; err != nil {
				log.Println(err)
				return nil, err
			}
			if err := g.redis.HMSet(c, keyInfo, "user_id", info.UserId, "create_date", info.CreateDate, "content", info.Content).Err(); err != nil {
				log.Println("list userinfo sat failed", err)
			}
		} else {
			fmt.Println("single comment info 缓存hit")
			vals := g.redis.HMGet(c, keyInfo, "user_id", "create_date", "content").Val()
			info.UserId, _ = strconv.ParseInt(vals[0].(string), 10, 64)
			info.CreateDate = vals[1].(string)
			info.Content = vals[2].(string)
		}

		user, err := g.userClient.GetUserInfo(c, &user_proto.GetUserInfoRequest{
			UserId: strconv.FormatInt(info.UserId, 10),
			MeId:   r.MeId,
		})
		if err != nil {
			log.Println("dial rpc GetUserInfo failed", err)
		}

		result := &comment_proto.CommentActionResponse{
			Id:         id,
			Content:    info.Content,
			CreateDate: info.CreateDate,
			User: &comment_proto.UserInfo{
				Id:              user.Id,
				Name:            user.Name,
				Avatar:          user.Avatar,
				BackgroundImage: user.BackgroundImage,
				Signature:       user.Signature,
				FollowCount:     user.FollowCount,
				FollowerCount:   user.FollowerCount,
				IsFollow:        user.IsFollow,
				TotalFavorited:  user.TotalFavorited,
				WorkCount:       user.WorkCount,
				FavoriteCount:   user.FavoriteCount,
			},
		}
		results = append(results, result)
	}

	return &comment_proto.CommentListResponse{CommentList: results}, nil
}
