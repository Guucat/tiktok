package model

type User struct {
	Id            int64  `json:"id"`                           // 用户id
	Username      string `json:"username" validate:"required"` // 用户名称
	Password      string `json:"password" validate:"required"`
	FollowCount   int64  `json:"follow_count,omitempty"`                      // 关注总数
	FollowerCount int64  `json:"follower_count,omitempty" `                   // 粉丝总数
	FavoriteVideo string `json:"favorite_video,omitempty" gorm:"default:(-)"` //点赞的视频
	IsFollow      bool   `json:"is_follow,omitempty" gorm:"default:(-)"`      // true-已关注，false-未关注
	Model
}
