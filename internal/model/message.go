package model

import (
	"gorm.io/gorm"
)

// Message 消息
type Message struct {
	gorm.Model
	ConversationID uint
	UserID         uint
	Content        string
}

// CreateMessage 新建消息
func CreateMessage(conversationID, userID uint, content string) (Message, error) {
	var message Message
	message = Message{ConversationID: conversationID, UserID: userID, Content: content}
	if err := db.Create(&message).Error; err != nil {
		return message, err
	}
	return message, nil
}

// FindMessagesByConversationID 根据对话ID查询消息列表
func FindMessagesByConversationID(conversationID interface{}) ([]Message, error) {
	var messages []Message
	if err := db.Where("conversation_id = ?", conversationID).Find(&messages).Error; err != nil {
		return messages, err
	}
	return messages, nil
}
