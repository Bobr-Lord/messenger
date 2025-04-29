package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gitlab.com/bobr-lord-messenger/gateway/internal/middleware"
	"log"
	"net/http"
)

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
	userHeader, exists := c.Get(middleware.UserIDKey)
	if !exists {
		c.Status(http.StatusUnauthorized)
		return
	}
	userID, ok := userHeader.(string)
	if !ok {
		c.Status(http.StatusUnauthorized)
		return
	}
	logrus.WithFields(logrus.Fields{
		middleware.RequestIDKey: requestID,
	}).Info("Handle Websocket")
	socketID := uuid.NewString()

	if err := h.redisCon.Set(c, "socket:"+userID, socketID, 0).Err(); err != nil {
		logrus.WithFields(logrus.Fields{
			middleware.RequestIDKey: requestID,
		}).Error(fmt.Sprintf("error setting socket for redis: %v", err))
		c.Status(http.StatusInternalServerError)
		return
	}
	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			middleware.RequestIDKey: requestID,
		}).Error(fmt.Sprintf("error upgrading connection: %v", err))
		return
	}
	defer conn.Close()
	h.connections[socketID] = conn

	logrus.WithFields(logrus.Fields{
		"socketID":              socketID,
		middleware.RequestIDKey: requestID,
	}).Info(fmt.Sprintf("User %s connected with socketID %s\n", userID, socketID))

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"socketID":              socketID,
				middleware.RequestIDKey: requestID,
			}).Error("ReadMessage error:", err)
			break
		}
		logrus.WithFields(logrus.Fields{
			"socketID":              socketID,
			middleware.RequestIDKey: requestID,
		}).Info(string(msg))

		if err := conn.WriteMessage(messageType, []byte("Message received")); err != nil {
			log.Println("Error sending message:", err)
			return
		}
	}
	delete(h.connections, socketID)
}
