package router

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kervinchang/chat/internal/handler"
	"github.com/kervinchang/chat/internal/middleware"
)

func Setup() *gin.Engine {
	router := gin.Default()

	// 时间格式处理函数
	router.SetFuncMap(template.FuncMap{
		"formatTime": func(t time.Time) string {
			return t.Format("2006/01/02 15:04:05")
		},
	})

	// 渲染模板文件目录
	router.LoadHTMLGlob("templates/*")

	// 静态资产文件目录
	router.StaticFile("/favicon.ico", "./assets/img/favicon.ico")
	router.Static("/js", "./assets/js")
	router.Static("/css", "./assets/css")
	router.Static("/fonts", "./assets/fonts")
	router.Static("/img", "./assets/img")

	// 根路由
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	// 登录
	router.POST("/login", handler.Login)

	// 聊天相关接口组
	chatAPI := router.Group("/chat")

	// 使用鉴权中间件
	chatAPI.Use(middleware.AuthMiddleware())

	// 聊天首页
	chatAPI.GET("", handler.Chat)

	// 新建对话
	chatAPI.POST("/conversation", handler.NewConversation)

	// 对话详情
	chatAPI.GET("/conversation/:id", handler.ShowConversation)

	// 修改对话名称
	chatAPI.PUT("/conversation/:id", handler.UpdateConversation)

	// 删除对话
	chatAPI.DELETE("/conversation/:id", handler.DeleteConversation)

	// 新建对话消息
	chatAPI.POST("/conversation/:id/message", handler.NewMessage)

	// 清除对话消息
	chatAPI.DELETE("/conversation/:id/message", handler.DeleteMessage)

	return router
}
