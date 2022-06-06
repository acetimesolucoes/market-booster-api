// https://www.youtube.com/watch?v=JP-D1In0juw&t=825s

package utils

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var (
// 	usr  = os.Getenv("MONGO_DB_NAME")
// 	pwd  = os.Getenv("MONGO_DB_PWD")
// 	host = os.Getenv("MONGO_DB_HOST")
// 	port = os.Getenv("MONGO_DB_PORT")
// 	db   = os.Getenv("MONGO_DB_DBNAME")
// )

var (
	usr  = "root"
	pwd  = "acetimesolucoes"
	host = "localhost"
	port = "27017"
	db   = "actm-business-crm"
)

func GetCollection(collection string) *mongo.Collection {

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", usr, pwd, host, port)

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
