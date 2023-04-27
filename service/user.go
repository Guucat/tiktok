package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	. "tiktok/dao/mysql"
	"tiktok/pkg/model"
)

func Register(u *model.User) error {
	if !SelectUserByName(u.Username) {
		return InsertUser(u)
	}
	return errors.New("user already exists")
}

func VerifyUser(name, pwd string) (*model.User, error) {
	return GetUserByNamePwd(name, pwd)
}

func GetFollowInfo(id int64, other int64, h gin.H) {
	h["follow_count"] = GetFollowCount(other)
	h["follower_count"] = GetFollowerCount(other)
	if id != other && id != 0 {
		h["is_follow"] = IsFollower(id, other)
	} else {
		h["is_follow"] = true
	}
}

func ExistUserWithId(id int64) string {
	return IsExistById(id)
}

func GetUserInfo(id, otherId string) (gin.H, error) {
	user := NewUserInfo()
	idn, _ := strconv.ParseInt(id, 10, 64)
	otherIdn, _ := strconv.ParseInt(otherId, 10, 64)
	name := ExistUserWithId(otherIdn)
	if name == "" {
		return user, errors.New("user doesn't exist")
	}

	user["id"] = otherIdn
	user["name"] = name
	GetFollowInfo(idn, otherIdn, user)
	GetWorkInfo(idn, user)
	return user, nil
}

// 获赞总数，作品总数，喜欢总数
func GetWorkInfo(id int64, h gin.H) {
	h["total_favorited"], h["favorite_count"], h["work_count"] = GetTotalWorkCount(id)
}

func NewUserInfo() gin.H {
	return gin.H{
		"id":             0,
		"name":           "",
		"follow_count":   0,
		"follower_count": 0,
		"is_follow":      false,

		"avatar":           "",
		"background_image": "",
		"signature":        "hello",
		"total_favorited":  "",
		"work_count":       0,
		"favorite_count":   0,
	}
}
