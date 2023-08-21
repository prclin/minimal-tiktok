package service

import (
	"fmt"
	"github.com/prclin/minimal-tiktok/dao"
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/response"
	"github.com/prclin/minimal-tiktok/util"
	"mime/multipart"
	"time"
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

// PostVideoToOSS 上传视频到OSS
func PostVideoToOSS(video *multipart.FileHeader) (string, error) {
	//获取bucket
	bucket, err := global.OSSClient.Bucket("lattice-storage")
	if err != nil {
		return "", err
	}
	//上传文件
	file, err := video.Open() //打开文件
	defer file.Close()
	if err != nil {
		return "", err
	}
	hash, err := util.FileHash(&file) //文件hash
	file.Seek(0, 0)                   //复位
	if err != nil {
		return "", err
	}
	fmt.Println(video.Filename)
	relativePath := fmt.Sprintf("video/%v-%v.mp4", hash, time.Now().UnixMilli())
	err = bucket.PutObject(relativePath, file) //上传
	if err != nil {
		return "", err
	}
	return "https://lattice-storage." + global.Configuration.OSS.EndPoint + "/" + relativePath, nil
}
