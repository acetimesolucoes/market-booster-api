package usecases

import (
	"github.com/acetimesolutions/marketbooster/application/repository"
	"github.com/acetimesolutions/marketbooster/domain"
)

type EnterpriseUsecase struct {
	EnterpriseRepository repository.Repository[domain.Enterprise]
}

func NewUsecase(enterpriseRepository repository.Repository[domain.Enterprise]) EnterpriseUsecase {
	return EnterpriseUsecase{
		EnterpriseRepository: enterpriseRepository,
	}
}

func (u *EnterpriseUsecase) FindAll(page int64, limit int64) ([]*domain.Enterprise, error) {

	enterprises, err := u.EnterpriseRepository.FindAll(page, limit)

	if err != nil {
		return nil, err
	}

	return enterprises, nil
}

func (u *EnterpriseUsecase) FindOneById(enterpriseId string) (domain.Enterprise, error) {

	enterprise, err := u.EnterpriseRepository.FindOneById(enterpriseId)

	if err != nil {
		return enterprise, err
	}

	return enterprise, nil
}

func (u *EnterpriseUsecase) Create(enterprise domain.Enterprise) error {

	err := u.EnterpriseRepository.Save(enterprise)

	if err != nil {
		return err
	}

	return nil
}

func (u *EnterpriseUsecase) Update(enterpriseId string, enterprise domain.Enterprise) error {

	err := u.EnterpriseRepository.Update(enterpriseId, enterprise)

	if err != nil {
		return err
	}

	return nil
}

func (u *EnterpriseUsecase) Delete(enterpriseId string) error {

	err := u.EnterpriseRepository.Delete(enterpriseId)

	if err != nil {
		return err
	}

	return nil
}
