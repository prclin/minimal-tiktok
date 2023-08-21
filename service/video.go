package service

import (
	"errors"
	"github.com/prclin/minimal-tiktok/dao"
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/entity"
	"github.com/prclin/minimal-tiktok/model/response"
	"github.com/prclin/minimal-tiktok/util"
	"gorm.io/gorm"
	"time"
)

func GetFeedList(userId uint64, latestTime time.Time) response.FeedResponse {
	//获取视频
	videos, err := dao.SelectVideosByCreateTime(latestTime, 30) //size暂时未提取到配置
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.Logger.Debug(err.Error())
		return response.FeedResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
	}

	//映射
	videoInfos := make([]response.VideoInfo, 0, len(videos))
	for _, video := range videos {
		user, err1 := dao.SelectUserById(video.UserId)
		if err1 != nil { //此处没查到记录也算错误
			global.Logger.Error(err1.Error())
			return response.FeedResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
		}
		videoInfos = append(videoInfos, response.VideoInfo{
			Video:      video,
			Author:     &response.UserInfo{User: user, IsFollow: util.Ternary(userId == 0, false, dao.IsFollow(userId, video.UserId))},
			IsFavorite: dao.IsFavorite(0, video.Id), //是否喜欢
		})
	}
	var nextTime int64
	if len(videos) != 0 {
		nextTime = videos[len(videos)-1].CreateTime.Unix()
	}
	return response.FeedResponse{
		Response:  response.Response{StatusCode: 0, StatusMsg: "获取成功"},
		NextTime:  nextTime,
		VideoList: videoInfos,
	}
}

func SaveVideo(video entity.Video) response.Response {
	err := dao.InsertVideo(global.Datasource, video)
	if err != nil {
		global.Logger.Debug(err.Error())
		return response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}
	}
	return response.Response{StatusCode: 0, StatusMsg: "投稿成功"}
}
