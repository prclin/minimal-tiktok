package dao

import (
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/entity"
)

func SelectUserInfoById(id uint64) entity.User {
	var user entity.User
	sql := "select id,username,name,avatar,background_image,signature,follow_count,follower_count,total_favorited,work_count,favorite_count from user where id=?"
	global.Datasource.Raw(sql, id).Scan(&user)
	return user
}
