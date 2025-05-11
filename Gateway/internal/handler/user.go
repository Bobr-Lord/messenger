package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	customErr "gitlab.com/bobr-lord-messenger/gateway/internal/errors"
	"gitlab.com/bobr-lord-messenger/gateway/internal/middleware"
	"gitlab.com/bobr-lord-messenger/gateway/internal/models"
	"net/http"
)

// GetMe godoc
// @Security BearerAuth
// @Summary      GetME
// @Description  получить свои данные
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.LoginResponse
// @Failure default {object} errors.ErrorResponse
// @Router       /user/me [get]
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
		errResp := customErr.NewErrorResponse(http.StatusInternalServerError, "internal server error")
		c.JSON(http.StatusOK, errResp)
		return
	}

	res, err := h.service.User.GetMe(id.(string))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"requestID": requestID,
		}).Errorf("error getting user: %v", err)
		code, msg := customErr.ParseCustomError(err)
		errResp := customErr.NewErrorResponse(code, msg)
		c.JSON(http.StatusOK, errResp)
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateMe godoc
// @Security BearerAuth
// @Summary      UpdateMe
// @Description  обновить свои данные
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        input  body  models.UpdateMeRequest  true  "Данные пользователя"
// @Failure default {object} errors.ErrorResponse
// @Router       /user/me [put]
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
		errResp := customErr.NewErrorResponse(http.StatusInternalServerError, "internal server error")
		c.JSON(http.StatusOK, errResp)
		return
	}
	var req models.UpdateMeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.WithFields(logrus.Fields{
			"requestID": requestID,
		}).Errorf("error parsing request: %v", err)
		errResp := customErr.NewErrorResponse(http.StatusInternalServerError, "internal server error")
		c.JSON(http.StatusOK, errResp)
		return
	}
	err := h.service.User.UpdateMe(id.(string), &req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"requestID": requestID,
		}).Errorf("error updating user: %v", err)
		code, msg := customErr.ParseCustomError(err)
		errResp := customErr.NewErrorResponse(code, msg)
		c.JSON(http.StatusOK, errResp)
		return
	}
	c.Status(200)
}

// GetUsers godoc
// @Security BearerAuth
// @Summary      GetUsers
// @Description  получить всех пользователей
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.GetUsersResponse
// @Failure default {object} errors.ErrorResponse
// @Router       /user/users [get]
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
		errResp := customErr.NewErrorResponse(http.StatusInternalServerError, "internal server error")
		c.JSON(http.StatusOK, errResp)
		return
	}
	res, err := h.service.User.GetUsers(id.(string))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"requestID": requestID,
		}).Errorf("error getting user: %v", err)
		code, msg := customErr.ParseCustomError(err)
		errResp := customErr.NewErrorResponse(code, msg)
		c.JSON(http.StatusOK, errResp)
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetUserByID godoc
// @Security BearerAuth
// @Summary      GetUserByID
// @Description  получить пользователя по id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        input  body  models.GetUserByIdRequest  true  "id пользователя"
// @Success      200  {object}  models.GetUserByIdResponse
// @Failure default {object} errors.ErrorResponse
// @Router       /user/id [get]
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
		errResp := customErr.NewErrorResponse(http.StatusInternalServerError, "internal server error")
		c.JSON(http.StatusOK, errResp)
		return
	}
	res, err := h.service.User.GetUserById(&req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"requestID": requestID,
		})
		code, msg := customErr.ParseCustomError(err)
		errResp := customErr.NewErrorResponse(code, msg)
		c.JSON(http.StatusOK, errResp)
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetUserByUsername godoc
// @Security BearerAuth
// @Summary      GetUserByUsername
// @Description  получить пользователя по username
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        input  body  models.GetUserByUsernameRequest  true  "username пользователя"
// @Success      200  {object}  models.GetUserByUsernameResponse
// @Failure default {object} errors.ErrorResponse
// @Router       /user/name [get]
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
		errResp := customErr.NewErrorResponse(http.StatusInternalServerError, "internal server error")
		c.JSON(http.StatusOK, errResp)
		return
	}
	res, err := h.service.User.GetUserByUsername(&req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"requestID": requestID,
		}).Errorf("error getting user: %v", err)
		code, msg := customErr.ParseCustomError(err)
		errResp := customErr.NewErrorResponse(code, msg)
		c.JSON(http.StatusOK, errResp)
		return
	}

	c.JSON(http.StatusOK, res)
}
