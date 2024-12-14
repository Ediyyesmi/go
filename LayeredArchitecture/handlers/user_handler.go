package handlers

import (
	"github.com/gofiber/fiber/v2"
	"project/services"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var input struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
	}

	err := h.Service.RegisterUser(input.Name, input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).SendString("User registered successfully")
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var input struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
	}

	isAuthenticated, err := h.Service.AuthenticateUser(input.Name, input.Password)
	if err != nil || !isAuthenticated {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid credentials")
	}

	return c.Status(fiber.StatusOK).SendString("Login successful")
}
