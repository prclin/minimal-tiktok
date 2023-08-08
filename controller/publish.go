package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prclin/minimal-tiktok/core"
	"github.com/prclin/minimal-tiktok/model/response"
	"github.com/prclin/minimal-tiktok/service"
	"strconv"
)

func init() {
	publish := core.ContextRouter.Group("/publish")
	publish.GET("/list", GetPublishList)
}

/*
GetPublishList 获取指定用户的视频发布列表

参数 token、user_id

只有已登录的用户才可以获取其他人的视频发布列表
*/
func GetPublishList(context *gin.Context) {
	//是否提供token
	token, ok := context.GetQuery("token")
	if !ok {
		context.JSON(200, &response.Response{StatusCode: 1, StatusMsg: "用户未登录"})
		return
	}
	fmt.Println(token)
	//是否提供user_id
	userIdStr, ok := context.GetQuery("user_id")
	if !ok {
		context.JSON(200, &response.Response{StatusCode: 1, StatusMsg: "参数不完整"})
		return
	}
	//user_id是否正确
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		context.JSON(200, &response.Response{StatusCode: 1, StatusMsg: "参数错误"})
		return
	}
	//获取发布列表
	resp := service.GetPublishList(userId)
	//返回发布列表
	context.JSON(200, response.VideoListResponse{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		VideoList: resp,
	})
}
