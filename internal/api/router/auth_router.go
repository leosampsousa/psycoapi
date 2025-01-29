package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leosampsousa/psycoapi/internal/controller"
)

func AuthRoute(router *gin.RouterGroup, ac *controller.AuthController) {
	router.POST("auth/login", ac.Login)
	router.POST("auth/login/available", ac.IsNewUsername)
	router.POST("auth/register", ac.Register)
}