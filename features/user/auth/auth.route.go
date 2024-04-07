package authRoutes

import (
	passwordRoutes "bybu/go-mongo-db/features/user/auth/password"
	"bybu/go-postgres/shared/module/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserAuthRoute(route fiber.Router) {
	authController := *NewUserAuthController();

	route.Post("/register", authController.register);
	route.Post("/login", authController.login);
	route.Post("/token", middleware.Validate.UserRole, authController.accessToken);

	passwordRoute := route.Group("/password");
	passwordRoutes.UserPasswordRoute(passwordRoute);
}