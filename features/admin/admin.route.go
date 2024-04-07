package adminRoutes

import (
	adminAuthRoutes "bybu/go-mongo-db/features/admin/auth"
	adminUsersRoutes "bybu/go-mongo-db/features/admin/users"

	"github.com/gofiber/fiber/v2"
)

func AdminRoute(route fiber.Router) {
	adminAuthRoute := route.Group("/auth");
	adminAuthRoutes.UserAuthRoute(adminAuthRoute);

	userManagementRoute := route.Group("/manage/users");
	adminUsersRoutes.AdminUsersRoute(userManagementRoute);
}