package service

import (
	"mime/multipart"
	"time"

	"github.com/prclin/minimal-tiktok/dao"
	"github.com/prclin/minimal-tiktok/model/entity"
	"github.com/prclin/minimal-tiktok/util"
)

func PublishAction(title string, videoData *multipart.FileHeader) bool {

	path, success := util.UploadToOss(videoData)
	if !success {
		return false
	}

	//视频信息存到数据库
	var video entity.Video
	video.Title = title
	video.PlayURL = path
	video.UserId = 114514
	video.CreateTime = time.Now()
	video.Extra = "这是一条视频 "

	saveSuccess := dao.SaveVideoInfo(&video)

	return saveSuccess
}
