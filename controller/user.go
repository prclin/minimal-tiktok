package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prclin/minimal-tiktok/core"
	"github.com/prclin/minimal-tiktok/model/response"
	"github.com/prclin/minimal-tiktok/service"
)

func init() {
	core.Router.POST("/any", func(context *gin.Context) {
		context.JSON(200, response.Response{StatusCode: 200, StatusMsg: "ok"})
	})
	//core.Router.NoRoute(func(context *gin.Context) {
	//	context.JSON(200, response.Response{StatusCode: 200, StatusMsg: "ok"})
	//})

	core.ContextRouter.POST("/user/register/", service.RegisterHandler)

}
