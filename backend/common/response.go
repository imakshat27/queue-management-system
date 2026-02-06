package common

import "github.com/gofiber/fiber/v2"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Respond(
	c *fiber.Ctx,
	status int,
	success bool,
	message string,
	data interface{},
) error {
	return c.Status(status).JSON(Response{
		Success: success,
		Message: message,
		Data:    data,
	})
}
