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
	host = "127.0.0.1"
	port = "27017"
	db   = "actm-business-crm"
)

func GetCollection(collection string) *mongo.Collection {

	fmt.Println("usr: ", usr)
	fmt.Println("pwd: ", pwd)
	fmt.Println("host: ", host)
	fmt.Println("port: ", port)

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin", usr, pwd, host, port, db)

	fmt.Println(uri)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	err = client.Ping(context.TODO(), nil)
	defer client.Disconnect(ctx)

	fmt.Println(err)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to MongoDB!")

	return client.Database(db).Collection(collection)
}
