package model

import comment_proto "tiktok/server/comment/api"

type Comment struct {
	Id         int64  `json:"id,omitempty"`                                    // 评论id
	UserId     int64  `json:"user_id,omitempty"`                               // 评论用户id
	VideoId    int64  `json:"video_id,omitempty"`                              //评论视频id
	Content    string `json:"content,omitempty"`                               // 评论内容
	CreateDate string `json:"create_date,omitempty" gorm:"column:create_time"` // 评论发布日期，格式 mm-dd
}

type CommentAction struct {
	VideoId    string
	MeId       string
	ActionType string
	Content    string
	ContentId  string
	TimeDate   string
}

type CommentResponse struct {
	Id         int64                   `json:"id,omitempty"`          // 视频评论id
	User       *comment_proto.UserInfo `json:"user"`                  // 评论用户信息
	Content    string                  `json:"content,omitempty"`     // 评论内容
	CreateDate string                  `json:"create_date,omitempty"` // 评论发布日期，格式 mm-dd
}
