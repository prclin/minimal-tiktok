package service

import (
	"errors"
	"github.com/prclin/minimal-tiktok/dao"
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/entity"
	"github.com/prclin/minimal-tiktok/model/response"
	"gorm.io/gorm"
)

func PublishComment(comment entity.Comment) response.CommentResponse {
	//插入评论
	tx := global.Datasource.Begin() //开启事务
	defer tx.Commit()               //提交事务
	cId, err := dao.InsertComment(tx, comment)
	if err != nil {
		tx.Rollback()
		global.Logger.Debug(err.Error())
		return response.CommentResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
	}

	//更新视频的点赞数
	err = dao.UpdateCommentCountBy(tx, comment.VideoId, 1)
	if err != nil {
		tx.Rollback()
		global.Logger.Debug(err.Error())
		return response.CommentResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
	}

	//获取评论
	cmt, err := dao.SelectCommentBy(tx, cId) //一定要处于同一个事务
	if err != nil {
		tx.Rollback()
		global.Logger.Debug(err.Error())
		return response.CommentResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
	}

	//获取用户信息
	user, err := dao.SelectUserById(comment.UserId)
	if err != nil {
		tx.Rollback()
		global.Logger.Debug(err.Error())
		return response.CommentResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
	}

	return response.CommentResponse{
		Response: response.Response{StatusCode: 0, StatusMsg: "评论成功"},
		Comment:  &response.CommentInfo{Comment: cmt, User: response.UserInfo{User: user, IsFollow: true}},
	}

}

func DeleteComment(userId, commentId uint64) response.CommentResponse {
	//获取评论用户id
	cUserId, err := dao.SelectCommentUserIdBy(commentId)
	if err != nil || userId != cUserId {
		if err != nil {
			global.Logger.Debug(err.Error())
		} else {
			global.Logger.Debugf("操作用户：%v,评论用户：%v", userId, cUserId)

		}
		return response.CommentResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
	}

	//开启事务
	tx := global.Datasource.Begin()
	defer tx.Commit()

	//删除评论
	err = dao.DeleteCommentBy(tx, commentId)
	if err != nil {
		tx.Rollback()
		global.Logger.Debug(err.Error())
		return response.CommentResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
	}
	//更新点赞数
	err = dao.UpdateCommentCountBy(tx, commentId, -1)
	if err != nil {
		tx.Rollback()
		global.Logger.Debug(err.Error())
		return response.CommentResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
	}

	return response.CommentResponse{Response: response.Response{StatusCode: 0, StatusMsg: "删除成功"}}
}

func GetCommentList(userId uint64, videoId uint64) response.CommentListResponse {
	//获取评论列表
	comments, err := dao.SelectCommentByVideoId(videoId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.Logger.Debug(err.Error())
		return response.CommentListResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
	}

	//映射
	commentInfos := make([]response.CommentInfo, 0, len(comments))
	for _, comment := range comments {
		//用户信息
		userInfoRes := GetUserInfo(userId, comment.UserId)
		if userInfoRes.StatusCode != 0 {
			global.Logger.Debug("在获取用户信息时发生错误!")
			return response.CommentListResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
		}
		//添加
		commentInfos = append(commentInfos, response.CommentInfo{Comment: comment, User: *userInfoRes.User})
	}

	return response.CommentListResponse{
		Response:    response.Response{StatusCode: 0, StatusMsg: "获取成功"},
		CommentList: commentInfos,
	}
}
