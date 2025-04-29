package handler

import "github.com/gin-gonic/gin"

func (h *Handler) CreateChat(c *gin.Context) {
	c.Status(200)
}

func (h *Handler) GetChats(c *gin.Context) {
	c.Status(200)
}

func (h *Handler) GetMessagesFromChat(c *gin.Context) {
	c.Status(200)
}
