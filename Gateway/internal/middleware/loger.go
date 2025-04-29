package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const RequestIDKey = "requestID"

func RequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := uuid.New()

		logrus.WithFields(logrus.Fields{
			RequestIDKey: requestId,
		}).Info("New request received")
		c.Set(RequestIDKey, requestId)
		c.Next()
	}
}
