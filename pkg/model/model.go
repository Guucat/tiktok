package model

import "time"

type Model struct {
	State      int8      `json:"state" gorm:"default:(-)"`
	CreateTime time.Time `json:"create_time" gorm:"default:(-)"`
	UpdateTime time.Time `json:"update_time" gorm:"default:(-)"`
}
