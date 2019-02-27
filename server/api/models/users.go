package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
}

func GetUsers() ([]*User, error) {
	var users []*User
	err := db.Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return users, nil
}
