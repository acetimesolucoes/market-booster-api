package utils

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	usr  = os.Getenv("MONGO_DB_NAME")
	pwd  = os.Getenv("MONGO_DB_PWD")
	host = os.Getenv("MONGO_DB_HOST")
	port = os.Getenv("MONGO_DB_PORT")
	db   = os.Getenv("MONGO_DB_DBNAME")
)

func GetCollection(collection string) *mongo.Collection {

	port, err := strconv.ParseInt(port, 0, 32)

	if err != nil {
		fmt.Print(err)
	}

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", usr, pwd, host, port)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	return client.Database(db).Collection(collection)
}
