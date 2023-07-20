package db

import (
	types "github.com/ItsKiani/mahfel/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorage interface {
	GetUserByID(string) (*types.User, error)
}

type MongoUserStorage struct {
	client *mongo.Client
}

func NewMongoUserStorage(client *mongo.Client) *MongoUserStorage {
	return &MongoUserStorage{
		client: client,
	}
}
