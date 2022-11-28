package usecase

import (
	"ddd-sample/domain/model"
	"ddd-sample/domain/repository"
)

type UserUsecase interface {
	Create(username string) (*model.User, error)
	FindByID(id int) (*model.User, error)
	FindAll() (*[]model.User, error)
	Update(id int, username string) (*model.User, error)
	Delete(id int) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepository: userRepo}
}

func (uc *userUsecase) Create(username string) (*model.User, error) {
	user, err := model.NewUser(username)
	if err != nil {
		return nil, err
	}

	createdUser, err := uc.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (uc *userUsecase) FindByID(id int) (*model.User, error) {
	foundUser, err := uc.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}

func (uc *userUsecase) FindAll() (*[]model.User, error) {
	foundUser, err := uc.userRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}

func (uc *userUsecase) Update(id int, username string) (*model.User, error) {
	targetUser, err := uc.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = targetUser.Set(username)
	if err != nil {
		return nil, err
	}

	updatedUser, err := uc.userRepository.Update(targetUser)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (uc *userUsecase) Delete(id int) error {
	targetUser, err := uc.userRepository.FindByID(id)
	if err != nil {
		return err
	}

	err = uc.userRepository.Delete(targetUser)
	if err != nil {
		return err
	}

	return nil
}
