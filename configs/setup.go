package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(EnvMongoURI()))

	if err != nil {
		log.Fatal("Error: ",err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Println("Connect to MongoDB")

	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("ManagementServer").Collection(collectionName)
	return collection
}