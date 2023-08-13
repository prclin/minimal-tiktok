package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prclin/minimal-tiktok/core"
	"github.com/prclin/minimal-tiktok/model/response"
	"github.com/prclin/minimal-tiktok/service"
	"net/http"
	"time"
)

func init() {
	core.Router.GET("/douyin/feed", feed)
}

func feed(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusOK, &response.Response{StatusCode: 0, StatusMsg: "用户未登录，请登录！"})
	}
	videos := service.GetVideo(c)
	var at int64
	if videos[0].CreatedAt.IsZero() {
		at = time.Now().UnixMilli()
	} else {
		at = videos[0].CreatedAt.UnixMilli()
	}
	c.JSON(http.StatusOK, &response.FeedResponse{StatusCode: 0, StatusMsg: "videoList", NextTime: at, VideoList: videos})
}
