package service

import (
	"github.com/prclin/minimal-tiktok/dao"
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/response"
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
