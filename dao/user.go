package dao

import (
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/entity"
	"gorm.io/gorm"
)

/*
InsertBy 插入用户
返回 id
*/
func InsertBy(tx *gorm.DB, user entity.User) (uint64, error) {
	var id uint64
	sql := "insert into user(username, password, name, avatar, background_image) value(?,?,?,?,?)"
	if err := tx.Exec(sql, user.Username, user.Password, user.Name, user.Avatar, user.BackgroundImage).Error; err != nil {
		return 0, err
	}
	if err := tx.Raw("select LAST_INSERT_ID()").Scan(&id).Error; err != nil {
		return 0, err
	}
	return id, nil
}

/*
SelectUserByUsername 通过用户名查找用户
*/
func SelectUserByUsername(username string) (entity.User, error) {
	var user entity.User
	sql := "select id,username,password from user where username=?"
	err := global.Datasource.Raw(sql, username).Scan(&user).Error
	return user, err
}

func SelectUserById(id uint64) (entity.User, error) {
	var user entity.User
	sql := "select id,username,name,avatar,background_image,signature,follow_count,follower_count,total_favorited,work_count,favorite_count from user where id=?"
	err := global.Datasource.Raw(sql, id).Scan(&user).Error
	return user, err
}

func SelectUserByIds(ids []uint64) ([]entity.User, error) {
	var users []entity.User
	sql := "select id,username,name,avatar,background_image,signature,follow_count,follower_count,total_favorited,work_count,favorite_count from user where id in ?"
	err := global.Datasource.Raw(sql, ids).Scan(&users).Error
	return users, err
}

func IsFollow(followerId, followeeId uint64) bool {
	var followed bool
	sql := "select count(*) from follow where follower_id=? and followee_id=?"
	global.Datasource.Raw(sql, followerId, followeeId).Scan(&followed)
	return followed
}

func SelectFolloweeBy(followerId uint64) ([]uint64, error) {
	var followees []uint64
	sql := "select followee_id from follow where follower_id=?"
	err := global.Datasource.Raw(sql, followerId).Scan(&followees).Error
	return followees, err
}

func SelectFollowerBy(followeeId uint64) ([]uint64, error) {
	var followers []uint64
	sql := "select follower_id from follow where followee_id=?"
	err := global.Datasource.Raw(sql, followeeId).Scan(&followers).Error
	return followers, err
}

func ExistsUser(id uint64) bool {
	var exist bool
	sql := "select count(id) from user where id=?"
	global.Datasource.Raw(sql, id).Scan(&exist)
	return exist
}
