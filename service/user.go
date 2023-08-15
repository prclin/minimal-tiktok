package service

import (
	"github.com/prclin/minimal-tiktok/dao"
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/entity"
	"github.com/prclin/minimal-tiktok/model/response"
	"github.com/prclin/minimal-tiktok/util"
	"math/rand"
	"strconv"
)

/*
RegisterUser

加密密码并补充默认用户信息，注册成功后返回用户id和token

无论是创建用户失败，或是token生成失败算注册不成功
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

func GetUserInfo(id uint64) response.User {
	user := dao.SelectUserInfoById(id)
	return response.User{User: user, IsFollow: false}
}
