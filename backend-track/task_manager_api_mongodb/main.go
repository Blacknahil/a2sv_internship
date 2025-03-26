package main

import (
	"context"
	"fmt"
	"log"
	"task_manager_api/router"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
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

	fmt.Println("Connected to MongoDb!")

	r := router.SetUpRouter(client)
	fmt.Println("Task Manager API is running on http://localhost:8080")
	r.Run(":8080")

}

//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjdj sj jk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
