package entity

import "time"

type Message struct {
	FromUserId uint64
	ToUserId   uint64
	Content    string
	CreateTime time.Time
}
