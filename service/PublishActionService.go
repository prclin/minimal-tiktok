package service

import (
	"mime/multipart"
	"time"

	"github.com/prclin/minimal-tiktok/dao"
	"github.com/prclin/minimal-tiktok/model/entity"
)

func PublishAction(title string, videoData *multipart.FileHeader) bool {

	//视频上传到minIO
	minIOClient := getMinIoClient()

	path, success := uploadVideo(minIOClient, videoData)

	if !success {
		return false
	}

	//视频信息存到数据库
	var video entity.Video
	video.Title = title
	video.PlayURL = path
	video.UserId = 114514
	video.CreateTime = time.Now()

	saveSuccess := dao.SaveVideoInfo(&video)

	return saveSuccess
}
