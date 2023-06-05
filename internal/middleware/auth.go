package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/kervinchang/chat/config"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 检查Cookie中是否包含Token
		if ck, err := ctx.Cookie("token"); err == nil {
			// Token存在于Cookie中的情况
			token, err := jwt.ParseWithClaims(ck, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
				// 验证 JWT 密钥
				return []byte(config.JwtKey), nil
			})
			if err == nil && token.Valid {
				// 将 userID 保存在上下文中
				claims := token.Claims.(*jwt.MapClaims)
				fmt.Println(claims)
				userID := (*claims)["userID"].(float64)
				ctx.Set("userID", userID)

				ctx.Next()
				return
			}
		}
		ctx.Redirect(http.StatusSeeOther, "/")
		ctx.Abort()
	}
}
