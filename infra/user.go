package infra

import (
	"ddd-sample/domain/model"
	"ddd-sample/domain/repository"

	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) repository.UserRepository {
	return &UserRepository{Conn: conn}
}

func (ur *UserRepository) Create(user *model.User) (*model.User, error) {
	err := ur.Conn.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) FindByID(id int) (*model.User, error) {
	var user model.User
	err := ur.Conn.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (ur *UserRepository) FindAll() (*[]model.User, error) {
	var user model.User
	err := ur.Conn.Find(&user).Error
	if err != nil {
		return nil, err
	}

	return &[]model.User{}, err
}

func (ur *UserRepository) Update(user *model.User) (*model.User, error) {
	err := ur.Conn.Model(&user).Updates(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) Delete(user *model.User) error {
	err := ur.Conn.Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}
