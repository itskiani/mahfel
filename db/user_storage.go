package db

import (
	types "github.com/ItsKiani/mahfel/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorage interface {
	GetUserByID(string) (*types.User, error)
}

type MongoUserStorage struct{}

func NewMongoUserStorage(c *mongo.Client) *MongoUserStorage {
	return &MongoUserStorage{
		client: client,
	}
}
