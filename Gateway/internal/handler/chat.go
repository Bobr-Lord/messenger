package handler

import "github.com/gin-gonic/gin"

// CreatePrivateChat godoc
// @Security BearerAuth
// @Summary      CreatePrivateChat
// @Description  создать приватный чат
// @Tags         chats
// @Accept       json
// @Produce      json
// @Param        input  body  models.CreatePrivateChatRequest  true  "Данные чата"
// @Success      200  {object}  models.CreatePrivateChatResponse
// @Failure default {object} errors.ErrorResponse
// @Router       /chat/private [post]
func (h *Handler) CreatePrivateChat(c *gin.Context) {}

// CreatePublicChat godoc
// @Security BearerAuth
// @Summary      CreatePublicChat
// @Description  создать группу
// @Tags         chats
// @Accept       json
// @Produce      json
// @Param        input  body  models.CreatePublicChatRequest  true  "Данные чата"
// @Success      200  {object}  models.CreatePrivateChatResponse
// @Failure default {object} errors.ErrorResponse
// @Router       /chat/public [post]
func (h *Handler) CreatePublicChat(c *gin.Context) {}

// GetMeChats godoc
// @Security BearerAuth
// @Summary      GetMeChats
// @Description  получить свои чаты
// @Tags         chats
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.GetMeChatsResponse
// @Failure default {object} errors.ErrorResponse
// @Router       /chat [get]
func (h *Handler) GetMeChats(c *gin.Context) {}

// GetChatUsers godoc
// @Security BearerAuth
// @Summary      GetChatUsers
// @Description  получить пользователей чата
// @Tags         chats
// @Accept       json
// @Produce      json
// @Param chat_id path string true "ID чата"
// @Success      200  {object}  models.GetChatUsersResponse
// @Failure default {object} errors.ErrorResponse
// @Router       /chat/{chat_id}/users [get]
func (h *Handler) GetChatUsers(c *gin.Context) {}
