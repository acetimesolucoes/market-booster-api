package repository

import (
	"context"
	"fmt"
	"time"

	"marketbooster/domain"
	"marketbooster/framework/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EnterpriseRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

func (r *EnterpriseRepository) init() {
	r.collection = utils.GetCollection("enterprises")
	r.ctx = context.Background()
}

func (r *EnterpriseRepository) FindAll(page int64, limit int64) (domain.Enterprises, error) {

	var enterprises domain.Enterprises

	var err error
	filter := bson.D{}

	opts := new(options.FindOptions)
	opts.Limit = &limit
	calc := (page-1)*limit + 1
	opts.Skip = &calc

	cursor, err := r.collection.Find(r.ctx, filter, opts)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for cursor.Next(r.ctx) {

		var enterprise domain.Enterprise
		err = cursor.Decode(&enterprise)

		if err != nil {
			return nil, err
		}

		enterprises = append(enterprises, &enterprise)
	}

	return enterprises, nil
}

func (r *EnterpriseRepository) FindOneById(enterpriseId string) (domain.Enterprise, error) {

	var err error

	oid, err := primitive.ObjectIDFromHex(enterpriseId)

	if err != nil {
		return domain.Enterprise{}, err
	}

	filter := bson.M{"_id": oid}

	result := r.collection.FindOne(r.ctx, filter)

	var enterprise domain.Enterprise
	err = result.Decode(&enterprise)

	if err != nil {
		return enterprise, err
	}

	return enterprise, nil

}

func (r *EnterpriseRepository) Save(enterprise domain.Enterprise) error {

	var err error

	enterprise.CreatedAt = time.Now()

	_, err = r.collection.InsertOne(r.ctx, enterprise)

	if err != nil {
		return err
	}

	return nil
}

func (r *EnterpriseRepository) Update(enterpriseId string, enterprise domain.Enterprise) error {

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

	_, err = r.collection.UpdateOne(r.ctx, filter, updated)

	if err != nil {
		return err
	}

	return nil
}

func (r *EnterpriseRepository) Delete(enterpriseId string) error {

	var err error
	var oid primitive.ObjectID

	oid, err = primitive.ObjectIDFromHex(enterpriseId)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": oid}

	_, err = r.collection.DeleteOne(r.ctx, filter)

	if err != nil {
		return err
	}

	return nil
}
