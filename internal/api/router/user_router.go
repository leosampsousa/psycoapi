package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leosampsousa/psycoapi/internal/api/middleware"
	"github.com/leosampsousa/psycoapi/internal/controller"
)

func UserRoute(router *gin.RouterGroup, uc *controller.UserController, md *middleware.AuthenticationMiddleware) {
	router.GET("/user", uc.GetUser)
	router.POST("/user", uc.Create)
}