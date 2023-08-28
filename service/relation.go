package service

import (
	"errors"
	"github.com/prclin/minimal-tiktok/dao"
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/response"
	"github.com/prclin/minimal-tiktok/util"
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

func UserFollowList(userId uint64) response.RelationResponse {
	//获取用户所有被关注者id
	followees, err := dao.SelectFolloweeBy(userId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.Logger.Debug(err.Error())
		return response.RelationResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
	}
	//获取所有被关注者信息
	users, err := dao.SelectUserByIds(followees)
	if err != nil {
		global.Logger.Debug(err.Error())
		return response.RelationResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
	}

	//映射
	userInfos := make([]response.UserInfo, len(users))
	for i := 0; i < len(users); i++ {
		userInfos[i] = response.UserInfo{User: users[i], IsFollow: true}
	}

	return response.RelationResponse{Response: response.Response{StatusCode: 0, StatusMsg: "ok"}, UserList: userInfos}
}

func UserFollowerList(userId uint64) response.RelationResponse {
	//获取用户所有粉丝id
	followers, err := dao.SelectFollowerBy(userId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.Logger.Debug(err.Error())
		return response.RelationResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
	}
	//获取所有粉丝信息
	users, err := dao.SelectUserByIds(followers)
	if err != nil {
		global.Logger.Debug(err.Error())
		return response.RelationResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
	}

	//映射
	userInfos := make([]response.UserInfo, len(users))
	for i := 0; i < len(users); i++ {
		userInfos[i] = response.UserInfo{User: users[i], IsFollow: dao.IsFollow(userId, users[i].Id)}
	}

	return response.RelationResponse{Response: response.Response{StatusCode: 0, StatusMsg: "ok"}, UserList: userInfos}
}

func UserFriendList(userId uint64) response.FriendResponse {
	//获取粉丝
	followers, err := dao.SelectFollowerBy(userId)
	//获取关注
	followees, err := dao.SelectFolloweeBy(userId)
	if err != nil {
		global.Logger.Debug(err.Error())
		return response.FriendResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
	}
	//获取朋友 互相关注为朋友
	friendIds := util.Intersection(followers, followees)
	//获取朋友信息
	friends, err := dao.SelectUserByIds(friendIds)
	if err != nil {
		global.Logger.Debug(err.Error())
		return response.FriendResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
	}

	//映射
	userInfos := make([]response.FriendInfo, len(friends))
	for i := 0; i < len(friends); i++ {
		send, _ := dao.SelectLastMessageBy(userId, friends[i].Id)
		receive, _ := dao.SelectLastMessageBy(friends[i].Id, userId)
		if send.CreateTime.After(receive.CreateTime) {

		}
		userInfos[i] = response.FriendInfo{
			User:     friends[i],
			IsFollow: true,
			Message:  util.Ternary(send.CreateTime.After(receive.CreateTime), send.Content, receive.Content),
			MsgType:  int64(util.Ternary(send.CreateTime.After(receive.CreateTime), 1, 0)),
		}
	}

	return response.FriendResponse{Response: response.Response{StatusCode: 0, StatusMsg: "ok"}, UserList: userInfos}
}
