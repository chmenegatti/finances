package main

import (
	"finances/src/config"
	"finances/src/database"
	"finances/src/handler"
	"finances/src/repositories"
	"finances/src/routes"
	"finances/src/services"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap/zapcore"
)

func main() {

	app := fiber.New()

	logFile := "messages.log"
	config.InitLogger(zapcore.DebugLevel, logFile)

	config.LoadConfig()

	database.Connect()

	userRepository := repositories.NewUserRepository(database.DB)
	userService := services.NewUserService(*userRepository)
	userHandler := handler.NewUserHandler(userService)
	routes.SetupUserRoutes(app, userHandler)

	err := app.Listen(":5000")
	if err != nil {
		return
	}

}
