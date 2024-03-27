package db

import (
	"context"
	"fmt"

	"github.com/rapinbook/hotel-reservation-go/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore interface {
	InsertUser(context.Context, *types.User) (primitive.ObjectID, error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   client.Database("hotel-reservation").Collection("user_profile"),
	}
}

func (s *MongoUserStore) InsertUser(c context.Context, user *types.User) (primitive.ObjectID, error) {
	// c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	res, err := s.coll.InsertOne(c, user)
	if err != nil {
		fmt.Printf("Cannot insert user %v", user)
	}
	id := res.InsertedID.(primitive.ObjectID)
	return id, nil
}
