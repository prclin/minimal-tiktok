package service

import (
	"github.com/prclin/minimal-tiktok/dao"
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/entity"
	"github.com/prclin/minimal-tiktok/model/response"
)

func FavoriteAction(favorite entity.Favorite, actionType uint8) response.Response {
	//视频是否存在
	exist := dao.ExistsVideo(favorite.VideoId)
	if !exist {
		return response.Response{StatusCode: 1, StatusMsg: "视频不存在"}
	}

	//操作
	var err error
	switch actionType {
	case 1:
		if !dao.IsFavorite(favorite.UserId, favorite.VideoId) { //已点赞不重复点赞
			err = dao.InsertFavorite(global.Datasource, favorite)
		}
	case 2:
		err = dao.DeleteFavorite(global.Datasource, favorite)
	}

	if err != nil {
		global.Logger.Debug(err.Error())
		return response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}
	}
	return response.Response{StatusCode: 0, StatusMsg: "操作成功"}
}
