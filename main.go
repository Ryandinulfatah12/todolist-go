package main

import (
    "todolist-go/config"
    "todolist-go/routes"
    "github.com/gofiber/fiber/v2"
    "log"
)

func main() {
    app := fiber.New()

    // Initialize database
    config.ConnectDB()

    // Setup routes
    routes.SetupRoutes(app)

    // Start server
    log.Fatal(app.Listen(":3000"))
}
