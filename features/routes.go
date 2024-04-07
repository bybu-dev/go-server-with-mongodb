package features

import (
	adminRoutes "bybu/go-mongo-db/features/admin"
	userRoutes "bybu/go-mongo-db/features/user"

	"github.com/gofiber/fiber/v2"
)

var Routes = func(app *fiber.App) {
	adminRoute := app.Group("/admin");
	adminRoutes.AdminRoute(adminRoute)

	userRoute := app.Group("/user");
	userRoutes.UserRoute(userRoute)

}