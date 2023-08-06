package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/entity"
	"github.com/prclin/minimal-tiktok/model/response"
	"net/http"
	"time"
)

func feed(c *gin.Context) {
	t := time.Now()
	layout := "2006-01-02 15:04:05 -0700"
	formattedTime := t.Format(layout)
	query := c.DefaultQuery("latest_time", formattedTime)
	parse, _ := time.Parse(layout, query)
	var at int64
	feedMax := 30 //最大视频数为30
	db := global.Datasource
	var videos []entity.Video
	tx := db.Table("video").Model(&entity.Video{}).Select("*").Joins("left join user on video.author_id = user.id").Where("video.created_at < ?", parse).Limit(feedMax).Scan(&videos).Debug()
	if tx != nil {
		fmt.Println(tx)
	}
	if videos[0].CreatedAt.IsZero() {
		at = time.Now().UnixMilli()
	} else {
		at = videos[0].CreatedAt.UnixMilli()
	}
	c.JSON(http.StatusOK, &response.FeedResponse{StatusCode: 0, StatusMsg: "videoList", NextTime: at, VideoList: videos})
}
