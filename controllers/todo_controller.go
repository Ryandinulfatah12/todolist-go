package controllers

import (
	"todolist-go/models"
	"todolist-go/repositories"
	"todolist-go/services"
	"todolist-go/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateTodoHandler(c *fiber.Ctx) error {
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return utils.SendError(c, err)
	}
	response := services.CreateTodoService(&todo)
	return c.JSON(response)
}

func GetTodosHandler(c *fiber.Ctx) error {
	response := services.GetTodosService()
	return c.JSON(response)
}

func GetTodoByIDHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.SendError(c, err)
	}
	response := services.GetTodoByIDService(uint(id))
	return c.JSON(response)
}

func UpdateTodoHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.SendError(c, err)
	}
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return utils.SendError(c, err)
	}
	todo.ID = uint(id)
	response := services.UpdateTodoService(&todo)
	return c.JSON(response)
}

func DeleteTodoHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.SendError(c, err)
	}
	todo, err := repositories.GetTodoByID(uint(id))
	if err != nil {
		return utils.SendError(c, err)
	}
	response := services.DeleteTodoService(&todo)
	return c.JSON(response)
}

func CompleteTodoHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.SendError(c, err)
	}
	response := services.CompleteTodoService(uint(id))
	return c.JSON(response)
}
