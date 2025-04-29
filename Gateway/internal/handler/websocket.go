package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
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
	//token := c.GetHeader("Authorization")
	//if token == "" {
	//	c.Status(http.StatusUnauthorized)
	//}
	//userId, err := jwt.ParseJWT(token)
	//if err != nil {
	//	c.Status(http.StatusUnauthorized)
	//	return
	//}
	//c.JSON(http.StatusOK, gin.H{
	//	"userId": userId,
	//})

}
