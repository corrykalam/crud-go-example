package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique; not null; type:varchar(25); column:username" validate:"required"`
	Email    string `json:"email" gorm:"unique; not null; type:varchar(100); column:email" validate:"required,email"`
	Password string `json:"password" gorm:"not null; type:varchar(200); column:password" validate:"required"`
	Age      int    `json:"age" gorm:"not null; column:age" validate:"required"`
	gorm.Model
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	pwd, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		fmt.Println("Failed to encrypt password: ", err)
		return err
	}

	u.Password = string(pwd)
	return nil
}
