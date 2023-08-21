package dao

import (
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/entity"
	"gorm.io/gorm"
	"time"
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

func UpdateCommentCountBy(tx *gorm.DB, id uint64, step int) error {
	sql := "update video set comment_count=comment_count+? where id=?"
	return tx.Exec(sql, step, id).Error
}

func SelectVideosByCreateTime(latestTime time.Time, size int) ([]entity.Video, error) {
	var videos []entity.Video
	sql := "select id, user_id, title, play_url, cover_url, favorite_count, comment_count, extra, create_time from video where UNIX_TIMESTAMP(create_time)<? order by create_time desc limit 0,?"
	err := global.Datasource.Raw(sql, latestTime.Unix(), size).Scan(&videos).Error
	return videos, err
}

func InsertVideo(tx *gorm.DB, video entity.Video) error {
	sql := "insert into video(user_id, title, play_url, cover_url) value (?,?,?,?)"
	return tx.Exec(sql, video.UserId, video.Title, video.PlayURL, video.CoverURL).Error
}
