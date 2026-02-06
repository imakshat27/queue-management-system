package handler

import (
	"log"
	"net/http"

	"github.com/chirayusahu/queue-management-system/backend/common"
	"github.com/chirayusahu/queue-management-system/backend/config"
	"github.com/chirayusahu/queue-management-system/backend/database"
	"github.com/chirayusahu/queue-management-system/backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var app *fiber.App

// Build the app once (cold start)
func setupApp() *fiber.App {
	if app != nil {
		return app
	}

	cfg := config.LoadConfig()

	app = fiber.New(fiber.Config{
		AppName: cfg.AppName,
	})

	database.Connect(cfg.DatabaseUrl)
	database.DB.AutoMigrate(&models.User{})

	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return common.Respond(c, 200, true, "API is healthy", nil)
	})

	app.All("*", func(c *fiber.Ctx) error {
		return common.Respond(
			c,
			fiber.StatusNotFound,
			false,
			"Cannot "+c.Method()+" "+c.OriginalURL(),
			nil,
		)
	})

	log.Println("Fiber app initialized (Vercel)")
	return app
}

func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()

	app := setupApp()

	adaptor.FiberApp(app).ServeHTTP(w, r)
}
