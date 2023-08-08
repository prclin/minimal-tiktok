package service

import (
	"github.com/prclin/minimal-tiktok/dao"
	"github.com/prclin/minimal-tiktok/model/response"
)

func GetUserInfo(id uint64) response.User {
	user := dao.SelectUserInfoById(id)
	return response.User{User: user, IsFollow: false}
}
