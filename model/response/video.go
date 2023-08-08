package response

import "github.com/prclin/minimal-tiktok/model/entity"

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

/*
Video 响应实体
*/
type Video struct {
	entity.Video
	Author     User `json:"author"`      // 视频作者信息
	IsFavorite bool `json:"is_favorite"` // true-已点赞，false-未点赞

}
