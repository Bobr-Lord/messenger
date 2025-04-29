package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gitlab.com/bobr-lord-messenger/gateway/internal/jwt"
	"gitlab.com/bobr-lord-messenger/gateway/internal/middleware"
	"net/http"
)

var (
	connections = make(map[string]*websocket.Conn)
	ctx         = context.Background()
	upgrader    = websocket.Upgrader{}
	redisConn   *redis.Client
)

func initRedis() {
	redisConn = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

// Websocket godoc
// @Security BearerAuth
// @Summary      WebSocket
// @Description  websocket коннект
// @Tags         websocket
// @Router       /ws [get]
func (h *Handler) Websocket(c *gin.Context) {
	requestID, exists := c.Get(middleware.RequestIDKey)
	if !exists {
		requestID = "unknown"
	}
	logrus.WithFields(logrus.Fields{
		middleware.RequestIDKey: requestID,
	}).Info("Handle Websocket")

	token := c.GetHeader("Authorization")
	if token == "" {
		logrus.WithFields(logrus.Fields{
			middleware.RequestIDKey: requestID,
		}).Error("Token is empty")
		c.Status(http.StatusUnauthorized)
		return
	}
	userId, err := jwt.ParseJWT(token)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		logrus.WithFields(logrus.Fields{
			"requestID": requestID,
		}).Errorf("Invalid token: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"userId": userId,
	})

}
