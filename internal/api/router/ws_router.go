package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leosampsousa/psycoapi/internal/ws"
)

func WSRoute(router *gin.RouterGroup, ws *ws.Manager) {
	router.GET("/ws", ws.ServeWS)
}