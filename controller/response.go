package controller

type Response struct {
	StatusCode int32  `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` // 返回状态描述
}

type User struct {
	Id              int64  `json:"id,omitempty"`             // 用户id
	Name            string `json:"name,omitempty"`           // 用户名称
	FollowCount     int64  `json:"follow_count,omitempty"`   // 关注总数
	FollowerCount   int64  `json:"follower_count,omitempty"` // 粉丝总数
	Avatar          string `json:"avatar"`                   //用户头像
	BackgroundImage string `json:"background_image"`         //用户个人页顶部大图
	Signature       string `json:"signature"`                //个人简介
	TotalFavorited  string `json:"total_favorited"`          //获赞数量
	WorkCount       int64  `json:"work_count"`               //作品数
	FavoriteCount   int64  `json:"favorite_count"`           //喜欢数
	IsFollow        bool   `json:"is_follow,omitempty"`      // true-已关注，false-未关注
}

type Video struct {
	Id            int64       `json:"id"`                       // 视频唯一标识
	Title         string      `json:"title"`                    // 视频标题
	Author        interface{} `json:"author"`                   // 视频作者信息
	PlayUrl       string      `json:"play_url,omitempty"`       // 视频播放地址
	CoverUrl      string      `json:"cover_url,omitempty"`      // 视频封面地址
	FavoriteCount int64       `json:"favorite_count,omitempty"` // 视频的点赞总数
	CommentCount  int64       `json:"comment_count,omitempty"`  // 视频的评论总数
	IsFavorite    bool        `json:"is_favorite,omitempty"`    // true-已点赞，false-未点赞
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`          // 视频评论id
	User       User   `json:"user"`                  // 评论用户信息
	Content    string `json:"content,omitempty"`     // 评论内容
	CreateDate string `json:"create_date,omitempty"` // 评论发布日期，格式 mm-dd
}

type Message struct {
	Id         int64  `json:"id"`                               // 用户id
	FromUserId int64  `json:"from_user_id" validate:"required"` // 消息发送者
	ToUserId   int64  `json:"to_user_id" validate:"required"`   //消息接收者
	Content    string `json:"content" validate:"required"`      // 消息内容
	CreateTime int64  `json:"create_time" validate:"required"`
}
