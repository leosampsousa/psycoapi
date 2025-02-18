package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leosampsousa/psycoapi/internal/controller"
)

func UserRoute(router *gin.RouterGroup, uc *controller.UserController) {
	router.GET("/user/:username", uc.GetUser)
	router.GET("/user/friends", uc.GetFriends)
	router.POST("/user/friends/:friendId", uc.AddFriend)
}