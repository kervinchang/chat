package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kervinchang/chat/config"
	"github.com/kervinchang/chat/internal/model"
)

func AuthorizeUser(username, password string) (string, error) {
	user, err := model.FindUserByUsernameAndPassword(username, password)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("用户身份验证失败")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	signedToken, err := token.SignedString([]byte(config.JwtKey))
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("token加密失败")
	}
	return signedToken, nil
}
