package service

import "tiktok/dao/mysql"

func CommentAction(videoId, commentId, userId, commentText interface{}) error {
	return mysql.AddComment(videoId, commentId, userId, commentText)
}

func DelCommentAction(commentId, videoId interface{}) error {
	return mysql.DeleteComment(commentId, videoId)
}
