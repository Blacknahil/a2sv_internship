package data

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

var (
	UserCollection *mongo.Collection
	TaskCollection *mongo.Collection
)

func IntializeMongoBD() {

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

	UserCollection = client.Database("taskdb2").Collection("users")
	TaskCollection = client.Database("taskdb2").Collection("tasks")

	err = createUniqueFieldIndex(ctx, UserCollection, "email")
	if err != nil {
		log.Fatalf("Failed to create unique index on email: %v", err)
	}

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

//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjMONGODB_URIk
// /dbjjhdsjhjhd
// jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbdjksjkjksjk
// adhjashjahjs
// /jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjshjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjdfjhdfsjhjhds
// sdjk
