package repository

import (
	"assessment2/pkg/user/dto"
	"assessment2/pkg/user/model"

	"gorm.io/gorm"
)

type RepositoryInterfaceUser interface {
	GetUserByEmail(username string) (model.User, error)
	GetUserById(id int) (model.User, error)
	CreateNewUser(user dto.UserDTO) error
}

type repositoryUser struct {
	db *gorm.DB
}

func InitRepositoryUser(db *gorm.DB) RepositoryInterfaceUser {
	db.AutoMigrate(&model.User{})
	return &repositoryUser{
		db: db,
	}
}

// GetUserById implements RepositoryInterfaceUser
func (r *repositoryUser) GetUserById(id int) (model.User, error) {
	var user model.User
	if err := r.db.Table("users").Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// GetUserByEmail implements RepositoryInterfaceUser
func (r *repositoryUser) GetUserByEmail(username string) (model.User, error) {
	var user model.User
	if err := r.db.Table("users").Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// Create New User implements RepositoryInterfaceUser
func (r *repositoryUser) CreateNewUser(user dto.UserDTO) error {
	if err := r.db.Table("users").Create(&user).Error; err != nil {
		return err
	}

	return nil
}
