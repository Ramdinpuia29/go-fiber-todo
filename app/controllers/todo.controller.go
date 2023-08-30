package controllers

import (
	"rdp/gotodos/app/models"
	"rdp/gotodos/app/repositories"

	"github.com/gofiber/fiber/v2"
)

func GetTodos(c *fiber.Ctx) error {
	var todos []models.Todo

	repositories.GetAllTodos(&todos)

	return c.Status(fiber.StatusOK).JSON(todos)
}

func GetTodoById(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo
	var resp models.Response

	err := repositories.GetTodoById(&todo, id)

	if err != nil {
		resp.Status = fiber.StatusNotFound
		resp.Message = fiber.ErrNotFound.Message
		resp.Data = nil

		return c.Status(fiber.StatusNotFound).JSON(resp)
	} else {
		resp.Status = fiber.StatusOK
		resp.Message = "Success"
		resp.Data = todo

		return c.Status(fiber.StatusOK).JSON(resp)
	}
}

func CreateTodo(c *fiber.Ctx) error {
	newTodo := new(models.Todo)
	var resp models.Response

	c.BodyParser(newTodo)

	if newTodo.Title == "" {
		resp.Status = fiber.StatusUnprocessableEntity
		resp.Message = "Title is required"
		resp.Data = nil
		return c.Status(fiber.StatusUnprocessableEntity).JSON(resp)
	} else {
		repositories.CreateTodo(newTodo)

		resp.Status = fiber.StatusCreated
		resp.Message = "Success"
		resp.Data = newTodo
		return c.Status(fiber.StatusCreated).JSON(resp)
	}
}

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo
	var resp models.Response
	updatedTodo := new(models.Todo)

	c.BodyParser(updatedTodo)

	err := repositories.GetTodoById(&todo, id)

	if err != nil {
		resp.Status = fiber.StatusNotFound
		resp.Message = fiber.ErrNotFound.Message
		resp.Data = nil

		return c.Status(fiber.StatusNotFound).JSON(resp)
	}

	if updatedTodo.Title == "" {
		resp.Status = fiber.StatusUnprocessableEntity
		resp.Message = "Title is required"
		resp.Data = nil
		return c.Status(fiber.StatusUnprocessableEntity).JSON(resp)
	} else {
		repositories.UpdateTodo(updatedTodo, id)

		resp.Status = fiber.StatusOK
		resp.Message = "Success"
		resp.Data = updatedTodo
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo
	var resp models.Response

	err := repositories.DeleteTodo(&todo, id)

	if err != nil {
		resp.Status = fiber.StatusNotFound
		resp.Message = fiber.ErrNotFound.Message
		resp.Data = nil

		return c.Status(fiber.StatusNotFound).JSON(resp)
	} else {
		resp.Status = fiber.StatusOK
		resp.Message = "Success"
		resp.Data = nil

		return c.Status(fiber.StatusOK).JSON(resp)
	}
}
