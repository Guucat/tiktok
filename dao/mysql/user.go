package mysql

import (
	"gorm.io/gorm"
	"log"
	"tiktok/model"
)

// true, if user exist.
func SelectUserByName(name string) bool {
	n := 0
	DB.Table("users").Select("count(*)").Where("username = ?", name).Find(&n)
	return n >= 1
}

func GetUserByNamePwd(name, pwd string) (*model.User, error) {
	var user *model.User
	err := DB.Where("username = ? and password = ?", name, pwd).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func InsertUser(u *model.User) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(u).Error; err != nil {
			log.Println("fail to insert user", err)
			return err
		}
		return nil
	})
	return err
}

// 关注数
func GetFollowCount(id int64) (n int64) {
	DB.Table("followers").Select("count(*)").Where("`from_user_id` = ?", id).Find(&n)
	return
}

// 粉丝数
func GetFollowerCount(id int64) (n int64) {
	DB.Table("followers").Select("count(*)").Where("`to_user_id` = ?", id).Find(&n)
	return
}

// 自己是否关注他
func IsFollower(yourId int64, otherId int64) bool {
	n := 0
	DB.Table("followers").Select("count(*)").Where("`from_user_id` = ? and `to_user_id` = ?", yourId, otherId).Find(&n)
	return n == 1
}

func IsExistById(id int64) string {
	name := ""
	DB.Table("users").Select("username").Where("id = ?", id).Find(&name)
	return name
}
