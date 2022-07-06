package models

import "gorm.io/gorm"

type Social struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	Name           string `json:"name" gorm:"column:name;type:varchar(100)"`
	SocialMediaUrl string `json:"social_media_url" gorm:"column:social_media_url;type:varchar(250)"`
	UserId         int    `json:"user_id" gorm:"column:user_id"`
	gorm.Model
}
