package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leosampsousa/psycoapi/internal/controller"
)

func UserRoute(router *gin.Engine, uc *controller.UserController) {
	router.GET("/user", uc.GetUser)
	router.POST("/user", uc.Create)
}