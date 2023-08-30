package main

import (
	"rdp/gotodos/app/models"
	"rdp/gotodos/app/routes"
	"rdp/gotodos/config"
	"rdp/gotodos/config/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	database.Migrate(&models.Todo{})

	app := fiber.New(fiber.Config{
		AppName: config.AppName,
	})

	routes.TodoRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("GO FIBER TODO APP!!!")
	})

	app.Listen("127.0.0.1:3000")
}
