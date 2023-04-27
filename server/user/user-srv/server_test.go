package user_srv

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
	"tiktok/pkg/model"
	user_proto "tiktok/server/user/api"
	"tiktok/server/user/dao"
	"time"
)

var client user_proto.UserClient

func init() {
	log.SetFlags(log.Lshortfile)
	conn, err := grpc.Dial("localhost:7777", grpc.WithInsecure())
	if err != nil {
		log.Fatal("fail to dial: ", err)
	}
	client = user_proto.NewUserClient(conn)

	go func() {
		dao.Init()
		lis, err := net.Listen("tcp", ":7777")
		if err != nil {
			log.Panicln("failed to listen: ", err)
		}
		s := grpc.NewServer()
		user_proto.RegisterUserServer(s, NewGrpcUserServer())
		if err = s.Serve(lis); err != nil {
			log.Panicln("failed to serve: ", err)
		}
		log.Println("测试grpc server退出")
	}()
	time.Sleep(2 * time.Second)
}

func TestGrpcUserServer_Register_Login(t *testing.T) {

	test := user_proto.RegisterRequest{Username: uuid.NewString()[:8], Password: "123456"}

	r, err := client.Register(context.Background(), &test)
	if err != nil {
		t.Fatal("注册失败", err)
	}

	id1 := r.UserId

	r, err = client.Register(context.Background(), &test)
	dao.GetMysqlCon().Table("users").Delete(model.User{}, id1)
	if err == nil {
		dao.GetMysqlCon().Table("users").Delete(model.User{}, r.UserId)
		t.Fatal("重复注册", err)
	}
}

func TestGrpcUserServer_Login(t *testing.T) {
	test := user_proto.LoginRequest{Username: uuid.NewString()[:8], Password: "123456"}
	t2 := user_proto.RegisterRequest{Username: test.Username, Password: test.Password}

	_, err := client.Login(context.Background(), &test)
	if err == nil {
		t.Fatal("可任意登陆", err)
	}

	reg, err := client.Register(context.Background(), &t2)
	if err != nil {
		t.Fatal("注册失败", err)
	}

	_, err = client.Login(context.Background(), &test)
	dao.GetMysqlCon().Table("users").Delete(model.User{}, reg.UserId)
	if err != nil {
		t.Fatal("无法正常登陆", err)
	}
}
