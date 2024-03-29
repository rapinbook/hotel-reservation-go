package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email     string             `BSON:"email" JSON:"email"`
	FirstName string             `BSON:"first_name" JSON:"first_name"`
	LastName  string             `BSON:"last_name" JSON:"last_name"`
	IsAdmin   bool               `BSON:"is_admin,omitempty" JSON:"is_admin,omitempty"`
}
