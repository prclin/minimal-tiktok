package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prclin/minimal-tiktok/core"
	"github.com/prclin/minimal-tiktok/model/response"
	"github.com/prclin/minimal-tiktok/service"
)

func init() {
	//创建路由组publish
	publish := core.ContextRouter.Group("/publish")
	publish.POST("/action/", PublishAction)
}

func PublishAction(context *gin.Context) {
	//token验证
	token := context.PostForm("token")

	log.Print(token)
	//参数解析并校验
	title := context.PostForm("title")
	if title == "" {
		context.JSON(200, response.Response{StatusCode: 1, StatusMsg: "请检查视频标题"})
		return
	}

	data, err := context.FormFile("data")
	if err != nil {
		context.JSON(200, response.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	//核心逻辑，投稿
	pubSuccess := service.PublishAction(title, data)

	//返回
	if pubSuccess {
		context.JSON(200, response.Response{
			StatusCode: 0,
			StatusMsg:  "投稿成功",
		})
	} else {
		context.JSON(200, response.Response{
			StatusCode: 1,
			StatusMsg:  "投稿失败",
		})
	}
}
