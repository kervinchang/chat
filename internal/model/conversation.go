package model

import (
	"time"

	"gorm.io/gorm"
)

// Conversation 对话
type Conversation struct {
	gorm.Model
	UserID   uint
	Name     string
	Messages []Message
}

const WELCOME = "网络一线牵，珍惜这段缘！你好，我是ChatGPT🤖️"

// CreateConversation 新建对话
func CreateConversation(userID uint, conversationName string) (Conversation, error) {
	conversation := Conversation{UserID: userID, Name: conversationName}
	if err := db.Create(&conversation).Error; err != nil {
		return conversation, err
	}

	message := Message{ConversationID: conversation.ID, Content: WELCOME}
	if err := db.Create(&message).Error; err != nil {
		return conversation, err
	}
	return conversation, nil
}

// FindConversationsByUserID 根据用户ID查询对话列表
func FindConversationsByUserID(userID interface{}) ([]Conversation, Conversation, error) {
	var conversations []Conversation
	conversation := Conversation{}
	if err := db.Where("user_id = ?", userID).Find(&conversations).Error; err != nil {
		return conversations, conversation, err
	}
	if len(conversations) > 0 {
		conversation = conversations[0]
		db.Preload("Messages").First(&conversation)
	}
	return conversations, conversation, nil
}

// FindConversationByIDAndUserID 根据ID和用户ID查询对话详情
func FindConversationByIDAndUserID(id, userID interface{}) (Conversation, error) {
	var conversation Conversation
	if err := db.Where("id = ? AND user_id = ?", id, userID).First(&conversation).Error; err != nil {
		return conversation, err
	}
	db.Preload("Messages").First(&conversation)
	return conversation, nil
}

// DeleteConversation 删除对话
func DeleteConversation(id interface{}) error {
	var conversation Conversation
	if err := db.Where("id = ?", id).First(&conversation).Error; err != nil {
		return err
	}
	if err := db.Delete(&conversation).Error; err != nil {
		return err
	}
	return nil
}

// UpdateName 修改对话标题
func (conversation *Conversation) UpdateName(name string) error {
	conversation.Name = name
	if err := db.Save(&conversation).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUpdatedAt 更新对话时间
func (conversation *Conversation) UpdateUpdatedAt(updatedAt time.Time) error {
	conversation.UpdatedAt = updatedAt
	if err := db.Save(&conversation).Error; err != nil {
		return err
	}
	return nil
}

// DeleteMessages 清除对话记录
func (conversation *Conversation) DeleteMessages() error {
	if err := db.Model(&conversation).Association("Messages").Clear(); err != nil {
		return err
	}
	return nil
}
