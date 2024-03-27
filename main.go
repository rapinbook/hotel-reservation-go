package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/rapinbook/hotel-reservation-go/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	// _, err = mongoClient.InsertUser(ctx, &types.User{
	// 	Email:     "John@foobar.com",
	// 	FirstName: "John",
	// 	LastName:  "Gunn",
	// })
	// if err != nil {
	// 	log.Fatalf("Cannot inserted row to databases")
	// }
	// _, err = mongoClient.InsertUser(ctx, &types.User{
	// 	Email:     "James@foobar.com",
	// 	FirstName: "James",
	// 	LastName:  "Gunn",
	// })
	// if err != nil {
	// 	log.Fatalf("Cannot inserted row to databases")
	// }
	// _, err = mongoClient.InsertUser(ctx, &types.User{
	// 	Email:     "Josh@foobar.com",
	// 	FirstName: "Josh",
	// 	LastName:  "Gunn",
	// })
	// if err != nil {
	// 	log.Fatalf("Cannot inserted row to databases")
	// }
	allUsers, err := mongoClient.GetUsers(ctx)
	if err != nil {
		fmt.Println(err)
	}
	// _, err := json.Marshal(allUsers)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	for _, s := range allUsers {
		b, err := json.Marshal(s)
		fmt.Println(string(b))
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Above show all rows in system")
	objID, err := primitive.ObjectIDFromHex("6603dfdcff4c801f4a9e6599")
	if err != nil {
		panic(err)
	}
	user, err := mongoClient.GetUserByID(ctx, objID)
	if err != nil {
		fmt.Println(err)
	}
	b, err := json.Marshal(user)
	fmt.Println(string(b))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Above show row filter from obj")
}
