package mysql

import "tiktok/model"

func GetUserByName(name string) (*model.User, error) {
	var user *model.User
	err := DB.Select("username").Where("username = ?", name).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByNamePwd(name, pwd string) (*model.User, error) {
	var user *model.User
	err := DB.Where("username = ? and password = ?", name, pwd).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func InsertUser(name string, pwd string) {
	DB.Create(&model.User{
		Username: name,
		Password: pwd,
	})
}

// 关注数
func GetFollowCount(id int64) (n int64) {
	DB.Table("followers").Select("count(*)").Where("`from` = ?", id).Find(&n)
	return
}

// 粉丝数
func GetFollowerCount(id int64) (n int64) {
	DB.Table("followers").Select("count(*)").Where("`to` = ?", id).Find(&n)
	return
}

// 自己是否关注他
func IsFollower(yourId int64, otherId int64) bool {
	n := 0
	DB.Table("followers").Select("count(*)").Where("`from` = ? and `to` = ?", yourId, otherId).Find(&n)
	return n == 1
}

func IsExistById(id int64) string {
	name := ""
	DB.Table("users").Select("username").Where("id = ?", id).Find(&name)
	return name
}
