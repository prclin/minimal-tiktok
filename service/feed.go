package service

import (
	"github.com/gin-gonic/gin"
	"github.com/prclin/minimal-tiktok/dao"
	"github.com/prclin/minimal-tiktok/model/entity"
)

func GetVideo(c *gin.Context) []entity.Video {
	videos := dao.GetVideos(c)
	return videos
}
