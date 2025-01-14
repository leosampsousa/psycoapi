package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leosampsousa/psycoapi/internal/api/router"
	"github.com/leosampsousa/psycoapi/internal/controller"
	"github.com/leosampsousa/psycoapi/internal/db"
	"github.com/leosampsousa/psycoapi/internal/repository"
	"github.com/leosampsousa/psycoapi/internal/service"
)

func main () {

	connection := db.GetInstance()
	db := db.New(connection)
	userController := handleUserDependencies(db)

	defer connection.Close()

	appRouter := gin.Default()
	router.UserRoute(appRouter, userController)

	appRouter.Run("localhost:8080")
}

func handleUserDependencies(db *db.Queries) *controller.UserController {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	return controller.NewUserController(userService)
} 