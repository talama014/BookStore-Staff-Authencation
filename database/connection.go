package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectUsers() *mongo.Collection {
	var client *mongo.Client
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:expamle@mongo:27017/"))

	clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:27017/?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("bookstore").Collection("user")
	return collection
}
