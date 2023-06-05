package service

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kervinchang/chat/internal/model"
	"github.com/sashabaranov/go-openai"
)

const (
	Bot     = 0
	GOODBYE = "我下线了，稍后再来提问吧～"
)

// ListConversation 对话列表
func ListConversation(userID interface{}) ([]model.Conversation, model.Conversation, error) {
	conversations, conversation, err := model.FindConversationsByUserID(userID)
	if err != nil {
		fmt.Println(err.Error())
		return conversations, conversation, errors.New("对话列表获取失败")
	}
	return conversations, conversation, nil
}

// NewConversation 新建对话
func NewConversation(userID interface{}, name string) (model.Conversation, error) {
	var conversation model.Conversation
	userIDFloat, ok := userID.(float64)
	if !ok {
		return conversation, errors.New("用户身份验证失败")
	}
	conversation, err := model.CreateConversation(uint(userIDFloat), name)
	if err != nil {
		fmt.Println(err.Error())
		return conversation, errors.New("对话创建失败")
	}
	return conversation, nil
}

// ShowConversation 获取对话详情
func ShowConversation(id string, userID interface{}) ([]model.Conversation, model.Conversation, error) {
	var conversation model.Conversation
	var conversations []model.Conversation
	conversation, err := model.FindConversationByIDAndUserID(id, userID)
	if err != nil {
		fmt.Println(err.Error())
		return conversations, conversation, errors.New("对话详情获取失败")
	}
	conversations, _, err = model.FindConversationsByUserID(userID)
	if err != nil {
		fmt.Println(err.Error())
		return conversations, conversation, errors.New("对话列表获取失败")
	}
	return conversations, conversation, nil
}

// UpdateConversation 更新对话信息
func UpdateConversation(id, userID interface{}, name string) error {
	conversation, err := model.FindConversationByIDAndUserID(id, userID)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("当前对话不存在或已被删除")
	}
	err = conversation.UpdateName(name)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("修改对话标题失败")
	}
	return nil
}

// DeleteConversation 删除对话
func DeleteConversation(id string) error {
	if err := model.DeleteConversation(id); err != nil {
		fmt.Println(err.Error())
		return errors.New("删除对话失败")
	}
	return nil
}

// NewMessage 新建消息
func NewMessage(ctx *gin.Context, conversationID, userID interface{}, content string) error {

	conversation, err := model.FindConversationByIDAndUserID(conversationID, userID)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("当前对话不存在或已被删除")
	}

	message, err := model.CreateMessage(conversation.ID, uint(userID.(float64)), content)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("消息保存失败")
	}

	if err := conversation.UpdateUpdatedAt(message.CreatedAt); err != nil {
		fmt.Println(err.Error())
		return errors.New("时间更新失败")
	}

	messages, err := model.FindMessagesByConversationID(conversationID)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("消息记录获取失败")
	}

	var chatCompletionMessages []openai.ChatCompletionMessage
	for _, message := range messages {
		var role string
		var content string
		if message.UserID == 0 {
			role = "assistant"
			content = message.Content
		} else {
			role = "user"
			content = message.Content
		}
		chatCompletionMessages = append(chatCompletionMessages, openai.ChatCompletionMessage{
			Role:    role,
			Content: content,
		})
	}

	// 在列表第一条插入一条 message
	chatCompletionMessages = append([]openai.ChatCompletionMessage{SystemMessage}, chatCompletionMessages...)
	replyContent, err := CreateChatCompletion(ctx, chatCompletionMessages)
	if err != nil {
		fmt.Println(err.Error())
		replyContent = GOODBYE
	}
	_, err = model.CreateMessage(conversation.ID, Bot, replyContent)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("消息保存失败")
	}
	return nil
}

// DeleteMessages 清除消息记录
func DeleteMessages(conversationID, userID interface{}) error {
	conversation, err := model.FindConversationByIDAndUserID(conversationID, userID)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("当前对话不存在或已被删除")
	}
	err = conversation.DeleteMessages()
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("消息删除失败")
	}
	return nil
}
