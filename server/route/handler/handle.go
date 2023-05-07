package handler

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net/http"
	"os"
	"strconv"
	"tiktok/pkg/model"
	comment_proto "tiktok/server/comment/api"
	user_proto "tiktok/server/user/api"
	"time"
)

type H struct {
	commentClient comment_proto.CommentClient
	userClient    user_proto.UserClient
}

func NewH() *H {
	ccComment, err := grpc.Dial("dns:///"+"service-comment-clusterip:7010", grpc.WithInsecure())
	if err != nil {
		log.Fatal("fail to dial: ", err)
	}

	kacp := keepalive.ClientParameters{
		Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
		Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
		PermitWithoutStream: true,             // send pings even without active streams
	}
	//myDnsResolver :=
	ccUser, err := grpc.Dial("dns:///"+"service-user-clusterip:7030",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithBlock(),
		grpc.WithResolvers(),
		grpc.WithKeepaliveParams(kacp),
		grpc.WithResolvers(),
	)
	if err != nil {
		log.Fatal("fail to dial: ", err)
	}

	return &H{
		commentClient: comment_proto.NewCommentClient(ccComment),
		userClient:    user_proto.NewUserClient(ccUser),
	}
}

func (h *H) Login(c *gin.Context) {
	host, _ := os.Hostname()
	log.Println("from route srv: ", host)
	name := c.Query("username")
	pwd := c.Query("password")

	in := user_proto.LoginRequest{
		Username: name,
		Password: pwd,
	}
	v, err := h.userClient.Login(c, &in)
	if err != nil {
		log.Println("dial user login srv failed", err)
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"StatusCode": 1,
			"StatusMsg":  "fail",
			"user_id":    nil,
			"token":      nil,
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"StatusCode": 0,
		"StatusMsg":  "success",
		"user_id":    v.UserId,
		"token":      v.Token,
	})
}

func (h *H) CommentList(c *gin.Context) {
	videoId := c.Query("video_id")
	userId, _ := c.Get("id")

	in := comment_proto.CommentListRequest{
		VideoId: videoId,
		MeId:    strconv.FormatInt(userId.(int64), 10),
	}
	v, err := h.commentClient.CommentList(c, &in)
	if err != nil {
		log.Println("dial comment list srv failed", err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"StatusCode": 1,
			"StatusMsg":  "fail",
			"comment":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"StatusCode":   0,
		"StatusMsg":    "success",
		"comment_list": v.CommentList,
	})
}

func (h *H) CommentAction(c *gin.Context) {
	actionType := c.Query("action_type")
	videoId := c.Query("video_id")
	commentText := c.Query("comment_text")
	userId, _ := c.Get("id")
	commentId := c.Query("comment_id")

	in := comment_proto.CommentActionRequest{
		CommentId:   commentId,
		VideoId:     videoId,
		CommentText: commentText,
		ActionType:  actionType,
		MeId:        strconv.FormatInt(userId.(int64), 10),
	}
	v, err := h.commentClient.CommentAction(c, &in)
	if err != nil {
		log.Println("dial comment action srv failed", err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"StatusCode": 1,
			"StatusMsg":  "fail to comment or del",
			"comment":    nil,
		})
		return
	}
	if actionType != "1" {
		c.JSON(http.StatusOK, map[string]interface{}{
			"StatusCode": 0,
			"StatusMsg":  "success",
			"comment":    nil,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"StatusCode": 0,
			"StatusMsg":  "success",
			"comment": model.CommentResponse{
				Id:         v.Id,
				User:       v.User,
				Content:    v.Content,
				CreateDate: v.CreateDate,
			},
		})
	}
}
