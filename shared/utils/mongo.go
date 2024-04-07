package utils

import (
	"bybu/go-mongo-db/shared/config"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongodb() *mongo.Client {
	config.VerifyEnvVariable();

	ctx, ctxErr := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxErr()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Env.GetMongoUrl()))
	if (err != nil) {
		log.Fatal("Error in connecting to database")
	}

	err = client.Ping(ctx, nil);
	if (err != nil) {
		log.Fatal("Error in pinging database")
	}

	fmt.Println("Connected to database successfully")
	return client;
}

var DB *mongo.Client = ConnectToMongodb()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("Paydeal").Collection(collectionName)
	return collection;
}