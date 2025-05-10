package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gitlab.com/bobr-lord-messenger/chat/internal/errors"
	"gitlab.com/bobr-lord-messenger/chat/internal/middleware"
	"gitlab.com/bobr-lord-messenger/chat/internal/models"
	"net/http"
)

// @Summary Создание Приватного чата
// @Tags API создание чата
// @Description Создание приватного чата
// @ID create-private-chat
// @Accept  json
// @Produce  json
// @Param input body models.CreatePrivateChatRequest true "credentials"
// @Success 200 {object} models.CreatePrivateChatResponse "data"
// @Failure default {object} errors.ErrorResponse
// @Router /chat/private [post]
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
		errResp := errors.NewErrorResponse(http.StatusUnauthorized, "id is required")
		c.JSON(http.StatusUnauthorized, errResp)
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
		errResp := errors.NewErrorResponse(http.StatusBadRequest, "invalid request body")
		c.JSON(http.StatusBadRequest, errResp)
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
		errCode, msg := errors.ParseCustomError(err)
		errResp := errors.NewErrorResponse(errCode, msg)
		c.JSON(http.StatusInternalServerError, errResp)
	}
	logrus.WithFields(logrus.Fields{
		"requestId": requestId,
	}).Infof("CreateChat response %+v", res)
	c.JSON(http.StatusCreated, res)
}

// @Summary Создание Группы
// @Tags API создание чата
// @Description Создание публичного чата
// @ID create-public-chat
// @Accept  json
// @Produce  json
// @Param input body models.CreatePublicChatRequest true "credentials"
// @Success 200 {object} models.CreatePublicChatResponse "data"
// @Failure default {object} errors.ErrorResponse
// @Router /chat/public [post]
func (h *Handler) CreatePublicChat(c *gin.Context) {
	requestId, ok := c.Get(middleware.RequestID)
	if !ok {
		requestId = "unknown"
	}
	logrus.WithFields(logrus.Fields{
		"requestId": requestId,
	}).Info("CreatePublicChat")
	id := c.GetHeader("id")
	id = "57fd4ddb-ab67-4385-9ef2-9b9137496fbd"
	if id == "" {
		logrus.WithFields(logrus.Fields{
			"requestId": requestId,
		}).Info("user ID required")
		errResp := errors.NewErrorResponse(http.StatusUnauthorized, "id is required")
		c.JSON(http.StatusUnauthorized, errResp)
		return
	}
	logrus.WithFields(logrus.Fields{
		middleware.RequestID: requestId,
	}).Info("id: " + id)
	var req *models.CreatePublicChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.WithFields(logrus.Fields{
			"requestId": requestId,
		}).Infof("invalid request body, %v", err)
		errResp := errors.NewErrorResponse(http.StatusBadRequest, "invalid request body")
		c.JSON(http.StatusBadRequest, errResp)
		return
	}
	logrus.WithFields(logrus.Fields{
		middleware.RequestID: requestId,
	}).Infof("request body: %+v", req)
	res, err := h.svc.Chat.CreatePublicChat(id, req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"requestId": requestId,
		}).Infof("create chat failed, %v", err)
		errResp := errors.NewErrorResponse(http.StatusInternalServerError, "create chat failed")
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	logrus.WithFields(logrus.Fields{
		"requestId": requestId,
	}).Infof("CreateChat response %+v", res)
	c.JSON(http.StatusCreated, res)
}

func (h *Handler) GetChats(c *gin.Context) {

}

func (h *Handler) GetChatHistory(c *gin.Context) {

}

func (h *Handler) AddChat(c *gin.Context) {

}
