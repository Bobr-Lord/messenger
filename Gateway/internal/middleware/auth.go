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
		//c.Set(UserIDKey, "1")
		//c.Next()
		//return
		requestId, ok := c.Get(RequestIDKey)
		if !ok {
			requestId = "unknown"
		}

		token := c.GetHeader("Authorization")
		if token == "" {
			logrus.WithFields(logrus.Fields{
				RequestIDKey: requestId,
			}).Info("empty token")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		userId, err := jwt.ParseJWT(token)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				RequestIDKey: requestId,
			}).Infof("invalid token: %v", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		logrus.WithFields(logrus.Fields{
			UserIDKey: userId,
		}).Infof("authorized user id: %v", userId)
		c.Set(UserIDKey, userId)
		c.Next()
	}
}
