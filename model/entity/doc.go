/*
Package entity 存放所有的数据实体结构体，与数据库中的表一一对应
*/
package entity

import "time"

type User struct {
	ID              int64
	CreateTime      time.Time
	Name            string // 用户名称
	FollowCount     int64  // 关注总数
	FollowerCount   int64  // 粉丝总数
	Avatar          string //用户头像
	BackgroundImage string //用户个人页顶部大图
	Signature       string //个人简介
	TotalFavorited  int64  //获赞数量
	WorkCount       int64  //作品数量
	FavoriteCount   int64  //点赞数量
	Password        string //密码
	Username        string //用户名
	Extra           string //额外字段
}
