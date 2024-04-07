package adminAuthRoutes

import (
	"github.com/gofiber/fiber/v2"
)

func UserAuthRoute(route fiber.Router) {
	adminAuthController := *NewAdminAuthController();

	route.Post("/create", adminAuthController.create);
	route.Post("/login", adminAuthController.login);
}