package auto

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username        string `json:"username" gorm:"type:varchar(200);;not null;uniqueIndex:user_name"`
	Password        string `json:"password" gorm:"not null;"`
	Avatar          string `json:"avatar"`                         // 用户头像
	FollowCount     int64  `json:"follow_count" gorm:"not null"`   // 关注总数
	FollowerCount   int64  `json:"follower_count" gorm:"not null"` // 粉丝总数
	BackgroundImage string `json:"background_image"`               //背景图片
	IsFollow        bool   `json:"is_follow"`                      // true-已关注，false-未关注 在数据库中这个字段没用
	Signature       string `json:"signature"`                      // 个人简介
	TotalFavorited  int64  `json:"total_favorited"`                // 获赞数量
	WorkCount       int64  `json:"work_count"`                     // 作品数
}

func (*User) TableName() string {
	return "user"
}