package types

import "time"

type User struct {
	Id        string    `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string    `bson:"first_name" json:"first_name"`
	LastName  string    `bson:"last_name" json:"last_name"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
