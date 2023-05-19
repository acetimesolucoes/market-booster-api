package user_use_cases

import (
	userRepository "marketbooster/application/repository"
	"marketbooster/domain"
)

func FindAll(page int64, limit int64) (domain.Users, error) {

	users, err := userRepository.FindAll(page, limit)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func FindOneById(userId string) (domain.Users, error) {

	user, err := userRepository.FindOneById(userId)

	if err != nil {
		return user, err
	}

	return user, nil
}

func Create(user domain.Users) error {

	err := userRepository.Save(user)

	if err != nil {
		return err
	}

	return nil
}

func Update(userId string, user domain.Users) error {

	err := userRepository.Update(userId, user)

	if err != nil {
		return err
	}

	return nil
}

func Delete(userId string) error {

	err := userRepository.Delete(userId)

	if err != nil {
		return err
	}

	return nil
}
