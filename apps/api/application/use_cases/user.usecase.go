package use_cases

import (
	"marketbooster/application/repository"
	"marketbooster/domain"
)

type UserUseCase struct {
}

func (uc *UserUseCase) FindAll(page int64, limit int64) (domain.Users, error) {

	users, err := new(repository.UserRepository).FindAll(page, limit)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uc *UserUseCase) FindOneById(userId string) (domain.User, error) {

	user, err := new(repository.UserRepository).FindOneById(userId)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (uc *UserUseCase) Create(user domain.User) error {

	err := new(repository.UserRepository).Save(user)

	if err != nil {
		return err
	}

	return nil
}

func (uc *UserUseCase) Update(userId string, user domain.User) error {

	err := new(repository.UserRepository).Update(userId, user)

	if err != nil {
		return err
	}

	return nil
}

func (uc *UserUseCase) Delete(userId string) error {

	err := new(repository.UserRepository).Delete(userId)

	if err != nil {
		return err
	}

	return nil
}
