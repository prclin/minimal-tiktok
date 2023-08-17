package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prclin/minimal-tiktok/core"
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/response"
	"github.com/prclin/minimal-tiktok/service"
	"github.com/prclin/minimal-tiktok/util"
	"net/http"
)

func init() {
	relation := core.ContextRouter.Group("/relation")
	relation.POST("/action", PostFollowAction)
}

/*
PostFollowAction 关注操作

参数 token toUserId actionType
*/
func PostFollowAction(c *gin.Context) {
	//参数校验
	var query struct {
		Token      string `form:"token" binding:"required"`
		ToUserId   uint64 `form:"to_user_id" binding:"required"`
		ActionType uint   `form:"action_type" binding:"required,min=1,max=2"` //1关注 2取关
	}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		global.Logger.Error(err.Error())
		c.JSON(http.StatusOK, response.Response{StatusCode: 1, StatusMsg: "参数错误"})
		return
	}

	claims, err := util.ParseToken(query.Token) //token解析
	if err != nil {
		global.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, response.Response{StatusCode: 1, StatusMsg: "参数错误"})
		return
	}
	followerId := claims.Id

	result := service.FollowAction(followerId, query.ToUserId, query.ActionType)
	c.JSON(http.StatusOK, result)
}
