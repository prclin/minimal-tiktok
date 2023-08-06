package entity

// User user类
type User struct {
	ID              uint64 `json:"id" gorm:"column:id;index;primary_key;NOT NULL"`
	Name            string `json:"name" gorm:"column:name;NOT NULL"`
	FollowCount     int64  `json:"follow_count" gorm:"column:follow_count;NOT NULL"`
	FollowerCount   int64  `json:"follower_count" gorm:"column:follower_count;NOT NULL"`
	IsFollow        bool   `json:"is_follow" gorm:"column:is_follow;NOT NULL"`
	Avatar          string `json:"avatar" gorm:"column:avatar;NOT NULL"`
	BackgroundImage string `json:"background_image" gorm:"column:background_image;NOT NULL"`
	Signature       string `json:"signature" gorm:"column:signature;NOT NULL"`
	TotalFavorited  string `json:"total_favorited" gorm:"column:total_favorited;NOT NULL"`
	WorkCount       int64  `json:"work_count" gorm:"column:work_count;NOT NULL"`
	FavoriteCount   int64  `json:"favorite_count" gorm:"column:favorite_count;NOT NULL"`
}
