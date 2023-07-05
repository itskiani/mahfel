package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	api "github.com/ItsKiani/mahfel/api/routes"
	"github.com/ItsKiani/mahfel/types"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const userColl = "users"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("DB_URL")))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	collection := db.Database(os.Getenv("DB_NAME")).Collection(userColl)

	user := types.User{
		FirstName: "Ali",
		LastName:  "Kiani",
		Email:     "codewithkiani@gmail.com",
	}

	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

	listenPort := flag.String("port", ":5000", "The port address for api server")
	flag.Parse()

	app := fiber.New()

	/** Routes */
	apiV1 := app.Group("/api/v1")
	apiV1.Get("/user", api.HandleGetUsers)

	app.Listen(*listenPort)
}
