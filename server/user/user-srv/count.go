package user_srv

import (
	"context"
	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/types/known/emptypb"
	user_proto "tiktok/server/user/api"
)

func (g *GrpcUserServer) IncrTotalFavorited(c context.Context, r *user_proto.OpCountRequest) (*emptypb.Empty, error) {
	return nil, incr(c, r.GetUserId(), "total_favorited", g.redis)

}
func (g *GrpcUserServer) DecrTotalFavorited(c context.Context, r *user_proto.OpCountRequest) (*emptypb.Empty, error) {
	return nil, decr(c, r.GetUserId(), "total_favorited", g.redis)
}
func (g *GrpcUserServer) IncrWorkCount(c context.Context, r *user_proto.OpCountRequest) (*emptypb.Empty, error) {
	return nil, incr(c, r.GetUserId(), "work_count", g.redis)
}
func (g *GrpcUserServer) DecrWorkCount(c context.Context, r *user_proto.OpCountRequest) (*emptypb.Empty, error) {
	return nil, decr(c, r.GetUserId(), "work_count", g.redis)
}
func (g *GrpcUserServer) IncrFavoriteCount(c context.Context, r *user_proto.OpCountRequest) (*emptypb.Empty, error) {
	return nil, incr(c, r.GetUserId(), "favorite_count", g.redis)
}
func (g *GrpcUserServer) DecrFavoriteCount(c context.Context, r *user_proto.OpCountRequest) (*emptypb.Empty, error) {
	return nil, decr(c, r.GetUserId(), "favorite_count", g.redis)
}
func (g *GrpcUserServer) IncrFollowCount(c context.Context, r *user_proto.OpCountRequest) (*emptypb.Empty, error) {
	return nil, incr(c, r.GetUserId(), "follow_count", g.redis)
}
func (g *GrpcUserServer) DecrFollowCount(c context.Context, r *user_proto.OpCountRequest) (*emptypb.Empty, error) {
	return nil, decr(c, r.GetUserId(), "follow_count", g.redis)
}
func (g *GrpcUserServer) IncrFollowerCount(c context.Context, r *user_proto.OpCountRequest) (*emptypb.Empty, error) {
	return nil, incr(c, r.GetUserId(), "follower_count_", g.redis)
}
func (g *GrpcUserServer) DecrFollowerCount(c context.Context, r *user_proto.OpCountRequest) (*emptypb.Empty, error) {
	return nil, decr(c, r.GetUserId(), "follower_count_", g.redis)
}

func incr(c context.Context, id string, filed string, redis *redis.Client) error {
	return redis.Incr(c, filed+"_"+id).Err()
}

func decr(c context.Context, id string, filed string, redis *redis.Client) error {
	return redis.Decr(c, filed+"_"+id).Err()
}
