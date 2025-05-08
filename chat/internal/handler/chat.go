package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gitlab.com/bobr-lord-messenger/chat/internal/middleware"
	"gitlab.com/bobr-lord-messenger/chat/internal/models"
	"net/http"
)

func (h *Handler) CreatePrivateChat(c *gin.Context) {
	requestId, ok := c.Get(middleware.RequestID)
	if !ok {
		requestId = "unknown"
	}
	logrus.WithFields(logrus.Fields{
		"requestId": requestId,
	}).Info("CreateChat")
	id := c.GetHeader("id")
	if id == "" {
		logrus.WithFields(logrus.Fields{
			"request_id": requestId,
		}).Info("user ID required")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "id is required"})
		return
	}
	logrus.WithFields(logrus.Fields{
		middleware.RequestID: requestId,
	}).Info("id: " + id)

	var req *models.CreatePrivateChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.WithFields(logrus.Fields{
			"requestId": requestId,
		}).Infof("invalid request body, %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	logrus.WithFields(logrus.Fields{
		"requestId": requestId,
	}).Infof("request body: %v", req)
	res, err := h.svc.Chat.CreatePrivateChat(id, req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"requestId": requestId,
		}).Infof("create chat failed, %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "create chat failed"})
		return
	}
	logrus.WithFields(logrus.Fields{
		"requestId": requestId,
	}).Infof("CreateChat response %+v", res)
	c.JSON(http.StatusCreated, res)
}

func (h *Handler) CreatePublicChat(c *gin.Context) {

}

func (h *Handler) GetChats(c *gin.Context) {

}

func (h *Handler) GetChatHistory(c *gin.Context) {

}

func (h *Handler) AddChat(c *gin.Context) {

}
