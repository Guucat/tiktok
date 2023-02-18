package controller

import (
	"github.com/gin-gonic/gin"
	"tiktok/service"
)

type FollowActionResponse struct {
	StatusCode int 32 json:"status_code"
	StatusMsg string json:"status_msg, omitempty"
}

func FollowAction(c *gin.Context){
	toUserId := c.Query("to_user_id") //所关注用户id
	actionType := c.Query("action_type") //1-关注，2-取消关注
	userID, _ := c.Get("id")

	/*

	 */
	if actionType == '1'{
		service.
	}
}
