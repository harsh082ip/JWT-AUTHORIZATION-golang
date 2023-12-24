package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func OpenCollection(client *mongo.Client, collectionname string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("notes").Collection(collectionname)
	return collection

}

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading a .env file")
	}

	MongoDb := os.Getenv("MONGODB_URL")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	// No need for a separate connection attempt, as mongo.Connect already establishes the connection.

	fmt.Println("Connected to MongoDB...")

	return client
}

var Client *mongo.Client = DBinstance()
