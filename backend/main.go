package main

import (
	"log"

	"github.com/chirayusahu/queue-management-system/backend/common"
	"github.com/chirayusahu/queue-management-system/backend/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: config.LoadConfig().AppName,
	})

	// Global Middlewares
	app.Use(logger.New())  // Logs requests
	app.Use(recover.New()) // Recovers from panics

	// Health check (good practice)
	app.Get("/health", func(c *fiber.Ctx) error {
		return common.Respond(c, fiber.StatusOK, true, "API is healthy", nil)
	})

	// Start server
	log.Println("Server running on http://localhost:" + config.LoadConfig().Port)
	log.Fatal(app.Listen(":" + config.LoadConfig().Port))
}
