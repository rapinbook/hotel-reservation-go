package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	HotelName string             `bson:"hotel_name" json:"hotel_name"`
	City      string             `bson:"city" json:"city"`
	Room      []*Room            `bson:"room" json:"room"`
	Rating    float64            `bson:"rating" json:"rating"`
}

// insert new room -> then insert hotel as well
type Room struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	HotelID primitive.ObjectID `bson:"hotel_id,omitempty" json:"hotel_id,omitempty"`
	Price   float64            `bson:"price" json:"price"`
	Size    string             `bson:"size" json:"size"`
}
