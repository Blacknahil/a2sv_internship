// djks
package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func IntializeMongoBD() mongo.Database {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoURI := os.Getenv("MONGODB_URI")

	if mongoURI == "" {
		log.Fatal("MONGO_URI not set in the .env file")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully Connected to MongoDB!")

	userCollection := client.Database("TaskAPIdb3").Collection("users")
	client.Database("TaskAPIdb3").Collection("tasks")

	err = createUniqueFieldIndex(ctx, userCollection, "email")
	if err != nil {
		log.Fatalf("Failed to create unique index on email: %v", err)
	}
	return *client.Database("TaskAPIdb3")

}

func createUniqueFieldIndex(ctx context.Context, collection *mongo.Collection, uniqueField string) error {

	// create the unique index on the email field

	indexModel := mongo.IndexModel{
		Keys:    bson.M{uniqueField: 1}, // create an index on the unique field
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}
	fmt.Printf("Unique index on '%v' created succesfully! \n", uniqueField)
	return nil
}
