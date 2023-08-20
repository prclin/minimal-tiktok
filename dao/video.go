package dao

import (
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/entity"
)

func SelectVideosByUserId(userId uint64) []entity.Video {
	//获取指定用户的video发布列表
	var videos []entity.Video
	global.Datasource.Raw("select id,title,play_url,cover_url,favorite_count,comment_count,extra,create_time from video where user_id=?", userId).Scan(&videos)
	return videos
}

func ExistsVideo(id uint64) bool {
	var exist bool
	sql := "select count(id) from video where id=?"
	global.Datasource.Raw(sql, id).Scan(&exist)
	return exist
}

func SelectVideosBy(ids []uint64) ([]entity.Video, error) {
	var videos []entity.Video
	sql := "select id,title,play_url,cover_url,favorite_count,comment_count,extra,create_time from video where id in ?"
	err := global.Datasource.Raw(sql, ids).Scan(&videos).Error
	return videos, err
}
