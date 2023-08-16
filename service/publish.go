package service

import (
	"github.com/prclin/minimal-tiktok/dao"
	"github.com/prclin/minimal-tiktok/model/response"
)

/*
GetPublishList 通过userId获取用户的视频发布列表
*/
func GetPublishList(erId, eeId uint64) []response.VideoInfo {
	//获取用户信息
	userInfo := GetUserInfo(erId, eeId)
	//获取用户视频列表
	videos := dao.SelectVideosByUserId(eeId)
	//映射
	rVideos := make([]response.VideoInfo, len(videos)) //避免切片扩容
	for _, video := range videos {
		rVideos = append(rVideos, response.VideoInfo{
			Video:      video,
			Author:     userInfo.User,
			IsFavorite: dao.IsFavorite(erId, video.Id), //是否喜欢
		})
	}
	return rVideos
}
