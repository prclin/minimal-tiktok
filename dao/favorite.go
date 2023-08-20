package dao

import (
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/entity"
	"gorm.io/gorm"
)

func IsFavorite(userId uint64, videoId uint64) bool {
	var favorited bool
	global.Datasource.Raw("select count(user_id) from favorite where user_id=? and video_id=?", userId, videoId).Scan(&favorited)
	return favorited
}

func InsertFavorite(tx *gorm.DB, favorite entity.Favorite) error {
	sql := "insert into favorite(user_id, video_id) value (?,?)"
	return tx.Exec(sql, favorite.UserId, favorite.VideoId).Error
}

func DeleteFavorite(tx *gorm.DB, favorite entity.Favorite) error {
	sql := "delete from favorite where user_id=? and video_id=?"
	return tx.Exec(sql, favorite.UserId, favorite.VideoId).Error
}

func SelectFavorites(userId uint64) ([]uint64, error) {
	var videoIds []uint64
	sql := "select video_id from favorite where user_id=?"
	err := global.Datasource.Raw(sql, userId).Scan(&videoIds).Error
	return videoIds, err
}
