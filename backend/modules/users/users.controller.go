package users

import (
	"github.com/chirayusahu/queue-management-system/backend/common"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {

	query := c.Queries()
	test := c.IP()

	return common.Respond(c, fiber.StatusOK, true, "Get All Users", fiber.Map{
		"query": query,
		"ip":    test,
	})
}

func DeleteUser(c *fiber.Ctx) error {

	UserId := c.Params("id")

	ApiResponse := fiber.Map{
		"user_id": UserId,
	}

	return common.Respond(c, fiber.StatusOK, true, "Delete User", ApiResponse)
}
