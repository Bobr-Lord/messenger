package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gitlab.com/bobr-lord-messenger/gateway/internal/errors"
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
	logrus.Info("register handler")
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.service.Auth.Register(&req)
	if err != nil {
		code, err := errors.ParseCustomError(err)
		c.JSON(code, gin.H{"error": err})
		return
	}
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
	c.Status(200)
}
