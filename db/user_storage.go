package db

import types "github.com/ItsKiani/mahfel/types"

type UserStorage interface {
	GetUserByID(string) (*types.User, error)
}

type MongoUserStorage struct{}
