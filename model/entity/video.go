package entity

import "time"

// Video video类
type Video struct {
	ID            uint64    `json:"id" gorm:"column:id;primary_key;NOT NULL"`
	Author        User      `gorm:"column:author_id;embedded;type:bigint;index;primary_key;NOT NULL"`
	PlayUrl       string    `json:"play_url" gorm:"column:play_url;NOT NULL"`
	CoverUrl      string    `json:"cover_url" gorm:"column:cover_url;NOT NULL"`
	FavoriteCount int64     `json:"favorite_count" gorm:"column:favorite_count;NOT NULL"`
	CommentCount  int64     `json:"comment_count" gorm:"column:comment_count;NOT NULL"`
	IsFavorite    bool      `json:"is_favorite" gorm:"column:is_favorite;NOT NULL"`
	Title         string    `json:"title" gorm:"column:title;NOT NULL"`
	CreatedAt     time.Time `json:"-" gorm:"column:created_at;NOT NULL"`
}
