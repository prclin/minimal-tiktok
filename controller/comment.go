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
	"strings"
)

func init() {
	comment := core.ContextRouter.Group("/comment")
	comment.POST("/action/", PostCommentAction)
	comment.GET("/list/", GetCommentList)
}

/*
PostCommentAction 评论操作

参数 token videoId actionType commentText commentId

commentText参数在actionType为1时指定，commentId在actionType为2时指定
*/
func PostCommentAction(c *gin.Context) {
	//参数校验
	var query struct {
		Token       string `form:"token" binding:"required"`
		VideoId     uint64 `form:"video_id" binding:"required,min=1"`
		ActionType  uint8  `form:"action_type" binding:"required,min=1,max=2"`
		CommentText string `form:"comment_text"`
		CommentId   uint64 `form:"comment_id"`
	}
	err := c.ShouldBindQuery(&query)
	if err != nil { //绑定错误
		global.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, response.CommentResponse{Response: response.Response{StatusCode: 1, StatusMsg: "参数错误"}})
		return
	}
	//token校验
	claims, err := util.ParseToken(query.Token)
	if err != nil {
		global.Logger.Debug(err != nil)
		c.JSON(http.StatusOK, response.CommentResponse{Response: response.Response{StatusCode: 1, StatusMsg: "参数错误"}})
		return
	}

	var res response.CommentResponse

	//操作校验
	if query.ActionType == 1 {
		if strings.TrimSpace(query.CommentText) == "" {
			global.Logger.Debug("发布评论时，评论内容不能为空!")
			c.JSON(http.StatusOK, response.CommentResponse{Response: response.Response{StatusCode: 1, StatusMsg: "参数错误"}})
			return
		}
		comment := entity.Comment{UserId: claims.Id, VideoId: query.VideoId, Content: query.CommentText}
		res = service.PublishComment(comment) //发布评论
	} else {
		if query.CommentId == 0 {
			global.Logger.Debug("删除评论时，评论id不能为0和空!")
			c.JSON(http.StatusOK, response.CommentResponse{Response: response.Response{StatusCode: 1, StatusMsg: "参数错误"}})
			return
		}
		res = service.DeleteComment(claims.Id, query.CommentId) //删除评论
	}

	c.JSON(http.StatusOK, res)
}

/*
GetCommentList 评论列表

参数 token videoId
*/
func GetCommentList(c *gin.Context) {
	var query struct {
		Token   string `form:"token" binding:"required"`
		VideoId uint64 `form:"video_id" binding:"required,min=1"`
	}
	err := c.ShouldBindQuery(&query)
	if err != nil { //绑定错误
		global.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, response.CommentResponse{Response: response.Response{StatusCode: 1, StatusMsg: "参数错误"}})
		return
	}

	claims, err := util.ParseToken(query.Token)
	if err != nil {
		global.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, response.CommentResponse{Response: response.Response{StatusCode: 1, StatusMsg: "登录信息错误"}})
		return
	}

	res := service.GetCommentList(claims.Id, query.VideoId)
	c.JSON(http.StatusOK, res)
}
