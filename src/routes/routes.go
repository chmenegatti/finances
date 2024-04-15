package routes

import (
	"finances/src/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App, handlers *handler.UserHandler) {
	userGroup := app.Group("/api/v1/users")
	userGroup.Post("/", handlers.CreateUser)
	userGroup.Get("/:id", handlers.GetUser)
	userGroup.Get("/", handlers.GetAllUsers)
	userGroup.Put("/:id", handlers.UpdateUser)
	userGroup.Delete("/:id", handlers.DeleteUser)
}
