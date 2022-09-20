package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/mapsapi/controllers"
)

func AuthRoutes(router fiber.Router) {
	// /api/v1/auth
	authRouter := router.Group("/auth")

	// /api/v1/auth/client
	authRouter.Post("/client", controllers.Login)
}
