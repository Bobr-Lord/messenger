package handler

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"gitlab.com/bobr-lord-messenger/gateway/internal/models"
)

func (h *Handler) HandleWebsocketRequest(msg *models.MessageDelivery) error {
	socketID, err := h.redisCon.Get(context.Background(), "socket").Result()
	if err != nil {
		logrus.Errorf("Error connecting to redis: %v", err)
		return err
	}
	conn, ok := h.connections[socketID]
	if !ok {
		logrus.Errorf("Error connecting to redis for socket %v", socketID)
		return err
	}
	if conn == nil {
		logrus.Errorf("Error connecting to redis for socket: %v", socketID)
		return err
	}
	jsMes, err := json.Marshal(msg)
	if err != nil {
		logrus.Errorf("Error marshalling json: %v", err)
		return err
	}
	if err := conn.WriteMessage(websocket.TextMessage, jsMes); err != nil {
		logrus.Errorf("Error sending ping message: %v", err)
		return err
	}

	logrus.Info("Successfully send mess")
	return nil
}
