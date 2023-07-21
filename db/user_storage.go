package db

import (
	"os"

	types "github.com/ItsKiani/mahfel/types"
	"go.mongodb.org/mongo-driver/mongo"
)

const userColl = "users"

type UserStorage interface {
	GetUserByID(string) (*types.User, error)
}

type MongoUserStorage struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoUserStorage(client *mongo.Client) *MongoUserStorage {
	return &MongoUserStorage{
		client:     client,
		collection: client.Database(os.Getenv("DB_NAME")).Collection(userColl),
	}
}
