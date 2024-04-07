package main

import (
	"bybu/go-mongo-db/features"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	features.Routes(app);

	if err := app.Listen(":8080"); err != nil {
		println(err)
	}
}