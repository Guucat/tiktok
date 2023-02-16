package service

import (
	"tiktok/dao/mysql"
	"tiktok/model"
)

func CommentAction(videoId, commentId, userId, commentText interface{}) error {
	return mysql.AddComment(videoId, commentId, userId, commentText)
}

func DelCommentAction(commentId, videoId interface{}) error {
	return mysql.DeleteComment(commentId, videoId)
}

func CommentList(videoId interface{}) (commentMessage []model.Comment, err error) {
	return mysql.GetCommentListByVideoId(videoId)
}
