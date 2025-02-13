package main

import (
	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"github.com/leosampsousa/psycoapi/internal/api/middleware"
	"github.com/leosampsousa/psycoapi/internal/api/router"
	"github.com/leosampsousa/psycoapi/internal/controller"
	"github.com/leosampsousa/psycoapi/internal/db"
	"github.com/leosampsousa/psycoapi/internal/repository"
	"github.com/leosampsousa/psycoapi/internal/service"
	"github.com/leosampsousa/psycoapi/internal/ws"
)

type Config struct {
	Host	string `env:"HOST" envDefault:"localhost"`
    Port    string `env:"APP_PORT" envDefault:"8080"`
    DBUrl   string `env:"DATABASE_URL" envDefault:"postgres://postgres:admin@localhost:5432/psycodb?sslmode=disable"`
	JwtSecret []byte `env:"JWT_SECRET" envDefault:"a1ebcfe43217edfdc81f5f191c7be2b3027a6f48ac9c48218fcb3431ad55a1f4e5e18af0ca43a4cfb14dbf31f85239da8b129f1f4252888fffa4c44fc5fa0ba35788ce34306c75afda74874e0c7a13181703c1eba5caa731cd2cedf297330bc0e2b1bf46710bc54942189735533e8be9f02d9b5373e6d67a0d58b2b59f74617b83dfb928d09a0b329d23eb81ba1a82f456394941d5a6cc3e51f6dd2da1a206fc76b453816c9eed901bf98b2ae414f57f308ee40797fb6bb3e7383b8313e246c669bfd686e90f0cf5c333db80f400d39cfac28949cf3d92b7718741af00a725abe507e5a8330f2fe62a92fbbde203118cc0b91753e8360eb2e9a3fa79c3f78aeb"`
}

func main () {

	cfg := Config{}
	env.Parse(&cfg)
	
	connection := db.GetInstance(cfg.DBUrl)
	db := db.New(connection)
	defer connection.Close()

	tokenService := service.NewTokenService(cfg.JwtSecret)
	authMiddleware := middleware.NewAuthenticationMiddleware(tokenService)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService) 

	authController := controller.NewAuthController(tokenService, userService)

	wsManager := ws.NewManager()

	appRouter := gin.Default()

	publicRoutes := appRouter.Group("") 
	{
		router.AuthRoute(publicRoutes, authController)
		router.WSRoute(publicRoutes, wsManager)
	}

	protectedRoutes := appRouter.Group("")
	protectedRoutes.Use(authMiddleware.ValidateJWT())					
	{
		router.UserRoute(protectedRoutes, userController, authMiddleware)
	}

	appRouter.Run(cfg.Host + ":" + cfg.Port)
}