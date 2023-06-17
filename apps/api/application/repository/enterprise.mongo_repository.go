package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/acetimesolutions/marketbooster/domain"
	"github.com/acetimesolutions/marketbooster/framework/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EnterpriseMongoRepository struct {
	Collection *mongo.Collection
	Context    context.Context
}

func NewEnterpriseMongoRepository() *EnterpriseMongoRepository {
	return &EnterpriseMongoRepository{
		Collection: utils.GetCollection("enterprises"),
		Context:    context.Background(),
	}
}

func (r *EnterpriseMongoRepository) FindAll(page int64, limit int64) ([]*domain.Enterprise, error) {

	var enterprises []*domain.Enterprise

	var err error
	filter := bson.D{}

	opts := new(options.FindOptions)
	opts.Limit = &limit
	calc := (page-1)*limit + 1
	opts.Skip = &calc

	cursor, err := r.Collection.Find(r.Context, filter, opts)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for cursor.Next(r.Context) {

		var enterprise domain.Enterprise
		err = cursor.Decode(&enterprise)

		if err != nil {
			return nil, err
		}

		enterprises = append(enterprises, &enterprise)
	}

	return enterprises, nil
}

func (r *EnterpriseMongoRepository) FindOneById(enterpriseId string) (domain.Enterprise, error) {

	var err error

	oid, err := primitive.ObjectIDFromHex(enterpriseId)

	if err != nil {
		return domain.Enterprise{}, err
	}

	filter := bson.M{"_id": oid}

	result := r.Collection.FindOne(r.Context, filter)

	var enterprise domain.Enterprise
	err = result.Decode(&enterprise)

	if err != nil {
		return enterprise, err
	}

	return enterprise, nil

}

func (r *EnterpriseMongoRepository) Save(enterprise domain.Enterprise) error {

	var err error

	enterprise.CreatedAt = time.Now()

	_, err = r.Collection.InsertOne(r.Context, enterprise)

	if err != nil {
		return err
	}

	return nil
}

func (r *EnterpriseMongoRepository) Update(enterpriseId string, enterprise domain.Enterprise) error {

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

	_, err = r.Collection.UpdateOne(r.Context, filter, updated)

	if err != nil {
		return err
	}

	return nil
}

func (r *EnterpriseMongoRepository) Delete(enterpriseId string) error {

	var err error
	var oid primitive.ObjectID

	oid, err = primitive.ObjectIDFromHex(enterpriseId)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": oid}

	_, err = r.Collection.DeleteOne(r.Context, filter)

	if err != nil {
		return err
	}

	return nil
}
