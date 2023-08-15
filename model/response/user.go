package response

import "github.com/prclin/minimal-tiktok/model/entity"

// UserEnrollResponse 用户登录注册响应
type UserEnrollResponse struct {
	Response
	UserId uint64
	Token  string
}

// User 用户信息响应
type User struct {
	entity.User
	IsFollow bool `json:"is_follow"` // true-已关注，false-未关注
}
