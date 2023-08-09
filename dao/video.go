package dao

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/entity"
	"time"
)

func GetVideos(c *gin.Context) []entity.Video {
	//获取最新时间
	t := time.Now()
	layout := "2006-01-02 15:04:05 -0700"
	formattedTime := t.Format(layout)
	query := c.DefaultQuery("latest_time", formattedTime)
	parse, _ := time.Parse(layout, query)
	feedMax := 30 //最大视频数为30
	db := global.Datasource
	var videos []entity.Video
	tx := db.Table("video").Model(&entity.Video{}).Select("*").Joins("left join user on video.author_id = user.id").Where("video.created_at < ?", parse).Limit(feedMax).Scan(&videos).Debug()
	if tx != nil {
		fmt.Println(tx)
	}
	return videos
}
