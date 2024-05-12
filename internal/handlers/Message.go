package handlers

import (
	"fmt"
	"net/http"

	"github.com/AliRamdhan/compstudioserver/internal/model"
	"github.com/AliRamdhan/compstudioserver/internal/services"
	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	messageService *services.MessageService
}

func NewMessageHandler(sh *services.MessageService) *MessageHandler {
	return &MessageHandler{messageService: sh}
}

func (sh *MessageHandler) CreateMessage(c *gin.Context) {
	var message model.Messages
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId := message.MessageUser
	serviceId := message.MessageService
	if userId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication first"})
		return
	}

	// Check if customerId is missing or invalid
	if serviceId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing service ID"})
		return
	}
	if err := sh.messageService.CreateMessage(&message, userId, serviceId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Message created successfully", "Message": message})
}

func (sh *MessageHandler) GetAllMessage(c *gin.Context) {
	messages, err := sh.messageService.GetAllMessage()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "All Messages", "Messages": messages})
}

func (sh *MessageHandler) GetMessageAndMarkAsRead(c *gin.Context) {
	serviceIdStr := c.Param("serviceId")
	var serviceId uint
	_, err := fmt.Sscanf(serviceIdStr, "%d", &serviceId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}

	// Get messages by service ID
	messages, err := sh.messageService.GetMessageByServiceId(serviceId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Messages not found"})
		return
	}

	// Mark messages as read for the service
	if err := sh.messageService.MarkMessageAsRead(serviceId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark messages as read"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Messages retrieved and marked as read successfully", "serviceId": serviceId, "Messages": messages})
}

func (sh *MessageHandler) GetAllMessagesByServiceId(c *gin.Context) {
	serviceIdStr := c.Param("serviceId")
	var serviceId uint
	_, err := fmt.Sscanf(serviceIdStr, "%d", &serviceId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}

	// Get last message by service ID
	messages, err := sh.messageService.GetAllMessagesByServiceId(serviceId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Messages not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Last message retrieved successfully", "serviceId": serviceId, "Messages": messages})
}

// func (sh *MessageHandler) GetLastMessages(c *gin.Context) {
// 	messages, err := sh.messageService.GetAllMessagesByServiceId()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Last messages retrieved successfully", "Messages": messages})
// }

func (sh *MessageHandler) UpdateMessage(c *gin.Context) {
	messageIdStr := c.Param("messageId")
	var messageId uint
	_, err := fmt.Sscanf(messageIdStr, "%d", &messageId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message ID"})
		return
	}

	var updateMessage model.Messages
	if err := c.ShouldBindJSON(&updateMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := sh.messageService.UpdateMessage(messageId, &updateMessage); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message updated successfully", "Message": updateMessage})
}

func (sh *MessageHandler) DeleteMessage(c *gin.Context) {
	messageIdStr := c.Param("messageId")
	var messageId uint
	_, err := fmt.Sscanf(messageIdStr, "%d", &messageId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message ID"})
		return
	}

	if err := sh.messageService.DeleteMessage(messageId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message deleted successfully"})
}
