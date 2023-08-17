package dao

import "gorm.io/gorm"

// InsertFollow 插入关注
func InsertFollow(tx *gorm.DB, followerId, followeeId uint64) error {
	sql := "insert into follow value (?,?)"
	return tx.Exec(sql, followerId, followeeId).Error
}

// DeleteFollow 删除关注
func DeleteFollow(tx *gorm.DB, followerId, followeeId uint64) error {
	sql := "delete from follow where follower_id =? and followee_id=?"
	return tx.Exec(sql, followerId, followeeId).Error
}
