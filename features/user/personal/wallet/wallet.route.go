package walletRoutes

import (
	"bybu/go-postgres/shared/module/middleware"

	"github.com/gofiber/fiber/v2"
)


func Route(router fiber.Router) {
	walletController := NewUserWalletController();

	router.Get("/fund", middleware.Validate.UserRole ,walletController.fund);
	router.Put("/withdraw", middleware.Validate.UserRole ,walletController.withdraw);
}