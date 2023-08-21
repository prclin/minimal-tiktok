package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prclin/minimal-tiktok/core"
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/response"
	"github.com/prclin/minimal-tiktok/service"
	"github.com/prclin/minimal-tiktok/util"
	"net/http"
	"strings"
	"time"
)

func init() {
	core.ContextRouter.GET("/feed/", GetFeedList)
}

/*
GetFeedList 视频feed流

参数 latestTime(可选) token(可选)
*/
func GetFeedList(c *gin.Context) {
	//参数校验
	var query struct {
		Token      string `form:"token"`
		LatestTime int64  `form:"latest_time"`
	}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		global.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, response.FeedResponse{Response: response.Response{StatusCode: 1, StatusMsg: "参数错误"}})
		return
	}

	//参数初始化
	var userId uint64
	if strings.TrimSpace(query.Token) != "" {
		claims, err1 := util.ParseToken(query.Token)
		if err1 != nil {
			global.Logger.Debug(err.Error())
			c.JSON(http.StatusOK, response.FeedResponse{Response: response.Response{StatusCode: 1, StatusMsg: "登录信息错误"}})
			return
		}
		userId = claims.Id
	}
	latestTime := util.Ternary(query.LatestTime == 0, time.Now(), time.Unix(query.LatestTime, 0))
	res := service.GetFeedList(userId, latestTime)
	c.JSON(http.StatusOK, res)
}
