package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDbClient() *mongo.Client {

	options := options.Client().ApplyURI("mongodb://root:root@localhost:27017?connect=direct")

	client, err := mongo.Connect(context.TODO(), options)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	return client
}
