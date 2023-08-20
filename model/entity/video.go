package entity

import "time"

/*
Video 视频实体
*/
type Video struct {
	Id            uint64 `json:"id"` // 视频唯一标识
	UserId        uint64
	Title         string `json:"title"`          // 视频标题
	PlayURL       string `json:"play_url"`       // 视频播放地址
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount uint32 `json:"favorite_count"` // 视频的点赞总数
	CommentCount  uint32 `json:"comment_count"`  // 视频的评论总数
	Extra         string
	CreateTime    time.Time
}
