package repository

import (
	"context"
	"time"
	"fmt"

	"marketbooster/domain"
	"marketbooster/framework/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {

}

var collection = utils.GetCollection("users")
var ctx = context.Background()

func (r *UserRepository) FindAll(page int64, limit int64) (domain.Users, error) {

	var users domain.Users

	var err error
	filter := bson.D{}

	opts := new(options.FindOptions)
	opts.Limit = &limit
	calc := (page-1)*limit + 1
	opts.Skip = &calc

	cursor, err := collection.Find(ctx, filter, opts)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for cursor.Next(ctx) {

		var user domain.Enterprise
		err = cursor.Decode(&user)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (r *UserRepository) FindOneById(userId string) (domain.Enterprise, error) {

	var err error

	oid, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return domain.Enterprise{}, err
	}

	filter := bson.M{"_id": oid}

	result := collection.FindOne(ctx, filter)

	var user domain.Enterprise
	err = result.Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil

}

func (r *UserRepository) Save(user domain.Enterprise) error {

	var err error

	user.CreatedAt = time.Now()

	_, err = collection.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Update(userId string, user domain.Enterprise) error {

	var err error

	oid, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.M{"_id": oid}

	// todo: concluir
	updated := bson.M{
		"$set": bson.M{
			"full_name": 		user.FullName,
			"email":         	user.Email,
			"enterprise_id": 	user.EnterpriseId,
			"updated_at":   	time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, filter, updated)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(userId string) error {

	var err error
	var oid primitive.ObjectID

	oid, err = primitive.ObjectIDFromHex(userId)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": oid}

	_, err = collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}
