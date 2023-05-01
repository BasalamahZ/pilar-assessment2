package usecase

import (
	"assessment2/pkg/user/dto"
	"assessment2/pkg/user/model"
	"assessment2/pkg/user/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UsecaseInterfaceUser interface {
	Register(req dto.UserDTO) error
	ViewProfile(username string) (model.User, error)
	CreateTokenUser(req dto.UserLogin) (model.User, error)
	GetUserById(id int) (model.User, error)
}

type usecaseUser struct {
	repository repository.RepositoryInterfaceUser
}

func InitUsecaseUser(repository repository.RepositoryInterfaceUser) UsecaseInterfaceUser {
	return &usecaseUser{
		repository: repository,
	}
}

// GetUserById implements UsecaseInterfaceUser
func (u *usecaseUser) GetUserById(id int) (model.User, error) {
	user, err := u.repository.GetUserById(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

// CreateTokenUser implements UsecaseInterfaceUser
func (u *usecaseUser) CreateTokenUser(req dto.UserLogin) (model.User, error) {
	user, err := u.repository.GetUserByEmail(req.Username)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return user, errors.New("wrong password")
	}

	return user, nil
}

// ViewProfile implements UsecaseInterfaceUser
func (u *usecaseUser) ViewProfile(username string) (model.User, error) {
	user, err := u.repository.GetUserByEmail(username)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Register implements UsecaseInterfaceUser
func (u *usecaseUser) Register(req dto.UserDTO) error {
	isUserExist, _ := u.repository.GetUserByEmail(req.Username)

	if isUserExist.ID != 0 {
		return errors.New("user already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	req.Password = string(hashedPassword)
	err_ := u.repository.CreateNewUser(req)
	if err_ != nil {
		return err_
	}

	return nil
}
