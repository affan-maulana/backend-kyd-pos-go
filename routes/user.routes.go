package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wpcodevo/golang-fiber-jwt/internal/middleware"
	userHandler "github.com/wpcodevo/golang-fiber-jwt/internal/modules/user"
)

func UserRoutes(router fiber.Router) {
	router.Get("/users/me", middleware.DeserializeUser, userHandler.GetMe)
}
