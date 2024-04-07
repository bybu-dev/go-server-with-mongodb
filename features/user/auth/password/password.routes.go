package passwordRoutes

import (
	"github.com/gofiber/fiber/v2"
)

func UserPasswordRoute(route fiber.Router) {
	passwordController := NewUserPasswordController();
	route.Post("/send-reset-code", passwordController.sendResetCode);
	route.Post("/reset-password", passwordController.resetPassword);
}