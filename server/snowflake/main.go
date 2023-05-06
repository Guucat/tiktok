package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	snowflake_proto "tiktok/server/snowflake/api"
	snowflake_srv "tiktok/server/snowflake/snowflake-srv"
)

func main() {
	lis, err := net.Listen("tcp", ":7020")
	if err != nil {
		log.Panicln("failed to listen: ", err)
	}
	s := grpc.NewServer()
	snowflake_proto.RegisterSnowflakeServer(s, snowflake_srv.NewSnowflakeServer())
	log.Fatal(s.Serve(lis))
}
