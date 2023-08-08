package service

import (
	"github.com/prclin/minimal-tiktok/dao"
	"github.com/prclin/minimal-tiktok/model/response"
)

/*
GetPublishList 通过userId获取用户的视频发布列表
*/
func GetPublishList(userId uint64) []response.Video {
	//获取用户信息
	author := GetUserInfo(userId)
	//获取用户视频列表
	videos := dao.SelectVideosByUserId(userId)

	//映射
	rVideos := make([]response.Video, len(videos)) //避免切片扩容
	for _, video := range videos {
		rVideos = append(rVideos, response.Video{
			Video:      video,
			Author:     author,
			IsFavorite: dao.IsFavorite(userId, video.Id), //是否喜欢
		})
	}
	return rVideos
}
