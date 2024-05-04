package services

import (
	"time"

	"github.com/AliRamdhan/compstudioserver/config"
	"github.com/AliRamdhan/compstudioserver/internal/model"
)

type MessageService struct{}

func NewMessageService() *MessageService {
	return &MessageService{}
}

func (sc *MessageService) CreateMessage(message *model.Messages) error {
	message.MessageISRead = false
	message.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return config.DB.Create(message).Error
}

func (sc *MessageService) GetAllMessage() ([]model.Messages, error) {
	var messages []model.Messages
	if err := config.DB.Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func (ts *MessageService) GetMessageByServiceId(serviceId uint) ([]model.Messages, error) {
	var messages []model.Messages
	if err := config.DB.Where("message_service = ?", serviceId).Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func (sc *MessageService) MarkMessageAsRead(serviceId uint) error {
	// Update message status
	var message model.Messages
	// if err := config.DB.Model(&message).Where("message_id = ?", messageId).Update("message_is_read", true).Error; err != nil {
	// 	return err
	// }
	if err := config.DB.Model(&message).Where("message_service = ?", serviceId).Update("message_is_read", true).Error; err != nil {
		return err
	}
	return nil
}

func (sc *MessageService) UpdateMessage(messageId uint, updateMessage *model.Messages) error {
	var existingMessage model.Messages
	if err := config.DB.First(&existingMessage, "message_id = ?", messageId).Error; err != nil {
		return err // Product not found
	}

	// Update fields of existing product with the new values
	existingMessage.MessageContent = updateMessage.MessageContent
	existingMessage.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	// Save the updated product
	return config.DB.Save(&existingMessage).Error
}

func (ps *MessageService) DeleteMessage(messageId uint) error {
	// Find the product with the given ID
	var message model.Messages
	if err := config.DB.First(&message, "message_id = ?", messageId).Error; err != nil {
		return err // message not found
	}
	// Delete the message
	return config.DB.Delete(&message).Error
}
