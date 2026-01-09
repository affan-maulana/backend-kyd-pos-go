package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	micro := fiber.New()
	app.Mount("/api", micro)

	// Setup all module routes
	AuthRoutes(micro)
	UserRoutes(micro)

	// Health check
	micro.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Server is running! JWT Authentication with Golang, Fiber, and GORM",
		})
	})

	// 404 handler
	micro.All("*", func(c *fiber.Ctx) error {
		path := c.Path()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("Path: %v does not exists on this server", path),
		})
	})
}
