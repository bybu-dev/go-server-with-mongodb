package adminUsersRoutes

import (
	"bybu/go-postgres/shared/module/middleware"

	"github.com/gofiber/fiber/v2"
)

func AdminUsersRoute(route fiber.Router) {
	var userManagementController = *newUserController();
	
	route.Post("/getall", middleware.Validate.AdminRole, userManagementController.getAllUsers);
	route.Post("/ban", middleware.Validate.AdminRole, userManagementController.banUser);
	route.Post("/unban", middleware.Validate.AdminRole, userManagementController.unbanUser);

}