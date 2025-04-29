package handler

import "github.com/gin-gonic/gin"

func (h *Handler) GetMe(c *gin.Context) {
	c.Status(200)
}

func (h *Handler) UpdateMe(c *gin.Context) {
	c.Status(200)
}

func (h *Handler) GetUsers(c *gin.Context) {
	c.Status(200)
}

func (h *Handler) AddContacts(c *gin.Context) {
	c.Status(200)
}

func (h *Handler) GetContacts(c *gin.Context) {
	c.Status(200)
}

func (h *Handler) GetUserByID(c *gin.Context) {
	c.Status(200)
}
func (h *Handler) GetUserByUsername(c *gin.Context) {
	c.Status(200)
}
