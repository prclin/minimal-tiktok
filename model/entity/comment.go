package entity

import "time"

type Comment struct {
	Id         uint64    `json:"id"`
	UserId     uint64    `json:"-"`
	VideoId    uint64    `json:"-"`
	Content    string    `json:"content"`
	Extra      string    `json:"-"`
	CreateTime time.Time `json:"create_time"`
}
