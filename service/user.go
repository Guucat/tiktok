package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	. "tiktok/dao/mysql"
	"tiktok/model"
)

func Register(name string, pwd string) error {
	_, err := GetUserByName(name)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			InsertUser(name, pwd)
			return nil
		}
		return err
	}
	return errors.New("user already exists")
}

func VerifyUser(name, pwd string) (*model.User, error) {
	return GetUserByNamePwd(name, pwd)
}

func GetFollowInfo(id int64, other int64, h gin.H) {
	h["follow_count"] = GetFollowCount(other)
	h["follower_count"] = GetFollowerCount(other)
	if id != other {
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
	return user, nil
}

func NewUserInfo() gin.H {
	return gin.H{
		"id":             0,
		"name":           "",
		"follow_count":   0,
		"follower_count": 0,
		"is_follow":      false,
	}
}
