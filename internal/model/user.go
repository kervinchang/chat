package model

import (
	"gorm.io/gorm"
)

// User 用户
type User struct {
	gorm.Model
	Username string // 用户名
	Password string // 密码
}

func FindUserByID(ID interface{}) {

}

// FindUserByUsernameAndPassword 根据用户名和密码查询用户
func FindUserByUsernameAndPassword(username, password string) (User, error) {
	var user User
	if err := db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
