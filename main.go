package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	api "github.com/ItsKiani/mahfel/api/routes"
	"github.com/ItsKiani/mahfel/db"
	"github.com/ItsKiani/mahfel/types"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const userColl = "users"

func main() {
	listenPort := flag.String("port", ":5000", "The port address for api server")
	flag.Parse()

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("DB_URL")))
	if err != nil {
		log.Fatal(err)
	}
	userHandler := api.NewUserHandler(db.NewMongoUserStorage(client))

	app := fiber.New()

	/** Routes */
	apiV1 := app.Group("/api/v1")
	apiV1.Get("/user", userHandler.HandleGetUsers)

	app.Listen(*listenPort)
}
