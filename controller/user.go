package controller

import (
	"github.com/gin-gonic/gin"
	. "tiktok/mid"
	"tiktok/mid/jwt"
	. "tiktok/mid/validate"
	"tiktok/model"
	s "tiktok/service"
)

// Register TODO 获取user id时会回表，效率低如何只操作1次数据库？
func Register(c *gin.Context) {
	data := gin.H{
		"user_id": "",
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
	err = s.Register(name, pwd)
	if err != nil {
		Fail(c, "Registration failed Account has been registered", data)
		return
	}
	user, _ := s.VerifyUser(name, pwd)
	data["user_id"] = user.Id
	data["token"], _ = jwt.GenToken(user.Id)
	Ok(c, "registered successfully", data)
}

func Login(c *gin.Context) {
	data := gin.H{
		"user_id": "",
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

func UserInfo(c *gin.Context) {
	data := gin.H{"user": nil}
	id, _ := c.Get("id")
	other := c.Query("user_id")

	user, err := s.GetUserInfo(id.(string), other)
	data["user"] = user
	if err != nil {
		Fail(c, err.Error(), data)
	}
	Ok(c, "user info", data)
}
