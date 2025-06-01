package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"gitlab.com/bobr-lord-messenger/gateway/internal/middleware"
	"gitlab.com/bobr-lord-messenger/gateway/internal/models"

	//"gitlab.com/bobr-lord-messenger/gateway/internal/models"
	"log"
	"net/http"
	"time"
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

	if err := h.redisCon.Set(c, "socket:"+userID, socketID, time.Minute*30).Err(); err != nil {
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
		if h.redisCon.Get(c, "socket:"+userID) == nil {
			logrus.WithFields(logrus.Fields{
				middleware.RequestIDKey: requestID,
			}).Info("connection closed")
			break
		}
		_, msg, err := conn.ReadMessage()
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

		if err := conn.WriteMessage(websocket.PongMessage, []byte("Message received")); err != nil {
			log.Println("Error sending message:", err)
			break
		}

		var in models.MessageDelivery

		if err := json.Unmarshal(msg, &in); err != nil {
			logrus.Info("Error json decoding message:", err)
			break
		}

		if err := h.prod.Send(c, []byte(in.ChatID), msg); err != nil {
			logrus.WithFields(logrus.Fields{
				middleware.RequestIDKey: requestID,
			}).Error(fmt.Sprintf("error sending message: %v", err))
			break
		}
	}
	delete(h.connections, socketID)
	h.redisCon.Del(c, "socket:"+userID)
}
