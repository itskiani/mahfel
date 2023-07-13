package api

import (
	"github.com/ItsKiani/mahfel/db"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStorage db.UserStorage
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userStorage.GetUserByID(id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON("Kiani")
}
