package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reply struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserId    int                `bson:"user_id" json:"user_id" validate:"required"`
	ThreadId  int                `bson:"thread_id" json:"thread_id" validate:"required"`
	Body      string             `bson:"body" json:"body" validate:"required"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
