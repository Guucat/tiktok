package user_srv

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"log"
	"strconv"
	"tiktok/pkg/cache"
	"tiktok/pkg/jwt"
	"tiktok/pkg/model"
	"time"

	user_proto "tiktok/server/user/api"
	"tiktok/server/user/dao"
)

type GrpcUserServer struct {
	user_proto.UnimplementedUserServer
	mysql *gorm.DB
	redis *redis.Client
}

func NewGrpcUserServer() *GrpcUserServer {
	return &GrpcUserServer{
		mysql: dao.GetMysqlCon(),
		redis: dao.GetRedisCon(),
	}
}

func (g *GrpcUserServer) Register(c context.Context, r *user_proto.RegisterRequest) (*user_proto.RegisterResponse, error) {
	user := model.User{}
	if err := g.mysql.Where("username = ?", r.Username).First(&user).Error; err == nil {
		return nil, errors.New("账号已被注册")
	}

	user = model.User{Username: r.Username, Password: r.Password}
	if err := g.mysql.Table("users").Create(&user).Error; err != nil {
		log.Println(err)
		return nil, errors.New("数据库插入失败")
	}

	token, _ := jwt.GenToken(user.Id)
	return &user_proto.RegisterResponse{
		UserId: user.Id,
		Token:  token,
	}, nil
}

func (g *GrpcUserServer) Login(c context.Context, r *user_proto.LoginRequest) (*user_proto.LoginResponse, error) {
	var id int64
	if err := g.mysql.Table("users").Select("id").Where("username = ? and password = ?", r.Username, r.Password).Find(&id).Limit(1).Error; err != nil || id == 0 {
		return nil, errors.New("账号或密码错误")
	}

	fmt.Println(id)
	token, _ := jwt.GenToken(id)
	return &user_proto.LoginResponse{
		UserId: id,
		Token:  token,
	}, nil
}

func (g *GrpcUserServer) GetUserInfo(c context.Context, r *user_proto.GetUserInfoRequest) (*user_proto.GetUserInfoResponse, error) {
	t1 := time.Now()
	id, _ := strconv.ParseInt(r.UserId, 10, 64)
	baseUser := model.UserBaseInfo{Id: id}
	v := g.redis.Get(c, "user_base_info_"+r.UserId)
	if v.Err() != nil {
		g.mysql.Table("users").Select("username", "avatar", "background_image", "signature").Find(&baseUser)
		b, _ := json.Marshal(&baseUser)
		g.redis.Set(c, "user_base_info_"+r.UserId, b, cache.RandExpiredTimeSec(8, 16))
		log.Println("baseUser缓存Miss")
	} else {
		bytes, _ := v.Bytes()
		_ = json.Unmarshal(bytes, &baseUser)
		log.Println("baseUser缓存Hit")
	}

	vals := g.redis.MGet(c,
		"total_favorited_"+r.UserId,
		"work_count_"+r.UserId,
		"favorite_count_"+r.UserId,
		"follow_count_"+r.UserId,
		"follower_count_"+r.UserId,
	).Val()
	var totalFavorited int64
	var workCount int64
	var favoriteCount int64
	var followCount int64
	var followerCount int64
	if vals[0] == nil {
		g.mysql.Table("users").Select("total_favorited").Where("id = ?", r.UserId).Find(&totalFavorited)
		g.redis.Set(c, "total_favorited_"+r.UserId, totalFavorited, cache.RandExpiredTimeSec(8, 16))
		log.Println("0缓存Miss")
	} else {
		totalFavorited, _ = strconv.ParseInt(vals[0].(string), 10, 64)
		log.Println("0缓存Hit")
	}
	if vals[1] == nil {
		g.mysql.Table("users").Select("work_count").Where("id = ?", r.UserId).Find(&workCount)
		g.redis.Set(c, "work_count_"+r.UserId, workCount, cache.RandExpiredTimeSec(8, 16))
		log.Println("1缓存Miss")
	} else {
		workCount, _ = strconv.ParseInt(vals[1].(string), 10, 64)
		log.Println("1缓存Hit")
	}
	if vals[2] == nil {
		g.mysql.Table("users").Select("favorite_count").Where("id = ?", r.UserId).Find(&favoriteCount)
		g.redis.Set(c, "favorite_count_"+r.UserId, favoriteCount, cache.RandExpiredTimeSec(8, 16))
		log.Println("2缓存Miss")
	} else {
		favoriteCount, _ = strconv.ParseInt(vals[2].(string), 10, 64)
		log.Println("2缓存Hit")
	}
	if vals[3] == nil {
		g.mysql.Table("users").Select("follow_count").Where("id = ?", r.UserId).Find(&followCount)
		g.redis.Set(c, "follow_count_"+r.UserId, followCount, cache.RandExpiredTimeSec(8, 16))
		log.Println("3缓存Miss")
	} else {
		followCount, _ = strconv.ParseInt(vals[3].(string), 10, 64)
		log.Println("3缓存Hit")
	}
	if vals[4] == nil {
		g.mysql.Table("users").Select("follower_count").Where("id = ?", r.UserId).Find(&followerCount)
		g.redis.Set(c, "follower_count_"+r.UserId, followerCount, cache.RandExpiredTimeSec(8, 16))
		log.Println("4缓存Miss")
	} else {
		followerCount, _ = strconv.ParseInt(vals[4].(string), 10, 64)
		log.Println("4缓存Hit")
	}
	// 社交模块rpc方法，暂未实现
	isFollow := false
	fmt.Println(time.Now().Sub(t1))
	return &user_proto.GetUserInfoResponse{
		Id:              baseUser.Id,
		Name:            baseUser.Username,
		Avatar:          baseUser.Avatar,
		BackgroundImage: baseUser.BackgroundImage,
		Signature:       baseUser.Signature,
		FollowCount:     followCount,
		FollowerCount:   followerCount,
		IsFollow:        isFollow,
		TotalFavorited:  totalFavorited,
		WorkCount:       workCount,
		FavoriteCount:   favoriteCount,
	}, nil
}
