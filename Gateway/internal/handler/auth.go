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
// @Failure default {object} errors.ErrorResponse
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
		errResp := errors.NewErrorResponse(http.StatusBadRequest, "Invalid register request")
		c.JSON(http.StatusBadRequest, errResp)
		return
	}
	res, err := h.service.Auth.Register(&req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			middleware.RequestIDKey: requestId,
		}).Error(err)
		code, err := errors.ParseCustomError(err)
		errResp := errors.NewErrorResponse(code, err)
		c.JSON(http.StatusBadRequest, errResp)
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
// @Description  авторизация и генерация jwt
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body  models.LoginRequest  true  "Данные регистрации"
// @Success      200  {object}  models.LoginResponse
// @Failure default {object} errors.ErrorResponse
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
		errResp := errors.NewErrorResponse(http.StatusBadRequest, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}
	logrus.WithFields(logrus.Fields{
		middleware.RequestIDKey: requestId,
	}).Infof("request: %+v", req)
	res, err := h.service.Auth.Login(&req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			middleware.RequestIDKey: requestId,
		}).Error(err)
		code, err := errors.ParseCustomError(err)
		errResp := errors.NewErrorResponse(code, err)
		c.AbortWithStatusJSON(code, errResp)
		return
	}
	logrus.WithFields(logrus.Fields{
		middleware.RequestIDKey: requestId,
	}).Info("Login success")
	c.JSON(http.StatusOK, res)
}
