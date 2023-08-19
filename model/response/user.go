package response

import "github.com/prclin/minimal-tiktok/model/entity"

// UserEnrollResponse 用户登录注册响应
type UserEnrollResponse struct {
	Response
	UserId uint64 `json:"user_id"`
	Token  string `json:"token"`
}

// UserInfoResponse 用户信息响应
type UserInfoResponse struct {
	Response
	User *UserInfo `json:"user"`
}

type UserInfo struct {
	entity.User
	IsFollow bool `json:"is_follow"` // true-已关注，false-未关注
}

type RelationResponse struct {
	Response
	UserList []UserInfo `json:"user_list"`
}
