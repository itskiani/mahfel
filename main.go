package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbUri = "mongodb://localhost:27017"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}

	listenPort := flag.String("port", ":5000", "The port address for api server")
	flag.Parse()

	app := fiber.New()
	app.Get("/", handleHome)

	app.Listen(*listenPort)
}

func handleHome(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "Its Working"})
}
