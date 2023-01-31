package model

import "time"

type Video struct {
	Id       int64     `json:"id"`
	AuthorId int64     `json:"author_id"`
	PlayUrl  string    `json:"play_url"`
	CoverUrl string    `json:"cover_url"`
	Title    string    `json:"title"`
	Time     time.Time `json:"time"`
}
