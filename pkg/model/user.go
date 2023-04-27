package model

type User struct {
	Id              int64  `json:"id"`                           // 用户id
	Username        string `json:"username" validate:"required"` // 用户名称
	Password        string `json:"password" validate:"required"`
	FollowCount     int64  `json:"follow_count,omitempty"`                      // 关注总数
	FollowerCount   int64  `json:"follower_count,omitempty" `                   // 粉丝总数
	FavoriteVideo   string `json:"favorite_video,omitempty" gorm:"default:(-)"` //点赞的视频
	Avatar          string `json:"avatar"`                                      //用户头像
	BackgroundImage string `json:"background_image"`                            //用户个人页顶部大图
	Signature       string `json:"signature"`                                   //个人简介
	TotalFavorited  int64  `json:"total_favorited"`                             //获赞数量
	WorkCount       int64  `json:"work_count"`                                  //作品数
	FavoriteCount   int64  `json:"favorite_count"`                              //喜欢数
	IsFollow        bool   `json:"is_follow,omitempty" gorm:"default:(-)"`      // true-已关注，false-未关注
	Model
}

type UserBaseInfo struct {
	Id              int64  `json:"id"`                           // 用户id
	Username        string `json:"username" validate:"required"` // 用户名称
	Avatar          string `json:"avatar"`                       //用户头像
	BackgroundImage string `json:"background_image"`             //用户个人页顶部大图
	Signature       string `json:"signature"`                    // 用户签名
}
