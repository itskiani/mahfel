package api

import (
	"github.com/ItsKiani/mahfel/db"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStorage db.MongoUserStorage
}

func HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(id)
}

func HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON("Kiani")
}
