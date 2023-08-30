package routes

import (
	"rdp/gotodos/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func TodoRoutes(app fiber.Router) {
	r := app.Group("/todos")

	r.Post("/", controllers.CreateTodo)
	r.Get("/", controllers.GetTodos)
	r.Get("/:id", controllers.GetTodoById)
	r.Patch("/:id", controllers.UpdateTodo)
	r.Delete("/:id", controllers.DeleteTodo)
}
