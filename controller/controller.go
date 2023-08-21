package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prclin/minimal-tiktok/core"
	"github.com/prclin/minimal-tiktok/model/response"
)

/*
Init 为了显示执行副作用引入
*/
func Init() {
	//just empty
}

func init() {
	core.ContextRouter.GET("/health/", func(context *gin.Context) {
		context.JSON(200, response.Response{StatusCode: 200, StatusMsg: "ok"})
	})
}
