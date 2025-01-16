package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leosampsousa/psycoapi/internal/controller"
)

func AuthRoute(router *gin.RouterGroup, ac *controller.AuthController) {
	router.POST("/login", ac.Login)
}