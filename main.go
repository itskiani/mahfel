package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/", handleHome)
	app.Listen(":5000")
}

func handleHome(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "Its Working"})
}
