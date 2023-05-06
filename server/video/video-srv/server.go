package video_srv

import (
	"context"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"log"
	"strconv"
	"tiktok/pkg/model"
	"tiktok/server/video/dao"

	video_proto "tiktok/server/video/api"
)

type GrpcVideoServer struct {
	video_proto.UnimplementedVideoServer
	mysql *gorm.DB
	redis *redis.Client
}

func NewGrpcVideoServer() *GrpcVideoServer {
	return &GrpcVideoServer{
		mysql: dao.GetMysqlCon(),
		redis: dao.GetRedisCon(),
	}
}

func (g *GrpcVideoServer) SaveVideoInfo(c context.Context, r *video_proto.SaveVideoInfoRequest) (*emptypb.Empty, error) {
	id, _ := strconv.ParseInt(r.Id, 10, 64)
	aId, _ := strconv.ParseInt(r.AuthorId, 10, 64)
	err := dao.InsertVideo(&model.Video{
		Id:       id,
		Title:    r.Title,
		AuthorId: aId,
		PlayUrl:  r.PlayUrl,
		CoverUrl: r.CoverUrl,
	})
	if err != nil {
		log.Panicln("上传视频信息保存失败")
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (g *GrpcVideoServer) GetVideoInfo(c context.Context, r *video_proto.VideoInfoRequest) (*video_proto.VideoInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoInfo not implemented")
}
func (g *GrpcVideoServer) IncrVideoLikeCount(c context.Context, r *video_proto.VideoId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrVideoLikeCount not implemented")
}
func (g *GrpcVideoServer) DecrVideoLikeCount(c context.Context, r *video_proto.VideoId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecrVideoLikeCount not implemented")
}
func (g *GrpcVideoServer) IncrVideoCommentCount(c context.Context, r *video_proto.VideoId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrVideoCommentCount not implemented")
}
func (g *GrpcVideoServer) DecrVideoCommentCount(c context.Context, r *video_proto.VideoId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecrVideoCommentCount not implemented")
}
