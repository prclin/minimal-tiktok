package service

import (
	"github.com/prclin/minimal-tiktok/dao"
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/entity"
)

func SendMessage(message entity.Message) bool {
	//接收消息的用户是否存在
	exist := dao.ExistsUser(message.ToUserId)
	if !exist {
		global.Logger.Warn("向不存在的用户发送消息")
		return false
	}
	//发送
	err := dao.InsertMessage(global.Datasource, message)
	if err != nil {
		global.Logger.Debug(err.Error())
		return false
	}
	return true
}
