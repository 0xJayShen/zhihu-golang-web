package models

import "github.com/jinzhu/gorm"

type Auth struct {
	gorm.Model

	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) uint {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {

		return auth.ID
	}

	return -1
}
