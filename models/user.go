package models

import (
	"github.com/jinzhu/gorm"
)

type Auth struct {
	gorm.Model

	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) uint {
	var auth Auth
	db.Where("username = ? AND password = ?",username,password).First(&auth)
	if auth.ID > 0 {
		return auth.ID
	}

	return 0
}
