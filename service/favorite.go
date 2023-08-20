package service

import (
	"errors"
	"github.com/prclin/minimal-tiktok/dao"
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/entity"
	"github.com/prclin/minimal-tiktok/model/response"
	"gorm.io/gorm"
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

func GetFavoriteList(acquirerId, userId uint64) response.VideoListResponse {
	//获取喜欢的视频id
	videoIds, err := dao.SelectFavorites(userId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.Logger.Debug(err.Error())
		return response.VideoListResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
	}

	//获取视频信息
	videos, err := dao.SelectVideosBy(videoIds)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.Logger.Debug(err.Error())
		return response.VideoListResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
	}

	//映射
	videoInfos := make([]response.VideoInfo, 0, len(videos))
	for _, video := range videos {
		videoInfos = append(videoInfos, response.VideoInfo{
			Video:      video,
			Author:     GetUserInfo(acquirerId, video.UserId).User,
			IsFavorite: dao.IsFavorite(acquirerId, video.Id),
		})
	}
	return response.VideoListResponse{
		Response:  response.Response{StatusCode: 0, StatusMsg: "获取成功"},
		VideoList: videoInfos,
	}
}
