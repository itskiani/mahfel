package api

import "github.com/gofiber/fiber/v2"

func HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(id)
}

func HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON("Kiani")
}
