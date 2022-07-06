package models

import "gorm.io/gorm"

type Comment struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	UserId  int    `json:"user_id" gorm:"column:user_id"`
	PhotoId int    `json:"photo_id" gorm:"column:photo_id"`
	Message string `json:"message" gorm:"column:message;type:varchar(250)"`
	gorm.Model
}
