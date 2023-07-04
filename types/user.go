package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email     string             `bson:"email" json:"email" validate:"email, required"`
	Password  string             `bson:"password" json:"_" validate:"required, min=6"`
	FirstName string             `bson:"first_name" json:"first_name" validate:"required, min=2, max=100"`
	LastName  string             `bson:"last_name" json:"last_name" validate:"required, min=2, max=100"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
