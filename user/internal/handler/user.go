package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gitlab.com/bobr-lord-messenger/user/internal/errors"
	"gitlab.com/bobr-lord-messenger/user/internal/middleware"
	"gitlab.com/bobr-lord-messenger/user/internal/models"
	"net/http"
)

func (h *Handler) GetMe(c *gin.Context) {
	requestId, ok := c.Get(middleware.RequestID)
	if !ok {
		requestId = "unknown"
	}
	id := c.GetHeader("id")
	logrus.WithFields(logrus.Fields{
		middleware.RequestID: requestId,
	}).Info("id: " + id)

	res, err := h.svc.User.GetMe(id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			middleware.RequestID: requestId,
		}).Info("err: " + err.Error())
		code, msg := errors.ParseCustomError(err)
		c.JSON(code, msg)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) UpdateMe(c *gin.Context) {
	requestId, ok := c.Get(middleware.RequestID)
	if !ok {
		requestId = "unknown"
	}
	id := c.GetHeader("id")
	logrus.WithFields(logrus.Fields{
		middleware.RequestID: requestId,
	}).Info("id: " + id)
	var req models.UpdateMeRequest
	if err := c.BindJSON(&req); err != nil {
		logrus.WithFields(logrus.Fields{
			middleware.RequestID: requestId,
		}).Info("err: " + err.Error())
		code, msg := errors.ParseCustomError(err)
		c.JSON(code, msg)
		return
	}
	err := h.svc.User.UpdateMe(id, &req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			middleware.RequestID: requestId,
		}).Info("err: " + err.Error())
		code, msg := errors.ParseCustomError(err)
		c.JSON(code, msg)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) GetUserById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) GetUserByName(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
