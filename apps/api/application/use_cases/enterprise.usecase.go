package enterprise_use_cases

import (
	enterpriseRepository "github.com/acetimesolutions/marketbooster/application/repository"
	"github.com/acetimesolutions/marketbooster/domain"
)

func FindAll(page int64, limit int64) (domain.Enterprises, error) {

	enterprises, err := enterpriseRepository.FindAll(page, limit)

	if err != nil {
		return nil, err
	}

	return enterprises, nil
}

func FindOneById(enterpriseId string) (domain.Enterprise, error) {

	enterprise, err := enterpriseRepository.FindOneById(enterpriseId)

	if err != nil {
		return enterprise, err
	}

	return enterprise, nil
}

func Create(enterprise domain.Enterprise) error {

	err := enterpriseRepository.Save(enterprise)

	if err != nil {
		return err
	}

	return nil
}

func Update(enterpriseId string, enterprise domain.Enterprise) error {

	err := enterpriseRepository.Update(enterpriseId, enterprise)

	if err != nil {
		return err
	}

	return nil
}

func Delete(enterpriseId string) error {

	err := enterpriseRepository.Delete(enterpriseId)

	if err != nil {
		return err
	}

	return nil
}
