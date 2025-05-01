package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	customErr "gitlab.com/bobr-lord-messenger/gateway/internal/errors"
	"gitlab.com/bobr-lord-messenger/gateway/internal/middleware"
	"gitlab.com/bobr-lord-messenger/gateway/internal/models"
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
			"error": "failed to get user id",
		})
		return
	}

	res, err := h.service.User.GetMe(id.(string))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"requestID": requestID,
		}).Errorf("error getting user: %v", err)
		code, msg := customErr.ParseCustomError(err)
		c.AbortWithStatusJSON(code, gin.H{
			"error": msg,
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) UpdateMe(c *gin.Context) {
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
			"error": "failed to get user id",
		})
		return
	}
	var req models.UpdateMeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.WithFields(logrus.Fields{
			"requestID": requestID,
		}).Errorf("error parsing request: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "error parsing request",
		})
	}
	err := h.service.User.UpdateMe(id.(string), &req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"requestID": requestID,
		}).Errorf("error updating user: %v", err)
		code, msg := customErr.ParseCustomError(err)
		c.AbortWithStatusJSON(code, gin.H{
			"error": msg,
		})
		return
	}
	c.Status(200)
}

func (h *Handler) GetUsers(c *gin.Context) {
	requestID, ok := c.Get(middleware.RequestIDKey)
	if !ok {
		requestID = "unknown"
	}
	logrus.WithFields(logrus.Fields{
		"requestID": requestID,
	}).Info("handler.GetUsers")
	id, ok := c.Get(middleware.UserIDKey)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"requestID": requestID,
		}).Error("handler.GetUsers: failed to get user id")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get user id",
		})
		return
	}
	res, err := h.service.User.GetUsers(id.(string))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"requestID": requestID,
		}).Errorf("error getting user: %v", err)
		code, msg := customErr.ParseCustomError(err)
		c.AbortWithStatusJSON(code, gin.H{
			"error": msg,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetUserByID(c *gin.Context) {
	requestID, ok := c.Get(middleware.RequestIDKey)
	if !ok {
		requestID = "unknown"
	}
	logrus.WithFields(logrus.Fields{
		"requestID": requestID,
	}).Info("handler.GetUserByID")
	var req models.GetUserByIdRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.WithFields(logrus.Fields{
			"requestID": requestID,
		}).Errorf("error parsing request: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "error parsing request",
		})
		return
	}
	res, err := h.service.User.GetUserById(&req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"requestID": requestID,
		})
		code, msg := customErr.ParseCustomError(err)
		c.AbortWithStatusJSON(code, gin.H{
			"error": msg,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *Handler) GetUserByUsername(c *gin.Context) {
	requestID, ok := c.Get(middleware.RequestIDKey)
	if !ok {
		requestID = "unknown"
	}
	logrus.WithFields(logrus.Fields{
		"requestID": requestID,
	}).Info("handler.GetUserByUsername")
	var req models.GetUserByUsernameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.WithFields(logrus.Fields{
			"requestID": requestID,
		}).Errorf("error parsing request: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "error parsing request",
		})
		return
	}
	res, err := h.service.User.GetUserByUsername(&req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"requestID": requestID,
		}).Errorf("error getting user: %v", err)
		code, msg := customErr.ParseCustomError(err)
		c.AbortWithStatusJSON(code, gin.H{
			"error": msg,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
