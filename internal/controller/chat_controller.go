package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leosampsousa/psycoapi/internal/dto"
	"github.com/leosampsousa/psycoapi/internal/service"
)

type ChatController struct {
	cs *service.ChatService
}

func NewChatController(cs *service.ChatService) *ChatController{
	return &ChatController{cs: cs}
}

func (cc *ChatController) GetAllChats(c *gin.Context) {
	ctx := c.Request.Context()
	username := c.MustGet("username").(string)

	chats, err := cc.cs.GetAllChats(ctx, username)
	if err != nil {
		c.IndentedJSON(err.Code, gin.H{"mensagem": err.Message})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"chats": chats})
}

func (cc *ChatController) GetChatMessages(c *gin.Context) {
	ctx := c.Request.Context()
	loggedUser := c.MustGet("loggedUser").(*dto.UserDTO)
	chatIdParam := c.Query("chatId")

	chatId, err := strconv.ParseInt(chatIdParam, 10, 32)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"mensagem": "parametros inválidos"})
		return
	}

	messages, errMessage := cc.cs.GetChatMessage(ctx, loggedUser.ID, int32(chatId))
	if (errMessage != nil) {
		c.IndentedJSON(errMessage.Code, gin.H{"mensagem": errMessage.Message})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"messages": messages})
}