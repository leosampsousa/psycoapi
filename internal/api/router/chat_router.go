package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leosampsousa/psycoapi/internal/controller"
)

func ChatRoute(router *gin.RouterGroup, cc *controller.ChatController) {
	router.GET("/chat", cc.GetAllChats)
	router.GET("/chat/messages", cc.GetChatMessages)
}