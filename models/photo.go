package models

import "gorm.io/gorm"

type Photo struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Title    string `json:"title" gorm:"column:title;type:varchar(100)"`
	Caption  string `json:"caption" gorm:"column:caption;type:varchar(250)"`
	PhotoUrl string `json:"photo_url" gorm:"column:photo_url;type:varchar(100)"`
	UserId   int    `json:"user_id" gorm:"column:user_id"`
	gorm.Model
}

type GetPhoto struct {
	Photos Photo
	Users  User
}
