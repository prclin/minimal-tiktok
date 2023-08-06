package response

import "github.com/prclin/minimal-tiktok/model/entity"

// FeedResponse 视频流返回
type FeedResponse struct {
	StatusCode int32          `json:"status_code"`
	StatusMsg  string         `json:"status_msg,omitempty"`
	NextTime   int64          `json:"next_time"`
	VideoList  []entity.Video `json:"video_list,omitempty"`
}
