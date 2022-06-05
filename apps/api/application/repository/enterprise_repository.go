package repository

import (
	"context"

	"acetime.com.br/business-crm/apps/api/domain"
	"acetime.com.br/business-crm/apps/api/framework/utils"
)

var collection = utils.GetCollection("enterprises")
var ctx = context.Background()

func Save(enterprise domain.Enterprise) error {

	var err error

	_, err = collection.InsertOne(ctx, enterprise)

	if err != nil {
		return err
	}

	return nil
}

func Find() (domain.Enterprises, error) {
	return nil, nil
}

func Update(enterpriseId string, enterprise domain.Enterprise) error {
	return nil
}

func Delete(enterpriseId string) error {
	return nil
}
