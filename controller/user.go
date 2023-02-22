package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"tiktok/dao/mysql"
	. "tiktok/mid"
	"tiktok/mid/jwt"
	. "tiktok/mid/validate"
	"tiktok/model"
	s "tiktok/service"
)

// Register TODO 获取user id时会回表，效率低如何只操作1次数据库 ok
func Register(c *gin.Context) {
	data := gin.H{
		"user_id": 0,
		"token":   "",
	}
	// Get and verify  user info
	name := c.Query("username")
	pwd := c.Query("password")
	err := Validate.Struct(&model.User{Username: name, Password: pwd})
	if err != nil {
		Fail(c, "Registration failed The user name or password is empty", data)
		return
	}

	// Register user
	user := &model.User{
		Username: name,
		Password: pwd,
	}
	if err = s.Register(user); err != nil {
		Fail(c, "Registration failed Account has been registered", data)
		return
	}

	data["user_id"] = user.Id
	data["token"], _ = jwt.GenToken(user.Id)
	Ok(c, "registered successfully", data)
}

func Login(c *gin.Context) {
	data := gin.H{
		"user_id": 0,
		"token":   "",
	}
	// // Get user info
	name := c.Query("username")
	pwd := c.Query("password")
	err := Validate.Struct(&model.User{Username: name, Password: pwd})
	if err != nil {
		Fail(c, "Login failed The user name or password is empty", data)
		return
	}

	// Verify the username and password
	user, err := s.VerifyUser(name, pwd)
	if err != nil {
		Fail(c, "Incorrect username or password", data)
		return
	}

	// Generate token
	data["token"], _ = jwt.GenToken(user.Id)
	data["user_id"] = user.Id
	Ok(c, "login successfully", data)
}

// UserInfo 获取用户信息
func UserInfo(c *gin.Context) {
	data := gin.H{"user": nil}
	authId, _ := c.Get("id")
	userId := c.Query("user_id")
	if strconv.FormatInt(authId.(int64), 10) != userId {
		log.Println("非法参数")
	}
	userDao, err := s.GetAuthorMessage(userId)
	if err != nil {
		Fail(c, err.Error(), data)
		return
	}

	isFollower := mysql.GetIsFollower(userId, userId)
	user := User{
		Id:              userDao.Id,
		Name:            userDao.Username,
		FollowCount:     userDao.FollowCount,
		FollowerCount:   userDao.FollowerCount,
		Avatar:          userDao.Avatar,
		BackgroundImage: userDao.BackgroundImage,
		Signature:       userDao.Signature,
		TotalFavorited:  userDao.TotalFavorited,
		WorkCount:       userDao.WorkCount,
		FavoriteCount:   userDao.FavoriteCount,
		IsFollow:        isFollower,
	}

	data["user"] = user
	Ok(c, "user info", data)
}
