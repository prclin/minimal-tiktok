package response

import "github.com/prclin/minimal-tiktok/model/entity"

// FeedResponse 视频流返回
type MessageResponse struct {
	StatusCode int32          `json:"status_code"`
	StatusMsg  string         `json:"status_msg,omitempty"`
	User       entity.UserDto `json:"user"`
}
