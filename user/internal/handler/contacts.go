package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) AddContact(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) GetContacts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
