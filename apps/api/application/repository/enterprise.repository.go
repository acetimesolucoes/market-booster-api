package repository

import (
	"context"
	"time"

	"acetime.com.br/business-crm/apps/api/domain"
	"acetime.com.br/business-crm/apps/api/framework/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection = utils.GetCollection("enterprises")
var ctx = context.Background()

func FindAll(page int64, limit int64) (domain.Enterprises, error) {

	var enterprises domain.Enterprises

	var err error
	filter := bson.D{}

	opts := new(options.FindOptions)
	opts.BatchSize = limit

	cursor, err := collection.Find(ctx, filter, opts)

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {

		var enterprise domain.Enterprise
		err = cursor.Decode(&enterprise)

		if err != nil {
			return nil, err
		}

		enterprises = append(enterprises, &enterprise)
	}

	return enterprises, nil
}

func FindOneById(enterpriseId string) (domain.Enterprise, error) {

	var err error

	oid, err := primitive.ObjectIDFromHex(enterpriseId)

	if err != nil {
		return domain.Enterprise{}, err
	}

	filter := bson.M{"_id": oid}

	result := collection.FindOne(ctx, filter)

	var enterprise domain.Enterprise
	err = result.Decode(&enterprise)

	if err != nil {
		return enterprise, err
	}

	return enterprise, nil

}

func Save(enterprise domain.Enterprise) error {

	var err error

	enterprise.CreatedAt = time.Now()

	_, err = collection.InsertOne(ctx, enterprise)

	if err != nil {
		return err
	}

	return nil
}

func Update(enterpriseId string, enterprise domain.Enterprise) error {

	var err error

	oid, _ := primitive.ObjectIDFromHex(enterpriseId)

	filter := bson.M{"_id": oid}

	// todo: concluir
	updated := bson.M{
		"$set": bson.M{
			"business_name": enterprise.BusinessName,
			"CNAE":          enterprise.CNAE,
			"fantasy_name":  enterprise.FantasyName,
			"is_filial":     enterprise.IsFilial,
			"updated_at":    time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, filter, updated)

	if err != nil {
		return err
	}

	return nil
}

func Delete(enterpriseId string) error {

	var err error
	var oid primitive.ObjectID

	oid, err = primitive.ObjectIDFromHex(enterpriseId)

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
