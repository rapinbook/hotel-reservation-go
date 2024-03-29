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

type HotelStore interface {
	InsertUser(context.Context, *types.Hotel) (primitive.ObjectID, error)
	GetUserByID(context.Context, primitive.ObjectID) (*types.Hotel, error)
	GetUsers(context.Context) ([]*types.Hotel, error)
	UpdateUser(context.Context) (primitive.ObjectID, error)
}

type RoomStore interface {
	InsertUser(context.Context, *types.Room) (primitive.ObjectID, error)
	GetUserByID(context.Context, primitive.ObjectID) (*types.Room, error)
	GetUsers(context.Context) ([]*types.Room, error)
	UpdateUser(context.Context) (primitive.ObjectID, error)
}

type MongoHotelStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

type MongoRoomStore struct {
	client *mongo.Client
	coll   *mongo.Collection
	MongoHotelStore
}

func NewMongoHotelStore(client *mongo.Client) *MongoHotelStore {
	return &MongoHotelStore{
		client: client,
		coll:   client.Database("hotel-reservation").Collection("hotel"),
	}
}

func NewMongoRoomStore(client *mongo.Client, hotelStore MongoHotelStore) *MongoRoomStore {
	return &MongoRoomStore{
		client:          client,
		coll:            client.Database("hotel-reservation").Collection("room"),
		MongoHotelStore: hotelStore,
	}
}

func (s *MongoHotelStore) InsertHotel(c context.Context, hotel *types.Hotel) (primitive.ObjectID, error) {
	res, err := s.coll.InsertOne(c, hotel)
	if err != nil {
		fmt.Printf("Cannot insert hotel %v", hotel)
	}
	id := res.InsertedID.(primitive.ObjectID)
	return id, nil
}

func (s *MongoHotelStore) GetHotels(c context.Context) ([]*types.Hotel, error) {
	var allHotel []*types.Hotel
	cur, err := s.coll.Find(c, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(c)
	for cur.Next(c) {
		var result *types.Hotel
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		allHotel = append(allHotel, result)
		// do something with result....
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return allHotel, nil
}

func (s *MongoHotelStore) GetHotelByID(c context.Context, objID primitive.ObjectID) (*types.Hotel, error) {
	var hotel *types.Hotel
	filter := bson.M{"_id": objID}
	err := s.coll.FindOne(c, filter).Decode(&hotel)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}
	return hotel, nil
}

func (s *MongoHotelStore) UpdateHotel(c context.Context, filter bson.M, update bson.M) error {
	_, err := s.coll.UpdateOne(c, filter, update)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}

	return err
}

func (s *MongoRoomStore) InsertRoom(c context.Context, room *types.Room) (primitive.ObjectID, error) {
	hotelID := room.HotelID
	// update []*Room inside hotel
	hotelRoom, err := s.MongoHotelStore.GetHotelByID(c, hotelID)
	if err != nil {
		log.Fatal(err)
	}
	updatedHotelRoom := append(hotelRoom.Room, room)
	filter := bson.M{"_id": hotelID}
	update := bson.M{"$set": bson.M{"room": updatedHotelRoom}}
	err = s.MongoHotelStore.UpdateHotel(c, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	res, err := s.coll.InsertOne(c, room)
	if err != nil {
		fmt.Printf("Cannot insert hotel %v", room)
	}
	id := res.InsertedID.(primitive.ObjectID)
	return id, nil
}

func (s *MongoRoomStore) GetRoom(c context.Context) ([]*types.Room, error) {
	var allRoom []*types.Room
	cur, err := s.coll.Find(c, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(c)
	for cur.Next(c) {
		var result *types.Room
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		allRoom = append(allRoom, result)
		// do something with result....
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return allRoom, nil
}

func (s *MongoRoomStore) GetRoomByID(c context.Context, objID primitive.ObjectID) (*types.Room, error) {
	var room *types.Room
	filter := bson.M{"_id": objID}
	err := s.coll.FindOne(c, filter).Decode(&room)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}
	return room, nil
}

func (s *MongoRoomStore) UpdateRoom(c context.Context, filter bson.M, update bson.M) error {
	_, err := s.coll.UpdateOne(c, filter, update)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}

	return err
}
