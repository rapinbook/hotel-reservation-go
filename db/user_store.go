package db

import (
	"context"
	"fmt"
	"log"

	"github.com/rapinbook/hotel-reservation-go/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore interface {
	InsertUser(context.Context, *types.User) (primitive.ObjectID, error)
	GetUserByID(context.Context, primitive.ObjectID) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	UpdateUser(context.Context) (primitive.ObjectID, error)
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

// func (s *MongoUserStore) GetUserByID(c context.Context, ID primitive.ObjectID) (*types.User, error) {
// 	return _, nil
// }

func (s *MongoUserStore) GetUsers(c context.Context) ([]*types.User, error) {
	var allUsers []*types.User
	cur, err := s.coll.Find(c, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(c)
	for cur.Next(c) {
		var result *types.User
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		allUsers = append(allUsers, result)
		// do something with result....
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return allUsers, nil
}

func (s *MongoUserStore) GetUserByID(c context.Context, objID primitive.ObjectID) (*types.User, error) {
	var user *types.User
	filter := bson.M{"_id": objID}
	err := s.coll.FindOne(c, filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}
	return user, nil
}

func (s *MongoUserStore) UpdateUser(c context.Context, filter bson.M, update bson.M) error {
	_, err := s.coll.UpdateOne(c, filter, update)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}

	return err
}
