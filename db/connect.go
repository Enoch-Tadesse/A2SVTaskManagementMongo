package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var DBName = "manager"

func ConnectToDB() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// set timeout to 5 second
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Connect to MongoDB
	var err error
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// ping the db to make sure it works
	if err := Client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Database Connected")

}
