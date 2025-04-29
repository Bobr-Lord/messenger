package handler

import "github.com/gin-gonic/gin"

func (h *Handler) UpdateMessageStatus(c *gin.Context) {
	c.Status(200)
}

func (h *Handler) GetUnsentMessages(c *gin.Context) {
	c.Status(200)
}
