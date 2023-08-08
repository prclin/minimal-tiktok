package dao

import "github.com/prclin/minimal-tiktok/global"

func IsFavorite(userId uint64, videoId uint64) bool {
	var count int
	global.Datasource.Raw("select count(user_id) from favorite where user_id=? and video_id=?", userId, videoId).Scan(&count)
	return count != 0
}
