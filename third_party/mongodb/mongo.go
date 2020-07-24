package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var cli *mongo.Client
var db *mongo.Database

// Connect to MongoDB
func Connect(uri string) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Ping right after connected
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	cli = client
	db = client.Database("kai-auth")
}

// GetDB return db instance
func GetDB() *mongo.Database {
	return db
}

// Disconnect from MongoDB
func Disconnect() {
	err := cli.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed!")
}
