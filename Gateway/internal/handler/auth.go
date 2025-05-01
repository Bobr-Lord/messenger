package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gitlab.com/bobr-lord-messenger/gateway/internal/errors"
	"gitlab.com/bobr-lord-messenger/gateway/internal/middleware"
	"gitlab.com/bobr-lord-messenger/gateway/internal/models"
	"net/http"
)

// Register godoc
// @Summary      Регистрация нового пользователя
// @Description  Регистрирует нового пользователя
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body  models.RegisterRequest  true  "Данные регистрации"
// @Success      200  {object}  models.RegisterResponse
// @Router       /auth/register [post]
func (h *Handler) Register(c *gin.Context) {
	requestId, ok := c.Get(middleware.RequestIDKey)
	if !ok {
		requestId = "unknown"
	}
	logrus.WithFields(logrus.Fields{
		middleware.RequestIDKey: requestId,
	}).Info("Register")

	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.WithFields(logrus.Fields{
			middleware.RequestIDKey: requestId,
		}).Errorf("Invalid register request: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.service.Auth.Register(&req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			middleware.RequestIDKey: requestId,
		}).Error(err)
		code, err := errors.ParseCustomError(err)
		c.AbortWithStatusJSON(code, gin.H{"error": err})
		return
	}
	logrus.WithFields(logrus.Fields{
		middleware.RequestIDKey: requestId,
	}).Info("Register success")
	c.JSON(http.StatusOK, res)
}

// Login godoc
// @Security BearerAuth
// @Summary      Login
// @Description  логинит нового пользователя и возвращает JWT токен
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body  models.LoginRequest  true  "Данные регистрации"
// @Success      200  {object}  models.LoginResponse
// @Router       /auth/login [post]
func (h *Handler) Login(c *gin.Context) {
	requestId, ok := c.Get(middleware.RequestIDKey)
	if !ok {
		requestId = "unknown"
	}
	logrus.WithFields(logrus.Fields{
		middleware.RequestIDKey: requestId,
	}).Info("Login")
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.WithFields(logrus.Fields{
			middleware.RequestIDKey: requestId,
		}).Errorf("Invalid login request: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.service.Auth.Login(&req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			middleware.RequestIDKey: requestId,
		}).Error(err)
		code, err := errors.ParseCustomError(err)
		c.AbortWithStatusJSON(code, gin.H{"error": err})
		return
	}
	logrus.WithFields(logrus.Fields{
		middleware.RequestIDKey: requestId,
	}).Info("Login success")
	c.JSON(http.StatusOK, res)
}
