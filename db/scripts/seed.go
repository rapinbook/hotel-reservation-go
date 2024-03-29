package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rapinbook/hotel-reservation-go/db"
	"github.com/rapinbook/hotel-reservation-go/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Database("hotel-reservation").Drop(ctx); err != nil {
		log.Fatal(err)
	}
	mongoUserStore := db.NewMongoUserStore(client)
	_, err = mongoUserStore.InsertUser(ctx, &types.User{
		Email:     "John@foobar.com",
		FirstName: "John",
		LastName:  "Gunn",
	})
	if err != nil {
		log.Fatalf("Cannot inserted row to databases")
	}
	_, err = mongoUserStore.InsertUser(ctx, &types.User{
		Email:     "admin@foobar.com",
		FirstName: "admin",
		LastName:  "admin",
	})
	if err != nil {
		log.Fatalf("Cannot inserted row to databases")
	}
	mongoHotelStore := db.NewMongoHotelStore(client)
	_, err = mongoHotelStore.InsertHotel(ctx, &types.Hotel{
		HotelName: "some hotel",
		City:      "bermuda",
		Room:      nil,
		Rating:    5,
	})
	if err != nil {
		log.Fatalf("Cannot inserted hotel row to databases")
	}
	parisHotelID, err := mongoHotelStore.InsertHotel(ctx, &types.Hotel{
		HotelName: "Any hotel",
		City:      "Paris",
		Room:      nil,
		Rating:    2.0,
	})
	if err != nil {
		log.Fatalf("Cannot inserted hotel row to databases")
	}
	mongoRoomStore := db.NewMongoRoomStore(client, *mongoHotelStore)
	roomID, err := mongoRoomStore.InsertRoom(ctx, &types.Room{
		HotelID: parisHotelID,
		Price:   747.42,
		Size:    "small",
	})
	if err != nil {
		log.Fatalf("Cannot inserted room row to databases")
	}
	fmt.Println(roomID.Hex())
	roomNumTwoID, err := mongoRoomStore.InsertRoom(ctx, &types.Room{
		HotelID: parisHotelID,
		Price:   4200.42,
		Size:    "large",
	})
	if err != nil {
		log.Fatalf("Cannot inserted room row to databases")
	}
	fmt.Println(roomNumTwoID.Hex())
}
