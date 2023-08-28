package dao

import (
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/entity"
	"gorm.io/gorm"
)

func InsertMessage(tx *gorm.DB, message entity.Message) error {
	sql := "insert into message(from_user_id, to_user_id, content) value (?,?,?)"
	return tx.Exec(sql, message.FromUserId, message.ToUserId, message.Content).Error
}

func SelectMessagesBy(from, to uint64) ([]entity.Message, error) {
	var messages []entity.Message
	sql := "select id,from_user_id, to_user_id, content, create_time from message where from_user_id=? and to_user_id=?"
	err := global.Datasource.Raw(sql, from, to).Scan(&messages).Error
	return messages, err
}

func SelectLastMessageBy(from, to uint64) (entity.Message, error) {
	var message entity.Message
	sql := "select id,from_user_id, to_user_id, content, create_time from message where from_user_id=? and to_user_id=? order by create_time limit 0,1"
	err := global.Datasource.Raw(sql, from, to).Scan(&message).Error
	return message, err
}
