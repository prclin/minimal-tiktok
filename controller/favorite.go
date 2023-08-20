package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prclin/minimal-tiktok/core"
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/entity"
	"github.com/prclin/minimal-tiktok/model/response"
	"github.com/prclin/minimal-tiktok/service"
	"github.com/prclin/minimal-tiktok/util"
	"net/http"
)

func init() {
	favorite := core.ContextRouter.Group("/favorite")
	favorite.POST("/action", PostFavoriteAction)
	favorite.GET("/list", GetFavoriteList)
}

/*
PostFavoriteAction 赞操作

参数 token videoId actionType
*/
func PostFavoriteAction(c *gin.Context) {
	//参数校验
	var query struct {
		Token      string `form:"token"`
		VideoId    uint64 `form:"video_id" binding:"required,min=1"`
		ActionType uint8  `form:"action_type" binding:"required,min=1,max=2"`
	}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		global.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, response.Response{StatusCode: 1, StatusMsg: "参数错误"})
		return
	}
	claims, err := util.ParseToken(query.Token)
	if err != nil {
		global.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, response.Response{StatusCode: 1, StatusMsg: "登录信息错误"})
		return
	}
	res := service.FavoriteAction(entity.Favorite{UserId: claims.Id, VideoId: query.VideoId}, query.ActionType)
	c.JSON(http.StatusOK, res)
}

/*
GetFavoriteList 获取喜欢列表

参数 userId token
*/
func GetFavoriteList(c *gin.Context) {
	//参数校验
	var query struct {
		Token  string `form:"token" binding:"required"`
		UserId uint64 `form:"user_id" binding:"required,min=1"`
	}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		global.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, response.VideoListResponse{Response: response.Response{StatusCode: 1, StatusMsg: "参数错误"}})
		return
	}
	claims, err := util.ParseToken(query.Token)
	if err != nil {
		global.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, response.VideoListResponse{Response: response.Response{StatusCode: 1, StatusMsg: "登录信息错误"}})
		return
	}
	//获取喜欢列表
	res := service.GetFavoriteList(claims.Id, query.UserId)
	c.JSON(http.StatusOK, res)
}
