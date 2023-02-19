package model

type Message struct {
	Id         int64  `json:"id"`                               // 用户id
	FromUserId int64  `json:"from_user_id" validate:"required"` // 消息发送者
	ToUserId   int64  `json:"to_user_id" validate:"required"`   //消息接收者
	Content    string `json:"content" validate:"required"`      // 消息内容
	Model
}
