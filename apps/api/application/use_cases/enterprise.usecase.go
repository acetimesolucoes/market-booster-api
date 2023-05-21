package use_cases

import (
	"marketbooster/application/repository"
	"marketbooster/domain"
)

type EnterpriseUseCase struct{}

func (uc *EnterpriseUseCase) FindAll(page int64, limit int64) (domain.Enterprises, error) {

	enterprises, err := new(repository.EnterpriseRepository).FindAll(page, limit)

	if err != nil {
		return nil, err
	}

	return enterprises, nil
}

func (uc *EnterpriseUseCase) FindOneById(enterpriseId string) (domain.Enterprise, error) {

	enterprise, err := new(repository.EnterpriseRepository).FindOneById(enterpriseId)

	if err != nil {
		return enterprise, err
	}

	return enterprise, nil
}

func (uc *EnterpriseUseCase) Create(enterprise domain.Enterprise) error {

	err := new(repository.EnterpriseRepository).Save(enterprise)

	if err != nil {
		return err
	}

	return nil
}

func (uc *EnterpriseUseCase) Update(enterpriseId string, enterprise domain.Enterprise) error {

	err := new(repository.EnterpriseRepository).Update(enterpriseId, enterprise)

	if err != nil {
		return err
	}

	return nil
}

func (uc *EnterpriseUseCase) Delete(enterpriseId string) error {

	err := new(repository.EnterpriseRepository).Delete(enterpriseId)

	if err != nil {
		return err
	}

	return nil
}
