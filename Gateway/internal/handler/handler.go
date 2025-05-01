package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "gitlab.com/bobr-lord-messenger/gateway/docs" // тут будет документация
	"gitlab.com/bobr-lord-messenger/gateway/internal/config"
	"gitlab.com/bobr-lord-messenger/gateway/internal/middleware"
	"gitlab.com/bobr-lord-messenger/gateway/internal/service"
	"net/http"
)

type Handler struct {
	service     *service.Service
	upgrader    *websocket.Upgrader
	connections map[string]*websocket.Conn
	redisCon    *redis.Client
	cfg         *config.Config
}

func NewHandler(srv *service.Service, redisConn *redis.Client, cfg *config.Config) *Handler {
	return &Handler{
		service:     srv,
		connections: make(map[string]*websocket.Conn),
		redisCon:    redisConn,
		upgrader: &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // разрешаем все соединения, для разработки ок
			},
		},
		cfg: cfg,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()
	r.Use(middleware.RequestMiddleware())
	r.GET("/ws", h.Websocket)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := r.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
	}
	user := r.Group("/user")
	{
		user.Use(middleware.AuthMiddleware())
		user.GET("/me", h.GetMe)
		user.PUT("/me", h.UpdateMe)
		user.GET("/users", h.GetUsers)
		user.GET("/id", h.GetUserByID)
		user.GET("/name", h.GetUserByUsername)
	}

	chat := r.Group("/chat")
	{
		chat.Use(middleware.AuthMiddleware())
		chat.POST("/", h.CreateChat)
		chat.GET("/", h.GetChats)
		chat.GET("/:id", h.GetMessagesFromChat)
	}

	message := r.Group("/message")
	{
		message.Use(middleware.AuthMiddleware())
		message.PUT("/upd", h.UpdateMessageStatus)
		message.GET("/:id", h.GetUnsentMessages)
	}

	return r
}
