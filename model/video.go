package model

type Video struct {
	Id            int64  `json:"id"`
	AuthorId      int64  `json:"author_id"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	Title         string `json:"title"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	Model
}
