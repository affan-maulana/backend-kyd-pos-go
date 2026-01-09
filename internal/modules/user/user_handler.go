package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wpcodevo/golang-fiber-jwt/internal/modules/entity"
)

func GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(entity.User)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}
