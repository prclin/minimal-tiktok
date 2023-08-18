package service

import (
	"errors"
	"github.com/prclin/minimal-tiktok/dao"
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/response"
	"gorm.io/gorm"
)

func FollowAction(followerId, followeeId uint64, action uint) response.Response {
	var err error
	switch action {
	case 1:
		err = dao.InsertFollow(global.Datasource, followerId, followeeId)
		break
	case 2:
		err = dao.DeleteFollow(global.Datasource, followerId, followeeId)
		break
	}
	if err != nil {
		global.Logger.Debug(err)
		return response.Response{StatusCode: 2, StatusMsg: "操作失败"}
	}
	return response.Response{StatusCode: 0, StatusMsg: "操作成功"}
}

func UserFollowList(userId uint64) response.FollowListResponse {
	//获取用户所有被关注者id
	followees, err := dao.SelectFolloweeBy(userId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.Logger.Debug(err.Error())
		return response.FollowListResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
	}
	//获取所有被关注者信息
	users, err := dao.SelectUserByIds(followees)
	if err != nil {
		global.Logger.Debug(err.Error())
		return response.FollowListResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
	}

	//映射
	userInfos := make([]response.UserInfo, len(users))
	for i := 0; i < len(users); i++ {
		follow, err1 := dao.IsFollow(users[i].Id, userId)
		if err1 != nil && !errors.Is(err1, gorm.ErrRecordNotFound) {
			return response.FollowListResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
		}
		userInfos[i] = response.UserInfo{User: users[i], IsFollow: follow}
	}

	return response.FollowListResponse{Response: response.Response{StatusCode: 0, StatusMsg: "ok"}, UserList: userInfos}
}

func UserFollowerList(userId uint64) response.FollowerListResponse {
	//获取用户所有粉丝id
	followers, err := dao.SelectFollowerBy(userId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.Logger.Debug(err.Error())
		return response.FollowerListResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
	}
	//获取所有粉丝信息
	users, err := dao.SelectUserByIds(followers)
	if err != nil {
		global.Logger.Debug(err.Error())
		return response.FollowerListResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
	}

	//映射
	userInfos := make([]response.UserInfo, len(users))
	for i := 0; i < len(users); i++ {
		follow, err1 := dao.IsFollow(userId, users[i].Id)
		if err1 != nil && !errors.Is(err1, gorm.ErrRecordNotFound) {
			return response.FollowerListResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
		}
		userInfos[i] = response.UserInfo{User: users[i], IsFollow: follow}
	}

	return response.FollowerListResponse{Response: response.Response{StatusCode: 0, StatusMsg: "ok"}, UserList: userInfos}
}
