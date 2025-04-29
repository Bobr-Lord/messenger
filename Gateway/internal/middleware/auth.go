package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gitlab.com/bobr-lord-messenger/gateway/internal/jwt"
	"net/http"
)

const UserIDKey = "userID"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(UserIDKey, "1")
		c.Next()
		return
		token := c.GetHeader("Authorization")
		if token == "" {
			logrus.Info("No Authorization header")
			c.Status(http.StatusUnauthorized)
			return
		}
		userId, err := jwt.ParseJWT(token)
		if err != nil {
			c.Status(http.StatusUnauthorized)
			return
		}
		c.Set(UserIDKey, userId)
		c.Next()
	}
}
