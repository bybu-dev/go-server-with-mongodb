package profileRoutes

import (
	"bybu/go-postgres/shared/module/middleware"

	"github.com/gofiber/fiber/v2"
)


func Route(router fiber.Router) {
	profileController := NewProfileController();

	router.Get("/", middleware.Validate.UserRole ,profileController.getUser);
	router.Put("/", middleware.Validate.UserRole ,profileController.updateUser);
}