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
	message := core.ContextRouter.Group("/message")
	message.POST("/action", PostMessage)
	message.GET("/chat", GetChatHistory)
}

/*
PostMessage 发送消息

参数 token toUserId actionType content
*/
func PostMessage(c *gin.Context) {
	//参数校验
	var query struct {
		Token      string `form:"token" binding:"required"`
		ToUserId   uint64 `form:"to_user_id" binding:"required,min=1"`
		ActionType uint8  `form:"action_type" binding:"required,min=1,max=1"`
		Content    string `form:"content" binding:"required"`
	}
	err := c.ShouldBindQuery(&query)
	if err != nil { //绑定失败
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

	isSent := service.SendMessage(entity.Message{FromUserId: claims.Id, ToUserId: query.ToUserId, Content: query.Content})
	if !isSent {
		c.JSON(http.StatusOK, response.Response{StatusCode: 1, StatusMsg: "发送失败"})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 0, StatusMsg: "发送成功"})

}

/*
GetChatHistory 获取聊天记录

参数 token toUserId
*/
func GetChatHistory(c *gin.Context) {
	//参数校验
	var query struct {
		Token    string `form:"token" binding:"required"`
		ToUserId uint64 `form:"to_user_id" binding:"required,min=1"`
	}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		global.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, response.ChatResponse{Response: response.Response{StatusCode: 1, StatusMsg: "参数错误"}})
		return
	}
	claims, err := util.ParseToken(query.Token)
	if err != nil {
		global.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, response.ChatResponse{Response: response.Response{StatusCode: 1, StatusMsg: "参数错误"}})
		return
	}
	history := service.GetChatHistory(claims.Id, query.ToUserId)
	c.JSON(http.StatusOK, history)
}
