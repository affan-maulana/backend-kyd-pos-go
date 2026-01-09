package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wpcodevo/golang-fiber-jwt/internal/middleware"
	authHandler "github.com/wpcodevo/golang-fiber-jwt/internal/modules/auth"
)

func AuthRoutes(router fiber.Router) {
	router.Route("/auth", func(authRouter fiber.Router) {
		authRouter.Post("/register", authHandler.SignUpUser)
		authRouter.Post("/login", authHandler.SignInUser)
		authRouter.Get("/logout", middleware.DeserializeUser, authHandler.LogoutUser)
	})
}
