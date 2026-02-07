package routes

import (
	"github.com/chirayusahu/queue-management-system/backend/common"
	"github.com/chirayusahu/queue-management-system/backend/modules/users"
	"github.com/gofiber/fiber/v2"
)

func V1Routes(app *fiber.App) {

	api := app.Group("/v1")

	api.Get("/", func(c *fiber.Ctx) error {
		return common.Respond(c, fiber.StatusOK, true, "V1 API is working!", nil)
	})

	users.UserRoutes(api)

}
