package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"tiktok/service"
)

// FollowAction 关注操作
func FollowAction(c *gin.Context) {
	toUserIdTemp := c.Query("to_user_id") //所关注用户id
	actionType := c.Query("action_type")  //操作类型，1-关注，2-取消关注
	userID, _ := c.Get("id")

	toUserId, _ := strconv.Atoi(toUserIdTemp) //传递的“所关注用户id”type为string，为方便后续sql语句使用，需转为int类型

	// 关注操作逻辑
	if actionType == "1" {
		if err := service.FollowAction(toUserId, userID); err != nil {
			log.Println("Following Error", err)
			c.JSON(http.StatusOK, map[string]interface{}{
				"status_code": 1,
				"status_msg":  "Fail to follow",
			})
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"status_code": 0,
			"status_msg":  "success to follow",
		})
	}

	// 取消关注操作逻辑
	if actionType == "2" {
		if err := service.CanFollowAction(toUserId, userID); err != nil {
			log.Println("Following Error", err)
			c.JSON(http.StatusOK, map[string]interface{}{
				"status_code": 1,
				"status_msg":  "Fail to cancel follow",
			})
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"status_code": 0,
			"status_msg":  "success to cancel follow",
		})
	}
}

//// FollowList 关注列表
//func FollowList(c *gin.Context)  {
//	userID, _ := c.Get("id")
//}
//// FollowerList 粉丝列表
//func FollowerList(c *gin.Context){
//	userID, _ := c.Get("id")
//}
