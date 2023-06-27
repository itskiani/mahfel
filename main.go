package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/ItsKiani/mahfel/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbUri = "mongodb://localhost:27017"
const dbName = "froum"
const userColl = "users"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	collection := client.Database(dbName).Collection(userColl)

	user := types.User{
		FirstName: "Ali",
		LastName:  "Kiani",
	}

	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

	listenPort := flag.String("port", ":5000", "The port address for api server")
	flag.Parse()

	app := fiber.New()
	app.Get("/", handleHome)

	app.Listen(*listenPort)
}

func handleHome(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "Its Working"})
}
