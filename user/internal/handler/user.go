package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetMe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) UpdateMe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) GetUserById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) GetUserByName(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
