package routes

import (
    "todolist-go/controllers"
    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/api")
    api.Post("/todos", controllers.CreateTodoHandler)
    api.Get("/todos", controllers.GetTodosHandler)
    api.Get("/todos/:id", controllers.GetTodoByIDHandler)
    api.Put("/todos/:id", controllers.UpdateTodoHandler)
    api.Delete("/todos/:id", controllers.DeleteTodoHandler)
    api.Patch("/todos/:id/complete", controllers.CompleteTodoHandler)
}
