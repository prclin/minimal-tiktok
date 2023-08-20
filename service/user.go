package service

import (
	"errors"
	"github.com/prclin/minimal-tiktok/dao"
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/entity"
	"github.com/prclin/minimal-tiktok/model/response"
	"github.com/prclin/minimal-tiktok/util"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
)

/*
RegisterUser

加密密码并补充默认用户信息，注册成功后返回用户id和token

无论是创建用户失败，或是token生成失败算注册不成功

此处事务其实不需要，练手而已
*/
func RegisterUser(user entity.User) response.UserEnrollResponse {
	user.Name = "user" + strconv.Itoa(rand.Intn(1000000))
	user.Avatar = "默认头像"
	user.BackgroundImage = "默认背景图"
	//密码加密
	user.Password = util.MD5(user.Password)

	//插入用户
	//开启事务
	tx := global.Datasource.Begin()
	//提交事务
	defer tx.Commit()
	id, err := dao.InsertBy(tx, user)
	//创建失败
	if err != nil {
		global.Logger.Debug(err.Error())
		tx.Rollback()
		return response.UserEnrollResponse{Response: response.Response{StatusCode: 1, StatusMsg: err.Error()}}
	}
	//生成token
	token, err := util.GenerateToken(id)
	if err != nil {
		global.Logger.Debug(err.Error())
		tx.Rollback()
		return response.UserEnrollResponse{Response: response.Response{StatusCode: 2, StatusMsg: err.Error()}}
	}
	//成功注册并签发token
	return response.UserEnrollResponse{Response: response.Response{StatusCode: 0, StatusMsg: "注册成功"}, UserId: id, Token: token}
}

/*
UserLogin 用户登录

加密参数密码并比对，成功颁发token

不区分用户存在与否
*/
func UserLogin(user entity.User) response.UserEnrollResponse {
	//查询数据
	rUser, err := dao.SelectUserByUsername(user.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.Logger.Debug(err.Error())
		return response.UserEnrollResponse{Response: response.Response{StatusCode: 2, StatusMsg: err.Error()}}
	}

	//密码校验
	if rUser.Password != util.MD5(user.Password) {
		global.Logger.Debug(user.Username + "登录密码错误!")
		return response.UserEnrollResponse{Response: response.Response{StatusCode: 1, StatusMsg: "用户名或密码错误"}}
	}

	//生成token
	token, err := util.GenerateToken(rUser.Id)
	if err != nil {
		global.Logger.Debug(err.Error())
		return response.UserEnrollResponse{Response: response.Response{StatusCode: 2, StatusMsg: err.Error()}}
	}

	//成功登录并签发token
	return response.UserEnrollResponse{Response: response.Response{StatusCode: 0, StatusMsg: "登录成功"}, UserId: rUser.Id, Token: token}
}

func GetUserInfo(erId, eeId uint64) response.UserInfoResponse {
	user, err := dao.SelectUserById(eeId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.UserInfoResponse{
				Response: response.Response{StatusCode: 1, StatusMsg: "用户不存在"},
			}
		} else {
			global.Logger.Debug(err)
			return response.UserInfoResponse{
				Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"},
			}
		}
	}

	return response.UserInfoResponse{
		Response: response.Response{StatusCode: 0, StatusMsg: "ok"},
		User: &response.UserInfo{
			User:     user,
			IsFollow: dao.IsFollow(erId, eeId),
		},
	}
}
