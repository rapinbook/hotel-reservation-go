package main

import (
	"context"
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
	mongoClient := db.NewMongoUserStore(client)
	_, err = mongoClient.InsertUser(ctx, &types.User{
		Email:     "foo@foobar.com",
		FirstName: "Josh",
		LastName:  "Gunn",
	})
	if err != nil {
		log.Fatalf("Cannot inserted row to databases")
	}
	// fmt.Printf()("Inserted row first run %v", insertedRow.Hex())

}
