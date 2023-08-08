package response

import "github.com/prclin/minimal-tiktok/model/entity"

// User 用户信息响应
type User struct {
	entity.User
	IsFollow bool `json:"is_follow"` // true-已关注，false-未关注
}
