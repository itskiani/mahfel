package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
)

func main() {
	listenPort := flag.String("port", ":5000", "The port address for api server")
	flag.Parse()

	app := fiber.New()
	app.Get("/", handleHome)

	app.Listen(*listenPort)
}

func handleHome(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "Its Working"})
}
