package handlers

import (
	"github.com/gofiber/fiber/v2"
	"v2/internal/database"
	"v2/internal/models"
	"v2/internal/utils"
)

func Register(c *fiber.Ctx) error {
	var req models.AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	user := models.User{
		Email:        req.Email,
		PasswordHash: utils.GeneratePassword(req.Password),
	}
	res := database.DB.Create(&user)
	if res.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": res.Error.Error(),
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "user created",
	})
}
