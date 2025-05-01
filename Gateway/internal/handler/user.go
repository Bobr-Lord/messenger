package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gitlab.com/bobr-lord-messenger/gateway/internal/middleware"
	"net/http"
)

func (h *Handler) GetMe(c *gin.Context) {
	requestID, ok := c.Get(middleware.RequestIDKey)
	if !ok {
		requestID = "unknown"
	}
	logrus.WithFields(logrus.Fields{
		"requestID": requestID,
	}).Info("handler.GetMe")

	id, ok := c.Get(middleware.UserIDKey)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"requestID": requestID,
		}).Error("handler.GetMe: failed to get user id")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get user id",
		})
		return
	}

	res, err := h.service.User.GetMe(id.(string))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"requestID": requestID,
		}).Errorf("error getting user: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": res,
	})
}

func (h *Handler) UpdateMe(c *gin.Context) {
	c.Status(200)
}

func (h *Handler) GetUsers(c *gin.Context) {
	c.Status(200)
}

func (h *Handler) AddContacts(c *gin.Context) {
	c.Status(200)
}

func (h *Handler) GetContacts(c *gin.Context) {
	c.Status(200)
}

func (h *Handler) GetUserByID(c *gin.Context) {
	c.Status(200)
}
func (h *Handler) GetUserByUsername(c *gin.Context) {
	c.Status(200)
}
