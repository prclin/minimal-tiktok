package service

import (
	"errors"
	"github.com/prclin/minimal-tiktok/dao"
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/entity"
	"github.com/prclin/minimal-tiktok/model/response"
	"gorm.io/gorm"
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

func GetChatHistory(from, to uint64) response.ChatResponse {
	messages, err := dao.SelectMessagesBy(from, to)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.Logger.Debug(err.Error())
		return response.ChatResponse{Response: response.Response{StatusCode: 2, StatusMsg: "服务器内部错误"}}
	}
	if messages == nil { //没有消息不返回0值
		messages = make([]entity.Message, 0)
	}
	return response.ChatResponse{Response: response.Response{StatusCode: 0, StatusMsg: "获取成功"}, MessageList: messages}
}
