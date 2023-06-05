package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kervinchang/chat/internal/service"
)

// Chat 聊天页面
func Chat(ctx *gin.Context) {

	userID, ok := ctx.Get("userID")
	if !ok {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	conversations, conversation, err := service.ListConversation(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.HTML(http.StatusOK, "chat.html", gin.H{
		"conversations": conversations,
		"conversation":  conversation,
	})
}

// NewConversation 新建对话
func NewConversation(ctx *gin.Context) {

	userID, ok := ctx.Get("userID")
	if !ok {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	name := ctx.PostForm("name")
	conversation, err := service.NewConversation(userID, name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/chat/conversation/"+strconv.Itoa(int(conversation.ID)))
}

// ShowConversation 对话详情
func ShowConversation(ctx *gin.Context) {
	id := ctx.Param("id")

	userID, ok := ctx.Get("userID")
	if !ok {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	conversations, conversation, err := service.ShowConversation(id, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.HTML(http.StatusOK, "chat.html", gin.H{
		"conversations": conversations,
		"conversation":  conversation,
	})
}

// UpdateConversation 修改对话标题
func UpdateConversation(ctx *gin.Context) {
	userID, ok := ctx.Get("userID")
	if !ok {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	id := ctx.Param("id")

	name := ctx.PostForm("name")

	if err := service.UpdateConversation(id, userID, name); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/chat/conversation/"+id)
}

// DeleteConversation 删除对话
func DeleteConversation(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := service.DeleteConversation(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/chat")
}

// NewMessage 新建消息
func NewMessage(ctx *gin.Context) {
	conversationID := ctx.Param("id")

	userID, ok := ctx.Get("userID")
	if !ok {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	content := ctx.PostForm("content")
	if len(strings.TrimSpace(content)) == 0 {
		ctx.Redirect(http.StatusSeeOther, "/chat/conversation/"+conversationID)
		return
	}

	err := service.NewMessage(ctx, conversationID, userID, content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/chat/conversation/"+conversationID)
}

// DeleteMessage 删除消息
func DeleteMessage(ctx *gin.Context) {
	conversationID := ctx.Param("id")

	userID, ok := ctx.Get("userID")
	if !ok {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	if err := service.DeleteMessages(conversationID, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/chat/conversation/"+conversationID)
}
