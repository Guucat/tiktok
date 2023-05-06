package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	user_proto "tiktok/server/user/api"
	"tiktok/server/user/dao"
	"tiktok/server/user/user-srv"
)

func main() {
	dao.Init()
	lis, err := net.Listen("tcp", ":7030")
	if err != nil {
		log.Panicln("failed to listen: ", err)
	}
	s := grpc.NewServer()
	user_proto.RegisterUserServer(s, user_srv.NewGrpcUserServer())
	log.Fatal(s.Serve(lis))
}
