package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/prclin/minimal-tiktok/core"
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/entity"
	"github.com/prclin/minimal-tiktok/model/response"
	"github.com/prclin/minimal-tiktok/service"
	"github.com/prclin/minimal-tiktok/util"
	"mime/multipart"
	"net/http"
	"strconv"
)

func init() {
	publish := core.ContextRouter.Group("/publish")
	publish.GET("/list/", GetPublishList)
	publish.POST("/action/", PostVideo)
}

/*
GetPublishList 获取指定用户的视频发布列表

参数 token、user_id

只有已登录的用户才可以获取其他人的视频发布列表
*/
func GetPublishList(c *gin.Context) {
	//参数校验
	var query struct {
		UserId string `form:"user_id" binding:"required"`
		Token  string `form:"token" binding:"required"`
	}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		global.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, response.VideoListResponse{
			Response:  response.Response{StatusCode: 1, StatusMsg: "参数错误"},
			VideoList: nil,
		})
		return
	}
	claims, err := util.ParseToken(query.Token) //token解析
	if err != nil {
		global.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, response.VideoListResponse{
			Response:  response.Response{StatusCode: 1, StatusMsg: "参数错误"},
			VideoList: nil,
		})
		return
	}
	erId := claims.Id                                    //获取查询者id
	eeId, err := strconv.ParseUint(query.UserId, 10, 64) //获取被查询者id
	if err != nil {
		global.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, response.VideoListResponse{
			Response:  response.Response{StatusCode: 1, StatusMsg: "参数错误"},
			VideoList: nil,
		})
		return
	}
	//获取发布列表
	resp := service.GetPublishList(erId, eeId)
	//返回发布列表
	c.JSON(200, response.VideoListResponse{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		VideoList: resp,
	})
}

/*
PostVideo 发布视频

参数 data token title
*/
func PostVideo(c *gin.Context) {
	//参数绑定
	var body struct {
		Data  *multipart.FileHeader `form:"data" binding:"required"`
		Token string                `form:"token" binding:"required"`
		Title string                `form:"title" binding:"required"`
	}
	err := c.ShouldBindWith(&body, binding.FormMultipart) //绑定multipart-form
	if err != nil {
		global.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, response.Response{StatusCode: 1, StatusMsg: "参数错误"})
		return
	}
	//token解析
	claims, err := util.ParseToken(body.Token)
	if err != nil {
		global.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, response.Response{StatusCode: 1, StatusMsg: "登录信息错误"})
		return
	}
	//上传视频
	filePath, err := service.PostVideoToOSS(body.Data)
	if err != nil {
		global.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, response.Response{StatusCode: 2, StatusMsg: "文件上传错误"})
		return
	}

	video := entity.Video{
		UserId:   claims.Id,
		Title:    body.Title,
		PlayURL:  filePath,
		CoverURL: "",
	}

	//保存投稿
	res := service.SaveVideo(video)
	c.JSON(http.StatusOK, res)
}
