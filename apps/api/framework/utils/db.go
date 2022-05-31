// https://www.youtube.com/watch?v=JP-D1In0juw&t=825s
package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var ctx = context.TODO()

func ConnectDB() {

	fmt.Println("Starting connection with mongodb...")

	connectionString := os.Getenv("MONGO_STRING_CONNECTION")
	fmt.Println(connectionString)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// collection = client.Database("tasker").Collection("tasks")
}

func GetClientDB() (*mongo.Client, error) {
	connectionString := os.Getenv("MONGO_STRING_CONNECTION")
	return mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
}
