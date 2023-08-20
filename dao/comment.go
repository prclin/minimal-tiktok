package dao

import (
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/entity"
	"gorm.io/gorm"
)

func InsertComment(tx *gorm.DB, comment entity.Comment) (uint64, error) {
	var id uint64 //主键
	sql := "insert into comment(user_id, video_id, content) value (?,?,?)"
	if err := tx.Exec(sql, comment.UserId, comment.VideoId, comment.Content).Error; err != nil {
		return 0, err
	}
	if err := tx.Raw("select LAST_INSERT_ID()").Scan(&id).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func SelectCommentBy(tx *gorm.DB, id uint64) (entity.Comment, error) {
	var comment entity.Comment
	sql := "select id, user_id,video_id, content, extra, create_time from comment where id=?"
	err := tx.Raw(sql, id).Scan(&comment).Error
	return comment, err
}

func SelectCommentUserIdBy(id uint64) (uint64, error) {
	var userId uint64
	sql := "select user_id from comment where id=?"
	err := global.Datasource.Raw(sql, id).Scan(&userId).Error
	return userId, err
}

func DeleteCommentBy(tx *gorm.DB, id uint64) error {
	sql := "delete from comment where id=?"
	err := tx.Exec(sql, id).Error
	return err
}
