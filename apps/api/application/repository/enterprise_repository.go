package repository

import (
	// "context"
	"fmt"
	// "log"
	// "os"

	// "acetime.com.br/business-crm/apps/api/domain"
	// "acetime.com.br/business-crm/apps/api/framework/utils"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
)

// var collection *mongo.Collection

func init() {
	fmt.Println("Starting enterprise repository")
	// databaseName := os.Getenv("MONGO_DB_NAME")
	// client, err := utils.GetClientDB()
	// if err != nil {
	// 	collection = client.Database(databaseName).Collection("enterprises")
	// }

}

// func Find() (nil error) {
// 	cur, err := collection.Find(context.Background(), bson.D{})

// 	if err != nil {
// 		return err
// 	}

// 	defer cur.Close(context.Background())

// 	for cur.Next(context.Background()) {
// 		result := struct {
// 			Foo string
// 			Bar int32
// 		}{}

// 		err := cur.Decode(&result)

// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		raw := cur.Current

// 		fmt.Println(raw)
// 	}

// 	if err := cur.Err(); err != nil {
// 		return err
// 	}

// 	return nil
// }

// type EnterpriseRepository interface {
// 	Insert(enterprise *domain.Enterprise) (*domain.Enterprise, error)
// 	Find() error
// }
