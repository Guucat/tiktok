package snowflake_srv

import (
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"google.golang.org/protobuf/types/known/emptypb"
	snowflake_proto "tiktok/server/snowflake/api"
)

type SnowflakeServer struct {
	snowflake_proto.UnimplementedSnowflakeServer
	node *snowflake.Node
}

func NewSnowflakeServer() *SnowflakeServer {
	d, _ := snowflake.NewNode(1)
	return &SnowflakeServer{
		node: d,
	}
}

func (s *SnowflakeServer) GetId(context.Context, *emptypb.Empty) (*snowflake_proto.IdResponse, error) {
	id := s.node.Generate().Int64()
	fmt.Println("success id", id)
	return &snowflake_proto.IdResponse{
		Id: id,
	}, nil
}
