package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Register godoc
// @Summary      Регистрация нового пользователя
// @Description  Регистрирует нового пользователя
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body  models.RegisterInput  true  "Данные регистрации"
// @Success      200  {object}  models.RegisterOutput
// @Router       /auth/register [post]
func (h *Handler) Register(c *gin.Context) {
	logrus.Info("register handler")
	c.Status(200)
}

// Login godoc
// @Security BearerAuth
// @Summary      Login
// @Description  логинит нового пользователя и возвращает JWT токен
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body  models.LoginInput  true  "Данные регистрации"
// @Success      200  {object}  models.LoginOutput
// @Router       /auth/login [post]
func (h *Handler) Login(c *gin.Context) {
	c.Status(200)
}
