package response

import "github.com/prclin/minimal-tiktok/model/entity"

type VideoListResponse struct {
	Response
	VideoList []VideoInfo `json:"video_list"`
}

// VideoInfo 响应实体
type VideoInfo struct {
	entity.Video
	Author     *UserInfo `json:"author"`      // 视频作者信息
	IsFavorite bool      `json:"is_favorite"` // true-已点赞，false-未点赞

}
