package handler

import (
	"finances/src/config"
	"finances/src/models"
	"finances/src/services"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		config.LogError("error parsing user", zap.String("error", err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	config.LogDebug("creating user", zap.Any("user", user))
	if err := h.userService.CreateUser(&user); err != nil {
		config.LogError("error creating user", zap.String("error", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	config.LogDebug("getting user", zap.String("id", id))
	user, err := h.userService.GetUser(id)
	if err != nil {
		config.LogError("error getting user", zap.String("error", err.Error()))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	config.LogDebug("getting all users")
	users, err := h.userService.GetAllUsers()
	if err != nil {
		config.LogError("error getting all users", zap.String("error", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		config.LogError("error parsing user", zap.String("error", err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	config.LogDebug("updating user", zap.String("id", id))
	if err := h.userService.UpdateUser(&user); err != nil {
		config.LogError("error updating user", zap.String("error", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	config.LogDebug("deleting user", zap.String("id", id))
	if err := h.userService.DeleteUser(id); err != nil {
		config.LogError("error deleting user", zap.String("error", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
