package entity

import "time"

/*
User 用户实体
*/
type User struct {
<<<<<<< HEAD
	Id              uint64    `json:"id"` // 用户id
	Username        string    `json:"-"`
	Password        string    `json:"-"`
	Name            string    `json:"name"`             // 用户名称
	Avatar          string    `json:"avatar"`           // 用户头像
	BackgroundImage string    `json:"background_image"` // 用户个人页顶部大图
	Signature       string    `json:"signature"`        // 个人简介
	FollowCount     uint32    `json:"follow_count"`     // 关注总数
	FollowerCount   uint32    `json:"follower_count"`   // 粉丝总数
	TotalFavorited  string    `json:"total_favorited"`  // 获赞数量
	WorkCount       uint32    `json:"work_count"`       // 作品数
	FavoriteCount   uint32    `json:"favorite_count"`   // 喜欢数
	Extra           string    `json:"-"`
	CreateTime      time.Time `json:"-"`
=======
	Id              uint64 `json:"id"` // 用户id
	Username        string
	Password        string
	Name            string `json:"name"`             // 用户名称
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	Signature       string `json:"signature"`        // 个人简介
	FollowCount     uint32 `json:"follow_count"`     // 关注总数
	FollowerCount   uint32 `json:"follower_count"`   // 粉丝总数
	TotalFavorited  string `json:"total_favorited"`  // 获赞数量
	WorkCount       uint32 `json:"work_count"`       // 作品数
	FavoriteCount   uint32 `json:"favorite_count"`   // 喜欢数
	Extra           string
	CreateTime      time.Time
>>>>>>> 8a44ffc (feat: rebase from master and add video feed support)
}
