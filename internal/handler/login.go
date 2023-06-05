package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kervinchang/chat/internal/service"
)

// Login 登录
func Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	token, err := service.AuthorizeUser(username, password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.SetCookie("token", token, 60*60*24, "", "", false, true)
	ctx.Redirect(http.StatusSeeOther, "/chat")
}
