package dao

import (
	"github.com/prclin/minimal-tiktok/model/entity"
	"gorm.io/gorm"
)

func InsertMessage(tx *gorm.DB, message entity.Message) error {
	sql := "insert into message(from_user_id, to_user_id, content) value (?,?,?)"
	return tx.Exec(sql, message.FromUserId, message.ToUserId, message.Content).Error
}
